package tests

import (
	"neb/src/base/client"
	"testing"
)

func TestApp(t *testing.T) {
	app := client.GetLocalTokens()

	for _, ticket := range app {
		t.Logf("Tickets: %v", *ticket)
	}
}
