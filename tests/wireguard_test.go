package tests

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pc0stas/conga/generators"
)

func TestWireGuard_GenerateBasic(t *testing.T) {
	tmp := t.TempDir()
	ifacePath := filepath.Join(tmp, "wg0")
	cfg := generators.WireGuardConfig{
		InterfaceName: ifacePath,
		PrivateKey:    "PRIV",
		Address:       "10.0.0.1/24",
		DNS:           "8.8.8.8",
		ListenPort:    "51820",
		Peers: []generators.WireGuardPeer{{PublicKey: "PUBKEY", AllowedIPs: "10.0.0.2/32"}},
	}

	if err := generators.WriteWireGuardConfig(cfg); err != nil {
		t.Fatalf("WriteWireGuardConfig failed: %v", err)
	}
	out := ifacePath + ".conf"
	b, _ := ioutil.ReadFile(out)
	s := string(b)
	if !strings.Contains(s, "[Interface]") || !strings.Contains(s, "Address = 10.0.0.1/24") {
		t.Fatalf("generated wg config missing interface section")
	}
}
