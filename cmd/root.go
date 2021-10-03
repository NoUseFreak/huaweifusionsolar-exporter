package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "huawei-fusionsolar",
	Short: "huawei-fusionsolar",
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("hfse")
	viper.SetDefault("username", "")
	viper.SetDefault("password", "")
	viper.SetDefault("api_endpoint", "https://eu5.fusionsolar.huawei.com/thirdData")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
