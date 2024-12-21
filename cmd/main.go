package main

import (
	"manabase-simulation/package/service"
	_ "net/http/pprof"
)

func main() {
	service.Start()
}
