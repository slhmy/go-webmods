package tests

import (
	"os"
	"testing"

	gwm_app "github.com/slhmy/go-webmods/app"
)

func TestConfigSchema(t *testing.T) {
	_ = os.Setenv("GWM_OVERRIDE_CONFIG_NAME", "override")
	_ = os.Setenv("SERVER_PORT", "443")
	port := gwm_app.Config().GetInt("server.port")
	if port != 443 {
		t.Errorf("Expected server.port to be `443`, got %d", port)
	}
	host := gwm_app.Config().GetString("server.host")
	if host != "docker.host.internal" {
		t.Errorf("Expected server.host to be `docker.host.internal`, got %s", host)
	}
}
