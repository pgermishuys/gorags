package options

import "testing"

func TestCommandLine_With_No_Arguments(t *testing.T) {
	optionSources, err := parseCommandLine([]string{})
	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}
	if len(optionSources) != 0 {
		t.Errorf("Expected no option sources, but received %+v", optionSources)
	}
}

func TestCommandLine_With_Single_Invalid_Argument(t *testing.T) {
	args := []string{"-", "port=8080"}
	_, err := parseCommandLine(args)

	if err == nil {
		t.Error("Expected an error, but received none")
	}
}

func TestCommandLine_With_Single_Long_Form_Argument(t *testing.T) {
	args := []string{"--port=8080"}
	optionSources, err := parseCommandLine(args)

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 1 {
		t.Errorf("Expected a single option source, but received %+v", optionSources)
	}

	optionSource := optionSources[0]

	expectedSource := "Command Line"
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

func TestCommandLine_With_Single_Long_Form_Argument_Without_Equals(t *testing.T) {
	args := []string{"--port", "8080"}
	optionSources, err := parseCommandLine(args)

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 1 {
		t.Errorf("Expected a single option source, but received %+v", len(optionSources))
	}

	optionSource := optionSources[0]

	expectedSource := "Command Line"
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

func TestCommandLine_With_Multiple_Long_Form_Argument_Without_Equals(t *testing.T) {
	args := []string{"--port", "8080", "--host", "localhost"}
	optionSources, err := parseCommandLine(args)

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 2 {
		t.Errorf("Expected two option sources, but received %+v", len(optionSources))
	}

	optionSource := optionSources[0]

	expectedSource := "Command Line"
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

	optionSource = optionSources[1]

	expectedSource = "Command Line"
	if optionSource.Source != expectedSource {
		t.Errorf("Expected source to be %s, but got %s", expectedSource, optionSource.Source)
	}

	expectedName = "host"
	if optionSource.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, optionSource.Name)
	}

	expectedValue = "localhost"
	if optionSource.Value != expectedValue {
		t.Errorf("Expected value to be %s, but got %s", expectedValue, optionSource.Value)
	}
}

func TestCommandLine_With_Single_Short_Form_Argument(t *testing.T) {
	args := []string{"--mem-db-"}
	optionSources, err := parseCommandLine(args)

	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}

	if len(optionSources) != 1 {
		t.Errorf("Expected a single option source, but received %+v", len(optionSources))
	}

	optionSource := optionSources[0]

	expectedSource := "Command Line"
	if optionSource.Source != expectedSource {
		t.Errorf("Expected source to be %s, but got %s", expectedSource, optionSource.Source)
	}

	expectedName := "mem-db"
	if optionSource.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, optionSource.Name)
	}

	expectedValue := "-"
	if optionSource.Value != expectedValue {
		t.Errorf("Expected value to be %s, but got %s", expectedValue, optionSource.Value)
	}
}
