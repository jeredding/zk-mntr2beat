package zookeeper

import (
    "github.com/elastic/beats/metricbeat/helper"
)

func init() {
    helper.Registry.AddModuler("zookeeper", New)
}

// New creates new instance of Moduler
func New() helper.Moduler {
    return &Moduler{}
}

type Moduler struct {}

func (r Moduler) Setup(mo *helper.Module) error {
    return nil
}
