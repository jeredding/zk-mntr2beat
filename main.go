package main

import (
	"os"

	"github.com/elastic/beats/metricbeat/beater"
	_ "github.com/jeredding/zk-mntr2beat/module/zookeeper/mntr"
	_ "github.com/jeredding/zk-mntr2beat/module/zookeeper"

	"github.com/elastic/beats/libbeat/beat"
)

var Name = "zk-mntr2beat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
