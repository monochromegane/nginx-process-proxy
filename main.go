package main

import (
	"flag"
	"fmt"
)

var service string
var serverName string

func init() {
	flag.StringVar(&service, "service", "", "service name")
	flag.StringVar(&serverName, "server-name", "", "server name")
	flag.Parse()
}

func main() {
	proxy, err := newProxy(serverName, service)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	proxy.run()
}
