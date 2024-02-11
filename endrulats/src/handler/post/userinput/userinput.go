package userinput

import (
	"net/http"

	"github.com/golangast/endrulats/internal/dbsql/comment"
	"github.com/golangast/endrulats/internal/dbsql/user"
	"github.com/labstack/echo/v4"
)

func UserInput(c echo.Context) error {
	nonce := c.Get("n")
	jsr := c.Get("jsr")
	cssr := c.Get("cssr")

	comments := new(comment.Comment)
	users := new(user.Users)

	if err := c.Bind(comments); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	u, err := users.GetUserByEmail(comments.Email, comments.Sitetoken)
	if err != nil {
		return err
	}

	if comments.Email == "" && comments.Language == "" && comments.Comment == "" && comments.Sitetoken == "" {
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{
			"U":     u,
			"C":     comments,
			"ST":    comments.Sitetoken,
			"nonce": nonce,
			"jsr":   jsr,
			"cssr":  cssr,
		})
	}

	// if err := comments.Validate(comments); err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	if err := comments.Create(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"U":     u,
		"ST":    u.SiteToken,
		"nonce": nonce,
		"jsr":   jsr,
		"cssr":  cssr,
	})

}
