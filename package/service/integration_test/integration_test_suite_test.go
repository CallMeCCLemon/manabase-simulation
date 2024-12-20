package integration_test

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"manabase-simulation/api"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var Client api.ManabaseSimulatorClient

func TestIntegrationTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegrationTest Suite")
}

var _ = BeforeSuite(func() {
	conn, err := grpc.NewClient("localhost:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	Expect(err).ToNot(HaveOccurred())
	Client = api.NewManabaseSimulatorClient(conn)
})
