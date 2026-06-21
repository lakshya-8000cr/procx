package cmd

import (
	"fmt"
	"strings"

	"procx/internal/linux"

	"github.com/spf13/cobra"
)

var fullEnv bool

var envCmd = &cobra.Command{
	Use:   "env <pid>",
	Short: "Show environment variables of a process",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		pid := args[0]

		envs, err := linux.GetEnvironment(pid)
		if err != nil {
			fmt.Println("failed to read environment:", err)
			return
		}

		fmt.Println()
		fmt.Println("PROCX   Process Environment")
		fmt.Println("PID     ", pid)
		fmt.Println()

		if fullEnv {
			for _, env := range envs {
				fmt.Println(env)
			}
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

			fmt.Printf(" %-10s %s\n", key, value)
		}

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