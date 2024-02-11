package get

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golangast/endrulats/internal/dbsql/user"
	"github.com/labstack/echo/v4"
)

func GetUserById(c echo.Context) error {
	// Connect to the database.
	users := new(user.Users)
	if err := c.Bind(users); err != nil {
		return err
	}

	if users.ID == "" && users.SiteToken == "" {
		return errors.New("user not found")
	}

	userss, err := users.GetUser(users.ID, users.SiteToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Marshal the `User` struct to JSON.
	jsonData, err := json.Marshal(userss)
	if err != nil {
		return err
	}

	// Write the JSON data to the response.
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, jsonData)
}
