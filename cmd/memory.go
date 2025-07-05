package cmd

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "used to check the resource memory usage",
	Run: func(cmd *cobra.Command, args []string) {
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			log.Fatalf("Error getting memory info: %v", err)
		}
		fmt.Printf("Memory Usage: %.2f%% (%s / %s)\n", vmStat.UsedPercent, formatGB(vmStat.Used), formatGB(vmStat.Total)) // fmt.Printf("Total Memory: %d, Memory used in percentage: %2.f, Memory used: %d", vmStat.Total, vmStat.UsedPercent, vmStat.Used)
	},
}

func init() {
	log.Println("memory init...")
}
