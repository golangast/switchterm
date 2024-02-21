/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/golangast/endrulats/internal/loggers"
	"github.com/spf13/cobra"
)

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "A brief description of your command",
	Long:  `go run . rest -t post -n dog -f name.string age.int`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rest called")
		logger := loggers.CreateLogger()

		types, _ := cmd.Flags().GetString("types")
		// name, _ := cmd.Flags().GetString("name")
		// fields, _ := cmd.Flags().GetString("fields")
		// checkfields, _ := cmd.Flags().GetString("checkfields")
		// folder, _ := cmd.Flags().GetString("folder")
		//folderdir := folder + "/"

		// field := GetPropDatatype(fields)
		// chfields := strings.Split(checkfields, " ")

		//d := Fielddata{Fields: field, CheckFields: chfields, Lowercasename: name, Uppercasename: cases.Title(language.Und, cases.NoLower).String(name)}

		switch types {
		case "":
			logger.Error(
				"trying to update route but no types",
				slog.String("types: ", types),
			)
		case "post", "Post", "POST":
			//poster.Poster(d, types, folderdir)

		default:
			// do something if none of the cases match
		}

	},
}

func init() {
	rootCmd.AddCommand(restCmd)
	restCmd.Flags().StringP("types", "t", "", "Set your types")
	restCmd.Flags().StringP("name", "n", "", "Set your name")
	restCmd.Flags().StringP("fields", "f", "", "Set your fields")
	restCmd.Flags().StringP("folder", "o", "", "Set your folder")
	restCmd.Flags().StringP("checkfields", "c", "", "fields that are going to be checked")

}
func GetPropDatatype(prop string) []string {
	var property []string
	var types []string
	var field []string
	var strright string
	s := strings.Split(prop, " ")

	for _, ss := range s {
		sss := strings.Replace(ss, "\"", "", -1)
		property = append(property, TrimDot(sss))
		strright = strings.Replace(TrimDotright(sss), ".", "", 1)
		types = append(types, strright)
	}

	for a, str1_word := range property {
		for b, str2_word := range types {
			if a == b {
				field = append(field, str1_word+" "+str2_word)
			}
		}
	}
	return field
}
func TrimDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}

type Fielddata struct {
	Fields        []string
	CheckFields   []string
	Lowercasename string
	Uppercasename string
}
