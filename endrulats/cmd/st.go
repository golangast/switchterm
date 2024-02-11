/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/golangast/endrulats/internal/dbsql/dbconn"
	"github.com/spf13/cobra"
)

// stCmd represents the st command
var stCmd = &cobra.Command{
	Use:   "st",
	Short: "to build and start program",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("st called")
		_, err := dbconn.DbConnection() //create db instance
		dbconn.ErrorCheck(err)
		fmt.Println("db connected")

		err, out, errout := Startprograms(`go build && ./genserv`)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println(out)
		fmt.Println("--- errs ---")
		fmt.Println(errout)

	},
}

func init() {
	rootCmd.AddCommand(stCmd)

}
func Startprograms(command string) (error, string, string) {
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
	return err, stdout.String(), stderr.String()
}
