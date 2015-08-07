package main

import "flag"

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
	proxy.reload()
}
