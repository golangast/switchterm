
package newch

import (
"net/http"

"github.com/labstack/echo/v4"
)

func Newch(c echo.Context) error {
//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
nonce := c.Get("n")
jsr := c.Get("jsr")
cssr := c.Get("cssr")

return c.Render(http.StatusOK, "newch.html", map[string]interface{}{
	"nonce": nonce,
	"jsr":   jsr,
	"cssr":  cssr,
})

}