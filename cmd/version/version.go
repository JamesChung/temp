package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/JamesChung/temp/flags"
	"github.com/JamesChung/temp/flags/lookup"
)

var (
	versionLookup    *lookup.Lookup
	persistentLookup *lookup.Lookup

	versionStep VersionSteps = new(versionRunner)
)

type VersionSteps interface {
	GetVersion() string
}

type versionRunner struct{}

func (v *versionRunner) GetVersion() string {
	return "v1.0.0"
}

func addVersionFlags(versionFlags *pflag.FlagSet) {
	versionFlags.Bool(flags.Version.CheckovVersion.String(), false, "Get checkov version")
}

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		PreRun: func(cmd *cobra.Command, _ []string) {
			persistentLookup = lookup.NewPersistent(cmd)
			versionLookup = lookup.New(cmd)
		},
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println()
			fmt.Println("Start of Version:")
			fmt.Printf("checkov-version: [%s]\n", versionLookup.Get(flags.Version.CheckovVersion))
			fmt.Printf("aws-profile: [%s]\n", versionLookup.Get(flags.Global.AWSProfile))
			fmt.Printf("GetVersion(): [%s]\n", versionStep.GetVersion())
		},
	}

	addVersionFlags(cmd.Flags())

	return cmd
}
