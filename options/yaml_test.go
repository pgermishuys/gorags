package options

import (
	"testing"
)

func Test_Yaml_With_No_Variables(t *testing.T) {
	optionSources, err := parseYaml([]byte{})
	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}
	if len(optionSources) != 0 {
		t.Errorf("Expected no option sources, but received %+v", optionSources)
	}
}

func Test_Yaml_With_Invalid_Variable(t *testing.T) {
	var data = `
	v: [A,
`
	_, err := parseYaml([]byte(data))

	if err == nil {
		t.Error("Expected an error, but received none")
	}
}

func Test_Yaml_With_Multiple_Variables(t *testing.T) {

	var data = `
host: localhost
port: 8080
`
	optionSources, err := parseYaml([]byte(data))

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 2 {
		t.Errorf("Expected two option sources, but received %+v", len(optionSources))
	}

	optionSource := optionSources[0]

	expectedSource := ConfigYamlSource
	if optionSource.Source != expectedSource {
		t.Errorf("Expected source to be %s, but got %s", expectedSource, optionSource.Source)
	}

	expectedName := "host"
	if optionSource.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, optionSource.Name)
	}

	expectedValue := "localhost"
	if optionSource.Value != expectedValue {
		t.Errorf("Expected value to be %s, but got %s", expectedValue, optionSource.Value)
	}

	optionSource = optionSources[1]

	expectedSource = ConfigYamlSource
	if optionSource.Source != expectedSource {
		t.Errorf("Expected source to be %s, but got %s", expectedSource, optionSource.Source)
	}

	expectedName = "port"
	if optionSource.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, optionSource.Name)
	}

	expectedValue = "8080"
	if optionSource.Value != expectedValue {
		t.Errorf("Expected value to be %s, but got %s", expectedValue, optionSource.Value)
	}
}
