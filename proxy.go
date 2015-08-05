package main

type proxy struct {
	service *Service
}

func newProxy(server, service string) (*proxy, error) {
	s, err := newService(server, service)
	if err != nil {
		return nil, err
	}
	return &proxy{s}, nil
}

func (p proxy) run() error {
	return nil
}
