package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/JamesChung/temp/cmd/version"
	"github.com/JamesChung/temp/flags"
)

func addPersistentFlags(pflags *pflag.FlagSet) {
	pflags.String(flags.Global.AWSProfile.String(), "", "example aws profile")
}

func addSubCommands(cmd *cobra.Command) {
	cmd.AddCommand(version.NewVersionCommand())
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "root",
	}

	addPersistentFlags(cmd.PersistentFlags())
	addSubCommands(cmd)

	return cmd
}
