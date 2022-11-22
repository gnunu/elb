package usecase

import (
	"fmt"
	"testing"
)

func TestQueryGeneral(t *testing.T) {
	us := NewUsecaseSet()
	u := NewUsecase("reid", []string{"CPU", "GPU"}, "balanced", []Endpoint{
		{Node: "node1", Pod: "pod1", Addr: "10.10.0.1", Port: "12345"},
		{Node: "node2", Pod: "pod2", Addr: "10.10.0.2", Port: "12345"},
	})
	us.Update(u)
	u = NewUsecase("reid2", []string{"GPU"}, "balanced", []Endpoint{
		{Node: "node3", Pod: "pod3", Addr: "10.10.0.3", Port: "12345"},
		{Node: "node4", Pod: "pod4", Addr: "10.10.0.4", Port: "12345"},
	})
	us.Update(u)
	up, ok := us.LookUp("reid")
	if ok {
		fmt.Printf("usecase reid found : %v\n", up)
	}
	_, ok = us.LookUp("reid3")
	if !ok {
		fmt.Printf("usecase reid3 not found\n")
	}
	us.List()
}
