package generate

import (
	"github.com/golangast/switchterm/switchtermer/generate/genserver"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switch/switchutility"
)

func Generate() {
	//list of selections
	listbash := []string{"server", "certificates", "tls", "dev", "create handler", "create data", "add data to handler"}

	//print directions
	switchutility.Directions()

	//print selection
	answerbash := switchselector.DigSingle(listbash, 1, "purple", "red")

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
	case "create handler":
		genserver.GenHandler()
	case "create data":
		genserver.CreateData()
	case "add data to handler":
		genserver.AddDataToHandler()
	}
}
