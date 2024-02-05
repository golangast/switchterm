package genserver

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/golangast/gentil/utility/term"
	"github.com/golangast/gentil/utility/text"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
)

func Genserver() {
	logger := loggers.CreateLogger()
	colortermer.ColorizeOutPut("purple", "purple", "CREATING SERVER! THIS MAY TAKE A FEW SECONDS...")
	fmt.Println("\n")

	out, errout, err := term.Shellout("go install golang.org/x/tools/cmd/gonew@latest && git clone https://github.com/golangast/genserv")
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
	colortermer.ColorizeOutPut("purple", "purple", "--PLEASE NOTE: THIS SERVER IN FOLDER /genserv HAS THE TLS COMMENTED OUT IN MAIN.GO AND YOU MAY NEED TO CHANGE THE MOD FILES AND IMPORTS")
	fmt.Println("\n")

	colortermer.ColorizeOutPut("black", "purple", "--PLEASE NOTE: THERE ARE ALSO IMAGES IN THE CSS FILES THAT THEIR DOMAIN NEEDS TO BE UPDATED.")
	fmt.Println("\n")

	colortermer.ColorizeOutPut("purple", "red", "--PLEASE NOTE: IN ORDER TO RUN IT WITH TLS PLEASE RUN go run `go env GOROOT`/src/crypto/tls/generate_cert.go --host=localhost --ecdsa-curve=P256 or run 'certificates'")
	fmt.Println("\n")

	colortermer.ColorizeOutPut("bpurple", "purple", "--PLEASE NOTE: REMEMBER TO CHANGE THE localhost TO YOUR DOMAIN BEFORE RUNNING!!")
	fmt.Println("\n")

	colortermer.ColorizeOutPut("magenta", "bpurple", "--PLEASE NOTE: REMEMBER TO COMMENT OUT e.Logger.Fatal(e.Start(':5002')) WHEN YOU WANT TLS")
	fmt.Println("\n")
	colortermer.ColorizeOutPut("purple", "purple", "DONE GENERATING AND PLESAE ENJOY!")
}

func Certificates() {

	logger := loggers.CreateLogger()

	fmt.Println("What is the domain name of your site without www or .com?")
	scannerdesc := bufio.NewScanner(os.Stdin)
	scannerdesc.Scan()
	dir := scannerdesc.Text()
	stripdir := strings.TrimSpace(dir)
	fmt.Println("What is your the domain name you are going to use for githhub?")
	scannergit := bufio.NewScanner(os.Stdin)
	scannergit.Scan()
	git := scannergit.Text()
	stripgit := strings.TrimSpace(git)
	d := domain.Domains{Domain: stripdir, Github: stripgit}

	d.Create()

	out, errout, err := term.Shellout("cd genserv/assets/certs && go run `go env GOROOT`/src/crypto/tls/generate_cert.go --host=" + stripdir + " --ecdsa-curve=P256")
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
	if _, err := os.Stat(stripdir); errors.Is(err, os.ErrNotExist) {
		if err != nil {
			logger.Error(
				"finding folder",
				slog.String("error: ", err.Error()),
			)
		}
		err = os.Rename("./genserv", stripdir)
		if err != nil {
			logger.Error(
				"renaming folder",
				slog.String("error: ", err.Error()),
			)
		}
	}
	colortermer.ColorizeOutPut("purple", "purple", "THE FILES CERT AND KEY ARE NOW GENERATED...PLEASE TURN ON TLS NOW")
	fmt.Println("\n")
	colortermer.ColorizeOutPut("purple", "purple", "AND THE FOLDER HAS BEEN RENAMED.  YOU STILL NEED TO UPDATE THE IMPORTS.")

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

	fmt.Println(domain)

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
	//`"golang.org/x/crypto/acme"`+"\n"+`"golang.org/x/crypto`+"\n"+`"crypto/tls"`+"\n"+`"time"`, ""
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
