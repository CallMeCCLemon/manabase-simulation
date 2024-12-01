package main

import (
	"context"
	"fmt"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"log"
	"manabase-simulation/api"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux(Cors())

	if err := api.RegisterManabaseSimulatorGraphql(mux); err != nil {
		log.Fatalln(err)
	}
	http.Handle("/graphql", mux)
	port := ":8888"
	log.Println(fmt.Sprintf("Listening on port %s", port))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s", port), nil))
}

// Cors is middelware function to provide CORS headers to response headers
func Cors() runtime.MiddlewareFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, error) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Max-Age", "1728000")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		return ctx, nil
	}
}
