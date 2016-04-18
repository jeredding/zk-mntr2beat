// +build integration

package mntr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/jeredding/zk-mntr2docker/module/zookeeper"
)

func TestConnect(t *testing.T) {

	config, _ := getZookeeperModuleConfig()

	module, mErr := helper.NewModule(config, zookeeper.New)
	assert.NoError(t, mErr)
	ms, msErr := helper.NewMetricSet("status", New, module)
	assert.NoError(t, msErr)

	// Setup metricset and metricseter
	err := ms.Setup()
	assert.NoError(t, err)
	err = ms.MetricSeter.Setup(ms)
	assert.NoError(t, err)

	// Check that host is correctly set
	assert.Equal(t, zookeeper.GetZookeeperEnvHost(), ms.Config.Hosts[0])

	data, err := ms.MetricSeter.Fetch(ms, ms.Config.Hosts[0])
	assert.NoError(t, err)

	// Check fields
	assert.Equal(t, 13, len(data))
}

type ZookeeperModuleConfig struct {
	Hosts  []string `config:"hostname"`
	Module string   `config:"module"`
}

func getZookeeperModuleConfig() (*common.Config, error) {
	return common.NewConfigFrom(ZookeeperModuleConfig{
		Module:   "zookeeper",
		Hostname: []string{zookeeper.GetZookeeperEnvHost()},
		Port:     []string{zookeeper.GetZookeeperEnvPort()},
	})
}
