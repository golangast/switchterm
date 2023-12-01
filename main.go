package main

"github.com/golangast/switchterm/configure"
"github.com/golangast/switchterm/switchtermer"

func main() {

	configure.GenConfigure()

	//commands
	cmds := configure.LoadConfig()

	//function to search or select a command
	answ := switchtermer.SwitchCol(cmds, 6, "red", "green")

	//returned command
	fmt.Println("you have chosen the command: ", answ)

}
