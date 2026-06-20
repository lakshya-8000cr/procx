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

    for _, pid := range pids {          // now we are looping over this inside we are calling the getprocessinfo from the process.go
	info, err := linux.GetProcessInfo(pid)  // this is the map 
	if err != nil {
		continue
	}

	fmt.Printf("%-8s %-20s %-20s %-12s %s\n",
		info.PID,
		info.Name,
		info.State,
		info.Memory,
		info.Threads,
	)   // Now we will print all the info we have asked 
}
	},
}

func init() {

	rootCmd.AddCommand(
		listCmd,
	)
}