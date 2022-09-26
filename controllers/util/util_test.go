package util

import (
	"net/netip"
	"testing"

	"github.com/gnunu/elb/protocol"
)

func TestQueryGeneral(t *testing.T) {
	hostip := [4]byte{10, 238, 135, 60}
	addr := netip.AddrPortFrom(netip.AddrFrom4(hostip), 9090)
	QueryGeneral(addr)
}

func TestSendUsecase(t *testing.T) {
	u := &protocol.Usecase{Name: "reid", Devices: "cpu,gpu", Policy: "balanced", Endpoints: "10.0.0.1:55555,10.0.0.2:55555"}

	SendUsecase(u)
}
