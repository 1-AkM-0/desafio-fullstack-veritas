package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	mu     sync.RWMutex
	tasks  = []Task{}
	lastID int64
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()
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
		mu.Lock()
		lastID++
		now := time.Now()
		newTask.CreatedAt = now
		newTask.UpdatedAt = now
		newTask.ID = lastID
		tasks = append(tasks, newTask)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newTask); err != nil {
			http.Error(w, "Erro no enconder", http.StatusInternalServerError)
		}
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)

	index := findTaskById(id)
	if err != nil {
		http.Error(w, "ID invalido na URL", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()

	if index == -1 {
		http.Error(w, "Id não encontrado", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var update Task
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			http.Error(w, "JSON invalido", http.StatusBadRequest)
			return
		}
		if err := update.ValidateStatus(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if update.Title != "" {
			tasks[index].Title = update.Title
		}
		if update.Status != "" {
			tasks[index].Status = update.Status
		}
		if update.Description != "" {
			tasks[index].Description = update.Description
		}

		tasks[index].UpdatedAt = time.Now()

		w.WriteHeader(http.StatusNoContent)

	case http.MethodDelete:
		tasks = append(tasks[:index], tasks[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
	}
}

func findTaskById(id int64) int {
	for index, task := range tasks {
		if task.ID == id {
			return index
		}
	}
	return -1
}
