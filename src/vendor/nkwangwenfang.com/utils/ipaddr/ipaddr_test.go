package ipaddr

import (
	"testing"
)

func TestLocalIP(t *testing.T) {
	t.Log(LocalIPv4())
	t.Log(LocalIPv4Bytes())
}
