package main

import (
	"strings"
	"testing"
)

func TestOut_OldProcIsNil(t *testing.T) {

	n := nginxConf{service: &Service{
		Server:   "server",
		Name:     "name",
		OldProcs: nil,
		NewProcs: []*Proc{
			&Proc{Port: 8000},
		},
	}}

	out, err := n.out()
	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}
	expect := "server localhost:8000;"
	if !strings.Contains(string(out), expect) {
		t.Errorf("upstream should contain %s, but not", expect)
	}
}

func TestOut_Cert(t *testing.T) {
	n := nginxConf{service: &Service{
		Server:  "server",
		Name:    "name",
		CertDir: "/etc/nginx",
		OldProcs: []*Proc{
			&Proc{Port: 8000},
		},
		NewProcs: []*Proc{
			&Proc{Port: 9000},
		},
	}}

	out, err := n.out()
	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}

	expect := "listen 443 ssl"
	if !strings.Contains(string(out), expect) {
		t.Errorf("server should contain %s, but not", expect)
	}

	expect = "ssl_certificate /etc/nginx/server.crt;"
	if !strings.Contains(string(out), expect) {
		t.Errorf("server should contain %s, but not", expect)
	}

	expect = "ssl_certificate_key /etc/nginx/server.key;"
	if !strings.Contains(string(out), expect) {
		t.Errorf("server should contain %s, but not", expect)
	}
}

func TestOut(t *testing.T) {

	n := nginxConf{service: &Service{
		Server: "server",
		Name:   "name",
		OldProcs: []*Proc{
			&Proc{Port: 8000},
		},
		NewProcs: []*Proc{
			&Proc{Port: 9000},
		},
	}}

	out, err := n.out()
	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}

	expect := "server localhost:8000 down;"
	if !strings.Contains(string(out), expect) {
		t.Errorf("upstream should contain %s, but not", expect)
	}

	expect = "server localhost:9000;"
	if !strings.Contains(string(out), expect) {
		t.Errorf("upstream should contain %s, but not", expect)
	}
}
