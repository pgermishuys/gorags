package options

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

// ConfigYamlSource ...
const ConfigYamlSource = "Config (Yaml)"

func parseYaml(data []byte) ([]OptionSource, error) {
	var optionSources []OptionSource
	var res interface{}

	err := yaml.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(res)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			structure := v.MapIndex(key)
			optionSources = append(optionSources, OptionSource{
				Source: ConfigYamlSource,
				Name:   strings.ToLower(fmt.Sprint(key.Interface())),
				Value:  strings.ToLower(fmt.Sprint(structure.Interface())),
			})
		}
	}

	return optionSources, nil
}
