package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}
	if *cfg != "" {
		viper.SetConfigFile(*cfg)
		viper.SetConfigType("yaml")
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.iam")
		viper.SetConfigName("config")
	}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Printf("used config file is: %s\n", viper.ConfigFileUsed())
}
