package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/api/idtoken"
	"log"
	"manabase-simulation/api"
	"manabase-simulation/package/logging"
	"net/http"
	"os"
)

func main() {
	logger := logging.CreateLogger()
	mux := runtime.NewServeMux(Cors(), AuthMiddleware(logger))

	if err := api.RegisterManabaseSimulatorGraphql(mux); err != nil {
		log.Fatalln(err)
	}
	http.Handle("/graphql", mux)
	port := ":8888"
	logger.Info(fmt.Sprintf("Listening on port %s", port))
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

// AuthMiddleware checks if the request has a valid token and sets the user in the context
func AuthMiddleware(logger *zap.Logger) runtime.MiddlewareFunc {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	if googleClientID == "" {
		panic("GOOGLE_CLIENT_ID is not set")
	}

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, error) {
		tokenValidator, err := idtoken.NewValidator(ctx)
		if err != nil {
			// handle error, stop execution
		}

		token := r.Header.Get("Authorization")
		if token == "" {
			return ctx, errors.New("no Authorization token found")
		}
		payload, err := tokenValidator.Validate(ctx, token, googleClientID)
		if err != nil {
			return ctx, err
		}

		logger.Info(fmt.Sprintf("User: %v", payload.Claims))
		return ctx, nil
	}
}
