package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

type optionSource struct {
	name  string
	value string
}

func parseCommandLine(args []string) ([]optionSource, error) {
	options := make([]optionSource, 0)
	for _, item := range args {
		token := item
		if strings.HasPrefix(token, "-") {
			key := token[1:]
			if len(key) == 0 {
				return []optionSource{}, errors.New("Missing argument value after '-'")
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
			options = append(options, optionSource{
				name:  key,
				value: value,
			})
		}
	}
	return options, nil
}

func main() {
	options, err := parseCommandLine(os.Args[1:])
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, option := range options {
		log.Printf("name: %s, value: %s\n", option.name, option.value)
	}
}
