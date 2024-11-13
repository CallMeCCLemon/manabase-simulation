
run:
	go build manabase-simulation
	./manabase-simulation
	rm manabase-simulation

test:
	go test ./...
