package options

import (
	"testing"
)

func Test_Parser_With_Existing_Config_File(t *testing.T) {
	optionSources, err := Parse([]string{"--config=test_config.yaml"}, TestOptions{}, "SAMPLE_")
	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}
	if len(optionSources) != 5 {
		t.Errorf("Expected no option sources, but received %+v", optionSources)
	}
}

func Test_Parser_With_Non_Existing_Config_File(t *testing.T) {
	_, err := Parse([]string{"--config=non_existing_file.yaml"}, TestOptions{}, "SAMPLE_")
	if err == nil {
		t.Errorf("Expected an error, but received none")
	}
}

func Test_Parser_Precedence_Resolution(t *testing.T) {
	optionSources := []OptionSource{
		OptionSource{
			Source: DefaultSource,
			Name:   "Port",
			Value:  "1",
		},
		OptionSource{
			Source: EnvironmentSource,
			Name:   "Port",
			Value:  "2",
		},
		OptionSource{
			Source: ConfigYamlSource,
			Name:   "Host",
			Value:  "192.168.1.4",
		},
		OptionSource{
			Source: EnvironmentSource,
			Name:   "Host",
			Value:  "192.168.1.2",
		},
		OptionSource{
			Source: DefaultSource,
			Name:   "Host",
			Value:  "192.168.1.1",
		},
		OptionSource{
			Source: CommandLineSource,
			Name:   "Host",
			Value:  "192.168.1.3",
		},
		OptionSource{
			Source: ConfigYamlSource,
			Name:   "Port",
			Value:  "4",
		},
		OptionSource{
			Source: CommandLineSource,
			Name:   "Port",
			Value:  "3",
		},
	}
	resolvedSources := resolvePrecedence(optionSources)
	receivedNumberOfSources := len(resolvedSources)
	t.Logf("%+v", resolvedSources)
	if receivedNumberOfSources != 2 {
		t.Errorf("Expected 2 options, got %d", receivedNumberOfSources)
	}

	optionSource := resolvedSources[0]

	expectedSource := CommandLineSource
	if optionSource.Source != expectedSource {
		t.Errorf("Expected source to be %s, but got %s", expectedSource, optionSource.Source)
	}

	expectedName := "Host"
	if optionSource.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, optionSource.Name)
	}

	expectedValue := "192.168.1.3"
	if optionSource.Value != expectedValue {
		t.Errorf("Expected value to be %s, but got %s", expectedValue, optionSource.Value)
	}

	optionSource = resolvedSources[1]

	expectedSource = CommandLineSource
	if optionSource.Source != expectedSource {
		t.Errorf("Expected source to be %s, but got %s", expectedSource, optionSource.Source)
	}

	expectedName = "Port"
	if optionSource.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, optionSource.Name)
	}

	expectedValue = "3"
	if optionSource.Value != expectedValue {
		t.Errorf("Expected value to be %s, but got %s", expectedValue, optionSource.Value)
	}
}
