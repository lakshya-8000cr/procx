package cmd

import (
	"fmt"

	"procx/internal/linux"

	"github.com/fatih/color"
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

		cyan := color.New(color.FgCyan, color.Bold)
		white := color.New(color.FgWhite, color.Bold)
		dim := color.New(color.FgWhite, color.Faint)
		green := color.New(color.FgGreen, color.Bold)
		red := color.New(color.FgRed, color.Bold)
		yellow := color.New(color.FgYellow)
		border := color.New(color.FgCyan, color.Faint)

		fmt.Println()
		border.Println("  ┌─────────────────────────────────────────────┐")
		fmt.Print("  │  ")
		cyan.Print("⬡  PROCX")
		dim.Print("  ·  ")
		white.Print("Port Inspector")
		dim.Println("                          │")
		border.Println("  └─────────────────────────────────────────────┘")
		fmt.Println()

		if info == nil {
			fmt.Print("  ")
			dim.Printf("%-10s", "port")
			yellow.Println(":" + port)
			fmt.Println()
			fmt.Print("  ")
			red.Print("✕  ")
			dim.Println("no process found using this port")
			fmt.Println()
			border.Println("  ─────────────────────────────────────────────")
			fmt.Println()
			return
		}

		fmt.Print("  ")
		dim.Printf("%-10s", "port")
		yellow.Println(":" + info.Port)

		fmt.Print("  ")
		dim.Printf("%-10s", "process")
		white.Println(info.Command)

		fmt.Print("  ")
		dim.Printf("%-10s", "pid")
		cyan.Println(info.PID)

		fmt.Print("  ")
		dim.Printf("%-10s", "user")
		green.Println(info.User)

		fmt.Println()
		border.Println("  ─────────────────────────────────────────────")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(portCmd)
}