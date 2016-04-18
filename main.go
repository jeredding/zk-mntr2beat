package main

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/metricbeat/beater"
	_ "github.com/elastic/beats/metricbeat/include"
	_ "github.com/jeredding/zkbeat/module/zookeeper"
	_ "github.com/jeredding/zkbeat/module/zookeeper/mntr"
	"os"
)

var Name = "zkbeat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
