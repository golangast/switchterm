package configure

import (
	"fmt"
	"strings"

	ff "github.com/golangast/gentil/utility/ff"
	temp "github.com/golangast/gentil/utility/temp"
	config "github.com/golangast/switchterm/configure/templates/config"
	"github.com/spf13/viper"
)

type Config struct {
	Path string
	File string
}

func LoadConfig() []string {
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {
		fmt.Println(err)
	}
	//get paths of asset folders from config file
	cmds := viper.GetStringSlice("cmd.basic.cmds")
	return cmds

}

func AddCommand(cmd string) {
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {
		fmt.Println(err)
	}
	cmds := viper.GetStringSlice("cmd.basic.cmds")
	strings.TrimSpace(cmd)

	cmds = append(cmds, cmd)
	viper.Set("cmd.basic.cmds", cmds)
	viper.WriteConfig()
}
func RemoveCommand(cmds []string) {
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {
		fmt.Println(err)
	}
	viper.Set("cmd.basic.cmds", cmds)
	viper.WriteConfig()
}

func GenConfigure() {
	sfile, err := ff.Filefolder("./config", "config.yaml")
	if err != nil {
		fmt.Print(err)
	}

	/* write to files*/
	temp.Writetemplate(config.Configtemp, sfile, nil)
	if err != nil {
		fmt.Print(err)
	}
}
