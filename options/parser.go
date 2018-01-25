package options

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func resolvePrecedence(optionSources []OptionSource) []OptionSource {
	sorted := toSortedMap(optionSources)
	effective := make(map[string]OptionSource, 0)
	for i := len(sorted); i > 0; i-- {
		for _, item := range sorted[i] {
			if _, ok := effective[item.Name]; !ok {
				effective[item.Name] = item
			}
		}
	}
	var done []OptionSource
	for _, item := range effective {
		done = append(done, item)
	}
	return done
}

func toSortedMap(optionSources []OptionSource) map[int][]OptionSource {
	sources := make(map[int][]OptionSource)
	for _, item := range optionSources {
		switch item.Source {
		case DefaultSource:
			if _, ok := sources[0]; !ok {
				sources[0] = make([]OptionSource, 0)
			}
			sources[0] = append(sources[0], item)
			break
		case EnvironmentSource:
			if _, ok := sources[1]; !ok {
				sources[1] = make([]OptionSource, 0)
			}
			sources[1] = append(sources[1], item)
			break
		case ConfigYamlSource:
			if _, ok := sources[2]; !ok {
				sources[2] = make([]OptionSource, 0)
			}
			sources[1] = append(sources[1], item)
			break
		case CommandLineSource:
			if _, ok := sources[3]; !ok {
				sources[3] = make([]OptionSource, 0)
			}
			sources[3] = append(sources[3], item)
			break
		}
	}
	return sources
}

func readConfigFile(configFileSource OptionSource) ([]OptionSource, error) {
	if _, err := os.Stat(configFileSource.Value); os.IsNotExist(err) {
		return []OptionSource{}, fmt.Errorf("The file at %s does not exist", configFileSource.Value)
	}
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
