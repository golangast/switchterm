/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"log/slog"

	"os"

	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/golangast/endrulats/internal/loggers"
	gentil "github.com/golangast/gentil/utility/ff"
	term "github.com/golangast/gentil/utility/term"
	text "github.com/golangast/gentil/utility/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// gengetCmd represents the genget command
var gengetCmd = &cobra.Command{
	Use:   "genget",
	Short: "generates a get route",
	Long:  `go run . genget -r about`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("assetdirectory") // name of config file (without extension)
		viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("./optimize/")    // path to look for the config file in
		err := viper.ReadInConfig()           // Find and read the config file
		check(err)
		//get paths of asset folders from config file
		repo := viper.GetString("opt.repo")
		logger := loggers.CreateLogger()

		route, _ := cmd.Flags().GetString("route")
		folder, _ := cmd.Flags().GetString("folder")
		folderdir := folder + "/"

		/* create folders/files*/
		handlerfile, err := gentil.Filefolder("./src/handler/get/"+folderdir+route, route+".go")
		if err != nil {
			logger.Error(
				"trying to create handler file",
				slog.String("error: ", err.Error()),
			)
		}

		/* create folders/files*/
		tempfile, err := gentil.Filefolder("./assets/templates/"+folderdir+route, route+".html")
		if err != nil {
			logger.Error(
				"trying to create tempfile file",
				slog.String("error: ", err.Error()),
			)
		}

		//replace imports
		//update router with route
		found := text.FindTextNReturn("./src/routes/routes.go", `e.GET("/`+route+`", `+route+`.`+cases.Title(language.Und, cases.NoLower).String(route)+`)`)
		if found != `e.GET("/`+route+`", `+route+`.`+cases.Title(language.Und, cases.NoLower).String(route)+`)` {
			err := text.UpdateText("./src/routes/routes.go", "//routes", `e.GET("/`+route+`", `+route+`.`+cases.Title(language.Und, cases.NoLower).String(route)+`)`+"\n"+`//routes`)
			if err != nil {
				logger.Error(
					"trying to update router.go",
					slog.String("error: ", err.Error()),
				)
			}
		}
		//update the route.go for the import
		found = text.FindTextNReturn("./src/routes/routes.go", repo+"/src/handler/get/"+route)
		if found != `"`+repo+`/src/handler/get/`+route+`"` {
			err := text.UpdateText("./src/routes/routes.go", "//imports", `"`+repo+`/src/handler/get/`+route+`"`+"\n"+`//imports`)
			if err != nil {
				logger.Error(
					"trying to update router.go for the import",
					slog.String("error: ", err.Error()),
				)
			}
		}

		handler := `
	package ` + route + `

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ` + cases.Title(language.Und, cases.NoLower).String(route) + `(c echo.Context) error {
	//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
	nonce := c.Get("n")
	jsr := c.Get("jsr")
	cssr := c.Get("cssr")

	return c.Render(http.StatusOK, "` + route + `.html", map[string]interface{}{
		"nonce": nonce,
		"jsr":   jsr,
		"cssr":  cssr,
	})

}`

		//footer/header map for {{template "footer" .}} {{end}}
		mb := make(map[string]string)
		headerb := fmt.Sprintf(`{{template "header" .}}%s`, "")
		footerb := fmt.Sprintf(`{{template "footer" .}}%s`, "")
		mb["route"] = route
		mb["footer"] = footerb
		mb["header"] = headerb
		//write it to the html file
		var Bodytemp = `
{{ .header }}
You created {{.route}}
<!-- write your code here -->
{{ .footer }}
`
		err = Writetemplate(Bodytemp, tempfile, mb)
		if err != nil {
			logger.Error(
				"trying to update router.html",
				slog.String("error: ", err.Error()),
			)
		}
		//write it to the html file
		err = Writetemplate(handler, handlerfile, nil)
		if err != nil {
			logger.Error(
				"trying to update the handler",
				slog.String("error: ", err.Error()),
			)
		}
		out, errout, err := term.Shellout(`go mod tidy && go mod vendor`)
		if err != nil {
			logger.Error(
				"trying to pull down dependencies"+errout+out,
				slog.String("error: ", err.Error()),
			)
		}

		handlerfile.Close()

	},
}

func init() {
	rootCmd.AddCommand(gengetCmd)
	gengetCmd.Flags().StringP("route", "r", "", "Set your route")
	gengetCmd.Flags().StringP("folder", "f", "", "Set your folder")
}

func Writetemplate(temp string, f *os.File, d map[string]string) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}
