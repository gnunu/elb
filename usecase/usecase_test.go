package usecase

import (
	"fmt"
	"testing"
)

func TestQueryGeneral(t *testing.T) {
	us := NewUsecaseSet()
	u := NewUsecase("reid", []string{"CPU", "GPU"}, "balanced", []string{"10.10.0.1:12345", "10.10.0.2:12345"})
	us.Update(u)
	u = NewUsecase("reid2", []string{"GPU"}, "balanced", []string{"10.10.0.3:12345", "10.10.0.4:12345"})
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
