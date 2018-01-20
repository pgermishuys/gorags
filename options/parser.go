package options

import (
	"log"
	"os"
)

// Parse command and environment variables
func Parse(option interface{}, prefix string) ([]OptionSource, error) {
	var options []OptionSource

	defaults, err := getDefaults(option)
	if err != nil {
		return options, err
	}

	options = append(options, defaults...)

	cmdLineOptions, err := parseCommandLine(os.Args[1:])
	if err != nil {
		return options, err
	}

	options = append(options, cmdLineOptions...)

	environmentVariables, err := parseEnvironmentVariables(option, prefix)
	if err != nil {
		return options, err
	}

	options = append(options, environmentVariables...)

	return options, nil
}

// Log ...
func Log(optionSources []OptionSource) {
	for _, option := range optionSources {
		log.Printf("Source: %s, Name: %s, Value: %s\n", option.Source, option.Name, option.Value)
	}
}
