package server

import (
	"log"
	"net/http"
	"testing"
)

func TestInitGinServer(t *testing.T) {
	server := InitGinServer()
	if server != nil {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Server closed under request")
			} else {
				log.Fatal("Server closed unexpect")
			}
		}
		log.Println("Server exiting")
	}
}
