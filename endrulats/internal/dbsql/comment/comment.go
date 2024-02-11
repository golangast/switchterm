package comment

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/golangast/endrulats/internal/dbsql/dbconn"
	"github.com/golangast/endrulats/internal/loggers"
)

func (u *Comment) Create() error {
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	// Create a statement to insert data into the `users` table.
	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `comment` (`email`, `language`, `comment`,  `sitetoken`) VALUES (?, ?,?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Insert data into the `users` table.
	_, err = stmt.ExecContext(context.Background(), u.Email, u.Comment, u.Language, u.Sitetoken)
	if err != nil {
		panic(err)
	}

	db.Close()
	return nil
}

// ValidateValuer implements validator.CustomTypeFunc
func (comment *Comment) Validate(comments *Comment) error {
	logger := loggers.CreateLogger()
	validate := validator.New()
	err := validate.Struct(comment)
	if err != nil {
		logger.Error(
			"trying to validate comment",
			slog.String("error: ", err.Error()),
		)
		return err
	}
	for _, err := range err.(validator.ValidationErrors) {
		if err != nil {
			logger.Error(
				"validating comment fields",
				slog.String("error: ", err.StructField()),
				slog.String("error: ", err.ActualTag()),
				slog.String("error: ", err.Kind().String()),
			)
			return nil
		}
		fmt.Println(err.Value())
		fmt.Println(err.Param())
	}
	return nil
	// save user to database
}

func (comment Comment) SetUserSitetoken(sitetoken string) error {
	//opening database
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}

	//prepare statement so that no sql injection
	stmt, err := db.Prepare("update comment set sitetoken=?")
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

type Comment struct {
	ID        string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Email     string `valid:"type(string),required" param:"email" query:"email" form:"email" json:"email" xml:"email" validate:"required,email" mod:"trim"`
	Language  string `valid:"type(string),required" param:"language" query:"language" form:"language" json:"language" xml:"language" validate:"required,language" mod:"trim"`
	Comment   string `valid:"type(string),required" param:"comment" query:"comment" form:"comment" json:"comment" xml:"comment" validate:"required,comment" mod:"trim"`
	Sitetoken string `valid:"type(string),required" param:"sitetoken" query:"sitetoken" form:"sitetoken" json:"sitetoken" xml:"sitetoken" validate:"required,sitetoken" mod:"trim"`
}
