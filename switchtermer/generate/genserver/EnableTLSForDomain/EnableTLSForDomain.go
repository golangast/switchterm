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
	if err := switchutility.UpdateCode("./"+chosendomain+"/main.go", `e.Logger.Fatal(e.Start(":5002"))`, `//#tls`, `e.Logger.Fatal(e.StartAutoTLS(":443"))`); err != nil {
		switchutility.Checklogger(err, "trying to update main.go tls part for development")
	}

	if err := switchutility.UpdateCode("./"+chosendomain+"/main.go", `var domain = "localhost"`, ``, `var domain = "`+chosendomain+`"`); err != nil {
		switchutility.Checklogger(err, "trying to update main.go csrt part for development")
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
	if err := switchutility.UpdateCode("./"+chosendomain+"/main.go", `e.Logger.Fatal(e.StartAutoTLS(":443"))`, `//#tls`, `e.Logger.Fatal(e.Start(":5002"))`); err != nil {
		switchutility.Checklogger(err, "trying to update main.go tls part for development")
	}

	if err := switchutility.UpdateCode("./"+chosendomain+"/main.go", `var domain = "`+chosendomain+`"`, ``, `var domain = "localhost"`); err != nil {
		switchutility.Checklogger(err, "trying to update main.go csrt part for development")
	}

}
