package cmd

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "used to get the details of cpu...",
	Run: func(cmd *cobra.Command, args []string) {
		cpuPercent, err := cpu.Percent(0, false)
		if err != nil {
			log.Fatalf("Error getting CPU info: %v", err)
		}
		fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
	},
}

func init() {
	log.Println("cpu init...")
}
