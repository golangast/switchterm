package configure

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Path string
	File string
}

func LoadConfig() []string {
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../config/") // path to look for the config file in
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {
		fmt.Println(err)
	}
	//get paths of asset folders from config file
	cmds := viper.GetStringSlice("cmd.basic.cmds")
	return cmds
	// viper.Set(f, t)
	// viper.WriteConfig()

}

func AddCommand(cmd string) {
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../config/") // path to look for the config file in
	err := viper.ReadInConfig()       // Find and read the config file
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
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../config/") // path to look for the config file in
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {
		fmt.Println(err)
	}
	viper.Set("cmd.basic.cmds", cmds)
	viper.WriteConfig()
}
