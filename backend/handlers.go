package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		var update Task
		idStr := r.URL.Path[len("/tasks/"):]
		id, err := strconv.ParseInt(idStr, 10, 64)
		index := findTaskById(id)
		if index == -1 {
			http.Error(w, "Id não encontrado", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "Erro convertendo string para int", http.StatusInternalServerError)
			return
		}

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

		w.Header().Set("Content-Type", "application/json")
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
