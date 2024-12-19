
all: build

build: generate test
	if [ ! -d "build" ]; then mkdir build; fi
	echo "Building main.go"
	go build -o build/main ./cmd/main.go

build-gql: test
	if [ ! -d "build" ]; then mkdir build; fi
	echo "Building main.go"
	go build -o build/gql-main ./gateway/main.go

run: fetch-data build
	echo "Now running main"
	./build/main

run-gql: build-gql
	echo "Running GraphQL Server"
	./build/gql-main

test: fetch-data
	go test ./...

fetch-data:
	if [ ! -d "data" ]; then mkdir data; fi
	echo "Checking for missing scryfall-db"
	if [ ! -f "data/scryfall-db.json" ]; then curl -o data/scryfall-db.json https://data.scryfall.io/oracle-cards/oracle-cards-20241217220246.json; fi

clean:
	rm -r ./build

generate:
	protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --graphql_out=.. "api/manabase-simulation.proto"

deploy: docker-build docker-push

docker-build: docker-build-server docker-build-gateway

docker-push: docker-push-server docker-push-gateway

docker-build-server:
	docker build --platform=linux/amd64 -t 100.69.236.43:32000/manabase-simulation-server:latest -f Dockerfile .

docker-push-server:
	docker push 100.69.236.43:32000/manabase-simulation-server:latest

docker-build-gateway:
	docker build --platform=linux/amd64 -t 100.69.236.43:32000/manabase-simulation-gql-gateway:latest -f gateway/Dockerfile .

docker-push-gateway:
	docker push 100.69.236.43:32000/manabase-simulation-gql-gateway:latest
