package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/solidsilver/hello-templ/templates"
)

var footerLinks = map[string]templ.SafeURL{
	"Home":   "/",
	"About":  "/about",
	"Search": "/search",
}

func main() {
	// component := hello("John")
	// header := templates.HeaderTemplate()

	http.Handle("/", templ.Handler(templates.Main("TestPage")))

	http.Handle("/footer", templ.Handler(templates.FooterTemplate(templates.FooterProps{
		Links: footerLinks,
	})))

	fmt.Println("Listening on :3001")
	http.ListenAndServe(":3001", nil)
}
