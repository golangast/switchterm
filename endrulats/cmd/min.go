package cmd

import (
	"fmt"

	"github.com/golangast/goservershell/optimize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var minCmd = &cobra.Command{
	Use:   "min",
	Short: "Used to minify files",
	Long:  `Used to minify files and directories`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.SetConfigName("assetdirectory") // name of config file (without extension)
		viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")              // path to look for the config file in
		err := viper.ReadInConfig()           // Find and read the config file
		check(err)
		//get paths of asset folders from config file
		cssin := viper.GetString("opt.cssin")
		jsin := viper.GetString("opt.jsin")
		cssout := viper.GetString("opt.cssout")
		jsout := viper.GetString("opt.jsout")
		imgin := viper.GetString("opt.imgin")
		// get all assets file
		err, files := optimize.Getfiles(cssin, ".css")
		check(err)
		err, jsfiles := optimize.Getfiles(jsin, ".js")
		check(err)
		err, imgfiles := optimize.GetImageFiles(imgin)
		check(err)

		//concatenate all assets
		optimize.Concat(files, cssout)
		optimize.Concat(jsfiles, jsout)

		//minify all assets
		optimize.Minifycss(cssout, cssout)
		optimize.Minifyjs(jsout, jsout)

		//optimize images
		if len(imgfiles) > 1 {
			for _, imgins := range imgfiles {
				go func() {
					optimize.Optimizer(imgins)
				}()
			}
		} else {
			for _, imgins := range imgfiles {
				optimize.Optimizer(imgins)
			}
		}

		fmt.Println("assets optimized...")
	},
}

func init() {
	rootCmd.AddCommand(minCmd)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
