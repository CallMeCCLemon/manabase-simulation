package main

import (
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"log"
	"manabase-simulation/api"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()

	if err := api.RegisterManabaseSimulatorGraphql(mux); err != nil {
		log.Fatalln(err)
	}
	http.Handle("/graphql", mux)
	log.Fatalln(http.ListenAndServe("mtg-mana-sim-app-server-service:8888", nil))
}
