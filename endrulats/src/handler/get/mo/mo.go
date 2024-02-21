
package mo

import (
"net/http"

"github.com/labstack/echo/v4"

"github.com/golangast/endrulats/internal/dbsql/mo"
// #import
)

func Mo(c echo.Context) error {
//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
nonce := c.Get("n")
jsr := c.Get("jsr")
cssr := c.Get("cssr")

mo:=mo.GetMo()
//#Data



return c.Render(http.StatusOK, "mo.html", map[string]interface{}{
	"nonce": nonce,
	"jsr":   jsr,
	"cssr":  cssr,
	"mo":mo,
// #tempdata
})

}