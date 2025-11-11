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
