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
