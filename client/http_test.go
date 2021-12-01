package client

import (
	"testing"
	"time"
)

func httpClientTest(t *testing.T) {
	t.Run("default http client", func(t *testing.T) {
		client := httpClient()
		if client == nil {
			t.Error("got nil http client")
		}

		dur := 10 * time.Second

		if client.Timeout != dur {
			t.Errorf("timeout doesn't match what's expected: %v", client.Timeout)
		}
	})
}
