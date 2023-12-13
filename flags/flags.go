package flags

import "github.com/JamesChung/temp/flags/lookup"

type flagName string

func (f flagName) String() string {
	return string(f)
}

type globalFlags struct {
	AWSProfile lookup.Flag
}

type versionFlags struct {
	CheckovVersion lookup.Flag
}

var (
	Global = globalFlags{
		AWSProfile: flagName("aws-profile"),
	}

	Version = versionFlags{
		CheckovVersion: flagName("checkov-version"),
	}
)
