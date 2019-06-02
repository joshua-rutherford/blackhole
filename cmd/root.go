package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute the command line interface.
func Execute() {
	command := &cobra.Command{
		Use:   "blackhole",
		Short: "A service for making blackholed requesets.",
	}
	command.AddCommand(Service())
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
