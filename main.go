package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"

	"github.com/a-h/templ"
	"github.com/solidsilver/hello-templ/model"
	"github.com/solidsilver/hello-templ/templates"
)

var footerLinks = map[string]templ.SafeURL{
	"Home":   "/",
	"About":  "/about",
	"Search": "/search",
}

var todoItems = []model.Todo{
	model.NewTodo("Buy milk", "Go to the store and buy milk"),
	model.NewTodo("Buy eggs", "Go to the store and buy eggs"),
	model.NewTodo("Buy bread", "Go to the store and buy bread"),
}

func View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {

	e := echo.New()
	// set zerolog as logger
	logger := zerolog.New(os.Stdout)
	logger.Level(zerolog.InfoLevel)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	e.Static("/", "public")

	// http.Handle("/", templ.Handler(templates.Main("TestPage")))
	e.GET("/footer", func(c echo.Context) error {
		return View(c, templates.FooterTemplate(templates.FooterProps{
			Links: footerLinks,
		}))
	})

	e.GET("/todos", GetTodos)
	e.POST("/todos", AddTodos)
	e.DELETE("/todos/:id", RemoveTodos)

	// http.Handle("/footer")

	e.Logger.Fatal(e.Start(":3001"))
}

func GetTodos(c echo.Context) error {
	return View(c, templates.Todos(todoItems))
}

func AddTodos(c echo.Context) error {
	name := c.FormValue("name")
	details := c.FormValue("details")
	newTodo := model.NewTodo(
		name,
		details,
	)
	todoItems = append(todoItems, newTodo)
	return View(c, templates.TodoItem(newTodo))
}

func RemoveTodos(c echo.Context) error {
	id := c.Param("id")
	for i, todo := range todoItems {
		if todo.Id == id {
			todoItems = append(todoItems[:i], todoItems[i+1:]...)
			break
		}
	}
	return nil
}
