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
	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
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
	fmt.Println("What is your the domain you are going to use for imports? (example/ or github.com/golangast/)")
	scannergit := bufio.NewScanner(os.Stdin)
	scannergit.Scan()
	git := scannergit.Text()
	stripgit := strings.TrimSpace(git)
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
	domain, err := d.GetDomain()
	if err != nil {
		logger.Error(
			"getting domain",
			slog.String("error: ", err.Error()),
		)
	}
	//ask for handler and route
	fmt.Println("What is the name of your handler?")
	scannerdesc := bufio.NewScanner(os.Stdin)
	scannerdesc.Scan()
	handlers := scannerdesc.Text()
	striphandler := strings.TrimSpace(handlers)
	fmt.Println("What is your route?")
	scannerroute := bufio.NewScanner(os.Stdin)
	scannerroute.Scan()
	routes := scannerroute.Text()
	striproutes := strings.TrimSpace(routes)

	//create handler
	handlerfile, err := ff.Filefolder(domain.Domain+"/src/handler/get/"+striphandler, striphandler+".go")
	if err != nil {
		logger.Error(
			"trying to create handler file",
			slog.String("error: ", err.Error()),
		)
	}
	//create template
	tempfile, err := ff.Filefolder(domain.Domain+"/assets/templates/"+striphandler, striphandler+".html")
	if err != nil {
		logger.Error(
			"trying to create tempfile file",
			slog.String("error: ", err.Error()),
		)
	}
	//update route in route folder
	found := text.FindTextNReturn(domain.Domain+"/src/routes/router.go", `e.GET("/`+striproutes+`", `+striproutes+`.`+cases.Title(language.Und, cases.NoLower).String(striproutes)+`)`)
	if found != `e.GET("/`+striproutes+`", `+striproutes+`.`+cases.Title(language.Und, cases.NoLower).String(striproutes)+`)` {
		err := text.UpdateText(domain.Domain+"/src/routes/router.go", "//getroutes", `e.GET("/`+striproutes+`", `+striproutes+`.`+cases.Title(language.Und, cases.NoLower).String(striproutes)+`)`+"\n"+`//getroutes`)
		if err != nil {
			logger.Error(
				"trying to update router.go",
				slog.String("error: ", err.Error()),
			)
		}
	}
	//update the route.go for the import
	found = text.FindTextNReturn(domain.Domain+"/src/routes/router.go", domain.Github+"/src/handler/get/"+striproutes)
	if found != `"`+domain.Github+`/src/handler/get/`+striproutes+`"` {
		err := text.UpdateText(domain.Domain+"/src/routes/router.go", "//imports", `"`+domain.Github+`/src/handler/get/`+striproutes+`"`+"\n"+`//imports`)
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
)

func ` + cases.Title(language.Und, cases.NoLower).String(striproutes) + `(c echo.Context) error {
//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
nonce := c.Get("n")
jsr := c.Get("jsr")
cssr := c.Get("cssr")

return c.Render(http.StatusOK, "` + striproutes + `.html", map[string]interface{}{
	"nonce": nonce,
	"jsr":   jsr,
	"cssr":  cssr,
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
	out, errout, err := term.Shellout(`go mod tidy && go mod vendor`)
	if err != nil {
		logger.Error(
			"trying to pull down dependencies"+errout+out,
			slog.String("error: ", err.Error()),
		)
	}

	handlerfile.Close()
}

func Writetemplate(temp string, f *os.File, d map[string]string) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}
