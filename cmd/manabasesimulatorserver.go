package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/health"
	"log"
	"manabase-simulation/api"
	"manabase-simulation/package/facade"
	"manabase-simulation/package/logging"
	"sync"
)

type manabaseSimulatorServer struct {
	api.UnimplementedManabaseSimulatorServer

	mu sync.Mutex // protects routeNotes
}

func newManabaseSimulatorServer() *manabaseSimulatorServer {
	s := &manabaseSimulatorServer{}
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

	return &api.ValidateDeckListResponse{
		IsValid: true,
	}, nil
}
