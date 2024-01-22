package model

import (
	"fmt"
	"math/rand"
	"time"
)

type Todo struct {
	Id           string
	Name         string
	Details      string
	Created      string
	LastModified string
}

func NewTodo(name string, details string) Todo {
	return Todo{
        Id:           fmt.Sprintf("%d",rand.Int()),
		Name:         name,
		Details:      details,
		Created:      time.Now().Local().Format("2006-01-02 15:04:05"),
		LastModified: time.Now().Local().Format("2006-01-02 15:04:05"),
	}
}

func (t *Todo) SetName(name string) {
	t.Name = name
	t.LastModified = time.Now().Local().Format("2006-01-02 15:04:05")
}

func (t *Todo) SetDetails(details string) {
	t.Details = details
	t.LastModified = time.Now().Local().Format("2006-01-02 15:04:05")
}
