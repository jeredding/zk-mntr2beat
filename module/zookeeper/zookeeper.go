package zookeeper

import (
    "github.com/elastic/beats/metricbeat/helper"
)

func init() {
    helper.Registry.AddModuler("zookeeper", Moduler{})
}

type Moduler struct {}

func (r Moduler) Setup() error {
    return nil
}
