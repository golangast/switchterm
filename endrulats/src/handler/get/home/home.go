package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	nonce := c.Get("n")
	jsr := c.Get("jsr")
	cssr := c.Get("cssr")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"nonce": nonce,
		"jsr":   jsr,
		"cssr":  cssr,
	})

}
