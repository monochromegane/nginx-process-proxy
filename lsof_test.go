package main

import "testing"

func TestPortByOutput(t *testing.T) {
	lsof := lsof{}

	output := `
COMMAND   PID    USER   FD     TYPE             DEVICE  SIZE/OFF     NODE NAME
program   95873 user    3u    IPv6 0x137fb155b78f4de3       0t0      TCP *:8000 (LISTEN)
program   95873 user    4u  KQUEUE                                       count=0, state=0x2
`
	port, err := lsof.portByOutput([]byte(output))
	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}
	if port != 8000 {
		t.Errorf("port should be 8000, but %v", port)
	}
}
