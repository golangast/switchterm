package post

import (
	"encoding/json"
	"net/http"

	"github.com/golangast/endrulats/internal/dbsql/user"
	"github.com/labstack/echo/v4"
)

func Posts(c echo.Context) error {
	// Bind the request body to a `User` struct.
	user := new(user.Users)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := user.Create()
	if err != nil {
		return err
	}

	// Marshal the `User` struct to JSON.
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Write the JSON data to the response.
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, jsonData)

}
