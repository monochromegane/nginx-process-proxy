package main

import "github.com/mitchellh/go-ps"

type Service struct {
	Server   string
	Name     string
	OldProcs []*Proc
}

func newService(server, name string) (*Service, error) {
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

	return service, nil
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
