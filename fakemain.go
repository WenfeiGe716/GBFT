package main

import (
	"os"

	"GBFT/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID, 1)

	server.Start()
}
