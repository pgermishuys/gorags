package options

import (
	"io/ioutil"
	"log"
)

// Parse ...
/*
	Precedence
	0 Default
	1 Config
	2 Environment Variable
	3 Command Line
*/
func Parse(args []string, option interface{}, prefix string) ([]OptionSource, error) {
	var options []OptionSource

	defaults, err := getDefaults(option)
	if err != nil {
		return options, err
	}

	options = append(options, defaults...)

	environmentVariables, err := parseEnvironmentVariables(option, prefix)
	if err != nil {
		return options, err
	}

	options = append(options, environmentVariables...)

	cmdLineOptions, err := parseCommandLine(args)
	if err != nil {
		return options, err
	}

	options = append(options, cmdLineOptions...)

	configFileOptionSource := getConfigFile(options)

	if configFileOptionSource != (OptionSource{}) {
		configFileOptionSources, err := readConfigFile(configFileOptionSource)
		if err != nil {
			return options, err
		}
		options = append(options, configFileOptionSources...)
	}

	return options, nil
}

func readConfigFile(configFileSource OptionSource) ([]OptionSource, error) {
	contents, err := ioutil.ReadFile(configFileSource.Value)
	if err != nil {
		return []OptionSource{}, err
	}
	return parseYaml(contents)
}

func getConfigFile(optionSources []OptionSource) OptionSource {
	for _, item := range optionSources {
		if item.Name == "config" {
			return item
		}
	}
	return OptionSource{}
}

// Log ...
func Log(optionSources []OptionSource) {
	for _, option := range optionSources {
		log.Printf("Source: %s, Name: %s, Value: %s\n", option.Source, option.Name, option.Value)
	}
}
