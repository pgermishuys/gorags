package options

import (
	"os"
	"testing"
)

func Test_EnvironmentVariables_With_No_EnvironmentVariables_Set(t *testing.T) {
	optionSources, err := parseEnvironmentVariables(TestOptions{}, "TEST_")

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 0 {
		t.Errorf("Expected no option sources, but received %+v", len(optionSources))
	}
}

func Test_EnvironmentVariables_With_A_Single_Environment_Variable_Set(t *testing.T) {
	os.Setenv("TEST_PORT", "8080")
	optionSources, err := parseEnvironmentVariables(TestOptions{}, "TEST_")

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 1 {
		t.Errorf("Expected one option source, but received %+v", len(optionSources))
	}

	optionSource := optionSources[0]

	expectedSource := EnvironmentSource
	if optionSource.Source != expectedSource {
		t.Errorf("Expected source to be %s, but got %s", expectedSource, optionSource.Source)
	}

	expectedName := "port"
	if optionSource.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, optionSource.Name)
	}

	expectedValue := "8080"
	if optionSource.Value != expectedValue {
		t.Errorf("Expected value to be %s, but got %s", expectedValue, optionSource.Value)
	}
}
