package service

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
	"log"
	"manabase-simulation/api"
	"manabase-simulation/package/facade"
	"manabase-simulation/package/logging"
	"manabase-simulation/package/simulation"
	"manabase-simulation/package/validation"
	"net"
	"os"
	"sync"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("server-port", 8889, "The server port")
)

func Start() {
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
	api.RegisterManabaseSimulatorServer(grpcServer, NewManabaseSimulatorServer(getDBConfig()))
	grpc_health_v1.RegisterHealthServer(grpcServer, NewHealthServer())
	reflection.Register(grpcServer)
	log.Println("Serving gRPC traffic now")
	grpcServer.Serve(lis)
}

type ManabaseSimulatorServer struct {
	api.UnimplementedManabaseSimulatorServer

	mu sync.Mutex // protects routeNotes

	cfg    postgres.Config
	Parser validation.Parser
}

func NewManabaseSimulatorServer(cfg postgres.Config) *ManabaseSimulatorServer {
	var p validation.Parser
	p = validation.NewDefaultParser(cfg)
	s := &ManabaseSimulatorServer{
		cfg:    cfg,
		Parser: p,
	}
	return s
}

func NewHealthServer() *health.Server {
	s := health.NewServer()
	return s
}

func (s *ManabaseSimulatorServer) SimulateDeck(ctx context.Context, in *api.SimulateDeckRequest) (*api.SimulateDeckResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	logger := logging.CreateLogger()
	logger.Info(fmt.Sprintf("SimulateDeckRequest: %s", in))

	deckList, _, err := s.Parser.Parse(in.DeckList)
	if err != nil {
		return nil, err
	}
	gameConfig := facade.ToInternalGameConfiguration(in.GameConfiguration)
	objective := facade.ToInternalTestObjective(in.Objective)

	checkpoints := simulation.Simulate(ctx, *deckList, gameConfig, objective)
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

func (s *ManabaseSimulatorServer) Echo(ctx context.Context, in *api.EchoRequest) (*api.EchoResponse, error) {
	log.Println(fmt.Sprintf("EchoRequest: %s", in.Message))
	return &api.EchoResponse{
		Message: in.Message,
	}, nil
}

func (s *ManabaseSimulatorServer) ValidateDeckList(ctx context.Context, in *api.ValidateDeckListRequest) (*api.ValidateDeckListResponse, error) {
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

func getDBConfig() postgres.Config {
	cfg := postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("host"), os.Getenv("username"), os.Getenv("password"), "app", os.Getenv("port")),
	}
	return cfg
}
