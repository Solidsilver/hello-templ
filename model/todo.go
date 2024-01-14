package model

import "time"

type Todo struct {
	Id           string
	Name         string
	Details      string
	Created      string
	LastModified string
}

func NewTodo(name string, details string) Todo {
	return Todo{
		Id:           time.Now().Local().Format("20060102150405"),
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
