package main

import "github.com/mitchellh/go-ps"

type Proc struct {
	Process ps.Process
	Port    int
}
