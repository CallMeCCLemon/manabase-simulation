package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"manabase-simulation/package/service"
	"os"

	"log"
	"manabase-simulation/api"
	"net"
	_ "net/http/pprof"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("server-port", 8889, "The server port")
)

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
	api.RegisterManabaseSimulatorServer(grpcServer, service.NewManabaseSimulatorServer(getDBConfig()))
	grpc_health_v1.RegisterHealthServer(grpcServer, service.NewHealthServer())
	reflection.Register(grpcServer)
	log.Println("Serving gRPC traffic now")
	grpcServer.Serve(lis)
}

func getDBConfig() postgres.Config {
	cfg := postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("host"), os.Getenv("username"), os.Getenv("password"), "app", os.Getenv("port")),
	}
	return cfg
}
