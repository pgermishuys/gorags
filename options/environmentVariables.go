package options

import (
	"os"
	"reflect"
	"strings"
)

// EnvironmentSource ...
const EnvironmentSource = "Environment"

func parseEnvironmentVariables(opts interface{}, envPrefix string) ([]OptionSource, error) {
	var options []OptionSource

	val := reflect.ValueOf(opts)
	indirect := reflect.Indirect(val)
	typeOfOpts := indirect.Type()

	for i := 0; i < typeOfOpts.NumField(); i++ {
		fieldName := typeOfOpts.Field(i).Name
		environmentVariableName := envPrefix + strings.ToUpper(fieldName)
		value := os.Getenv(environmentVariableName)
		if len(value) > 0 {
			options = append(options, OptionSource{
				Source: EnvironmentSource,
				Name:   strings.ToLower(fieldName),
				Value:  value,
			})
		}
	}
	return options, nil
}
