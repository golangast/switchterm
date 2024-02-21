package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/golangast/endrulats/internal/dbsql/dbconn"
	"github.com/golangast/endrulats/internal/security/cookies"
	"github.com/golangast/endrulats/internal/security/crypt"
	"github.com/golangast/endrulats/internal/security/jwt"
	"github.com/labstack/echo/v4"
)

func (u *Users) Exists() error {
	var exists bool
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=?)", u.Email)
	err = stmts.Scan(&exists)
	if err != nil {
		return err
	}
	db.Close()

	return nil

}
func Exists(email, password, sitetoken string) (bool, error) {
	var passwordhash string
	db, err := dbconn.DbConnection()
	if err != nil {
		return false, err
	}

	stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=? AND sitetoken=?)", email, sitetoken)
	err = stmts.Scan(&passwordhash)
	if err != nil {
		return false, err
	}

	err = crypt.CheckPassword([]byte(passwordhash), []byte(password))
	if err != nil {
		return false, err
	}

	db.Close()

	return true, nil

}
func (u *Users) CheckLogin(c echo.Context, email, sitetokens, passwordraw string) (error, string) {
	db, err := dbconn.DbConnection()
	if err != nil {
		return err, "wrong input"
	}
	var exists string
	ctx, cancel := context.WithTimeout(context.Background(), 56666*time.Millisecond)
	defer cancel()
	stmts := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email=? AND sitetoken=? AND passwordraw=?)", email, sitetokens, passwordraw)
	err = stmts.Scan(&exists)
	if err != nil {
		return err, "wrong input"
	}

	if exists == "0" {
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

		// cookie, err := cookies.ReadCookie(c, u.SessionName)
		// if err != nil {
		// 	return err, "wrong input"
		// }

		err = cookies.WriteCookie(c, u.SessionName, u.SessionKey)
		if err != nil {
			return err, "cookie didnt write"
		}

		fmt.Println(u)
		// if cookie.Name != u.SessionName && cookie.Value != u.SessionKey {
		// 	return err, "wrong input"
		// }

		//header
		hkey := c.Response().Header().Get("headerkey")

		crypt.CheckPassword([]byte(hkey), []byte("goservershell"))

		db.Close()

		return nil, ""
	}

}
func (u *Users) CheckUser(c echo.Context, email, sitetokens string) (error, string) {
	db, err := dbconn.DbConnection()
	if err != nil {
		return err, "wrong input"
	}
	var exists string
	ctx, cancel := context.WithTimeout(context.Background(), 56666*time.Millisecond)
	defer cancel()
	stmts := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email=? AND sitetoken=?)", email, sitetokens)
	err = stmts.Scan(&exists)
	if err != nil {
		return err, "wrong input"
	}

	if exists == "0" {
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

		db.Close()

		return nil, ""
	}

}
func (u *Users) Create() error {

	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	// Create a statement to insert data into the `users` table.
	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `users` (`email`, `passwordhash`, `passwordraw`, `isdisabled`, `sessionkey`, `sessionname`,  `sessiontoken`, `sitetoken`) VALUES (?, ?,?, ?,?,?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Insert data into the `users` table.
	_, err = stmt.ExecContext(context.Background(), u.Email, u.PasswordHash, u.PasswordRaw, "true", u.SessionKey, u.SessionName, u.SessionToken, u.SiteToken)
	if err != nil {
		panic(err)
	}

	db.Close()
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
	db, err := dbconn.DbConnection()
	if err != nil {
		return *user, err
	}
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

	//get from database
	stmt, err := db.Prepare("SELECT * FROM users WHERE id = ? AND SiteToken = ?")
	if err != nil {
		return u, err
	}
	err = stmt.QueryRow(id, idkey).Scan(email, passwordhash, isdisabled, sessionkey, sessionname, sessiontoken, sitetoken)
	if err != nil {
		return u, err
	}
	u = Users{ID: id, Email: email, PasswordHash: passwordhash, Isdisabled: isdisabled, SessionKey: sessionkey, SessionName: sessionname, SessionToken: sessiontoken, SiteToken: sitetoken}
	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		// close db when not in use
		return u, nil

	case nil:
		fmt.Println("nil!!!!!!!!!!!!")

		// close db when not in use
		return u, nil

	default:

		fmt.Println("default!!!!!!!!!!!!")

		return u, nil
	}

}

// https://golangbot.com/mysql-select-single-multiple-rows/
func (user Users) GetUserByEmail(email, idkey string) (Users, error) {
	// ctx, cancelfunc := context.WithTimeout(context.Background(), 500*time.Second)
	// defer cancelfunc()
	var (
		id           string
		passwordhash string
		passwordraw  string
		isdisabled   string
		sessionkey   string
		sessionname  string
		sessiontoken string
		sitetoken    string
		u            Users
	)
	db, err := dbconn.DbConnection()
	if err != nil {
		return u, err
	}

	//get from database
	stmt, err := db.Prepare("SELECT * FROM users WHERE email = ? AND sitetoken = ?")
	if err != nil {
		return u, err
	}
	err = stmt.QueryRow(email, idkey).Scan(&id, &email, &passwordhash, &passwordraw, &isdisabled, &sessionkey, &sessionname, &sessiontoken, &sitetoken)
	if err != nil {
		return u, err
	}
	u = Users{ID: id, Email: email, PasswordHash: passwordhash, Isdisabled: isdisabled, SessionKey: sessionkey, SessionName: sessionname, SessionToken: sessiontoken, SiteToken: sitetoken}
	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		// close db when not in use
		return u, nil

	case nil:
		fmt.Println("was nil !!!!!!!!!!!!!", email)

		// close db when not in use
		return u, nil

	default:

		fmt.Println("default!!!!!!!!!!!!")

		return u, nil
	}

}

func (user Users) SetUserSitetoken(sitetoken string) error {
	//opening database
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}

	//prepare statement so that no sql injection
	stmt, err := db.Prepare("update users set sitetoken=?")
	if err != nil {
		return err
	}

	//execute qeury
	_, err = stmt.Exec(sitetoken)
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
