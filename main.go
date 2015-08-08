package main

import (
	"flag"
	"fmt"
)

var service string
var serverName string
var process int
var dest string
var start string
var notify string

func init() {
	flag.StringVar(&service, "service", "", "service name")
	flag.StringVar(&serverName, "server-name", "", "server name")
	flag.IntVar(&process, "process", 1, "process number")
	flag.StringVar(&dest, "dest", "", "output file path")
	flag.StringVar(&start, "start", "", "start command")
	flag.StringVar(&notify, "notify", "", "notify command")
	flag.Parse()
}

func main() {

	proxy := proxy{
		server:  serverName,
		service: service,
		process: process,
		dest:    dest,
		start:   start,
		notify:  notify,
	}
	err := proxy.genAndNotify()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
