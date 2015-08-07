package main

import "github.com/mitchellh/go-ps"

type Proc struct {
	Process ps.Process
	Port    int
}

func (p Proc) Pid() int {
	return p.Process.Pid()
}
