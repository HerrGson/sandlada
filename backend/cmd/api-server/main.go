package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := getConfigFromEnv()
	server := newServer(config)

	log.Printf("Starting server on port %s\n", config.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// newServer sets up a server on given port with a router.
func newServer(config Config) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: newRouter(config),
	}
}
