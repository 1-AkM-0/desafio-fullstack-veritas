package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func resetTasks() {
	tasks = []Task{}
}

func TestGetTask(t *testing.T) {
	t.Run("Get sem tasks", func(t *testing.T) {
		resetTasks()

		req, _ := http.NewRequest("GET", "/tasks", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tasksHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("status code esperado %v recebido %v", status, http.StatusOK)
		}
		if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
			t.Errorf("esperado Content-Type application/json recebido %v", contentType)
		}

		var got []Task
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if len(got) != 0 {
			t.Errorf("esperado lista vazia, recebido %v", got)
		}
	})
	t.Run("Get com multiplas tasks", func(t *testing.T) {
		resetTasks()

		want := []Task{{ID: 1, Title: "Passar no desafio", Description: "passar no desafio da veritas", Status: "Em Progresso"}, {ID: 2, Title: "Nao sei", Description: "To sem criatividade", Status: "Concluído"}}
		tasks = want
		req, _ := http.NewRequest("GET", "/tasks", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tasksHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler retornou codigo errado: got %v want %v", status, http.StatusOK)
		}
		if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
			t.Errorf("esperado Content-Type application/json recebido %v", contentType)
		}

		var got []Task
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, want) {
			t.Errorf("esperado %v recebido %v", want, got)
		}
	})
}

func TestPOSTTasks(t *testing.T) {
	t.Run("POST task válida", func(t *testing.T) {
		resetTasks()

		taskPayload := []byte(`{"title":"Minha task de teste", "description":"Teste", "status":"A Fazer"}`)

		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskPayload))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(tasksHandler)

		handler.ServeHTTP(rr, req)
		task := tasks[0]
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler retornou %v esperado %v ", status, http.StatusCreated)
		}
		if err := task.Validate(); err != nil {
			t.Errorf("esperado erro de validação, mas nao retornou nada")
		}

		if tasks[0].Title != "Minha task de teste" {
			t.Errorf("titulo recebido %s esperado %s", tasks[0].Title, "Minha task de teste")
		}
	})
	t.Run("POST task invalida", func(t *testing.T) {
		resetTasks()
		taskPayload := []byte(`{"title":"", "description":"", "status":""}`)

		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskPayload))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(tasksHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler retornou %v esperado %v ", status, http.StatusBadRequest)
		}
		if len(tasks) != 0 {
			t.Errorf("recebido %d tasks esperado 0", len(tasks))
		}
	})
}

func TestUPDATETaks(t *testing.T) {
	t.Run("PUT com id válido", func(t *testing.T) {
		resetTasks()

		tasks = append(tasks, Task{ID: 1, Title: "Task que vai ser atualizada", Description: "", Status: "A Fazer"})

		body := []byte(`{"status":"Em Progresso"}`)

		req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(taskHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler retornou %v esperado %v", status, http.StatusNoContent)
		}
	})
	t.Run("PUT com id inválido", func(t *testing.T) {
		resetTasks()

		tasks = append(tasks, Task{ID: 1, Title: "Task que vai ser atualizada", Description: "", Status: "A Fazer"})

		body := []byte(`{"status":"Em Progresso"}`)

		req, _ := http.NewRequest("PUT", "/tasks/2", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(taskHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler retornou %v esperado %v", status, http.StatusNotFound)
		}
	})
}

func TestDELETETaks(t *testing.T) {
	t.Run("DELETE task id válido", func(t *testing.T) {
		resetTasks()
		tasks = append(tasks, Task{ID: 1, Title: "Task que vai ser deletada", Description: "", Status: "A Fazer"})

		req, _ := http.NewRequest("DELETE", "/tasks/1", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(taskHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler retornou %v esperado %v", status, http.StatusNoContent)
		}
		if len(tasks) != 0 {
			t.Errorf("recebido %d tasks esperado 0", len(tasks))
		}
	})
	t.Run("DELETE task id inválido", func(t *testing.T) {
		resetTasks()
		tasks = append(tasks, Task{ID: 1, Title: "Task que vai ser deletada", Description: "", Status: "A Fazer"})

		req, _ := http.NewRequest("DELETE", "/tasks/2", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(taskHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler retornou %v esperado %v", status, http.StatusNotFound)
		}
	})
}
