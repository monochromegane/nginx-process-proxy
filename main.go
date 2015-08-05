package main

import "flag"

var service string

func init() {
	flag.StringVar(&service, "service", "", "service name")
	flag.Parse()
}

func main() {

}
