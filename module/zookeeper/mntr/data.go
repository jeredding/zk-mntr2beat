package mntr

import (
    "fmt"
    "time"
    "bufio"

    "strconv"
    "github.com/elastic/beats/libbeat/common"
)

// type ZookeeperMntrStats struct {
//     VersionString            string `json:"version_string"`
//     AvgLatency               uint64 `json:"avg_latency"`
//     MinLatency               uint64 `json:"min_latency"`
//     MaxLatency               uint64 `json:"max_latency"`
//     PacketsReceived          uint64 `json:"packets_received"`
//     PacketsSent              uint64 `json:"packets_sent"`
//     NumAliveConnections      uint64 `json:"num_alive_connections"`
//     OutstandingRequests      unit64 `json:"outstanding_requests"`
//     ServerState              unit64 `json:"server_state"`
//     ZnodeCount               unit64 `json:"znode_count"`
//     WatchCount               unit64 `json:"watch_count"`
//     EphemeralsCount          unit64 `json:"ephemerals_count"`
//     ApproximateDataSize      unit64 `json:"approximate_data_size"`
//     OpenFileDescriptorCount  unit64 `json:"open_file_descriptor_count"`
//     MaxFileDescriptorCount   unit64 `json:"max_file_descriptor_count"`
//     Followers                unit64 `json:"followers,omitempty"`
//     SyncedFollowers          unit64 `json:"synced_followers,omitempty"`
//     PendingSyncs             unit64 `json:"pending_syncs,omitempty"`
//     ctime                    time.Time
// }

func eventMapping(response io.Reader) common.MapStr {

    var (
        versionString           string
        avgLatency              uint64
        minLatency              uint64
        maxLatency              uint64
        packetsReceived         uint64
        packetsSent             uint64
        numAliveConnections     uint64
        outstandingRequests     unit64
        serverState             unit64
        znodeCount              unit64
        watchCount              unit64
        ephemeralsCount         unit64
        approximateDataSize     unit64
        openFileDescriptorCount unit64
        maxFileDescriptorCount  unit64
        followers               unit64
        syncedFollowers         unit64
        pendingSyncs            unit64
    )

    /*
    zk_version      3.5.1-alpha-1693007, built on 07/28/2015 07:19 GMT
    zk_avg_latency  0
    zk_max_latency  1789
    zk_min_latency  0
    zk_packets_received     22152032
    zk_packets_sent 30959914
    zk_num_alive_connections        1033
    zk_outstanding_requests 0
    zk_server_state leader
    zk_znode_count  242609
    zk_watch_count  940522
    zk_ephemerals_count     8565
    zk_approximate_data_size        372143564
    zk_open_file_descriptor_count   1083
    zk_max_file_descriptor_count    1048576
    zk_followers    5
    zk_synced_followers     2
    zk_pending_syncs        0
    */

    var re *regexp.Regexp
    scanner := bufio.NewScanner(response)
    for scanner.Scan() {

        re = regexp.MustCompile("zk_version +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            versionString, _ = matches[1]
        }
        re = regexp.MustCompile("zk_avg_latency +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            avgLatency, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_max_latency +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            maxLatency, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_min_latency +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            minLatency, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_packets_received +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            packetsReceived, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_packets_sent +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            packetsSent, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_num_alive_connections +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            numAliveConnections, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_outstanding_requests +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            outstandingRequests, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_server_state +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            serverState, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_znode_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            znodeCount, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_watch_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            watchCount, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_ephemerals_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            ephemeralsCount, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_approximate_data_size +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            approximateDataSize, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_open_file_descriptor_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            openFileDescriptorCount, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_max_file_descriptor_count +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            maxFileDescriptorCount, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_followers +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            znodeCount, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_synced_followers +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            syncedFollowers, _ = strconv.ParseInt(matches[1], 10, 64)
        }
        re = regexp.MustCompile("zk_pending_syncs +(.*$)")
        if matches := re.FindStringSubmatch(scanner.Text()); matches != nil {
            pendingSyncs, _ = strconv.ParseInt(matches[1], 10, 64)
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