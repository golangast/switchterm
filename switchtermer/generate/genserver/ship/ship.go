package ship

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Ship() {

	d := domain.Domains{}

	dd, err := d.GetDomain()
	switchutility.Checklogger(err, "trying to get domains for running database server")

	var do []string
	for _, v := range dd {
		do = append(do, v.Domain)
	}

	chosendomain := switchselector.MenuInstuctions(do, 1, "purple", "purple", "Which website are you going to run the database server for?")
	rr := Rander()
	rrr := Rander()

	if err := switchutility.ReplaceLine(chosendomain+"/assets/db/rqlite/bin/config.json", `username`, `"username": "`+rr+`",`); err != nil {
		switchutility.Checklogger(err, "trying to update config.json")
	}
	if err := switchutility.ReplaceLine(chosendomain+"/assets/db/rqlite/bin/config.json", `password`, `"password": "`+rrr+`",`); err != nil {
		switchutility.Checklogger(err, "trying to update config.json")
	}
	if err := switchutility.ReplaceLine(chosendomain+"/internal/dbsql/dbconn/dbconn.go", `conn, err := gorqlite.Open(`, `conn, err := gorqlite.Open("http://`+rr+`:`+rrr+`@localhost:4001/")`); err != nil {
		switchutility.Checklogger(err, "trying to update config.json")
	}

	colortermer.ColorizeOutPut("dpurple", "purple", "if you want resources on the db server lookup the following\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "https://github.com/rqlite/rqlite and https://github.com/rqlite/gorqlite\n")

	if err := switchutility.ShellBash("cd "+chosendomain+"/bin && chmod 755 ./rqlited && cd .. && go build -o bin/ main.go   ", "trying to run database server bash command"); err != nil {
		switchutility.Checklogger(err, "running database server")
	}

	colortermer.ColorizeOutPut("dpurple", "purple", `your binary and database are in the /bin folder now.  
	└──bin
		├──main - go binary
		├──assets/assetdirectory.yaml - used for optimizing assets
		├──config.json - used for auth for database
		├──rqbench - 
		├──rqlite - used for sql cli
		├──rqlited - used for database
	 `)
	colortermer.ColorizeOutPut("dpurple", "purple", "Remember you will need to access the certificates for the rqlite database\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "if you are using digitalocean use this https://www.digitalocean.com/community/tutorials/how-to-use-certbot-standalone-mode-to-retrieve-let-s-encrypt-ssl-certificates-on-ubuntu-20-04\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "and remember to use the fullchain.pem as the public pem and to set the permission fo the files so rqlite can use them!\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "like using this but change (yourdomain) sudo chmod 755 /etc/letsencrypt/live/yourdomain/*.pem \n")
	colortermer.ColorizeOutPut("dpurple", "purple", "to start the rqlite database use this as a template\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "./rqlited -auth config.json -http-addr yourdomain.com:25060 -http-cert /etc/letsencrypt/live/yourdomain.com/fullchain.pem -http-key /etc/letsencrypt/live/yourdomain.com/privkey.pem  ~/node.1\n")
}

func Rander() string {
	randomNumber := rand.Intn(1000)
	Randnum := strconv.Itoa(randomNumber)

	return fmt.Sprintf(Randnum)
}
