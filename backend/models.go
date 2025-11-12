package main

import (
	"errors"
	"sync"
	"time"
)

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var statuses = map[string]bool{
	"todo":  true,
	"doing": true,
	"done":  true,
}

func (t *Task) ValidateStatus() error {
	if !statuses[t.Status] {
		return errors.New("status inválido")
	}
	return nil
}

func (t *Task) validateTitle() error {
	if t.Title == "" {
		return errors.New("título é obrigatório")
	}
	return nil
}

func (t *Task) Validate() error {
	if err := t.ValidateStatus(); err != nil {
		return err
	}
	if err := t.validateTitle(); err != nil {
		return err
	}
	return nil
}

type TaskStore struct {
	mu     sync.RWMutex
	tasks  []Task
	lastID int64
}

func (ts *TaskStore) CreateTask(newTask Task) Task {
	ts.mu.Lock()
	ts.lastID++
	now := time.Now()
	newTask.CreatedAt = now
	newTask.UpdatedAt = now
	newTask.ID = ts.lastID
	ts.tasks = append(ts.tasks, newTask)
	ts.mu.Unlock()

	return newTask
}

func (ts *TaskStore) GetTasks() []Task {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.tasks
}

func (ts *TaskStore) FindTaskByID(id int64) int {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for index, task := range ts.tasks {
		if task.ID == id {
			return index
		}
	}
	return -1
}

func (ts *TaskStore) UpdateTask(index int, update Task) {
	if update.Title != "" {
		ts.tasks[index].Title = update.Title
	}
	if update.Status != "" {
		ts.tasks[index].Status = update.Status
	}
	if update.Description != "" {
		ts.tasks[index].Description = update.Description
	}

	ts.tasks[index].UpdatedAt = time.Now()
}

func (ts *TaskStore) DeleteTask(index int) {
	ts.tasks = append(ts.tasks[:index], ts.tasks[index+1:]...)
}
