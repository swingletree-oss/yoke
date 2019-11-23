package swingletree

import (
	"fmt"
)

type Response struct {
	Data   interface{}
	Errors []Error
}

type ErrorResponse struct {
	Errors []Error
}

type Error struct {
	Id     string
	Status int
	Code   string
	Title  string
	Detail string
	Meta   map[string]string
}

func (e Error) String() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Detail)
}
