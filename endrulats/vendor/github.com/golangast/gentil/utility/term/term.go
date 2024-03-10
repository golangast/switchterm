package term

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// used for terminal
func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// kill a process
func killProcessByName(procname string) int {
	kill := exec.Command("taskkill", "/im", procname, "/T", "/F")
	err := kill.Run()
	if err != nil {
		return -1
	}
	return 0
}
func Runapplinuxmac() (string, string, error) {
	err, out, errout := Shellout(`go run *.go`)
	return out, err, errout
}
func Runwindows() (string, string, error) {
	err, out, errout := Shellout(`go run .`)
	return out, err, errout
}
func Reload() (string, string, error) {
	err, out, errout := Shellout("pwd && cd app && go mod tidy && go mod vendor && go install && go build")
	return out, err, errout

}

// pulls down dependencies + installs echo
func Pulldowneverything(p string) {
	out, errout, err := Shellout("cd .. && cd " + p + " && go mod init " + p + "&& go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)

	outs, errouts, err := Shellout("cd .. && cd " + p + " && go install github.com/labstack/echo/v4 && go install github.com/labstack/echo/v4/middleware && go get github.com/labstack/gommon/log && go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
}
