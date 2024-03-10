package genserver

import (
	"strings"

	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Genserver() {
	//ask questions for domain and git
	stripdir := switchutility.InputScan("What is the domain name of your site without www or .com?")
	stripgit := switchutility.InputScan("What is your the domain you are going to use for imports? (example github.com/golangast)")
	if !strings.HasSuffix(stripgit, "/") {
		stripgit = stripgit + "/"
	}
	//store data
	d := domain.Domains{Domain: stripdir, Github: stripgit}
	exists, err := d.Exists(d.Domain)
	switchutility.Checklogger(err, "getting domain and checking if it exists")
	if !exists {
		d.Create()
	} else {
		colortermer.ColorizeOutPut("purple", "purple", "DOMAIN ALREADY EXISTS IN YOUR DATABASE"+"\n")
	}

	colortermer.ColorizeOutPut("purple", "purple", "CREATING SERVER! THIS MAY TAKE A FEW SECONDS..."+"\n")
	//run command to generate server
	switchutility.ShellBash("go install golang.org/x/tools/cmd/gonew@latest && gonew github.com/golangast/genserv "+stripgit+stripdir+" && cd "+stripdir+
		" && go mod init "+stripgit+stripdir+" && go mod tidy && go mod vendor", "generating server pull down")

	switchutility.ShellBash("cd "+stripdir+" && go mod tidy && go mod vendor && go build", "generating server pull down")

	colortermer.ColorizeOutPut("purple", "purple", `
	PLEASE NOTE: .....

	1. THIS SERVER HAS THE TLS COMMENTED OUT IN MAIN.GO
	2. YOU WILL NEED TO CHANGE THE IMPORTS NAME FROM GITHUB.COM/GOLANGAST/GENSERVER TO YOURS
	3. THERE ARE ALSO IMAGES IN THE CSS FILES THAT THEIR DOMAIN NEEDS TO BE UPDATED.
	4. REMEMBER TO CHANGE THE localhost TO YOUR DOMAIN BEFORE RUNNING!!
	5. RUN UNDER GENERATE COMMAND CERTIFICATES TO ISSUE DOMAIN CERTS
	6. RUN UNDER GENERATE TLS TO SECURE DOMAIN
	7. RUN DEV TO START DEVING
	8. RUN HANDLER TO GENERATE A HANDLER
	9. RUN DATA TO CREATE DATA FOR THE HANDLER
	10. RUN ADD DATA TO THE HANDLER.
	11. THE PROJECT USES https://github.com/rqlite/gorqlite FOR THE DATABASE
	12. REMEMBER TO RUN GO MOD TIDY AND GO MOD VENDOR BEFORE RUNNING IT IN YOUR PROJECT FOLDER

	`)
	colortermer.ColorizeOutPut("purple", "purple", "DONE GENERATING AND PLEASE ENJOY!")
}
