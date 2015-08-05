package main

type Service struct {
	Server string
	Name   string
}

func newService(server, name string) (*Service, error) {
	service := &Service{Server: server, Name: name}
	return service, nil
}
