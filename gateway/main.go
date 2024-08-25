package main

import (
	"fmt"
	"log"
	"net/http"

	common "github.com/darkphotonKN/slowfood-microservice-common"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error when attempting to load env file: %s", err)
	}

	var (
		gateWayPort = common.EnvString("GATEWAY_PORT", "3100")
	)

	// start gateway server
	mux := http.NewServeMux()
	httpHandler := NewHandler()
	httpHandler.registerRoutes(mux)

	fmt.Printf("GATEWAY Server started. Listening on port %s", gateWayPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", gateWayPort), mux); err != nil {
		log.Fatalf("Failed to start server, err: %s", err)
	}
}
