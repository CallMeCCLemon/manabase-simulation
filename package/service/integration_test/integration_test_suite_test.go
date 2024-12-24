package integration_test

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"manabase-simulation/api"
	"manabase-simulation/package/service"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var Client api.ManabaseSimulatorClient

const (
	IntegrationTestLabel = "integration"
)

func TestIntegrationTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegrationTest Suite")
}

var _ = BeforeSuite(func() {
	err := godotenv.Load()
	if err != nil {
		Fail(err.Error())
	}

	go func() {
		service.Start()
	}()

	grpcEndpoint := os.Getenv("GRPC_HOST")
	if grpcEndpoint == "" {
		log.Println("No GRPC_HOST environment variable found. Defaulting to localhost:8889")
		grpcEndpoint = "localhost:8889"
	}

	conn, err := grpc.NewClient(grpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	Expect(err).ToNot(HaveOccurred())
	Client = api.NewManabaseSimulatorClient(conn)
})
