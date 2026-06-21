package cmd

import (
	"fmt"
	"time"

	"procx/internal/linux"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// now we are making our core diagnse command
var diagnoseCmd = &cobra.Command{
	Use:   "diagnose <pid>",
	Short: "Diagnose a Linux process",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		pid := args[0]

		cyan := color.New(color.FgCyan, color.Bold)
		white := color.New(color.FgWhite, color.Bold)
		dim := color.New(color.FgWhite, color.Faint)
		green := color.New(color.FgGreen, color.Bold)
		yellow := color.New(color.FgYellow)
		red := color.New(color.FgRed, color.Bold)
		border := color.New(color.FgCyan, color.Faint)
		magenta := color.New(color.FgMagenta)

		spinnerFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		steps := []string{
			"reading /proc/" + pid + "/status ...",
			"scanning file descriptors ...",
			"checking memory usage ...",
			"resolving executable path ...",
			"analysing threads ...",
			"building diagnosis ...",
		}

		for i := 0; i < 18; i++ {
			frame := spinnerFrames[i%len(spinnerFrames)]
			step := steps[i*len(steps)/18]
			fmt.Printf("\r  %s  %s",
				cyan.Sprint(frame),
				dim.Sprint(step),
			)
			time.Sleep(60 * time.Millisecond)
		}
		fmt.Print("\r\033[K")

		bars := []string{"▰▱▱▱▱▱▱", "▰▰▱▱▱▱▱", "▰▰▰▱▱▱▱", "▰▰▰▰▱▱▱", "▰▰▰▰▰▱▱", "▰▰▰▰▰▰▱", "▰▰▰▰▰▰▰"}
		for _, f := range bars {
			fmt.Printf("\r  %s  %s",
				cyan.Sprint(f),
				dim.Sprint("finalizing ..."),
			)
			time.Sleep(60 * time.Millisecond)
		}
		fmt.Print("\r\033[K")

		diagnosis, err := linux.Diagnose(pid)
		if err != nil {
			fmt.Println("failed to diagnose process:", err)
			return
		}

		fmt.Println()
		border.Println("  ┌─────────────────────────────────────────────┐")
		fmt.Print("  │  ")
		cyan.Print("⬡  PROCX")
		dim.Print("  ·  ")
		white.Print("Process Diagnosis")
		dim.Println("                        │")
		border.Println("  └─────────────────────────────────────────────┘")
		fmt.Println()

		// ── process info ──
		border.Println("  ─────────────────────────────────────────────")
		fmt.Println()

		fmt.Print("  ")
		dim.Printf("%-14s", "process")
		white.Println(diagnosis.Process.Name)

		fmt.Print("  ")
		dim.Printf("%-14s", "pid")
		cyan.Println(diagnosis.Process.PID)

		fmt.Print("  ")
		dim.Printf("%-14s", "uptime")
		green.Println(diagnosis.Uptime)

		fmt.Print("  ")
		dim.Printf("%-14s", "state")
		switch diagnosis.Process.State {
		case "R (running)":
			green.Println("● " + diagnosis.Process.State)
		case "Z (zombie)":
			red.Println("✕ " + diagnosis.Process.State)
		default:
			color.New(color.FgBlue, color.Faint).Println("○ " + diagnosis.Process.State)
		}

		fmt.Print("  ")
		dim.Printf("%-14s", "ram")
		yellow.Println(diagnosis.Process.Memory)

		fmt.Print("  ")
		dim.Printf("%-14s", "threads")
		cyan.Println(diagnosis.Threads)

		fmt.Print("  ")
		dim.Printf("%-14s", "fds")
		if diagnosis.FDs > 800 {
			red.Printf("%d", diagnosis.FDs)
			dim.Println("  ⚠ high")
		} else if diagnosis.FDs > 500 {
			yellow.Printf("%d", diagnosis.FDs)
			dim.Println("  · monitor")
		} else {
			green.Println(diagnosis.FDs)
		}

		fmt.Print("  ")
		dim.Printf("%-14s", "executable")
		magenta.Println(diagnosis.Executable)

		fmt.Print("  ")
		dim.Printf("%-14s", "directory")
		magenta.Println(diagnosis.WorkingDir)

		fmt.Print("  ")
		dim.Printf("%-14s", "command")
		white.Println(diagnosis.CommandLine)

		fmt.Println()
		border.Println("  ─────────────────────────────────────────────")

		if len(
			diagnosis.Warnings,
		) > 0 {

			fmt.Println()

			fmt.Print("  ")
			red.Println("WARNINGS")

			fmt.Println()

			for _, w := range diagnosis.Warnings {

				fmt.Print("  ")
				red.Print("⚠  ")
				yellow.Println(w)
			}

			fmt.Println()
			border.Println("  ─────────────────────────────────────────────")
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}