package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/solidsilver/hello-templ/templates"
)

func main() {
	// component := hello("John")
	header := templates.HeaderTemplate("bob")

	http.Handle("/", templ.Handler(header))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
