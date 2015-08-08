package main

import (
	"os/exec"
	"strings"
	"syscall"
)

type proxy struct {
	server  string
	service string
	process int
	dest    string
	start   string
	notify  string
}

func (p proxy) reload() error {
	s, err := newService(p.server, p.service, p.process)
	if err != nil {
		return err
	}

	// generate nginx config
	err = nginxConf{s}.generate(dest)
	if err != nil {
		return err
	}

	// start new service
	err = s.startNewService(p.start)
	if err != nil {
		return err
	}

	// reload nginx
	commands := strings.Split(p.notify, " ")
	err = exec.Command(commands[0], commands[1:]...).Run()
	if err != nil {
		return err
	}

	// signal to current service
	s.signalToOldProcs(syscall.SIGQUIT)

	return nil
}
