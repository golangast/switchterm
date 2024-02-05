package generate

import (
	"github.com/golangast/switchterm/switchtermer/generate/genserver"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switch/switchutility"
)

func Generate() {
	listbash := []string{"server", "certificates", "tls", "dev"}

	//print directions
	switchutility.Directions()

	answerbash := switchselector.DigSingle(listbash, 1, "green", "red")

	switch answerbash {

	case "server":
		genserver.Genserver()
	case "certificates":
		genserver.Certificates()
	case "tls":
		genserver.TLS()
	case "dev":
		genserver.Dev()
	}
}
