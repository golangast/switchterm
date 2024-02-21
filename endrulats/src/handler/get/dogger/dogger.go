package dogger

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/golangast/endrulats/internal/dbsql/dog"
// #import
)

func Dogger(c echo.Context) error {
	// needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
	nonce := c.Get("n")
	jsr := c.Get("jsr")
	cssr := c.Get("cssr")

	dog:=dog.GetDog()
//#Data

	return c.Render(http.StatusOK, "dogger.html", map[string]interface{}{
		"nonce": nonce,
		"jsr":   jsr,
		"cssr":  cssr,
		"dog":dog,
// #tempdata
	})

}
