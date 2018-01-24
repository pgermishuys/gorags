package options

import (
	"fmt"
	"reflect"
)

func ensureExistence(optionSources []OptionSource, instance interface{}) error {
	val := reflect.ValueOf(instance)
	for _, item := range optionSources {
		_, found := val.Type().FieldByName(item.Name)
		if !found {
			return fmt.Errorf("The field %s does not exist", item.Name)
		}
	}
	return nil
}
