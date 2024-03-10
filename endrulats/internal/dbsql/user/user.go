package user

import (
	"database/sql"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/golangast/endrulats/internal/security/cookies"
	"github.com/golangast/endrulats/internal/security/crypt"
	"github.com/golangast/endrulats/internal/security/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rqlite/gorqlite"
)

func (u *Users) Exists() (error, bool) {
	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		return err, false
	}
	qr, err := conn.QueryOneParameterized(
		gorqlite.ParameterizedStatement{
			Query:     "SELECT 1 from users where email = ?",
			Arguments: []interface{}{u.Email},
		},
	)
	if err != nil {
		return err, false
	}
	var exist bool

	for qr.Next() {
		err := qr.Scan(&exist)
		if err != nil {
			fmt.Println(err)
		}
	}
	switch qr.Err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("nil!!!!!!!!!!!!")
	default:
		fmt.Println("default!!!!!!!!!!!!")
	}

	conn.Close()

	return nil, exist

}
func Exists(email, password, sitetoken string) (bool, error) {
	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		return false, err
	}

	qr, err := conn.QueryOneParameterized(
		gorqlite.ParameterizedStatement{
			Query:     "SELECT PasswordHash from users where email = ? AND sitetoken=?",
			Arguments: []interface{}{email, sitetoken},
		},
	)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	var PasswordHash string

	for qr.Next() {
		err := qr.Scan(&PasswordHash)
		if err != nil {
			return false, err
		}
	}
	switch qr.Err {
	case sql.ErrNoRows:
		return false, err
	case nil:
		return false, nil
	default:
		err = crypt.CheckPassword([]byte(PasswordHash), []byte(password))
		if err != nil {
			return false, err
		}
		return true, nil
	}

}
func (u *Users) CheckLogin(c echo.Context, email, sitetokens, passwordraw string) (error, string) {

	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		return err, ""
	}
	qr, err := conn.QueryOneParameterized(
		gorqlite.ParameterizedStatement{
			Query:     "SELECT 1 from users where email = ? AND sitetoken=? AND passwordraw=?",
			Arguments: []interface{}{email, sitetokens, passwordraw},
		},
	)
	if err != nil {
		return err, ""
	}

	var exist bool

	for qr.Next() {
		err := qr.Scan(&exist)
		if err != nil {
			return err, ""

		}
	}
	defer conn.Close()

	switch qr.Err {
	case sql.ErrNoRows:
		return err, ""

	case nil:
		return err, ""

	default:

		if exist == false {
			return err, "wrong input"
		} else {
			u, err := u.GetUserByEmail(email, sitetokens)
			if err != nil {
				return err, "wrong input"
			}

			err = crypt.CheckPassword([]byte(u.PasswordHash), []byte(u.PasswordRaw))
			if err != nil {
				return err, "wrong input"
			}

			err = cookies.WriteCookie(c, u.SessionName, u.SessionKey)
			if err != nil {
				return err, "cookie didnt write"
			}

			//header
			hkey := c.Response().Header().Get("headerkey")

			crypt.CheckPassword([]byte(hkey), []byte("goservershell"))

			return nil, ""
		}
	}

}
func (u *Users) CheckUser(c echo.Context, email, sitetokens string) (error, string) {

	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		return err, ""
	}
	qr, err := conn.QueryOneParameterized(
		gorqlite.ParameterizedStatement{
			Query:     "SELECT 1 from users where email = ? AND sitetoken=?",
			Arguments: []interface{}{email, sitetokens},
		},
	)
	if err != nil {
		return err, ""
	}
	var exist bool

	for qr.Next() {
		err := qr.Scan(&exist)
		if err != nil {
			return err, ""

		}
	}
	defer conn.Close()

	switch qr.Err {
	case sql.ErrNoRows:
		return err, ""

	case nil:
		return err, ""

	default:
		if exist == false {
			return err, "wrong input"
		} else {
			u, err := u.GetUserByEmail(email, sitetokens)
			if err != nil {
				return err, "wrong input"
			}

			err = crypt.CheckPassword([]byte(u.PasswordHash), []byte(u.PasswordRaw))
			if err != nil {
				return err, "wrong input"
			}

			cookie, err := cookies.ReadCookie(c, u.SessionName)
			if err != nil {
				return err, "wrong input"
			}

			fmt.Println(u)
			if cookie.Name != u.SessionName && cookie.Value != u.SessionKey {
				return err, "wrong input"
			}

			//header
			hkey := c.Response().Header().Get("headerkey")

			crypt.CheckPassword([]byte(hkey), []byte("goservershell"))

			return nil, ""
		}
	}

}
func (u *Users) Create() error {

	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		return err
	}
	_, err = conn.WriteParameterized(
		[]gorqlite.ParameterizedStatement{
			{
				Query:     "INSERT INTO `users` (`email`, `passwordhash`, `passwordraw`, `isdisabled`, `sessionkey`, `sessionname`,  `sessiontoken`, `sitetoken`) VALUES (?, ?,?, ?,?,?, ?, ?)",
				Arguments: []interface{}{u.Email, u.PasswordHash, u.PasswordRaw, "true", u.SessionKey, u.SessionName, u.SessionToken, u.SiteToken},
			},
		},
	)
	if err != nil {
		return err
	}

	defer conn.Close()

	return nil
}

func (u *Users) JWT() error {
	t, err := jwt.CreateJWT(u.SessionName, u.SessionKey)
	if err != nil {
		return err
	}
	u.SessionToken = t
	return nil
}
func (u *Users) SessionKeys(c echo.Context) error {
	err := cookies.WriteCookie(c, u.SessionName, u.SessionKey)
	if err != nil {
		return err
	}
	return nil
}

// ValidateValuer implements validator.CustomTypeFunc
func (users *Users) Validate(user *Users) error {

	// use a single instance of Validate, it caches struct info
	//var validate *validator.Validate

	validate := validator.New()

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		fmt.Println("------ List of tag fields with error ---------")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
		}

	}
	return nil
	// save user to database
}

func (user *Users) GetUser(id, idkey string) (Users, error) {
	var (
		email        string
		passwordhash string
		isdisabled   string
		sessionkey   string
		sessionname  string
		sessiontoken string
		sitetoken    string
		u            Users
	)
	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		return u, err
	}
	qrr, err := conn.QueryOneParameterized(
		gorqlite.ParameterizedStatement{
			Query:     "SELECT 1 from users where id = ? AND sitetoken=?",
			Arguments: []interface{}{id, idkey},
		},
	)
	if err != nil {
		return u, err
	}

	var exist bool

	for qrr.Next() {
		err := qrr.Scan(&exist)
		if err != nil {
			return u, err
		}
	}
	defer conn.Close()
	switch qrr.Err {
	case sql.ErrNoRows:
		return u, qrr.Err
	case nil:
		return u, qrr.Err
	default:

		if !exist {
			qr, err := conn.QueryOneParameterized(
				gorqlite.ParameterizedStatement{
					Query:     "SELECT * FROM users WHERE id = ? AND SiteToken = ?",
					Arguments: []interface{}{id, idkey},
				},
			)
			if err != nil {
				return u, err
			}
			for qr.Next() {
				err := qr.Scan(&id, &email, &passwordhash, &isdisabled, &sessionkey, &sessionname, &sessiontoken, &sitetoken)
				if err != nil {
					return u, err
				}
			}
			u = Users{ID: id, Email: email, PasswordHash: passwordhash, Isdisabled: isdisabled, SessionKey: sessionkey, SessionName: sessionname, SessionToken: sessiontoken, SiteToken: sitetoken}
			switch err {
			case sql.ErrNoRows:
				return u, err

			case nil:
				return u, err

			default:
				return u, nil
			}
		}
		return u, nil
	}

}

// https://golangbot.com/mysql-select-single-multiple-rows/
func (user Users) GetUserByEmail(email, idkey string) (Users, error) {

	var (
		id           float64
		passwordhash string
		passwordraw  string
		isdisabled   string
		sessionkey   string
		sessionname  string
		sessiontoken string
		sitetoken    string
		u            Users
	)
	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		return u, err
	}
	qr, err := conn.QueryOneParameterized(
		gorqlite.ParameterizedStatement{
			Query:     "SELECT * FROM users WHERE email = ? AND SiteToken = ?",
			Arguments: []interface{}{email, idkey},
		},
	)
	if err != nil {
		return u, err
	}

	for qr.Next() {
		err := qr.Scan(&id, &email, &passwordhash, &passwordraw, &isdisabled, &sessionkey, &sessionname, &sessiontoken, &sitetoken)
		if err != nil {
			return u, err
		}
	}
	sid := fmt.Sprintf("%f", id)
	u = Users{ID: sid, Email: email, PasswordHash: passwordhash, PasswordRaw: passwordraw, Isdisabled: isdisabled, SessionKey: sessionkey, SessionName: sessionname, SessionToken: sessiontoken, SiteToken: sitetoken}
	switch err {
	case sql.ErrNoRows:
		// close db when not in use
		return u, err

	case nil:
		return u, err

	default:
		return u, nil
	}

}

func (user Users) SetUserSitetoken(sitetoken string) error {
	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		return err
	}
	_, err = conn.WriteParameterized(
		[]gorqlite.ParameterizedStatement{
			{
				Query:     "update users set sitetoken=? ",
				Arguments: []interface{}{sitetoken},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

type Users struct {
	ID           string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Email        string `valid:"type(string),required" param:"email" query:"email" form:"email" json:"email" xml:"email" validate:"required,email" mod:"trim"`
	PasswordHash string `valid:"type(string),required" param:"passwordhash" query:"passwordhash" form:"passwordhash" json:"passwordhash" xml:"passwordhash"`
	PasswordRaw  string `valid:"type(string)" param:"passwordraw" query:"passwordraw" form:"passwordraw" json:"passwordraw" xml:"passwordraw" validate:"required" scrub:"password" mod:"trim"`
	Isdisabled   string `valid:"type(string)" param:"isdisabled" query:"isdisabled" form:"isdisabled" json:"isdisabled" xml:"isdisabled"`
	SessionKey   string `valid:"type(string)" param:"sessionkey" query:"sessionkey" form:"sessionkey" json:"sessionkey" xml:"sessionkey"`
	SessionName  string `valid:"type(string)" param:"sessionname" query:"sessionname" form:"sessionname" json:"sessionname" xml:"sessionname"`
	SessionToken string `valid:"type(string)" param:"sessiontoken" query:"sessiontoken" form:"sessiontoken" json:"sessiontoken" xml:"sessiontoken"`
	SiteToken    string `valid:"type(string),required" param:"sitetoken" query:"sitetoken" form:"sitetoken" json:"sitetoken" xml:"sitetoken"`
}
