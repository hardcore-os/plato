package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ConfigPath string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&ConfigPath,
		"config",
		"./plato.yaml",
		"config file (default is ./plato.yaml)")
}

var rootCmd = &cobra.Command{
	Use:   "plato",
	Short: "这是一个超牛逼的IM系统",
	Run:   Plato,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Plato(cmd *cobra.Command, args []string) {

}

func initConfig() {

}
