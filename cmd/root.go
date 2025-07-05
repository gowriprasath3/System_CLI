package cmd

import "github.com/spf13/cobra"

var roorCmd = &cobra.Command{
	Use:   "system",
	Short: "used to check system resource",
}

func Execute() {
	cobra.CheckErr(roorCmd.Execute())
}

func init() {
	roorCmd.AddCommand(cpuCmd)
	roorCmd.AddCommand(memoryCmd)
	roorCmd.AddCommand(diskCmd)
}
