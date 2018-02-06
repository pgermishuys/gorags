package options

import (
	"testing"
)

func Test_Parser_With_Existing_Config_File(t *testing.T) {
	options := TestOptions{}

	optionSources, err := Parse([]string{"--config=test_config.yaml", "--port=2113", "--host=192.168.1.1"}, &options, "SAMPLE_")
	if err != nil {
		t.Errorf("Expected no error, but received %s", err.Error())
	}
	if len(optionSources) != 3 {
		t.Errorf("Expected three option sources, but received %+v", len(optionSources))
	}
	expectedHost := "192.168.1.1"
	if options.Host != expectedHost {
		t.Errorf("Expected host to be %s but got %s", expectedHost, options.Host)
	}
	expectedPort := 2113
	if options.Port != expectedPort {
		t.Errorf("Expected port to be %d but got %d", expectedPort, options.Port)
	}
}

func Test_Parser_With_Non_Existing_Config_File(t *testing.T) {
	_, err := Parse([]string{"--config=non_existing_file.yaml"}, TestOptions{}, "SAMPLE_")
	if err == nil {
		t.Errorf("Expected an error, but received none")
	}
}
