package tests

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pc0stas/conga/generators"
)

func TestDocker_GenerateBasic(t *testing.T) {
	cfg := generators.DockerComposeConfig{
		Version: "3.8",
		Services: []generators.DockerService{{Name: "app", Image: "alpine:latest"}},
	}
	tmp := t.TempDir()
	out := filepath.Join(tmp, "docker-compose.yml")
	if err := generators.WriteDockerCompose(cfg, out); err != nil {
		t.Fatalf("WriteDockerCompose failed: %v", err)
	}
	b, _ := ioutil.ReadFile(out)
	s := string(b)
	if !strings.Contains(s, "app:") || !strings.Contains(s, "image: alpine") {
		t.Fatalf("generated docker compose missing expected service")
	}
}
