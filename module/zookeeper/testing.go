/*

Helper functions for testing used in the zookeeper metricsets

*/
package zookeeper
import (
    "os"
)

func GetZookeeperEnvHost() string {
    host := os.Getenv("ZOOKEEPER_HOST")

    if len(host) == 0 {
        host = "127.0.0.1"
    }
    return host
}