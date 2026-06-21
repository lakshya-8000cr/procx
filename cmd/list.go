package cmd

import (
	"fmt"
	"os"
	"time"

	"procx/internal/linux"

	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{

	Use: "list",

	Run: func(
		cmd *cobra.Command,
		args []string,
	) {

		pids, err := linux.GetPIDs() // getting all the pids from the process.go file from internals

		if err != nil {
			fmt.Println(err)
			return
		}

		 // just adding some fancy for terminal uk
		cyan := color.New(color.FgCyan, color.Bold)
		green := color.New(color.FgGreen, color.Bold)
		yellow := color.New(color.FgYellow)
		white := color.New(color.FgWhite, color.Bold)
		dim := color.New(color.FgWhite, color.Faint)
		red := color.New(color.FgRed, color.Bold)
		border := color.New(color.FgCyan, color.Faint)

		spinnerFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

		fmt.Println()

		for i := 0; i < 18; i++ {
			frame := spinnerFrames[i%len(spinnerFrames)]
			fmt.Printf("\r  %s  %s",
				color.New(color.FgCyan, color.Bold).Sprint(frame),
				color.New(color.FgWhite, color.Faint).Sprint("scanning /proc ..."),
			)
			time.Sleep(60 * time.Millisecond)
		}

		fmt.Print("\r\033[K")

		headerStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00d7ff")).
			Background(lipgloss.Color("#0a0e14")).
			PaddingLeft(2).
			PaddingRight(2)

		boxStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#1e3a5f")).
			Padding(0, 1)

		statStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#3fb950")).
			Bold(true)

		_ = headerStyle
		_ = boxStyle
		_ = statStyle

		fmt.Println()
		border.Println("  ┌─────────────────────────────────────────────────────────────────────┐")
		fmt.Print("  │  ")
		cyan.Print("⬡  PROCX")
		dim.Print("  ·  ")
		white.Print("Running Processes")
		dim.Print("  ·  ")
		color.New(color.FgGreen, color.Faint).Printf("found %d processes", len(pids))
		dim.Println("                        │")
		border.Println("  └─────────────────────────────────────────────────────────────────────┘")

		fmt.Println()

		dim.Printf(
			"  %-8s %-20s %-20s %-14s %s\n",
			"PID",
			"NAME",
			"STATE",
			"RAM",
			"THREADS",
		)
		dim.Println("  ─────────────────────────────────────────────────────────────────────")
		fmt.Println()

		var runningCount, sleepingCount, zombieCount int

		for i, pid := range pids {
			info, err := linux.GetProcessInfo(pid)
			if err != nil {
				continue
			}

			if i%3 == 0 {
				time.Sleep(8 * time.Millisecond)
			}

			pidStr := color.New(color.FgCyan, color.Faint).Sprintf("%-8s", info.PID)
			nameStr := white.Sprintf("%-20s", info.Name)

			var stateStr string
			switch info.State {
			case "R (running)":
				stateStr = green.Sprintf("%-20s", "● "+info.State)
				runningCount++
			case "S (sleeping)":
				stateStr = color.New(color.FgBlue, color.Faint).Sprintf("%-20s", "○ "+info.State)
				sleepingCount++
			case "Z (zombie)":
				stateStr = red.Sprintf("%-20s", "✕ "+info.State)
				zombieCount++
			default:
				stateStr = yellow.Sprintf("%-20s", "◌ "+info.State)
			}

			memStr := yellow.Sprintf("%-14s", info.Memory)
			thrStr := dim.Sprintf("%s", info.Threads)

			fmt.Printf("  %s %s %s %s %s\n",
				pidStr,
				nameStr,
				stateStr,
				memStr,
				thrStr,
			)
		}

		fmt.Println()
		border.Println("  ─────────────────────────────────────────────────────────────────────")
		fmt.Println()

		fmt.Print("  ")
		green.Print("● ")
		dim.Printf("%-4d running", runningCount)
		fmt.Print("   ")
		color.New(color.FgBlue, color.Faint).Print("○ ")
		dim.Printf("%-4d sleeping", sleepingCount)

		if zombieCount > 0 {
			fmt.Print("   ")
			red.Print("✕ ")
			red.Printf("%d zombie", zombieCount)
		}

		fmt.Print("   ")
		dim.Printf("total: %d", len(pids))
		fmt.Println()
		fmt.Println()

		if zombieCount > 0 {
			fmt.Print("  ")
			red.Printf("⚠  %d zombie process detected — run ", zombieCount)
			cyan.Print("procx reaper")
			red.Println(" to clean up")
			fmt.Println()
		}

		dim.Println("  scanning complete")

		frames2 := []string{"▰▱▱▱▱", "▰▰▱▱▱", "▰▰▰▱▱", "▰▰▰▰▱", "▰▰▰▰▰"}
		for _, f := range frames2 {
			fmt.Printf("\r  %s  %s",
				color.New(color.FgCyan).Sprint(f),
				dim.Sprint("done"),
			)
			time.Sleep(80 * time.Millisecond)
		}
		fmt.Println()
		fmt.Println()

		os.Exit(0)

	},
}

func init() {

	rootCmd.AddCommand(
		listCmd,
	)
}