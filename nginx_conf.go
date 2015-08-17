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
# If we receive X-Forwarded-Proto, pass it through; otherwise, pass along the
# scheme used to connect to this server
map $http_x_forwarded_proto $proxy_x_forwarded_proto {
  default $http_x_forwarded_proto;
  ''      $scheme;
}

# If we receive Upgrade, set Connection to "upgrade"; otherwise, delete any
# Connection header that may have been passed to this server
map $http_upgrade $proxy_connection {
  default upgrade;
  '' close;
}

gzip_types text/plain text/css application/javascript application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript;

log_format vhost '$host $remote_addr - $remote_user [$time_local] '
                 '"$request" $status $body_bytes_sent '
                 '"$http_referer" "$http_user_agent"';

access_log /var/log/nginx/access.log vhost;
error_log  /var/log/nginx/error.log;

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
	server localhost:{{.Port}} down;
{{end}}
{{range .NewProcs}}
	server localhost:{{.Port}};
{{end}}
}

{{ if eq .CertDir "" }}
server {
	server_name {{.Server}};
	location / {
		proxy_pass http://{{.Name}};
	}
}
{{ else }}
server {
	server_name {{.Server}};
	listen 443 ssl

	ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
	ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:AES:CAMELLIA:DES-CBC3-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK:!aECDH:!EDH-DSS-DES-CBC3-SHA:!EDH-RSA-DES-CBC3-SHA:!KRB5-DES-CBC3-SHA;

	ssl_prefer_server_ciphers on;
	ssl_session_timeout 5m;
	ssl_session_cache shared:SSL:50m;

	ssl_certificate {{ .CertDir }}/{{ (printf "%s.crt" .Server) }};
	ssl_certificate_key {{ .CertDir }}/{{ (printf "%s.key" .Server) }};

	add_header Strict-Transport-Security "max-age=31536000";

	location / {
		proxy_pass http://{{.Name}};
	}
}
{{ end }}
`
}
