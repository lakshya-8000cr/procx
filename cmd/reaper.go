package cmd

import (
	"fmt"
	"strings"

	"procx/internal/linux"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var fixZombies bool

var reaperCmd = &cobra.Command{
	Use:   "reaper",
	Short: "Find zombie processes",
	Run: func(cmd *cobra.Command, args []string) {
		zombies, err := linux.FindZombies()
		if err != nil {
			color.New(color.FgRed).Printf("\n  ✖  failed to scan processes: %v\n\n", err)
			return
		}

		width := 60

		fmt.Println()
		printRBorder("top", width)
		printRHeader(width)
		printRBorder("mid", width)

		if len(zombies) == 0 {
			color.New(color.Faint).Print("  │  ")
			color.New(color.FgGreen).Print("✔  ")
			color.New(color.FgWhite).Print("No zombie processes found.")
			color.New(color.Faint).Printf("%-*s│\n", 28, "")
			printRBorder("bot", width)
			fmt.Println()
			return
		}

		// column headers
		color.New(color.Faint).Print("  │  ")
		color.New(color.Bold, color.Faint).Printf("  %-8s %-8s %-20s %-8s", "PID", "PPID", "PROCESS", "STATE")
		color.New(color.Faint).Printf("  │\n")

		printRBorder("mid", width)

		// rows
		for _, z := range zombies {
			color.New(color.Faint).Print("  │  ")
			color.New(color.FgRed).Print("☠  ")
			color.New(color.FgWhite).Printf("%-8s ", z.PID)
			color.New(color.Faint).Printf("%-8s ", z.PPID)
			color.New(color.FgYellow).Printf("%-20s ", z.Name)
			color.New(color.FgRed, color.Bold).Printf("%-8s", z.State)
			color.New(color.Faint).Print("  │\n")
		}

		printRBorder("mid", width)

		// summary
		color.New(color.Faint).Print("  │  ")
		color.New(color.FgRed).Printf("  ● %d zombie process(es) found", len(zombies))
		color.New(color.Faint).Printf("%-*s│\n", 26, "")

		printRBorder("bot", width)

		// fix mode
		if fixZombies {
			fmt.Println()
			color.New(color.Faint).Print("  ")
			color.New(color.BgYellow, color.FgBlack, color.Bold).Print(" CLEANUP ")
			color.New(color.FgWhite).Println("  Sending SIGTERM to parent process(es)...")
			fmt.Println()

			for _, z := range zombies {
				err := linux.KillParent(z.PPID)
				if err != nil {
					color.New(color.Faint).Print("  ")
					color.New(color.FgRed).Printf("  ✖  parent PID %-6s  %v\n", z.PPID, err)
					continue
				}
				color.New(color.Faint).Print("  ")
				color.New(color.FgGreen).Printf("  ✔  SIGTERM → parent PID %-6s  (zombie PID %s)\n", z.PPID, z.PID)
			}
		}

		fmt.Println()
	},
}

func init() {
	reaperCmd.Flags().BoolVar(
		&fixZombies,
		"fix",
		false,
		"terminate parent processes of zombies",
	)
	rootCmd.AddCommand(reaperCmd)
}

// ── border helpers ───────────────────────────────────────────

func printRBorder(kind string, w int) {
	dim := color.New(color.Faint)
	switch kind {
	case "top":
		dim.Printf("  ╭%s╮\n", strings.Repeat("─", w))
	case "mid":
		dim.Printf("  ├%s┤\n", strings.Repeat("─", w))
	case "bot":
		dim.Printf("  ╰%s╯\n", strings.Repeat("─", w))
	}
}

func printRHeader(w int) {
	_ = w
	color.New(color.Faint).Print("  │  ")
	color.New(color.BgRed, color.FgWhite, color.Bold).Print(" PROCXY ")
	color.New(color.Faint).Print("  ")
	color.New(color.Bold, color.FgWhite).Print("Zombie Inspector")
	color.New(color.Faint).Printf("%-*s│\n", 33, "")
}