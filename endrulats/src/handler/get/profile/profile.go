package profile

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// #import
)

func Profile(c echo.Context) error {

	nonce := c.Get("n")
	jsr := c.Get("jsr")
	cssr := c.Get("cssr")

	//#Data

	return c.Render(http.StatusOK, "profile.html", map[string]interface{}{
		"nonce": nonce,
		"jsr":   jsr,
		"cssr":  cssr,
		//#tempdata
	})
}
