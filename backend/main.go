package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	router := server.RegisterRoutes()
	fmt.Println("Server rodando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
