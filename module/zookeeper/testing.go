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

func GetZookeeperEnvPort() string {
	port := os.Getenv("ZOOKEEPER_PORT")

	if len(port) == 0 {
		port = "2181"
	}
	return port
}
