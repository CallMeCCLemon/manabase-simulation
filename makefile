
all: build

build: test
	if [ ! -d "build" ]; then mkdir build; fi
	echo "Building main.go"
	go build -o build/main ./cmd/main.go

run: fetch-data build
	echo "Now running main"
	./build/main

test: fetch-data
	go test ./...

fetch-data:
	if [ ! -d "data" ]; then mkdir data; fi
	echo "Checking for missing scryfall-db"
	if [ ! -f "data/scryfall-db.json" ]; then curl -o data/scryfall-db.json https://data.scryfall.io/oracle-cards/oracle-cards-20241112220333.json; fi

clean:
	rm -r ./build
