package tests

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pc0stas/conga/generators"
)

func TestNginx_GenerateHTTP(t *testing.T) {
	cfg := generators.NginxConfig{
		Domain:   "example.test",
		UseHTTPS: false,
		Routes: []generators.Route{{Path: "/", Type: "static", Directory: "/var/www/html"}},
		CertPath: "/etc/nginx/certs/example.test",
	}

	tmp := t.TempDir()
	out := filepath.Join(tmp, "default.conf")
	if err := generators.WriteNginxConfig(cfg, out); err != nil {
		t.Fatalf("WriteNginxConfig failed: %v", err)
	}
	b, _ := ioutil.ReadFile(out)
	s := string(b)
	if !strings.Contains(s, "server_name example.test") {
		t.Fatalf("expected domain in config")
	}
	if strings.Contains(s, "ssl_certificate") {
		t.Fatalf("did not expect ssl_certificate in HTTP config")
	}
}

func TestNginx_GenerateHTTPS(t *testing.T) {
	cfg := generators.NginxConfig{
		Domain:   "example.test",
		UseHTTPS: true,
		Routes: []generators.Route{{Path: "/api", Type: "proxy", Proxy: "localhost:3000", Headers: true}},
		CertPath: "/etc/nginx/certs/example.test",
	}

	tmp := t.TempDir()
	out := filepath.Join(tmp, "default.conf")
	if err := generators.WriteNginxConfig(cfg, out); err != nil {
		t.Fatalf("WriteNginxConfig failed: %v", err)
	}
	b, _ := ioutil.ReadFile(out)
	s := string(b)
	if !strings.Contains(s, "listen 443 ssl") || !strings.Contains(s, "ssl_certificate /etc/nginx/certs/example.test.crt") {
		t.Fatalf("expected TLS directives in config")
	}
}
