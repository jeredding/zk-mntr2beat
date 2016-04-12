package mntr

import (
    "net"
    "time"
    "bytes"
    "io/ioutil"
    "github.com/elastic/beats/libbeat/logp"
    "github.com/elastic/beats/libbeat/common"
    "github.com/elastic/beats/metricbeat/helper"
    _ "github.com/jeredding/zk-mntr2beat/module/zookeeper"
)

func init() {
    helper.Registry.AddMetricSeter("zookeeper", "mntr", New)
}


func New() helper.MetricSeter {
    return &MetricSeter{}
}

type MetricSeter struct {
    Hostname string
    Port     string
    Timeout  time.Duration
}

func (m *MetricSeter) Setup(ms *helper.MetricSet) error {

    // Additional configuration options
    config := struct {
        Hostname  string `config:"hostname"`
        Port      string `config:"port"`
        Timeout   string `config:"timeout"`
    }{
        Hostname: "127.0.0.1",
        Port:     "2181",
        Timeout:  "60s",
    }

    if err := ms.Module.ProcessConfig(&config); err != nil {
        return err
    }

    m.Hostname = config.Hostname
    m.Port = config.Port
    m.Timeout , _ = time.ParseDuration(config.Timeout)
    return nil
}


func (m *MetricSeter) Fetch(ms *helper.MetricSet, host string) (event common.MapStr, err error) {
    // this is basically implementing https://github.com/samuel/go-zookeeper/blob/master/zk/flw.go#L262
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(m.Hostname, m.Port), m.Timeout)
    if err != nil {
        logp.Err("Error connecting to host (%s:%s): %v", m.Hostname, m.Port, err)
        return nil, err
    }

    defer conn.Close()
    // timeout attempting to write...
    conn.SetReadDeadline(time.Now().Add(m.Timeout))

    // now, lets write!
    _ ,err = conn.Write([]byte("mntr"))
    if err != nil {
        logp.Err("Error sending mntr command: %v", err)
        return nil, err
    }

    // lets read, but only for so long...
    conn.SetReadDeadline(time.Now().Add(m.Timeout))

    // gobble gobble...
    result, err := ioutil.ReadAll(conn)

    if err != nil {
        logp.Err("Error reading mntr command: %v", err)
        return nil, err
    }

    return eventMapping(bytes.NewReader(result)), nil
}
