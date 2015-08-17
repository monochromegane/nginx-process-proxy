package main

import (
	"flag"
	"fmt"
)

var service string
var serverName string
var process int
var dest string
var certDir string
var start string
var reload string

func init() {
	flag.StringVar(&serverName, "server-name", "", "server name")
	flag.StringVar(&service, "service", "", "service name")
	flag.StringVar(&start, "start", "", "start service command")
	flag.StringVar(&dest, "dest", "", "nginx config path")
	flag.StringVar(&certDir, "cert-dir", "", "nginx certs directory")
	flag.StringVar(&reload, "reload", "", "nginx reload command")
	flag.IntVar(&process, "process", 1, "process number")
	flag.Parse()
}

func main() {

	proxy := proxy{
		server:  serverName,
		service: service,
		process: process,
		dest:    dest,
		certDir: certDir,
		start:   start,
		notify:  reload,
	}
	err := proxy.reload()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
