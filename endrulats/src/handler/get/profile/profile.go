package profile

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Profile(c echo.Context) error {

	nonce := c.Get("n")
	jsr := c.Get("jsr")
	cssr := c.Get("cssr")

	return c.Render(http.StatusOK, "profile.html", map[string]interface{}{
		"nonce": nonce,
		"jsr":   jsr,
		"cssr":  cssr,
	})
}
