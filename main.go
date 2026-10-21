package main

import (
	"fmt"
	"log/slog"
	"runtime"

	"github.com/shirou/gopsutil/v4/mem"
)

func main() {

	os := runtime.GOOS
	arch := runtime.GOARCH

	fmt.Printf("Current environment: %s/%s\n\n", os, arch)

	//Memory
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		slog.Error("could not get memory information", "error", err)
	} else {
		fmt.Printf("usedMemory %.2f, totalMemory %.2f, percentage used: %.2f\n", toGB(memoryStats.Used), toGB(memoryStats.Total), memoryStats.UsedPercent)
	}

}

func toGB(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024 / 1024
}
