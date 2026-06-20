package cmd

import (
	"fmt"

	"procx/internal/linux"

	"github.com/spf13/cobra"
)

// now we are making our core diagnse command
var diagnoseCmd = &cobra.Command{
	Use:   "diagnose <pid>",
	Short: "Diagnose a Linux process",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		pid := args[0]

		diagnosis, err := linux.Diagnose(pid)
		if err != nil {
			fmt.Println("failed to diagnose process:", err)
			return
		}

		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println(" PROCX   Process Diagnosis")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()

		fmt.Printf(" %-12s %s\n", "Process", diagnosis.Process.Name)
		fmt.Printf(" %-12s %s\n", "PID", diagnosis.Process.PID)
		fmt.Printf(" %-12s %s\n", "State", diagnosis.Process.State)
		fmt.Printf(" %-12s %s\n", "RAM", diagnosis.Process.Memory)
		fmt.Printf(" %-12s %d\n", "Threads", diagnosis.Threads)
		fmt.Printf(" %-12s %d\n", "FDs", diagnosis.FDs)

		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}