package options

import "reflect"

// OptionSource ...
type OptionSource struct {
	Source string
	Name   string
	Value  string
	Type   reflect.Type
}
