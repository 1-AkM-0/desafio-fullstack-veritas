package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)
	fmt.Println("Server rodando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
