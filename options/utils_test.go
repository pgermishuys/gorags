package options

import "testing"

func Test_EnsureExistence_With_An_OptionSource_Which_Does_Not_Exist(t *testing.T) {
	optionSources := []OptionSource{
		OptionSource{
			Source: "Command Line",
			Name:   "Port",
			Value:  "8080",
		},
		OptionSource{
			Source: "Command Line",
			Name:   "NonExistent",
			Value:  "BogusValue",
		},
	}
	err := ensureExistence(optionSources, TestOptions{})
	if err == nil {
		t.Errorf("Expected a single error, but received none")
	}
}

func Test_EnsureExistence_With_OptionSources_That_Does_Exist(t *testing.T) {
	optionSources := []OptionSource{
		OptionSource{
			Source: "Command Line",
			Name:   "Host",
			Value:  "localhost",
		},
		OptionSource{
			Source: "Command Line",
			Name:   "Port",
			Value:  "8080",
		},
	}
	err := ensureExistence(optionSources, TestOptions{})
	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}
}
