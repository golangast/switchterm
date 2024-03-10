package comment

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/rqlite/gorqlite"
)

func (u *Comment) Create() error {
	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		panic(err)
	}

	// statements := make([]string, 0)
	// pattern := "INSERT INTO comment(email, language, comment, sitetoken) VALUES (%s, '%s', '%s', '%s')"

	// using parameterized statements
	_, err = conn.WriteParameterized(
		[]gorqlite.ParameterizedStatement{
			{
				Query:     "INSERT INTO comment(email, language, comment, sitetoken) VALUES(?, ?, ?,?)",
				Arguments: []interface{}{u.Email, u.Comment, u.Language, u.Sitetoken},
			},
		},
	)
	// statements = append(statements, fmt.Sprintf(pattern, u.Email, u.Comment, u.Language, u.Sitetoken))
	// results, err := conn.Write(statements)
	if err != nil {
		return err
	}

	return nil
}

// ValidateValuer implements validator.CustomTypeFunc
func (comment *Comment) Validate(comments *Comment) error {
	validate := validator.New()
	err := validate.Struct(comment)
	if err != nil {
		return err
	}

	for _, err := range err.(validator.ValidationErrors) {
		if err != nil {
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}

	}
	return nil
	// save user to database
}

func (comment Comment) SetUserSitetoken(sitetoken string) error {

	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	if err != nil {
		panic(err)
	}

	statements := make([]string, 0)
	pattern := "update comment set sitetoken ('%s')"
	statements = append(statements, fmt.Sprintf(pattern, sitetoken))
	results, err := conn.Write(statements)
	if err != nil {
		fmt.Println(err)
	}

	for n, v := range results {
		fmt.Printf("for result %d, %d rows were affected\n", n, v.RowsAffected)
		if v.Err != nil {
			fmt.Printf("   we have this error: %s\n", v.Err.Error())
		}
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
