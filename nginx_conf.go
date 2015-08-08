package main

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

type nginxConf struct {
	service *Service
}

func (n nginxConf) generate(file string) error {
	out, err := n.out()
	if err != nil {
		return err
	}
	return n.write(out, file)
}

func (n nginxConf) write(content []byte, file string) error {
	ioutil.WriteFile(file, content, 0644)
	return nil
}

func (n nginxConf) out() ([]byte, error) {
	var out bytes.Buffer
	tmpl, err := template.New("nginx").Parse(n.template())
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&out, n.service)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func (n nginxConf) template() string {
	return `
proxy_http_version 1.1;
proxy_buffering off;
proxy_set_header Host $http_host;
proxy_set_header Upgrade $http_upgrade;
proxy_set_header Connection $proxy_connection;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
proxy_set_header X-Forwarded-Proto $proxy_x_forwarded_proto;

server {
	listen 80;
	server_name _;
	return 503;
}

upstream {{.Name}} {
{{range .OldProcs}}
	server localhost:{{.Port}} weight=0;
{{end}}
{{range .NewProcs}}
	server localhost:{{.Port}};
{{end}}
}

server {
	server_name {{.Server}}
	location / {
		proxy_pass http://{{.Name}}
	}
}
`
}
