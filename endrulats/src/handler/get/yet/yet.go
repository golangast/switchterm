package yet

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/golangast/endrulats/internal/dbsql/yet"
	// #import
)

func Yet(c echo.Context) error {
	// needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
	nonce := c.Get("n")
	jsr := c.Get("jsr")
	cssr := c.Get("cssr")

	yet := yet.GetYet()
	//#Data

	return c.Render(http.StatusOK, "yet.html", map[string]interface{}{
		"nonce": nonce,
		"jsr":   jsr,
		"cssr":  cssr,
		"yet":yet,
// #tempdata
	})

}
