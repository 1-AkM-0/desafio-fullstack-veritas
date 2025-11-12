package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestGetTask(t *testing.T) {
	t.Run("Get sem tasks", func(t *testing.T) {
		s := NewServer()
		req, _ := http.NewRequest("GET", "/tasks", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.handleGetTasks)

		handler.ServeHTTP(rr, req)

		assertContentType(t, rr.Header().Get("Content-Type"), "application/json")

		var got []Task
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if len(got) != 0 {
			t.Errorf("esperado lista vazia, recebido %v", got)
		}

		assertStatus(t, rr.Code, http.StatusOK)
	})
	t.Run("Get com multiplas tasks", func(t *testing.T) {
		s := NewServer()

		want := []Task{
			{ID: 1, Title: "Passar no desafio", Description: "passar no desafio da veritas", Status: "doing"},
			{ID: 2, Title: "Nao sei", Description: "To sem criatividade", Status: "done"},
		}
		s.store.tasks = want
		req, _ := http.NewRequest("GET", "/tasks", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.handleGetTasks)

		handler.ServeHTTP(rr, req)

		var got []Task
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, want) {
			t.Errorf("esperado %v recebido %v", want, got)
		}

		assertContentType(t, rr.Header().Get("Content-Type"), "application/json")
		assertStatus(t, rr.Code, http.StatusOK)
	})
}

func TestPOSTTasks(t *testing.T) {
	t.Run("POST task válida", func(t *testing.T) {
		s := NewServer()
		taskPayload := []byte(`{"title":"Minha task de teste", "description":"Teste", "status":"todo"}`)

		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskPayload))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(s.handleCreateTask)

		handler.ServeHTTP(rr, req)
		task := s.store.tasks[0]

		if err := task.Validate(); err != nil {
			t.Errorf("esperado erro de validação, mas nao retornou nada")
		}

		if s.store.tasks[0].Title != "Minha task de teste" {
			t.Errorf("titulo recebido %s esperado %s", s.store.tasks[0].Title, "Minha task de teste")
		}

		assertContentType(t, rr.Header().Get("Content-Type"), "application/json")
		assertStatus(t, rr.Code, http.StatusCreated)
	})
	t.Run("POST task invalida", func(t *testing.T) {
		s := NewServer()
		taskPayload := []byte(`{"title":"", "description":"", "status":""}`)

		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskPayload))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(s.handleCreateTask)
		handler.ServeHTTP(rr, req)

		if len(s.store.tasks) != 0 {
			t.Errorf("recebido %d tasks esperado 0", len(s.store.tasks))
		}
		assertStatus(t, rr.Code, http.StatusBadRequest)
	})
}

func TestUPDATETaks(t *testing.T) {
	t.Run("PUT com id válido", func(t *testing.T) {
		s := NewServer()
		s.store.tasks = append(s.store.tasks, Task{ID: 1, Title: "Task que vai ser atualizada", Description: "", Status: "todo"})

		body := []byte(`{"status":"doing"}`)

		req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
		req.SetPathValue("id", "1")
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.handleUpdateTasks)
		handler.ServeHTTP(rr, req)

		assertStatus(t, rr.Code, http.StatusNoContent)
	})
	t.Run("PUT com id inválido", func(t *testing.T) {
		s := NewServer()
		s.store.tasks = append(s.store.tasks, Task{ID: 1, Title: "Task que vai ser atualizada", Description: "", Status: "todo"})

		body := []byte(`{"status":"doing"}`)

		req, _ := http.NewRequest("PUT", "/tasks/2", bytes.NewBuffer(body))
		req.SetPathValue("id", "2")
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.handleUpdateTasks)
		handler.ServeHTTP(rr, req)

		assertStatus(t, rr.Code, http.StatusNotFound)
	})
}

func TestDELETETaks(t *testing.T) {
	t.Run("DELETE task id válido", func(t *testing.T) {
		s := NewServer()
		s.store.tasks = append(s.store.tasks, Task{ID: 1, Title: "Task que vai ser deletada", Description: "", Status: "todo"})

		req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
		req.SetPathValue("id", "1")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.handleDeleteTask)
		handler.ServeHTTP(rr, req)

		assertStatus(t, rr.Code, http.StatusNoContent)

		if len(s.store.tasks) != 0 {
			t.Errorf("recebido %d tasks esperado 0", len(s.store.tasks))
		}
	})
	t.Run("DELETE task id inválido", func(t *testing.T) {
		s := NewServer()
		s.store.tasks = append(s.store.tasks, Task{ID: 1, Title: "Task que vai ser deletada", Description: "", Status: "todo"})

		req, _ := http.NewRequest("DELETE", "/tasks/2", nil)
		req.SetPathValue("id", "2")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.handleDeleteTask)
		handler.ServeHTTP(rr, req)

		assertStatus(t, rr.Code, http.StatusNotFound)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("mensagem de status incorreta, recebida %d, esperada %d", got, want)
	}
}

func assertContentType(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("content-type incorreto, recebido %s, esperado %s", got, want)
	}
}
