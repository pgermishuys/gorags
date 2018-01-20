package main

import (
	"log"

	"github.com/pgermishuys/gorags/options"
)

type SampleOptions struct {
	Host string
	Port int
}

func main() {
	opts, err := options.Parse(SampleOptions{
		Host: "127.0.0.1",
	}, "SAMPLE_")

	if err != nil {
		log.Fatal(err)
	}

	options.Log(opts)
}
