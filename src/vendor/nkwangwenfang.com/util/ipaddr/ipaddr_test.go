package ipaddr

import (
	"testing"
)

func TestLocalIP(t *testing.T) {
	t.Log(LocalIP())
}
