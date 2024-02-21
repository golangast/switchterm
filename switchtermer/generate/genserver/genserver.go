package genserver

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/golangast/gentil/utility/ff"
	"github.com/golangast/gentil/utility/term"
	"github.com/golangast/gentil/utility/text"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/db/data"
	"github.com/golangast/switchterm/switchtermer/db/dbconn"
	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/db/handler"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Genserver() {
	logger := loggers.CreateLogger()
	fmt.Println("What is the domain name of your site without www or .com?")
	scannerdesc := bufio.NewScanner(os.Stdin)
	scannerdesc.Scan()
	dir := scannerdesc.Text()
	stripdir := strings.TrimSpace(dir)
	fmt.Println("What is your the domain you are going to use for imports? (example github.com/golangast)")
	scannergit := bufio.NewScanner(os.Stdin)
	scannergit.Scan()
	git := scannergit.Text()
	stripgit := strings.TrimSpace(git)
	stripgit = stripgit + "/"
	d := domain.Domains{Domain: stripdir, Github: stripgit}

	d.Create()
	colortermer.ColorizeOutPut("purple", "purple", "CREATING SERVER! THIS MAY TAKE A FEW SECONDS...")
	fmt.Println("\n")

	out, errout, err := term.Shellout("go install golang.org/x/tools/cmd/gonew@latest && gonew github.com/golangast/genserv " + stripgit + stripdir + " && cd " + stripdir +
		" && go mod init " + stripgit + stripdir + " && go mod tidy && go mod vendor")
	if err != nil {
		logger.Error(
			"pulling down genserv",
			slog.String("error: ", err.Error()),
		)
	}
	if out != "" {
		fmt.Println(out)
	}
	if errout != "" {

		fmt.Println(errout)
	}
	colortermer.ColorizeOutPut("purple", "purple", "1. PLEASE NOTE: THIS SERVER HAS THE TLS COMMENTED OUT IN MAIN.GO")
	fmt.Println("\n")
	colortermer.ColorizeOutPut("purple", "purple", "2. YOU WILL NEED TO CHANGE THE IMPORTS NAME FROM GITHUB.COM/GOLANGAST/GENSERVER TO YOURS")
	fmt.Println("\n")

	colortermer.ColorizeOutPut("black", "purple", "3.THERE ARE ALSO IMAGES IN THE CSS FILES THAT THEIR DOMAIN NEEDS TO BE UPDATED.")
	fmt.Println("\n")

	colortermer.ColorizeOutPut("purple", "red", "4. IN ORDER TO RUN IT WITH TLS PLEASE RUN go run `go env GOROOT`/src/crypto/tls/generate_cert.go --host=localhost --ecdsa-curve=P256 or run 'certificates'")
	fmt.Println("\n")

	colortermer.ColorizeOutPut("bpurple", "purple", "5. REMEMBER TO CHANGE THE localhost TO YOUR DOMAIN BEFORE RUNNING!!")
	fmt.Println("\n")

	colortermer.ColorizeOutPut("magenta", "bpurple", "6. REMEMBER TO COMMENT OUT e.Logger.Fatal(e.Start(':5002')) WHEN YOU WANT TLS")
	fmt.Println("\n")
	colortermer.ColorizeOutPut("purple", "purple", "DONE GENERATING AND PLESAE ENJOY!")
}

func Certificates() {

	logger := loggers.CreateLogger()

	d := domain.Domains{}

	domain, err := d.GetDomain()
	if err != nil {
		logger.Error(
			"getting domain",
			slog.String("error: ", err.Error()),
		)
	}

	d.Create()

	out, errout, err := term.Shellout("cd genserv/assets/certs && go run `go env GOROOT`/src/crypto/tls/generate_cert.go --host=" + domain.Domain + " --ecdsa-curve=P256")
	if err != nil {
		logger.Error(
			"running certification",
			slog.String("error: ", err.Error()),
		)
	}
	if out != "" {
		fmt.Println(out)
	}
	if errout != "" {

		fmt.Println("--- errs ---")
		fmt.Println(errout)
	}

	colortermer.ColorizeOutPut("purple", "purple", "THE FILES CERT AND KEY ARE NOW GENERATED...PLEASE TURN ON TLS NOW")
	fmt.Println("\n")

}
func TLS() {

	logger := loggers.CreateLogger()
	d := domain.Domains{}

	domain, err := d.GetDomain()
	if err != nil {
		logger.Error(
			"getting domain",
			slog.String("error: ", err.Error()),
		)
	}

	found := text.FindTextNReturn("./"+domain.Domain+"/main.go", `//#tls`)
	if found != `//#tls` {
		err := text.UpdateText("./"+domain.Domain+"/main.go", `//#tls`, tls+"\n"+`//#tls`)
		if err != nil {
			logger.Error(
				"trying to update main.go tls para",
				slog.String("error: ", err.Error()),
			)
		}
	}
	foundimport := text.FindTextNReturn("./"+domain.Domain+"/main.go", `// #importtls`)
	if foundimport != `// #importtls` {
		err := text.UpdateText("./"+domain.Domain+"/main.go", `// #importtls`, `"golang.org/x/crypto/acme/autocert"`+" \n"+`"golang.org/x/crypto/acme"`+" \n"+`"crypto/tls"`+" \n"+`"time"`+" \n"+`//#importtls`)
		if err != nil {
			logger.Error(
				"trying to update main.go import",
				slog.String("error: ", err.Error()),
			)
		}
	}
	foundport := text.FindTextNReturn("./"+domain.Domain+"/main.go", `e.Logger.Fatal(e.Start(":5002"))`)
	if foundport != `e.Logger.Fatal(e.Start(":5002"))` {
		err := text.UpdateText("./"+domain.Domain+"/main.go", `e.Logger.Fatal(e.Start(":5002"))`, "")
		if err != nil {
			logger.Error(
				"trying to update main.go removing port",
				slog.String("error: ", err.Error()),
			)
		}
		err = text.UpdateText("./"+domain.Domain+"/main.go", `e.Static("/", "assets/optimized")`, "")
		if err != nil {
			logger.Error(
				"trying to update main.go removing assets",
				slog.String("error: ", err.Error()),
			)
		}
	}

	out, errout, err := term.Shellout(" cd " + domain.Domain + " && go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		logger.Error(
			"trying to tls the app",
			slog.String("error: ", err.Error()),
		)
	}
	if out != "" {
		fmt.Println(out)
	}
	if errout != "" {

		fmt.Println("--- errs ---")
		fmt.Println(errout)
	}
}

func Dev() {
	logger := loggers.CreateLogger()
	d := domain.Domains{}

	domain, err := d.GetDomain()
	if err != nil {
		logger.Error(
			"getting domain",
			slog.String("error: ", err.Error()),
		)
	}

	found := text.FindTextNReturn("./"+domain.Domain+"/main.go", tls)
	if found != `#tls` {
		err := text.UpdateText("./"+domain.Domain+"/main.go", tls, `e.Static("/", "assets/optimized")`+"\n"+`e.Logger.Fatal(e.Start(":5002"))`)
		if err != nil {
			logger.Error(
				"trying to update main.go",
				slog.String("error: ", err.Error()),
			)
		}
	}
	foundport := text.FindTextNReturn("./"+domain.Domain+"/main.go", `"golang.org/x/crypto/acme/autocert"`)
	if foundport != `"golang.org/x/crypto/acme/autocert"` {
		err := text.UpdateText("./"+domain.Domain+"/main.go", `"golang.org/x/crypto/acme/autocert"`, "")
		if err != nil {
			logger.Error(
				"trying to update main.go removing import",
				slog.String("error: ", err.Error()),
			)
		}
		err = text.UpdateText("./"+domain.Domain+"/main.go", `"golang.org/x/crypto/acme"`, "")
		if err != nil {
			logger.Error(
				"trying to update main.go removing import",
				slog.String("error: ", err.Error()),
			)
		}
		err = text.UpdateText("./"+domain.Domain+"/main.go", `"golang.org/x/crypto`, "")
		if err != nil {
			logger.Error(
				"trying to update main.go removing import",
				slog.String("error: ", err.Error()),
			)
		}
		err = text.UpdateText("./"+domain.Domain+"/main.go", `"crypto/tls"`, "")
		if err != nil {
			logger.Error(
				"trying to update main.go removing import",
				slog.String("error: ", err.Error()),
			)
		}
		err = text.UpdateText("./"+domain.Domain+"/main.go", `"time"`, "")
		if err != nil {
			logger.Error(
				"trying to update main.go removing import",
				slog.String("error: ", err.Error()),
			)
		}
	}

	out, errout, err := term.Shellout(" cd " + domain.Domain + " && go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		logger.Error(
			"trying to tls the app",
			slog.String("error: ", err.Error()),
		)
	}
	if out != "" {
		fmt.Println(out)
	}
	if errout != "" {

		fmt.Println("--- errs ---")
		fmt.Println(errout)
	}
}

var tls = `
e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
e.Use(middleware.Recover())
e.Use(middleware.Logger())
e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	HTML5:      true,
	Root:       "assets/optimized/",
	Filesystem: http.FS(AssetsOptimize),
}))
autoTLSManager := autocert.Manager{
	Prompt: autocert.AcceptTOS,
	// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
	Cache:      autocert.DirCache("/var/www/.cache"),
	HostPolicy: autocert.HostWhitelist("genserver.com"),
}

s := http.Server{
	Addr:    ":443",
	Handler: e, // set Echo as handler
	TLSConfig: &tls.Config{
		Certificates:   nil, // <-- s.ListenAndServeTLS will populate this field
		GetCertificate: autoTLSManager.GetCertificate,
		NextProtos:     []string{acme.ALPNProto},
	},
	ReadTimeout: 30 * time.Second, // use custom timeouts
}
keyspem, err := assets.Keypem.ReadFile("certs/key.pem")
if err != nil {
	fmt.Println(err)
}
certspem, err := assets.Certpem.ReadFile("certs/cert.pem")
if err != nil {
	fmt.Println(err)
}
if err := s.ListenAndServeTLS(string(certspem), string(keyspem)); err != http.ErrServerClosed {
	e.Logger.Fatal(err)
}`

func GenHandler() {
	logger := loggers.CreateLogger()

	//get folder and github
	d := domain.Domains{}
	dd, err := d.GetAllDomain()
	if err != nil {
		logger.Error(
			"getting domains",
			slog.String("error: ", err.Error()),
		)
	}
	//get the name of the folder
	var ds []string
	for _, v := range dd {
		ds = append(ds, v.Domain)

	}
	//show them to allow selection
	fmt.Println("which domain do you prefer to use?")
	domainanswer := switchselector.Menu(ds, 1, "purple", "purple")

	//get the github by using the domain name
	var dg string
	for _, v := range dd {
		if v.Domain == domainanswer {
			dg = v.Github
		}
	}
	//ask for handler and route
	fmt.Println("What is the name of your handler?")
	scannerhandler := bufio.NewScanner(os.Stdin)
	scannerhandler.Scan()
	handlers := scannerhandler.Text()
	striphandler := strings.TrimSpace(handlers)

	fmt.Println("What is your single segment route? important! do not start with a / (example do not /dog just use dog)")
	scannerroute := bufio.NewScanner(os.Stdin)
	scannerroute.Scan()
	routes := scannerroute.Text()
	striproutes := strings.TrimSpace(routes)

	h := handler.Handler{Domain: domainanswer, Handle: striphandler, Segment: striproutes}
	err = h.Create()
	if err != nil {
		logger.Error(
			"trying to create handler in Database",
			slog.String("error: ", err.Error()),
		)
	}
	//create handler
	handlerfile, err := ff.Filefolder(domainanswer+"/src/handler/get/"+striphandler, striphandler+".go")
	if err != nil {
		logger.Error(
			"trying to create handler file",
			slog.String("error: ", err.Error()),
		)
	}
	//create template
	tempfile, err := ff.Filefolder(domainanswer+"/assets/templates/"+striphandler, striphandler+".html")
	if err != nil {
		logger.Error(
			"trying to create tempfile file",
			slog.String("error: ", err.Error()),
		)
	}
	//update route in route folder
	found := text.FindTextNReturn(domainanswer+"/src/routes/router.go", `e.GET("/`+striproutes+`", `+striproutes+`.`+cases.Title(language.Und, cases.NoLower).String(striproutes)+`)`)
	if found != `e.GET("/`+striproutes+`", `+striproutes+`.`+cases.Title(language.Und, cases.NoLower).String(striproutes)+`)` {
		err := text.UpdateText(domainanswer+"/src/routes/router.go", "//getroute", `e.GET("/`+striproutes+`", `+striproutes+`.`+cases.Title(language.Und, cases.NoLower).String(striproutes)+`)`+"\n"+`//getroute`)
		if err != nil {
			logger.Error(
				"trying to update router.go",
				slog.String("error: ", err.Error()),
			)
		}
	}
	//update the route.go for the import
	found = text.FindTextNReturn(domainanswer+"/src/routes/router.go", dg+domainanswer+"/src/handler/get/"+striproutes)
	if found != `"`+dg+domainanswer+`/src/handler/get/`+striproutes+`"` {
		err := text.UpdateText(domainanswer+"/src/routes/router.go", "// importroute", `"`+dg+domainanswer+`/src/handler/get/`+striproutes+`"`+"\n"+`// importroute`)
		if err != nil {
			logger.Error(
				"trying to update router.go for the import",
				slog.String("error: ", err.Error()),
			)
		}
	}

	handler := `
package ` + striproutes + `

import (
"net/http"

"github.com/labstack/echo/v4"

// #import
)

func ` + cases.Title(language.Und, cases.NoLower).String(striproutes) + `(c echo.Context) error {
//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
nonce := c.Get("n")
jsr := c.Get("jsr")
cssr := c.Get("cssr")

//#Data



return c.Render(http.StatusOK, "` + striproutes + `.html", map[string]interface{}{
	"nonce": nonce,
	"jsr":   jsr,
	"cssr":  cssr,
	// #tempdata
})

}`

	//footer/header map for {{template "footer" .}} {{end}}
	mb := make(map[string]string)
	headerb := fmt.Sprintf(`{{template "header" .}}%s`, "")
	footerb := fmt.Sprintf(`{{template "footer" .}}%s`, "")
	mb["route"] = striproutes
	mb["footer"] = footerb
	mb["header"] = headerb
	//write it to the html file
	var Bodytemp = `
{{ .header   }}
You created {{.route}}
<!-- write your code here -->
<!-- #data -->
{{ .footer }}
`
	err = Writetemplate(Bodytemp, tempfile, mb)
	if err != nil {
		logger.Error(
			"trying to update router.html",
			slog.String("error: ", err.Error()),
		)
	}
	//write it to the html file
	err = Writetemplate(handler, handlerfile, nil)
	if err != nil {
		logger.Error(
			"trying to update the handler",
			slog.String("error: ", err.Error()),
		)
	}
	//clean up the imports
	out, errout, err := term.Shellout(`cd ` + domainanswer + ` && go mod tidy && go mod vendor`)
	if err != nil {
		logger.Error(
			"trying to pull down dependencies"+errout+out,
			slog.String("error: ", err.Error()),
		)
	}

	handlerfile.Close()
}

func CreateData() {
	logger := loggers.CreateLogger()

	f := data.Fields{}
	fmt.Println("Whats Name of your data structure?")
	scannerdesc := bufio.NewScanner(os.Stdin)
	scannerdesc.Scan()
	datastructor := scannerdesc.Text()
	f.Name = strings.TrimSpace(datastructor)

	fmt.Println("Whats the fields of the data? Please field.type then space then field.type like the example Name.string age.int")
	scannerfield := bufio.NewScanner(os.Stdin)
	scannerfield.Scan()
	datafields := scannerfield.Text()

	ff := GetPropDatatype(datafields)
	f.Field = append(f.Field, ff...)
	err := f.Create()
	if err != nil {
		logger.Error(
			"store the data structure",
			slog.String("error: ", err.Error()),
		)
	}

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"opening up the connection to the database",
			slog.String("error: ", err.Error()),
		)
	}
	_, err = db.Query("CREATE TABLE IF NOT EXISTS `domains` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `domain` VARCHAR(64) NULL, `github` VARCHAR(255) NOT NULL)")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Close()

}
func AddDataToHandler() {
	logger := loggers.CreateLogger()

	f := data.Fields{}
	fff, err := f.GetAll()
	if err != nil {
		logger.Error(
			"getting all fields for handler",
			slog.String("error: ", err.Error()),
		)
	}

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
	if err != nil {
		logger.Error(
			"getting all handlers",
			slog.String("error: ", err.Error()),
		)
	}

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
	if err != nil {
		logger.Error(
			"trying to get git by domain",
			slog.String("error: ", err.Error()),
		)
	}

	fields, types := GetField(fd)

	elementMap := make(map[string]string)
	for i := 0; i < len(fields); i++ {
		elementMap[fields[i]] = types[i]
	}

	mapfield := make(map[string]string)
	for i := 0; i < len(fields); i++ {
		mapfield[cases.Title(language.Und, cases.NoLower).String(fields[i])] = fields[i]
	}
	fds := strings.Join(fields, ",&")

	da := Data{Name: fieldanswer, MapData: elementMap, Fields: fds, MapFields: mapfield, Github: github, Domain: do}

	//https://go.dev/play/p/SBqlAIHlVoF
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
	dbfile, err := ff.Filefolder(do+"/internal/dbsql/"+fieldanswer, fieldanswer+".go")
	if err != nil {
		logger.Error(
			"trying to create handler file",
			slog.String("error: ", err.Error()),
		)
	}

	err = WritetemplateStruct(Datavarstemp, dbfile, da)
	if err != nil {
		logger.Error(
			"trying to create internal/dbsql file",
			slog.String("error: ", err.Error()),
		)
	}

	found := text.FindTextNReturn(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", cases.Title(language.Und, cases.NoLower).String(fieldanswer)+"()")
	if !strings.Contains(found, cases.Title(language.Und, cases.NoLower).String(fieldanswer)+"()") {
		err := text.UpdateText(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", "//#Data", fieldanswer+":="+fieldanswer+".Get"+cases.Title(language.Und, cases.NoLower).String(fieldanswer)+"()"+"\n"+`//#Data`)
		if err != nil {
			logger.Error(
				"trying to update handler for calling function",
				slog.String("error: ", err.Error()),
			)
		}
	}
	datafound := text.FindTextNReturn(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", `"`+fieldanswer+`":`+fieldanswer)
	if !strings.Contains(datafound, `"`+fieldanswer+`":`+fieldanswer) {
		err := text.UpdateText(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", "// #tempdata", `"`+fieldanswer+`":`+fieldanswer+`,`+"\n"+`// #tempdata`)
		if err != nil {
			logger.Error(
				"trying to update handler for template data",
				slog.String("error: ", err.Error()),
			)
		}
	}

	importfound := text.FindTextNReturn(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", fieldanswer)
	if !strings.Contains(importfound, `"`+github+do+"/internal/dbsql/"+fieldanswer+`"`) {
		err := text.UpdateText(do+"/src/handler/get/"+handleranswer+"/"+handleranswer+".go", "// #import", `"`+github+do+"/internal/dbsql/"+fieldanswer+`"`+"\n"+`// #import`)
		if err != nil {
			logger.Error(
				"trying to update handler for import",
				slog.String("error: ", err.Error()),
			)
		}
	}
	templatefound := text.FindTextNReturn(do+"/assets/templates/"+handleranswer+"/"+handleranswer+".html", handleranswer)
	if !strings.Contains(templatefound, `"`+github+do+"/internal/dbsql/"+handleranswer+`"`) {
		err := text.UpdateText(do+"/assets/templates/"+handleranswer+"/"+handleranswer+".html", "<!-- #data -->", `{{.`+handleranswer+`}}`+"\n"+`<!-- #data -->`)
		if err != nil {
			logger.Error(
				"trying to update handler for import",
				slog.String("error: ", err.Error()),
			)
		}
	}

}
func TrimDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}
func GetPropDatatype(prop string) []string {
	var property []string
	var types []string
	var field []string
	var strright string
	s := strings.Split(prop, " ")

	for _, ss := range s {
		sss := strings.Replace(ss, "\"", "", -1)
		property = append(property, TrimDot(sss))
		strright = strings.Replace(TrimDotright(sss), ".", "", 1)
		types = append(types, strright)
	}

	for a, str1_word := range property {
		for b, str2_word := range types {
			if a == b {
				field = append(field, str1_word+" "+str2_word)
			}
		}
	}
	return field
}

func GetProp(prop string) []string {
	var property []string
	var types []string
	var field []string
	var strright string
	s := strings.Split(prop, " ")

	for _, ss := range s {
		sss := strings.Replace(ss, "\"", "", -1)
		property = append(property, TrimDot(sss))
		strright = strings.Replace(TrimDotright(sss), ".", "", 1)
		types = append(types, strright)
	}

	for a, str1_word := range property {
		for b, str2_word := range types {
			if a == b {
				field = append(field, str1_word+" "+str2_word)
			}
		}
	}
	return field
}
func GetField(fields string) ([]string, []string) {
	var field []string
	var property []string
	s := strings.Split(fields, " ")
	for i, ss := range s {
		if i%2 == 0 {
			//get even
			field = append(field, TrimDotright(ss))
		} else {
			property = append(property, TrimDotright(ss))

		}

	}
	return field, property
}

// type Fields struct {
// 	Name  string
// 	Field []string
// }

func Writetemplate(temp string, f *os.File, d map[string]string) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}

type Data struct {
	Name      string
	MapData   map[string]string
	Fields    string
	MapFields map[string]string
	Github    string
	Domain    string
}

func WritetemplateStruct(temp string, f *os.File, d Data) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}
