package main

import (
	"encoding/json"
	"net/http"
)

var tasks = []Task{}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, "Erro ao serializar as tasks", http.StatusInternalServerError)
			return
		}
	}
}
