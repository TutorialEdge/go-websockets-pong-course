package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/TutorialEdge/go-websockets-pong-course/internal/transport/http"
)

func Run() error {

	handler := transportHTTP.New()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go WebSocket Pong Clone")
	if err := Run(); err != nil {
		fmt.Println("Failed to start WebSocket Pong Server")
		fmt.Println(err.Error())
	}
}
