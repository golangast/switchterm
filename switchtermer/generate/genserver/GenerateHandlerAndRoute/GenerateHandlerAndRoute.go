package generatehandlerandroute

import (
	"fmt"

	"github.com/golangast/gentil/utility/ff"
	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/db/handler"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateHandlerAndRoute() {

	//get folder and github
	d := domain.Domains{}
	dd, err := d.GetAllDomain()
	switchutility.Checklogger(err, "getting domains for GenerateHandlerAndRoute")

	//get the name of the folder
	var ds []string
	for _, v := range dd {
		ds = append(ds, v.Domain)

	}
	//show them to allow selection
	domainanswer := switchselector.MenuInstuctions(ds, 1, "purple", "purple", "which domain do you prefer to use?")

	//get the github by using the domain name
	var dg string
	for _, v := range dd {
		if v.Domain == domainanswer {
			dg = v.Github
		}
	}
	//ask for handler and route
	striphandler := switchutility.InputScanDirections("What is the name of your handler?")
	striproutes := switchutility.InputScanDirections("What is your single segment route? important! do not start with a / (example do not /dog just use dog)")

	h := handler.Handler{Domain: domainanswer, Handle: striphandler, Segment: striproutes}

	if err := h.Create(); err != nil {
		switchutility.Checklogger(err, "trying to create handler in Database")
	}

	//create handler
	handlerfile, err := ff.Filefolder(domainanswer+"/src/handler/get/"+striphandler, striphandler+".go")
	switchutility.Checklogger(err, "trying to create handler file")

	//create template
	tempfile, err := ff.Filefolder(domainanswer+"/assets/templates/"+striphandler, striphandler+".html")
	switchutility.Checklogger(err, "trying to create tempfile file")

	//updating routes
	if err := switchutility.UpdateText(domainanswer+"/src/routes/router.go", `e.GET("/`+striproutes+`", `+striproutes+`.`+cases.Title(language.Und, cases.NoLower).String(striproutes)+`)`, `//getroute`, `e.GET("/`+striproutes+`", `+striproutes+`.`+cases.Title(language.Und, cases.NoLower).String(striproutes)+`)`+"\n"+`//getroute`); err != nil {
		switchutility.Checklogger(err, "trying to update router.go")
	}

	//update import
	if err := switchutility.UpdateText(domainanswer+"/src/routes/router.go", dg+domainanswer+"/src/handler/get/"+striproutes, "// importroute", `"`+dg+domainanswer+`/src/handler/get/`+striproutes+`"`+"\n"+`// importroute`); err != nil {
		switchutility.Checklogger(err, "trying to update router.go for the import")
	}

	handlermap := make(map[string]string)
	handlermap["route"] = striproutes
	mb := make(map[string]string)
	headerb := fmt.Sprintf(`{{template "header" .}}%s`, "")
	footerb := fmt.Sprintf(`{{template "footer" .}}%s`, "")
	mb["route"] = striproutes
	mb["footer"] = footerb
	mb["header"] = headerb

	//write it to the html file
	if err := switchutility.Writetemplate(Bodytemp, tempfile, mb); err != nil {
		switchutility.Checklogger(err, "trying to update router.html")
	}
	if err := switchutility.Writetemplate(handlertemp, handlerfile, handlermap); err != nil {
		switchutility.Checklogger(err, "trying to update the handler")
	}

	switchutility.ShellBash(`cd `+domainanswer+` && go mod tidy && go mod vendor`, "trying to pull down dependencies")

	defer tempfile.Close()
	defer handlerfile.Close()
}

var Bodytemp = `
{{ .header   }}
You created {{.route}}
<!-- write your code here -->
<!-- #data -->
{{ .footer }}
`
var handlertemp = `
package {{.route}}

import (
"net/http"

"github.com/labstack/echo/v4"

// #import
)

func {{title .route}}(c echo.Context) error {
//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
nonce := c.Get("n")
jsr := c.Get("jsr")
cssr := c.Get("cssr")

//#Data



return c.Render(http.StatusOK, "{{.route}}.html", map[string]interface{}{
	"nonce": nonce,
	"jsr":   jsr,
	"cssr":  cssr,
	// #tempdata
})

}`
