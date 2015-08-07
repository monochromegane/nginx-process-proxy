package main

import (
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

type Service struct {
	Server   string
	Name     string
	OldProcs []*Proc
	NewProcs []*Proc
}

func newService(server, name string, process int) (*Service, error) {
	service := &Service{Server: server, Name: name}
	procs, err := service.currentProcs()
	if err != nil {
		return nil, err
	}

	// set current processes
	for _, proc := range procs {
		port, err := lsof{}.portByPid(proc.Pid())
		if err != nil {
			return nil, err
		}
		service.OldProcs = append(service.OldProcs, &Proc{
			Process: proc,
			Port:    port,
		})
	}

	// set new processes
	newPort := service.startPortNumber()
	for i := 0; i < process; i++ {
		service.NewProcs = append(service.NewProcs, &Proc{
			Port: newPort,
		})
		newPort++
	}

	return service, nil
}

func (s Service) startPortNumber() int {
	// TODO use option for start port number, and good implementation...
	if len(s.OldProcs) > 0 && strings.HasPrefix(strconv.Itoa(s.OldProcs[0].Pid()), "8") {
		return 9000
	}
	return 8000
}

func (s Service) currentProcs() ([]ps.Process, error) {
	all, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	var procs []ps.Process
	for _, p := range all {
		if p.Executable() == s.Name {
			procs = append(procs, p)
		}
	}
	return procs, nil
}
