package generate

import (
	"github.com/golangast/switchterm/switchtermer/generate/generategrid"
	"github.com/golangast/switchterm/switchtermer/generate/genserver"
	enabletlsfordomain "github.com/golangast/switchterm/switchtermer/generate/genserver/EnableTLSForDomain"
	generatehandlerandroute "github.com/golangast/switchterm/switchtermer/generate/genserver/GenerateHandlerAndRoute"
	integratedatastructurewithhandler "github.com/golangast/switchterm/switchtermer/generate/genserver/IntegrateDataStructureWithHandler"
	setupdatastructureanddomaintable "github.com/golangast/switchterm/switchtermer/generate/genserver/SetupDataStructureAndDomainTable"
	"github.com/golangast/switchterm/switchtermer/generate/genserver/rundbserver"
	"github.com/golangast/switchterm/switchtermer/generate/genserver/ship"
	"github.com/golangast/switchterm/switchtermer/generate/optimizer"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Generate() {
	//list of selections
	listbash := []string{"server", "tls", "dev", "handler", "data", "add data to handler", "run db server locally only", "ship", "ask", "grid", "optimize"}

	//print directions
	switchutility.Directions()

	//print selection
	answerbash := switchselector.DigSingle(listbash, 2, "purple", "red")

	//choose selection
	switch answerbash {
	case "server":
		genserver.Genserver()
	case "tls":
		enabletlsfordomain.EnableTLSForDomain()
	case "dev":
		enabletlsfordomain.Dev()
	case "handler":
		generatehandlerandroute.GenerateHandlerAndRoute()
	case "data":
		setupdatastructureanddomaintable.SetupDataStructureAndDomainTable()
	case "add data to handler":
		integratedatastructurewithhandler.IntegrateDataStructureWithHandler()
	case "run db server locally only":
		rundbserver.Rundbserver()
	case "ship":
		ship.Ship()
	case "grid":
		generategrid.Grid()
	case "optimize":
		optimizer.Optimizes()

	}
}
