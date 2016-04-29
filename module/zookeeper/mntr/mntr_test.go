// +build integration

package mntr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"time"

	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/module/zookeeper"
)

func TestConnect(t *testing.T) {

	config := helper.ModuleConfig{
		Hosts: []string{zookeeper.GetZookeeperEnvHost() + ":" + zookeeper.GetZookeeperEnvPort()},
	}

	timeout, _ := 5 * time.Second
	module := &helper.Module{
		Config:  config,
		Timeout: timeout,
	}

	ms, msErr := helper.NewMetricSet("mntr", New, module)
	assert.NoError(t, msErr)

	var err error

	// Setup metricset
	err = ms.Setup()
	assert.NoError(t, err)

	// Load events
	event, err := ms.MetricSeter.Fetch(ms, module.Config.Hosts[0])
	assert.NoError(t, err)

	// Check values
	version := event["version_string"].(string)
	avgLatency := event["avg_latency"].(int)
	maxLatency := event["max_latency"].(int)
	numAliveConnections := event["num_alive_connections"].(int)

	assert.Equal(t, version, "3.4.8--1, built on 02/06/2016 03:18 GMT")
	assert.True(t, avgLatency >= 0)
	assert.True(t, maxLatency >= 0)
	assert.True(t, numAliveConnections > 0)

	// Check fields
	assert.Equal(t, 18, len(event))

}
