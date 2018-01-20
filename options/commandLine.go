package options

import (
	"errors"
	"strings"
)

func parseCommandLine(args []string) ([]OptionSource, error) {
	var options []OptionSource
	for _, item := range args {
		token := item
		if strings.HasPrefix(token, "-") {
			key := token[1:]
			if len(key) == 0 {
				return []OptionSource{}, errors.New("Missing argument value after '-'")
			}
			var value string
			if (strings.HasPrefix(key, "-") && strings.HasSuffix(key, "-")) ||
				(strings.HasPrefix(key, "-") && strings.HasSuffix(key, "+")) {
				value = key[len(key)-1 : len(key)]
				key = key[1 : len(key)-1]
			} else if strings.HasPrefix(key, "-") && strings.Contains(key, "=") {
				var index = strings.Index(key, "=")
				value = string(key[index+1:])
				key = string(key[1:index])
			} else {

			}
			options = append(options, OptionSource{
				Source: "Command Line",
				Name:   strings.ToLower(key),
				Value:  value,
			})
		}
	}
	return options, nil
}
