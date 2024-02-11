package createuser

import (
	"net/http"

	"github.com/golangast/endrulats/internal/dbsql/user"
	"github.com/golangast/endrulats/internal/security/cookies"
	"github.com/golangast/endrulats/internal/security/jwt"
	"github.com/golangast/endrulats/internal/security/tokens"
	"github.com/labstack/echo/v4"
)

func Createuser(c echo.Context) error {
	nonce := c.Get("n")
	jsr := c.Get("jsr")
	cssr := c.Get("cssr")

	users := new(user.Users)

	if err := c.Bind(users); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if users.Email == "" && users.PasswordRaw == "" && users.SiteToken == "" {
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{
			"EX":    "",
			"M":     "",
			"U":     users,
			"ST":    users.SiteToken,
			"nonce": nonce,
			"jsr":   jsr,
			"cssr":  cssr,
		})
	}

	if err := users.Validate(users); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err, exist := users.CheckLogin(c, users.Email, users.SiteToken, users.PasswordRaw)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	sessionkey := tokens.Timername()
	sessionname := tokens.Timername()

	users.SessionKey = sessionkey
	users.SessionName = sessionname
	users.PasswordHash = tokens.CreateHash(users.PasswordRaw)

	sessiontoken, err := jwt.CreateJWT(sessionname, sessionkey)
	if err != nil {
		return err
	}

	users.SessionToken = sessiontoken

	err = cookies.WriteCookie(c, sessionname, sessionkey)
	if err != nil {
		return err
	}

	// err = email.EmailVerify(users.Email, users.SiteToken)
	// if err != nil {
	// 	return err
	// }
	if err := users.JWT(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := users.Create(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"EX":    exist,
		"M":     "",
		"U":     users,
		"ST":    users.SiteToken,
		"nonce": nonce,
		"jsr":   jsr,
		"cssr":  cssr,
	})

}
