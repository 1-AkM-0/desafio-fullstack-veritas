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

		req, _ := http.NewRequest("GET", "/api/tasks", nil)

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
		req, _ := http.NewRequest("GET", "/api/tasks", nil)

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

		req, _ := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(taskPayload))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(tasksHandler)

		handler.ServeHTTP(rr, req)
		task := tasks[0]
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler retornou %q esperado %q ", status, http.StatusCreated)
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

		req, _ := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(taskPayload))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(tasksHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler retornou %q esperado %q ", status, http.StatusBadRequest)
		}
		if len(tasks) != 0 {
			t.Errorf("recebido %d tasks esperado 0", len(tasks))
		}
	})
}
