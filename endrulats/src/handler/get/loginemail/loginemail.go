package loginemail

import (
	"net/http"

	"github.com/golangast/endrulats/internal/dbsql/user"
	"github.com/labstack/echo/v4"
)

func LoginEmail(c echo.Context) error {

	users := new(user.Users)

	email := c.Param("email")
	SiteToken := c.Param("sitetoken")

	userss, err := users.GetUserByEmail(email, SiteToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "loginemail.html", map[string]interface{}{
		"users": userss,
	})

}
