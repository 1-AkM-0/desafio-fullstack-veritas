package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	store *TaskStore
}

func NewServer() *Server {
	return &Server{
		store: &TaskStore{
			tasks:  []Task{},
			lastID: 0,
		},
	}
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", s.handleGetTasks)
	mux.HandleFunc("POST /tasks", s.handleCreateTask)
	mux.HandleFunc("PUT /tasks/{id}", s.handleUpdateTasks)
	mux.HandleFunc("DELETE /tasks/{id}", s.handleDeleteTask)

	return mux
}

func (s *Server) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	if err := newTask.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdTask := s.store.CreateTask(newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdTask); err != nil {
		log.Printf("Erro ao encodar resposta JSON para o cliente, %v", err)
		return
	}
}

func (s *Server) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := s.store.GetTasks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Printf("Erro ao encodar resposta JSON para o cliente, %v", err)
		return
	}
}

func (s *Server) handleUpdateTasks(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	index := s.store.FindTaskByID(id)
	if err != nil {
		http.Error(w, "ID invalido na URL", http.StatusBadRequest)
		return
	}
	if index == -1 {
		http.Error(w, "Id não encontrado", http.StatusNotFound)
		return
	}
	var update Task
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}
	if err := update.ValidateStatus(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.store.UpdateTask(index, update)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	index := s.store.FindTaskByID(id)
	if err != nil {
		http.Error(w, "ID invalido na URL", http.StatusBadRequest)
		return
	}
	if index == -1 {
		http.Error(w, "Id não encontrado", http.StatusNotFound)
		return
	}
	s.store.DeleteTask(index)
	w.WriteHeader(http.StatusNoContent)
}
