package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func getAuthors(c echo.Context) error {
	return c.JSON(http.StatusOK, Authors)
}

func addAuthor(c echo.Context) error {
	newAuthor := Author{}
	content, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &newAuthor)
	if err != nil {
		panic(err)
	}

	Authors = append(Authors, newAuthor)

	return c.JSON(http.StatusOK, newAuthor)

}

func deleteAuthor(c echo.Context) error {
	deletedAuthorId := c.Param("id")
	deletedAuthor := Author{}
	content, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	deletedAuthorIndex := -1

	for key, value := range Authors {
		if strconv.Itoa(int(value.Id)) == deletedAuthorId {
			deletedAuthorIndex = key
			fmt.Println(deletedAuthorIndex)
		}
	}

	if deletedAuthorIndex == -1 {
		panic("author with this id doesn't exist")
	} else {
		deletedAuthor = Authors[deletedAuthorIndex]
		Authors = append(Authors[:deletedAuthorIndex], Authors[deletedAuthorIndex+1:]...)
	}

	return c.JSON(http.StatusOK, deletedAuthor)

}

func updateAuthor(c echo.Context) error {
	updatedAuthor := Author{}
	updatedAuthorId := c.Param("id")
	updatedAuthorIndex := -1
	content, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &updatedAuthor)
	for key, value := range Authors {
		if strconv.Itoa(int(value.Id)) == updatedAuthorId {
			Authors[key] = updatedAuthor
			updatedAuthorIndex = key
		}
	}

	return c.JSON(http.StatusOK, Authors[updatedAuthorIndex])

}
