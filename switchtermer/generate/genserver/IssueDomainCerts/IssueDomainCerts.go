package issuedomaincerts

import (
	"fmt"

	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func IssueDomainCerts() {

	d := domain.Domains{}

	dd, err := d.GetDomain()
	switchutility.Checklogger(err, "trying to get domains for running database server")

	var do []string
	for _, v := range dd {
		do = append(do, v.Domain)
	}

	chosendomain := switchselector.MenuInstuctions(do, 1, "purple", "purple", "Which website are you going to run the database server for?")

	if err := d.Create(); err != nil {
		switchutility.Checklogger(err, "creating domain")
	}
	if err := switchutility.ShellBash("cd genserv/assets/certs && go run `go env GOROOT`/src/crypto/tls/generate_cert.go --host="+chosendomain+" --ecdsa-curve=P256", "creating certs"); err != nil {
		switchutility.Checklogger(err, "running certification")
	}

	colortermer.ColorizeOutPut("purple", "purple", "THE FILES CERT AND KEY ARE NOW GENERATED...PLEASE TURN ON TLS NOW")
	fmt.Println("\n")

}
