package main

import (
	"embed"
	"log"
	"os"
	"os/exec"
)

//go:embed bash/*
var bash embed.FS

func main() {
	v := []string{"/dev/stdin", "Jon", "Lundy"}
	Execute("bash/b.bash", "bash", v)
}

func Execute(file, types string, commands []string) (*exec.Cmd, error) {

	f, e := os.Open(file)
	if e != nil {
		log.Fatal(e.Error())
	}
	defer f.Close()
	cmd := exec.Command(types, commands...)
	cmd.Stdin = f
	cmd.Stdout = os.Stdout

	if cmd.Err != nil {
		return cmd, cmd.Err
	}

	if e := cmd.Run(); e != nil {
		log.Fatal(e.Error())
	}

	return cmd, nil
}
