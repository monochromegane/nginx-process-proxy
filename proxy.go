package main

type proxy struct {
	server  string
	service string
	process int
}

func (p proxy) reload() error {
	// generate nginx config
	s, err := newService(p.server, p.service, p.process)
	if err != nil {
		return err
	}
	conf := nginxConf{s}
	err = conf.generate("default.conf")
	if err != nil {
		return err
	}

	// reload nginx
	return nil
}

func (p proxy) run() error {
	return nil
}
