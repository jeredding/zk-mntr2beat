package mntr

import (
    "io"
    "bufio"
    "strconv"
    "regexp"
    "github.com/elastic/beats/libbeat/common"
)

func eventMapping(response io.Reader) common.MapStr {

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

        re = regexp.MustCompile("zk_version +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            versionString = matches[1]
        }
        re = regexp.MustCompile("zk_avg_latency +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            avgLatency, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_max_latency +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            maxLatency, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_min_latency +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            minLatency, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_packets_received +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            packetsReceived, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_packets_sent +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            packetsSent, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_num_alive_connections +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            numAliveConnections, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_outstanding_requests +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            outstandingRequests, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_server_state +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            serverState, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_znode_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            znodeCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_watch_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            watchCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_ephemerals_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            ephemeralsCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_approximate_data_size +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            approximateDataSize, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_open_file_descriptor_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            openFileDescriptorCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_max_file_descriptor_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            maxFileDescriptorCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_followers +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            znodeCount, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_synced_followers +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            syncedFollowers, _ = strconv.Atoi(matches[1])
        }
        re = regexp.MustCompile("zk_pending_syncs +(.*$)")
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