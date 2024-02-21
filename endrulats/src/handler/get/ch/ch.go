
package ch

import (
"net/http"

"github.com/labstack/echo/v4"

// #import
)

func Ch(c echo.Context) error {
//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
nonce := c.Get("n")
jsr := c.Get("jsr")
cssr := c.Get("cssr")

//#Data

return c.Render(http.StatusOK, "ch.html", map[string]interface{}{
	"nonce": nonce,
	"jsr":   jsr,
	"cssr":  cssr,
	//#tempdata
})

}