package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gin-example",
	Short: "Gin example cli",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogging()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gin-example.yaml)")

	RootCmd.PersistentFlags().String("log-level", "info", "one of debug, info, warn, error, or fatal")
	RootCmd.PersistentFlags().String("log-format", "text", "specify output (text or json)")

	viper.BindPFlags(RootCmd.PersistentFlags()) // bind all pflags to viper
	viper.BindPFlags(RootCmd.Flags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".gin-example")                    // name of config file (without extension)
	viper.AddConfigPath("$HOME")                           // adding home directory as first search path
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_")) // use "_" in place of "-" in environment variables
	viper.AutomaticEnv()                                   // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.WithField("config_file", viper.ConfigFileUsed()).Debug("using config file")
	}
}
