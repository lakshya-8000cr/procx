package cmd

import (
	"fmt"

	"procx/internal/linux"

	"github.com/spf13/cobra"
)

var fixZombies bool

var reaperCmd = &cobra.Command{
	Use:   "reaper",
	Short: "Find zombie processes",
	Run: func(cmd *cobra.Command, args []string) {
		zombies, err := linux.FindZombies()
		if err != nil {
			fmt.Println("failed to scan processes:", err)
			return
		}

		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println(" PROCX   Zombie Inspector")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()

		if len(zombies) == 0 {
			fmt.Println(" No zombie processes found.")
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			return
		}

		fmt.Printf(" %-8s %-8s %-20s %s\n", "PID", "PPID", "PROCESS", "STATE")
		fmt.Println()

		for _, z := range zombies {
			fmt.Printf(
				" %-8s %-8s %-20s %s\n",
				z.PID,
				z.PPID,
				z.Name,
				z.State,
			)
		}

		fmt.Println()
		fmt.Printf(" Found %d zombie process(es).\n", len(zombies))
		fmt.Println()
		if fixZombies {
	fmt.Println()
	fmt.Println(" Attempting cleanup by sending SIGTERM to parent process(es)...")
	fmt.Println()

	for _, z := range zombies {
		err := linux.KillParent(z.PPID)
		if err != nil {
			fmt.Printf(" Failed to terminate parent PID %s: %v\n", z.PPID, err)
			continue
		}

		fmt.Printf(" Sent SIGTERM to parent PID %s for zombie PID %s\n", z.PPID, z.PID)
	}
}
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
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