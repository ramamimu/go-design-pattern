package main

import (
	"fmt"
	"hello_world/api"
	"log"
)

func main() {
	port := ":9090"

	server := api.NewServer(port)
	fmt.Println("listening on port", port)
	log.Fatal(server.Start())
}
