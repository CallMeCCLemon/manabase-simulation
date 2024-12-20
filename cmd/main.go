package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"manabase-simulation/package/facade"
	"manabase-simulation/package/logging"
	"manabase-simulation/package/simulation"
	"manabase-simulation/package/util/test"
	"manabase-simulation/package/validation"

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

type manabaseSimulatorServer struct {
	api.UnimplementedManabaseSimulatorServer

	mu sync.Mutex // protects routeNotes

	cfg postgres.Config
}

func newManabaseSimulatorServer(cfg postgres.Config) *manabaseSimulatorServer {
	s := &manabaseSimulatorServer{
		cfg: cfg,
	}
	return s
}

func newHealthServer() *health.Server {
	s := health.NewServer()
	return s
}

func (s *manabaseSimulatorServer) SimulateDeck(ctx context.Context, in *api.SimulateDeckRequest) (*api.SimulateDeckResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	logger := logging.CreateLogger()
	logger.Info(fmt.Sprintf("SimulateDeckRequest: %s", in))

	deckList := facade.ToInternalDeckList(in.DeckList)
	gameConfig := facade.ToInternalGameConfiguration(in.GameConfiguration)
	objective := facade.ToInternalTestObjective(in.Objective)

	checkpoints := simulate(ctx, deckList, gameConfig, objective)
	externalCheckpoints := make([]*api.ResultCheckpoint, len(checkpoints))
	for i, c := range checkpoints {
		externalCheckpoints[i] = facade.ToExternalResultCheckpoint(c)
	}

	response := &api.SimulateDeckResponse{
		Message:     "The server did the thing!",
		SuccessRate: checkpoints[len(checkpoints)-1].GetSuccessRate(),
		Checkpoints: externalCheckpoints,
	}
	logger.Info(fmt.Sprintf("SimulateDeckResponse SuccessRate: %f, Message: %s", response.SuccessRate, response.Message))
	return response, nil
}

func (s *manabaseSimulatorServer) Echo(ctx context.Context, in *api.EchoRequest) (*api.EchoResponse, error) {
	log.Println(fmt.Sprintf("EchoRequest: %s", in.Message))
	return &api.EchoResponse{
		Message: in.Message,
	}, nil
}

func (s *manabaseSimulatorServer) ValidateDeckList(ctx context.Context, in *api.ValidateDeckListRequest) (*api.ValidateDeckListResponse, error) {
	log.Println(fmt.Sprintf("ValidateDeckListRequest: %s", in.DeckList))

	parser := validation.NewDefaultParser(s.cfg)

	_, invalidCards, err := parser.Parse(in.DeckList)
	if err != nil {
		return nil, err
	}

	if len(invalidCards) > 0 {
		externalInvalidCards := make([]*api.InvalidCard, len(invalidCards))
		for i, c := range invalidCards {
			externalInvalidCards[i] = facade.ToExternalInvalidCard(c)
		}
		return &api.ValidateDeckListResponse{
			IsValid:      false,
			InvalidCards: externalInvalidCards,
		}, nil
	}

	return &api.ValidateDeckListResponse{
		IsValid: true,
	}, nil
}

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
