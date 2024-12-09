package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"manabase-simulation/package/facade"

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
	log.Println(fmt.Sprintf("SimulateDeckRequest: %s", in))

	deckList := facade.ToInternalDeckList(in.DeckList)
	gameConfig := facade.ToInternalGameConfiguration(in.GameConfiguration)
	objective := facade.ToInternalTestObjective(in.Objective)

	result := simulate(ctx, deckList, gameConfig, objective)
	response := &api.SimulateDeckResponse{
		Message:     "The server did the thing!",
		SuccessRate: result,
	}
	log.Println(fmt.Sprintf("SimulateDeckResponse SuccessRate: %f, Message: %s", response.SuccessRate, response.Message))
	return response, nil
}

func (s *manabaseSimulatorServer) Echo(ctx context.Context, in *api.EchoRequest) (*api.EchoResponse, error) {
	log.Println(fmt.Sprintf("EchoRequest: %s", in.Message))
	return &api.EchoResponse{
		Message: in.Message,
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
	api.RegisterManabaseSimulatorServer(grpcServer, newManabaseSimulatorServer())
	grpc_health_v1.RegisterHealthServer(grpcServer, newHealthServer())
	reflection.Register(grpcServer)
	log.Println("Serving gRPC traffic now")
	grpcServer.Serve(lis)
}

func simulate(ctx context.Context, decklist model.DeckList, configuration model.GameConfiguration, objective model.TestObjective) float32 {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	logger := CreateLogger()
	logger.Info(decklist.ToString())
	logger.Info(objective.ToString())

	now := time.Now()

	successCount := 0
	iterations := 1000

	c := make(chan bool, 100)
	wg := new(sync.WaitGroup)

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go start(decklist, configuration, objective, c, wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for result := range c {
		if result {
			successCount++
		}
	}
	successRate := float32(successCount) / float32(iterations) * 100.0

	// Capture results to be consumes.
	logger.Info(fmt.Sprintf("Success count: %d", successCount))
	logger.Info(fmt.Sprintf("Success Rate: %f", successRate))
	logger.Info(fmt.Sprintf("Time taken: %s", time.Since(now)))

	return successRate
}

func start(deckList model.DeckList, gameConfiguration model.GameConfiguration, objective model.TestObjective, c chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- SimulateDeck(deckList, gameConfiguration, objective)
}

// CreateLogger Creates a new logger with the default configuration.
func CreateLogger() *zap.Logger {
	// Create a custom logger configuration
	config := zap.NewProductionConfig()

	// Set the output to stdout
	config.OutputPaths = []string{"stdout"}

	// Set the error output to stderr
	config.ErrorOutputPaths = []string{"stderr"}

	// Configure the encoder to use a human-readable format (good for testing)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Optional: Colorize log levels

	// Build the logger
	logger, err := config.Build(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		// Redirect the core output to GinkgoWriter
		return zapcore.NewCore(
			zapcore.NewConsoleEncoder(config.EncoderConfig),
			zapcore.AddSync(ginkgo.GinkgoWriter), // Send logs to GinkgoWriter
			zapcore.InfoLevel,
		)
	}))
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	// Use the logger
	return logger
}

// SimulateDeck Simulates a deck against a given objective with the provided configuration.
func SimulateDeck(deckList model.DeckList, gameConfiguration model.GameConfiguration, objective model.TestObjective) bool {
	//logger := CreateLogger()
	//logger.Debug("Starting deck simulation")

	// Generate Randomized Deck
	deck := deckList.GenerateDeck()
	hand := model.NewDeck()
	board := model.NewBoardState()

	// TODO: Add validations like Validate deck is >= 60 cards

	// Draw Initial Hand
	for range gameConfiguration.InitialHandSize {
		hand = deck.DrawCard(hand)
	}

	// For turnNumber to target turn
	for turnNumber := range objective.TargetTurn {
		// If turnNumber = 1 and on the play, skip draw
		if turnNumber == 0 && gameConfiguration.OnThePlay {
			// Skip your draw
			//logger.Debug("Playing first, skipping draw")
		} else {
			hand = deck.DrawCard(hand)
		}

		hand = board.PlayLand(hand, objective, turnNumber+1)
	}

	// Compute if target is met (possibly using backtracking?)
	// Computation can start with the most restrictive lands by sorting based on number of colors it taps for.
	isMet, _ := board.ValidateTestObjective(objective)
	return isMet
}
