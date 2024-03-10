package cookies

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func WriteCookie(c echo.Context, sessionname, sessionkey string) error {
	cookie := new(http.Cookie)
	cookie.Name = sessionname
	cookie.Value = sessionkey
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	// return c.String(http.StatusOK, "write a cookie")

	return nil
}
func ReadCookie(c echo.Context, sessionname string) (*http.Cookie, error) {
	cookie, err := c.Cookie(sessionname)
	if err != nil {
		return cookie, err
	}
	return cookie, nil
}
