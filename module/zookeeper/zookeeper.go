package zookeeper

import (
	"bytes"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"
	"io"
	"io/ioutil"
	"net"
	"time"
)

func init() {
	helper.Registry.AddModuler("zookeeper", New)
}

// New creates new instance of Moduler
func New() helper.Moduler {
	return &Moduler{}
}

type Moduler struct{}

func (r Moduler) Setup(mo *helper.Module) error {
	return nil
}

// generic socket communication for four-letter commands.  returns a reader object to
// pass to a mapping that'll dissect the output
func RunCommand(command string, connectionString string, timeout time.Duration) (responseReader io.Reader, err error) {
	// this is basically implementing https://github.com/samuel/go-zookeeper/blob/master/zk/flw.go#L262

	conn, err := net.DialTimeout("tcp", connectionString, timeout)
	if err != nil {
		logp.Err("Error connecting to host (%s): %v", connectionString, err)
		return nil, err
	}

	defer conn.Close()
	// timeout attempting to write...
	conn.SetReadDeadline(time.Now().Add(timeout))

	// now, lets write!
	_, err = conn.Write([]byte(command))
	if err != nil {
		logp.Err("Error sending %s command: %v", command, err)
		return nil, err
	}

	// lets read, but only for so long...
	conn.SetReadDeadline(time.Now().Add(timeout))

	// gobble gobble...
	result, err := ioutil.ReadAll(conn)

	neterr, ok := err.(net.Error)
	if ok && neterr.Timeout() {
		err = nil // timeout isn't an error in this case
	}
	if err != nil {
		logp.Err("Error reading  %s command: %v", command, err)
		return nil, err
	}

	return bytes.NewReader(result), nil
}
