package ddos

import (
	"testing"
	"time"
)

func TestDDOS(t *testing.T) {
	t.Run("ddos", func(t *testing.T) {
		d, _ := New("http://ocs-mt-dev.devel.pro/ocs-auth-config-api/v1/subsystems/auth/config", 1000)
		d.Run()
		i := 0
		for {
			time.Sleep(1 * time.Second)
			i++
			t.Logf("iter %d: %d/%d", i, d.successRequest, d.amountRequests)
		}
	})
}
