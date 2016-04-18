package mntr

import (
    "bufio"
    "bytes"
    "net"
    "regexp"
    "strconv"
    "time"
    "github.com/elastic/beats/libbeat/logp"
    "github.com/elastic/beats/libbeat/common"
    "github.com/elastic/beats/metricbeat/helper"
    "github.com/jeredding/zkbeat/module/zookeeper"
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
    outputReader := zookeeper.RunCommand("mntr", net.JoinHostPort(m.Hostname, m.Port), m.Timeout)
    return mntrEventMapping(outputReader), nil
}

func mntrEventMapping(response io.Reader) common.MapStr {

    var (
        versionString           string
        avgLatency              int
        minLatency              int
        maxLatency              int
        packetsReceived         int
        packetsSent             int
        numAliveConnections     int
        outstandingRequests     int
        serverState             int
        znodeCount              int
        watchCount              int
        ephemeralsCount         int
        approximateDataSize     int
        openFileDescriptorCount int
        maxFileDescriptorCount  int
        followers               int
        syncedFollowers         int
        pendingSyncs            int
    )

    //  zk_version      3.5.1-alpha-1693007, built on 07/28/2015 07:19 GMT
    //  zk_avg_latency  0
    //  zk_max_latency  1789
    //  zk_min_latency  0
    //  zk_packets_received     22152032
    //  zk_packets_sent 30959914
    //  zk_num_alive_connections        1033
    //  zk_outstanding_requests 0
    //  zk_server_state leader
    //  zk_znode_count  242609
    //  zk_watch_count  940522
    //  zk_ephemerals_count     8565
    //  zk_approximate_data_size        372143564
    //  zk_open_file_descriptor_count   1083
    //  zk_max_file_descriptor_count    1048576
    //  zk_followers    5
    //  zk_synced_followers     2
    //  zk_pending_syncs        0

    var re *regexp.Regexp
    scanner := bufio.NewScanner(response)
    for scanner.Scan() {

        re = regexp.MustCompile("zk_version\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            versionString = matches[1]
        }
        re = regexp.MustCompile("zk_avg_latency\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            avgLatency, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_max_latency\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            maxLatency, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_min_latency\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            minLatency, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_packets_received\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            packetsReceived, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_packets_sent\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            packetsSent, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_num_alive_connections\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            numAliveConnections, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_outstanding_requests\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            outstandingRequests, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_server_state\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            serverState, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_znode_count\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            znodeCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_watch_count\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            watchCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_ephemerals_count\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            ephemeralsCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_approximate_data_size\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            approximateDataSize, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_open_file_descriptor_count\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            openFileDescriptorCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_max_file_descriptor_count\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            maxFileDescriptorCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_followers\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            znodeCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_synced_followers\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            syncedFollowers, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_pending_syncs\\s+(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            pendingSyncs, _ = strconv.Atoi(matches[1])
        }
    }

    event := common.MapStr{
        "version_string":             versionString,
        "avg_latency":                avgLatency,
        "min_latency":                minLatency,
        "max_latency":                maxLatency,
        "packets_received":           packetsReceived,
        "packets_sent":               packetsSent,
        "num_alive_connections":      numAliveConnections,
        "outstanding_requests":       outstandingRequests,
        "server_state":               serverState,
        "znode_count":                znodeCount,
        "watch_count":                watchCount,
        "ephemerals_count":           ephemeralsCount,
        "approximate_data_size":      approximateDataSize,
        "open_file_descriptor_count": openFileDescriptorCount,
        "max_file_descriptor_count":  maxFileDescriptorCount,
        "followers":                  followers,
        "synced_followers":           syncedFollowers,
        "pending_syncs":              pendingSyncs,
    }
    return event
}