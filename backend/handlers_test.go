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
		resetTasks()

		want := []Task{
			{ID: 1, Title: "Passar no desafio", Description: "passar no desafio da veritas", Status: "doing"},
			{ID: 2, Title: "Nao sei", Description: "To sem criatividade", Status: "done"},
		}
		tasks = want
		req, _ := http.NewRequest("GET", "/tasks", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tasksHandler)

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
		resetTasks()

		taskPayload := []byte(`{"title":"Minha task de teste", "description":"Teste", "status":"todo"}`)

		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskPayload))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(tasksHandler)

		handler.ServeHTTP(rr, req)
		task := tasks[0]

		if err := task.Validate(); err != nil {
			t.Errorf("esperado erro de validação, mas nao retornou nada")
		}

		if tasks[0].Title != "Minha task de teste" {
			t.Errorf("titulo recebido %s esperado %s", tasks[0].Title, "Minha task de teste")
		}

		assertContentType(t, rr.Header().Get("Content-Type"), "application/json")
		assertStatus(t, rr.Code, http.StatusCreated)
	})
	t.Run("POST task invalida", func(t *testing.T) {
		resetTasks()
		taskPayload := []byte(`{"title":"", "description":"", "status":""}`)

		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskPayload))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(tasksHandler)
		handler.ServeHTTP(rr, req)

		if len(tasks) != 0 {
			t.Errorf("recebido %d tasks esperado 0", len(tasks))
		}
		assertStatus(t, rr.Code, http.StatusBadRequest)
	})
}

func TestUPDATETaks(t *testing.T) {
	t.Run("PUT com id válido", func(t *testing.T) {
		resetTasks()

		tasks = append(tasks, Task{ID: 1, Title: "Task que vai ser atualizada", Description: "", Status: "todo"})

		body := []byte(`{"status":"doing"}`)

		req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(taskHandler)
		handler.ServeHTTP(rr, req)

		assertStatus(t, rr.Code, http.StatusNoContent)
	})
	t.Run("PUT com id inválido", func(t *testing.T) {
		resetTasks()

		tasks = append(tasks, Task{ID: 1, Title: "Task que vai ser atualizada", Description: "", Status: "todo"})

		body := []byte(`{"status":"doing"}`)

		req, _ := http.NewRequest("PUT", "/tasks/2", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(taskHandler)
		handler.ServeHTTP(rr, req)

		assertStatus(t, rr.Code, http.StatusNotFound)
	})
}

func TestDELETETaks(t *testing.T) {
	t.Run("DELETE task id válido", func(t *testing.T) {
		resetTasks()
		tasks = append(tasks, Task{ID: 1, Title: "Task que vai ser deletada", Description: "", Status: "todo"})

		req, _ := http.NewRequest("DELETE", "/tasks/1", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(taskHandler)
		handler.ServeHTTP(rr, req)

		assertStatus(t, rr.Code, http.StatusNoContent)

		if len(tasks) != 0 {
			t.Errorf("recebido %d tasks esperado 0", len(tasks))
		}
	})
	t.Run("DELETE task id inválido", func(t *testing.T) {
		resetTasks()
		tasks = append(tasks, Task{ID: 1, Title: "Task que vai ser deletada", Description: "", Status: "todo"})

		req, _ := http.NewRequest("DELETE", "/tasks/2", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(taskHandler)
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
