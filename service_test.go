package main

import "testing"

func TestStartPortNumber(t *testing.T) {

	type assert struct {
		currentProcs []*Proc
		expect       int
	}

	asserts := []assert{
		assert{
			currentProcs: nil,
			expect:       8000,
		},
		assert{
			currentProcs: []*Proc{&Proc{Port: 8000}},
			expect:       9000,
		},
		assert{
			currentProcs: []*Proc{&Proc{Port: 9000}},
			expect:       8000,
		},
	}

	for _, a := range asserts {
		port := Service{OldProcs: a.currentProcs}.startPortNumber()
		if port != a.expect {
			t.Errorf("port should be %d, but %d", a.expect, port)
		}
	}

}
