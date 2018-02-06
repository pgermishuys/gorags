package main

import (
	"log"
	"os"

	"github.com/pgermishuys/gorags/options"
)

// SampleOptions ...
type SampleOptions struct {
	Host string
	Port int
}

func main() {
	sampleOptions := SampleOptions{
		Host: "127.0.0.1",
		Port: 8080,
	}
	opts, err := options.Parse(os.Args[1:], &sampleOptions, "SAMPLE_")

	if err != nil {
		log.Fatal(err)
	}

	options.Log(opts)
	log.Printf("Parsed: %+v\n", sampleOptions)
}
