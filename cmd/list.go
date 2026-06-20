package cmd

import (
	"fmt"

	"procx/internal/linux"

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

fmt.Println()   // you all can see pretty printing is started
fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")  
fmt.Println(" PROCX   Running Processes")
fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

fmt.Printf(
	" %-8s %-20s %-12s %-14s %s\n",
	"PID",
	"NAME",
	"STATE",
	"RAM",
	"THREADS",
)

fmt.Println()

for _, pid := range pids {
	info, err := linux.GetProcessInfo(pid)
	if err != nil {
		continue
	}

	fmt.Printf(
		" %-8s %-20s %-12s %-14s %s\n",
		info.PID,
		info.Name,
		info.State,
		info.Memory,
		info.Threads,
	)
}

fmt.Println()
fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")


	},
}

func init() {

	rootCmd.AddCommand(
		listCmd,
	)
}