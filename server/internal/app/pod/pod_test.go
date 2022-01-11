package pod

import (
	"net"
	"testing"
	"time"
)

func TestListenAndServe(t *testing.T) {
	go ListenAndServe()

	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", "localhost:8080", timeout)
	if err != nil {
		t.Errorf("Server could not be reached on port 8080. Got error %s", err)
	}
}
