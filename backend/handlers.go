package main

import (
	"encoding/json"
	"net/http"
	"time"
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
	case http.MethodPost:
		var newTask Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, "JSON invalido", http.StatusBadRequest)
			return
		}

		if err := newTask.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		now := time.Now()
		newTask.CreatedAt = now
		newTask.UpdatedAt = now
		newTask.ID = int64(len(tasks) + 1)
		tasks = append(tasks, newTask)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newTask); err != nil {
			http.Error(w, "Erro no enconder", http.StatusInternalServerError)
		}
	}
}
