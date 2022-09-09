package cmd

import (
	"github.com/hardcore-os/plato/gateway"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
	Use: "gateway",
	Run: GatewayHandle,
}

func GatewayHandle(cmd *cobra.Command, args []string) {
	gateway.RunMain(ConfigPath)
}
