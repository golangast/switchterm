package generate

import (
	"github.com/golangast/switchterm/switchtermer/generate/genserver"
	enabletlsfordomain "github.com/golangast/switchterm/switchtermer/generate/genserver/EnableTLSForDomain"
	generatehandlerandroute "github.com/golangast/switchterm/switchtermer/generate/genserver/GenerateHandlerAndRoute"
	integratedatastructurewithhandler "github.com/golangast/switchterm/switchtermer/generate/genserver/IntegrateDataStructureWithHandler"
	issuedomaincerts "github.com/golangast/switchterm/switchtermer/generate/genserver/IssueDomainCerts"
	setupdatastructureanddomaintable "github.com/golangast/switchterm/switchtermer/generate/genserver/SetupDataStructureAndDomainTable"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Generate() {
	//list of selections
	listbash := []string{"server", "certificates", "tls", "dev", "handler", "data", "add data to handler"}

	//print directions
	switchutility.Directions()

	//print selection
	answerbash := switchselector.DigSingle(listbash, 1, "purple", "red")

	//choose selection
	switch answerbash {
	case "server":
		genserver.Genserver()
	case "certificates":
		issuedomaincerts.IssueDomainCerts()
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
	}
}
