package cmd

import (
	"fmt"
	"strings"
	"time"

	"procx/internal/linux"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var fullEnv bool

var envCmd = &cobra.Command{
	Use:   "env <pid>",
	Short: "Show environment variables of a process",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		pid := args[0]


		//ahh again some pretty print
		cyan := color.New(color.FgCyan, color.Bold)
		white := color.New(color.FgWhite, color.Bold)
		dim := color.New(color.FgWhite, color.Faint)
		green := color.New(color.FgGreen)
		yellow := color.New(color.FgYellow)
		border := color.New(color.FgCyan, color.Faint)

		spinnerFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		for i := 0; i < 12; i++ {
			frame := spinnerFrames[i%len(spinnerFrames)]
			fmt.Printf("\r  %s  %s",
				cyan.Sprint(frame),
				dim.Sprint("reading /proc/"+pid+"/environ ..."),
			)
			time.Sleep(60 * time.Millisecond)
		}
		fmt.Print("\r\033[K")

		envs, err := linux.GetEnvironment(pid)
		if err != nil {
			fmt.Println("failed to read environment:", err)
			return
		}

		bars := []string{"▰▱▱▱▱", "▰▰▱▱▱", "▰▰▰▱▱", "▰▰▰▰▱", "▰▰▰▰▰"}
		for _, f := range bars {
			fmt.Printf("\r  %s  %s",
				cyan.Sprint(f),
				dim.Sprint("parsing env vars ..."),
			)
			time.Sleep(70 * time.Millisecond)
		}
		fmt.Print("\r\033[K")

		fmt.Println()
		border.Println("  ┌─────────────────────────────────────────────┐")
		fmt.Print("  │  ")
		cyan.Print("⬡  PROCX")
		dim.Print("  ·  ")
		white.Print("Process Environment")
		dim.Println("                      │")
		border.Println("  └─────────────────────────────────────────────┘")
		fmt.Println()

		fmt.Print("  ")
		dim.Printf("%-10s", "pid")
		cyan.Println(pid)
		fmt.Println()

		if fullEnv {
			border.Println("  ─────────────────────────────────────────────")
			fmt.Println()
			for _, env := range envs {
				parts := strings.SplitN(env, "=", 2)
				if len(parts) != 2 {
					continue
				}
				fmt.Print("  ")
				yellow.Printf("%-20s", parts[0])
				dim.Print(" = ")
				white.Println(parts[1])
			}
			fmt.Println()
			border.Println("  ─────────────────────────────────────────────")
			fmt.Println()
			return
		}

		important := map[string]bool{ // instead of all the unccessecary info we will show only neccessary info

			// thats why we have removed the for loop over the entries, but if you wanna see full then use flag
			"USER":  true,
			"HOME":  true,
			"SHELL": true,
			"TERM":  true,
		}

		border.Println("  ─────────────────────────────────────────────")
		fmt.Println()

		found := 0
		for _, env := range envs {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) != 2 {
				continue
			}

			key := parts[0]
			value := parts[1]

			if !important[key] {
				continue
			}

			fmt.Print("  ")
			green.Printf("%-10s", key)
			dim.Print(" = ")
			white.Println(value)
			found++
		}

		fmt.Println()
		border.Println("  ─────────────────────────────────────────────")
		fmt.Println()
		fmt.Print("  ")
		dim.Printf("showing %d key variables", found)
		fmt.Print("  ·  ")
		dim.Print("use ")
		cyan.Print("--full")
		dim.Println(" to see all")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(envCmd)

	envCmd.Flags().BoolVar(
		&fullEnv,
		"full",
		false,
		"show all environment variables",
	)
}