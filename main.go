package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {

	//hostInfo, err := host.Info()
	//if err != nil {
	//	log.Fatal("could not get host information")
	//}

	platform, _, version, err := host.PlatformInformation()
	if err != nil {
		log.Fatal("could not get platform information")
	}

	displayOS := platform
	if platform == "darwin" {
		displayOS = "macOS"
	}

	fmt.Printf("%s %s\n", displayOS, version)

	//Memory
	memoryStats, err := mem.VirtualMemory()

	if err != nil {
		slog.Error("could not get memory information", "error", err)
	} else {
		totalMemory := toGB(memoryStats.Total)
		fmt.Printf(
			"usedMemory %.2f, totalMemory %.2f, percentage used: %.2f\n",
			toGB(memoryStats.Used),
			totalMemory,
			memoryStats.UsedPercent,
		)
	}

}

func toGB(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024 / 1024
}
