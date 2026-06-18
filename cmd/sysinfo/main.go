package main

import (
	"fmt"
	"os"

	"sys-info-app/pkg/sys"
)

func main() {
	// Initializes the correct collector automatically based on OS build target
	collector := sys.NewCollector()

	info, err := collector.GetHardwareInfo()
	if err != nil {
		fmt.Printf("Failed to collect system information: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Manufacturer:  %s\n", info.Manufacturer)
	fmt.Printf("Model:         %s\n", info.Model)
	fmt.Printf("Serial Number: %s\n", info.SerialNumber)
	fmt.Printf("UUID:          %s\n\n", info.UUID)
}
