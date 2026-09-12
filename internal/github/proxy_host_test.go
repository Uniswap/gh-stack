package github

import "testing"

func TestNewClientDefaultsToProxyHost(t *testing.T) {
	t.Setenv("GH_HOST", "")
	t.Setenv("GH_ENTERPRISE_TOKEN", "test-token")

	for _, resolved := range []string{"", "github.com"} {
		c, err := NewClient(resolved, "Uniswap", "universe")
		if err != nil {
			t.Fatalf("NewClient(%q): %v", resolved, err)
		}
		if c.host != ProxyHost {
			t.Errorf("NewClient(%q) host = %q, want %q", resolved, c.host, ProxyHost)
		}
	}
}

func TestExplicitGHHostWins(t *testing.T) {
	// An explicit GH_HOST is always honoured, including GH_HOST=github.com,
	// which is the documented way to opt out of the proxy.
	t.Setenv("GH_HOST", "github.com")
	t.Setenv("GH_TOKEN", "test-token")
	c, err := NewClient("github.com", "Uniswap", "universe")
	if err != nil {
		t.Fatal(err)
	}
	if c.host != "github.com" {
		t.Errorf("host = %q, want github.com", c.host)
	}
}

func TestGHESHostUnaffected(t *testing.T) {
	t.Setenv("GH_HOST", "")
	t.Setenv("GH_ENTERPRISE_TOKEN", "test-token")
	c, err := NewClient("github.mycompany.com", "o", "r")
	if err != nil {
		t.Fatal(err)
	}
	if c.host != "github.mycompany.com" {
		t.Errorf("host = %q, want github.mycompany.com", c.host)
	}
}
