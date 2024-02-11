package generate

import (
	"github.com/golangast/switchterm/switchtermer/generate/genserver"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switch/switchutility"
)

func Generate() {
	//list of selections
	listbash := []string{"server", "certificates", "tls", "dev", "gethandler"}

	//print directions
	switchutility.Directions()

	//print selection
	answerbash := switchselector.DigSingle(listbash, 1, "green", "red")

	//choose selection
	switch answerbash {
	case "server":
		genserver.Genserver()
	case "certificates":
		genserver.Certificates()
	case "tls":
		genserver.TLS()
	case "dev":
		genserver.Dev()
	case "gethandler":
		genserver.GenHandler()
	}
}
