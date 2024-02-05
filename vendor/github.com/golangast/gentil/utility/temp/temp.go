package temp

import (
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// write any template to file
func Writetemplate(temp string, f *os.File, d map[string]string) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}
