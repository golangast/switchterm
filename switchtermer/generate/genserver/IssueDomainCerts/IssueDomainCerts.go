package issuedomaincerts

import (
	"fmt"

	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func IssueDomainCerts() {

	d := domain.Domains{}

	domain, err := d.GetDomain()
	switchutility.Checklogger(err, "getting all handler")

	if err := d.Create(); err != nil {
		switchutility.Checklogger(err, "creating domain")
	}
	if err := switchutility.ShellBash("cd genserv/assets/certs && go run `go env GOROOT`/src/crypto/tls/generate_cert.go --host="+domain.Domain+" --ecdsa-curve=P256", "creating certs"); err != nil {
		switchutility.Checklogger(err, "running certification")
	}

	colortermer.ColorizeOutPut("purple", "purple", "THE FILES CERT AND KEY ARE NOW GENERATED...PLEASE TURN ON TLS NOW")
	fmt.Println("\n")

}
