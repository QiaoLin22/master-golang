package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/user", handleEchoGetUser)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if apiErr, ok := err.(APIError); ok {
			c.JSON(apiErr.Status, map[string]any{"error": apiErr.Msg})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]any{"error": "internal server error"})
	}
}

type APIError struct {
	Status int
	Msg    string
}

func (e APIError) Error() string {
	return e.Msg
}

func NotFoundError() APIError {
	return APIError{
		Status: http.StatusNotFound,
		Msg:    "Not Found",
	}
}

func handleEchoGetUser(c echo.Context) error {
	user, err := getUser()
	if err != nil {
		return NotFoundError()
	}
	return c.JSON(http.StatusOK, user)
}

type User struct{}

func getUser() (*User, error) {
	return nil, nil
}
