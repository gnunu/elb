package util

import (
	"context"
	"fmt"
	"net/netip"
	"os"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func QueryGeneral(addr netip.AddrPort) {
	client, err := api.NewClient(api.Config{
		Address: fmt.Sprintf("http://%s", addr.String()),
	})

	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, "up", time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)
}

func GetCPU(addr netip.AddrPort) float32 {
	return 0.0
}

func GetMem(addr netip.AddrPort) float32 {
	return 0.0
}
