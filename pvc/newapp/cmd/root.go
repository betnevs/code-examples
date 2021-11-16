package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"


)

var (
	cfgFile     string
	projectBase string
	userLicense string
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "short",
	Long:  "long",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().StringVar(&projectBase, "project", "", "config file")
	rootCmd.PersistentFlags().String("project", "", "config file")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("$HOME")
		viper.SetConfigName(".cobra")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("can not read config:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
