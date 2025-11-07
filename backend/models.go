package main

import (
	"errors"
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
	"A Fazer":      true,
	"Em Progresso": true,
	"Concluídas":   true,
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("título é obrigatório")
	}
	if !statuses[t.Status] {
		return errors.New("status inválido")
	}
	return nil
}
