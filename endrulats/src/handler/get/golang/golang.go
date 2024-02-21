
package golang

import (
"net/http"

"github.com/labstack/echo/v4"

"github.com/golangast/endrulats/internal/dbsql/golang"
// #import
)

func Golang(c echo.Context) error {
//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
nonce := c.Get("n")
jsr := c.Get("jsr")
cssr := c.Get("cssr")

golang:=golang.GetGolang()
//#Data



return c.Render(http.StatusOK, "golang.html", map[string]interface{}{
	"nonce": nonce,
	"jsr":   jsr,
	"cssr":  cssr,
	"golang":golang,
// #tempdata
})

}