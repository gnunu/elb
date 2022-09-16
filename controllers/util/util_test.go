package util

import (
	"net/netip"
	"testing"
)

func TestQueryGeneral(t *testing.T) {
	hostip := [4]byte{10, 238, 135, 60}
	addr := netip.AddrPortFrom(netip.AddrFrom4(hostip), 9090)
	QueryGeneral(addr)
}
