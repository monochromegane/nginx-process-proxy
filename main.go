package main

import (
	"flag"
	"fmt"
)

var service string
var serverName string
var process int

func init() {
	flag.StringVar(&service, "service", "", "service name")
	flag.StringVar(&serverName, "server-name", "", "server name")
	flag.IntVar(&process, "process", 1, "process number")
	flag.Parse()
}

func main() {
	proxy := proxy{
		server:  serverName,
		service: service,
		process: process,
	}
	err := proxy.reload()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
