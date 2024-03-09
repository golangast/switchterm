package enabletlsfordomain

import (
	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func EnableTLSForDomain() {

	//getting domain data
	d := domain.Domains{}

	dd, err := d.GetDomain()
	switchutility.Checklogger(err, "trying to get domains for running database server")

	var do []string
	for _, v := range dd {
		do = append(do, v.Domain)
	}

	chosendomain := switchselector.MenuInstuctions(do, 1, "purple", "purple", "Which website are you going to turn on tls for?")

	//add tls code
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", `//#tls`, `//#tls`, tls); err != nil {
		switchutility.Checklogger(err, "trying to update main.go tls part")
	}
	//update imports
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", `// #importtls`, `// #importtls`, `"golang.org/x/crypto/acme/autocert"`+" \n"+`"golang.org/x/crypto/acme"`+" \n"+`"crypto/tls"`+" \n"+`"time"`); err != nil {
		switchutility.Checklogger(err, "trying to update main.go import")
	}
	//update port
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", `e.Logger.Fatal(e.Start(":5002"))`, `e.Logger.Fatal(e.Start(":5002"))`, ""); err != nil {
		switchutility.Checklogger(err, "trying to update main.go removing port")
	}
	//update route
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", `e.Static("/", "assets/optimized")`, `e.Static("/", "assets/optimized")`, ""); err != nil {
		switchutility.Checklogger(err, "trying to update main.go removing assets")
	}
	//pull it all down
	if err := switchutility.ShellBash(" cd "+chosendomain+" && go mod tidy && go mod vendor && go install && go build", "pulling down for tls"); err != nil {
		switchutility.Checklogger(err, "pulling down for tls")
	}

}
func Dev() {
	// getting domain data
	d := domain.Domains{}

	dd, err := d.GetDomain()
	switchutility.Checklogger(err, "trying to get domains for running database server")

	var do []string
	for _, v := range dd {
		do = append(do, v.Domain)
	}

	chosendomain := switchselector.MenuInstuctions(do, 1, "purple", "purple", "Which website are you going to turn on dev for?")

	//updating routes
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", tls, `#tls`, `e.Static("/", "assets/optimized")`+"\n"+`e.Logger.Fatal(e.Start(":5002"))`); err != nil {
		switchutility.Checklogger(err, "trying to update main.go tls part for development")
	}
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", `"golang.org/x/crypto/acme/autocert"`, `"golang.org/x/crypto/acme/autocert"`, ``); err != nil {
		switchutility.Checklogger(err, "trying to update main.go removing import for development")
	}
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", ``, `"golang.org/x/crypto/acme"`, ``); err != nil {
		switchutility.Checklogger(err, "trying to update main.go removing import for development")
	}
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", ``, `"golang.org/x/crypto`, ``); err != nil {
		switchutility.Checklogger(err, "trying to update main.go removing import for development")
	}
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", ``, `"crypto/tls"`, ``); err != nil {
		switchutility.Checklogger(err, "trying to update main.go removing import for development")
	}
	if err := switchutility.UpdateText("./"+chosendomain+"/main.go", ``, `"time"`, ``); err != nil {
		switchutility.Checklogger(err, "trying to update main.go removing import for development")
	}
	if err := switchutility.ShellBash(" cd "+chosendomain+" && go mod tidy && go mod vendor && go install && go build", "pulling down to dev"); err != nil {
		switchutility.Checklogger(err, "trying to pull down for dev")
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
