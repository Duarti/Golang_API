package main

import (
	"github.com/labstack/echo"
)

type Author struct {
	Id   uint   `json:"id,omitempty"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.GET("/authors", getAuthors)
	e.POST("/authors", addAuthor)
	e.DELETE("authors/:id", deleteAuthor)
	e.PATCH("authors/:id", updateAuthor)
	e.Start(":8000")
}

// HANDLER FUNCTIONS
