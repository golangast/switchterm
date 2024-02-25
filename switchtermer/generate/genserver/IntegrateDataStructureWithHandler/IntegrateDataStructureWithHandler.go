package integratedatastructurewithhandler

import (
	"strings"

	"github.com/golangast/gentil/utility/ff"
	"github.com/golangast/switchterm/switchtermer/db/data"
	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/db/handler"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func IntegrateDataStructureWithHandler() {

	f := data.Fields{}
	fff, err := f.GetAll()
	switchutility.Checklogger(err, "getting all fields for handler")

	var fs []string
	for _, v := range fff {
		fs = append(fs, v.Name)

	}

	fieldanswer := switchselector.MenuInstuctions(fs, 1, "purple", "purple", "Which data structure are you gonna use?")

	var fd string
	for _, v := range fff {
		fd = v.Fields

	}

	h := handler.Handler{}
	hh, err := h.GetAll()
	switchutility.Checklogger(err, "getting all handler")

	var hs []string
	for _, v := range hh {
		hs = append(hs, v.Handle)

	}

	handleranswer := switchselector.MenuInstuctions(hs, 1, "purple", "purple", "Whats Name of the handler you want to use?")

	var do string
	for _, v := range hh {
		if v.Handle == handleranswer {
			do = v.Domain
		}
	}
	d := domain.Domains{Domain: do}

	github, err := d.GetGitByDomain()
	switchutility.Checklogger(err, "trying to get git by domain")

	fields, types := switchutility.GetField(fd)

	elementMap := make(map[string]string)
	for i := 0; i < len(fields); i++ {
		elementMap[fields[i]] = types[i]
	}

	mapfield := make(map[string]string)
	for i := 0; i < len(fields); i++ {
		mapfield[cases.Title(language.Und, cases.NoLower).String(fields[i])] = fields[i]
	}
	fds := strings.Join(fields, ",&")

	da := switchutility.Data{Name: fieldanswer, MapData: elementMap, Fields: fds, MapFields: mapfield, Github: github, Domain: do}

	dbfile, err := ff.Filefolder(do+"/internal/dbsql/"+fieldanswer, fieldanswer+".go")
	switchutility.Checklogger(err, "trying to create handler file")

	if err := switchutility.WritetemplateStruct(Datavarstemp, dbfile, da); err != nil {
		switchutility.Checklogger(err, "trying to create internal/dbsql file")
	}

	if err := switchutility.UpdateText(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", cases.Title(language.Und, cases.NoLower).String(fieldanswer)+"()", "//#Data", fieldanswer+":="+fieldanswer+".Get"+cases.Title(language.Und, cases.NoLower).String(fieldanswer)+"()"); err != nil {
		switchutility.Checklogger(err, "trying to update handler for calling function")
	}

	if err := switchutility.UpdateText(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", `"`+fieldanswer+`":`+fieldanswer, "// #tempdata", `"`+fieldanswer+`":`+fieldanswer+`,`); err != nil {
		switchutility.Checklogger(err, "trying to update handler for template data")
	}

	if err := switchutility.UpdateText(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", fieldanswer, "// #import", `"`+github+do+"/internal/dbsql/"+fieldanswer+`"`); err != nil {
		switchutility.Checklogger(err, "trying to update handler for import")
	}

	if err := switchutility.UpdateText(do+"/assets/templates/"+handleranswer+"/"+handleranswer+".html", handleranswer, "<!-- #data -->", `{{.`+handleranswer+`}}`); err != nil {
		switchutility.Checklogger(err, "trying to update handler for import")
	}
}

// https://go.dev/play/p/SBqlAIHlVoF
var Datavarstemp = `
package {{.Name}}

import (
	"fmt"
	"{{.Github}}{{.Domain}}/internal/dbsql/dbconn"
)


type {{title .Name}} struct {
	ID     int   ` + "`" + `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id" ` + "`" + `
	{{range $k, $v := .MapData }}
	{{title $k}} {{$v}}  ` + "`" + `param:"{{$k}}" query:"{{$k}}" header:"{{$k}}" form:"{{$k}}" json:"{{$k}}" xml:"{{$k}}" ` + "`" + `
	{{end}}
	
	
	}

func Get{{title .Name}}() []{{title .Name}} {

data, err := dbconn.DbConnection()
if err != nil {
	fmt.Println(err)
}

//variables used to store data from the query
var (
	id int
	{{range $k, $v := .MapData }}
	{{lower $k}} {{$v}}
	{{end}}
	{{.Name}}s  []{{title .Name}} 
	)//used to store all users
	//https://go.dev/play/p/82imTtvHWzb
_, err = data.Query("CREATE TABLE IF NOT EXISTS {{.Name}} (id INTEGER PRIMARY KEY AUTOINCREMENT, {{$first := true}} {{range $k, $v := .MapData }}{{if $first}}{{$first = false}}{{else}} , {{end}}{{lower $k}} {{$v | replace "string" "text"}} NULL {{end}})")
if err != nil {
	fmt.Println(err)
}

//get from database
rows, err := data.Query("select * from {{.Name}}")
if err != nil {
	fmt.Println(err)
}



//cycle through the rows to collect all the data
for rows.Next() {
	err := rows.Scan(&id, &{{ lower .Fields}} )
	if err != nil {
		fmt.Println(err)
	}
	
	u := {{title .Name}}{ID: id,
		{{range $k, $v := .MapFields }}
		{{$k}}: {{lower $v}},
		{{end}}
		} 
		{{.Name}}s = append({{.Name}}s, u)
	}


//close everything
rows.Close()
data.Close()
return {{.Name}}s
}
`
