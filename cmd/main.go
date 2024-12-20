package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"manabase-simulation/package/logging"
	"manabase-simulation/package/simulation"
	"manabase-simulation/package/util/test"

	"log"
	"manabase-simulation/api"
	"manabase-simulation/package/model"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 8889, "The server port")
)

const (
	// CheckpointInterval is the number of iterations between checkpoints.
	CheckpointInterval = 100

	// TotalCheckpoints is the total number of checkpoints to generate.
	TotalCheckpoints = 30
)

//cfg := postgres.Config{
//		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", address, username, password, dbname, port),
//	}

func main() {
	flag.Parse()
	log.Println(fmt.Sprintf("Starting Listening on port %d", *port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	api.RegisterManabaseSimulatorServer(grpcServer, newManabaseSimulatorServer(test.GetDBConfig()))
	grpc_health_v1.RegisterHealthServer(grpcServer, newHealthServer())
	reflection.Register(grpcServer)
	log.Println("Serving gRPC traffic now")
	grpcServer.Serve(lis)
}

func simulate(ctx context.Context, decklist model.DeckList, configuration model.GameConfiguration, objective model.TestObjective) []model.ResultCheckpoint {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	logger := logging.CreateLogger()
	logger.Info(decklist.ToString())
	logger.Info(objective.ToString())

	now := time.Now()

	successCount := 0
	iterations := CheckpointInterval * TotalCheckpoints

	resultChannel := make(chan bool, 100)
	wg := new(sync.WaitGroup)

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go start(ctx, decklist, configuration, objective, resultChannel, wg)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	completedIterations := 0
	checkpoints := make([]model.ResultCheckpoint, TotalCheckpoints)
	for result := range resultChannel {
		completedIterations++
		if result {
			successCount++
		}

		if completedIterations%CheckpointInterval == 0 {
			checkpoints[(completedIterations/CheckpointInterval)-1] = model.ResultCheckpoint{
				Iterations: int32(completedIterations),
				Successes:  int32(successCount),
			}
		}

	}

	// Capture results to be consumes.
	logger.Info(fmt.Sprintf("Success count: %d", checkpoints[len(checkpoints)-1].Successes))
	logger.Info(fmt.Sprintf("Success Rate: %f", checkpoints[len(checkpoints)-1].GetSuccessRate()))
	logger.Info(fmt.Sprintf("Checkpoints: %v", checkpoints))
	logger.Info(fmt.Sprintf("Time taken: %s", time.Since(now)))

	return checkpoints
}

func start(ctx context.Context, deckList model.DeckList, gameConfiguration model.GameConfiguration, objective model.TestObjective, c chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- simulation.SimulateDeck(ctx, deckList, gameConfiguration, objective)
}
