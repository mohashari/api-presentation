package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type user User

func postSomething(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func putSomething(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func patchSomething(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func getSomething(c echo.Context) error {
	return c.String(http.StatusOK, "Hello ")
}

func deleteSomething(c echo.Context) error {
	return c.String(http.StatusOK, "deleted")
}

func main() {

	e := echo.New()
	e.POST("/", postSomething)
	e.PUT("/", putSomething)
	e.PATCH("/", patchSomething)
	e.DELETE("/", deleteSomething)
	e.GET("/", getSomething)
	e.Logger.Info(e.Start(":8090"))
}
