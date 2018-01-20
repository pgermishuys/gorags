package options

import (
	"os"
	"reflect"
	"strings"
)

func parseEnvironmentVariables(opts interface{}, envPrefix string) ([]OptionSource, error) {
	var options []OptionSource

	val := reflect.ValueOf(opts)

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		environmentVariableName := envPrefix + strings.ToUpper(fieldName)
		value := os.Getenv(environmentVariableName)
		if len(value) > 0 {
			options = append(options, OptionSource{
				Source: "Environment",
				Name:   strings.ToLower(fieldName),
				Value:  value,
			})
		}
	}
	return options, nil
}
