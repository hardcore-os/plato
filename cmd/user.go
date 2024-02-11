package cmd

import (
	"github.com/hardcore-os/plato/domain/user"
	"github.com/spf13/cobra"
)

func init() {
	userCmd.AddCommand(domainCMD)
	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "这是用户模块的命令，通常有api和domain两个子命令",
}

var domainCMD = &cobra.Command{
	Use: "domain",
	Run: DomainHandle,
}

func DomainHandle(cmd *cobra.Command, args []string) {
	user.RunMain(ConfigPath)
}
