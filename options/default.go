package options

import (
	"fmt"
	"reflect"
	"strings"
)

func getDefaults(opts interface{}) ([]OptionSource, error) {
	var options []OptionSource

	val := reflect.ValueOf(opts)
	indirect := reflect.Indirect(val)
	typeOfOpts := indirect.Type()

	for i := 0; i < typeOfOpts.NumField(); i++ {
		name := typeOfOpts.Field(i).Name
		value := fmt.Sprintf("%v", indirect.Field(i))
		if len(value) > 0 {
			options = append(options, OptionSource{
				Source: "Default",
				Name:   strings.ToLower(name),
				Value:  value,
			})
		}
	}
	return options, nil
}
