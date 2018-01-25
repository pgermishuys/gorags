package options

import "testing"

type TestOptions struct {
	Host string
	Port int
}

func Test_Defaults_With_No_Explicit_Defaults(t *testing.T) {
	optionSources, err := getDefaults(TestOptions{})

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 2 {
		t.Errorf("Expected two option sources, but received %+v", len(optionSources))
	}
}

func Test_Defaults_With_A_Single_Default_Explicitly_Set(t *testing.T) {
	optionSources, err := getDefaults(TestOptions{
		Port: 8080,
	})

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 2 {
		t.Errorf("Expected two option sources, but received %+v", len(optionSources))
	}

	optionSource := optionSources[0]

	expectedSource := DefaultSource
	if optionSource.Source != expectedSource {
		t.Errorf("Expected source to be %s, but got %s", expectedSource, optionSource.Source)
	}

	expectedName := "host"
	if optionSource.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, optionSource.Name)
	}

	expectedValue := ""
	if optionSource.Value != expectedValue {
		t.Errorf("Expected value to be %s, but got %s", expectedValue, optionSource.Value)
	}

	optionSource = optionSources[1]

	expectedSource = DefaultSource
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
