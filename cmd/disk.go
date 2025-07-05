package cmd

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "used to check the resource disk usage",
	Run: func(cmd *cobra.Command, args []string) {
		diskStats, err := disk.Partitions(false)
		if err != nil {
			log.Fatalf("Error getting disk partitions: %v", err)
		}

		var total uint64
		var used uint64

		fmt.Println("\nDisk Drives:")
		for _, part := range diskStats {
			// On Windows, filter for drives like C:\, D:\ only
			if runtime.GOOS == "windows" && !strings.HasSuffix(part.Device, ":") {
				continue
			}
			usage, err := disk.Usage(part.Mountpoint)
			if err == nil && usage.Total > 0 {
				fmt.Printf("- %s: %.2f%% used (%s / %s)\n", part.Mountpoint, usage.UsedPercent, formatGB(usage.Used), formatGB(usage.Total))
				total += usage.Total
				used += usage.Used
			}
		}

		if total > 0 {
			diskUsage := float64(used) / float64(total) * 100
			fmt.Printf("\nOverall Disk Usage: %.2f%% (%s / %s)\n", diskUsage, formatGB(used), formatGB(total))
		} else {
			fmt.Println("No usable disk partitions found.")
		}
	},
}

func formatGB(b uint64) string {
	return fmt.Sprintf("%.2f GB", float64(b)/1024/1024/1024)
}
