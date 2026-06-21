package cmd

import (
	"fmt"

	"procx/internal/linux"

	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port <port>",
	Short: "Inspect which process is using a port",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		port := args[0]

		info, err := linux.InspectPort(port)
		if err != nil {
			fmt.Println("failed to inspect port:", err)
			return
		}

		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println(" PROCX   Port Inspector")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()

		if info == nil {
			fmt.Printf(" %-10s %s\n", "Port", port)
			fmt.Println()
			fmt.Println(" No process found using this port")
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			return
		}

		fmt.Printf(" %-10s %s\n", "Port", info.Port)
		fmt.Printf(" %-10s %s\n", "Process", info.Command)
		fmt.Printf(" %-10s %s\n", "PID", info.PID)
		fmt.Printf(" %-10s %s\n", "User", info.User)

		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(portCmd)
}