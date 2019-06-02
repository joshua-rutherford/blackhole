package cmd

import (
	"github.com/joshua-rutherford/blackhole/server"
	"github.com/spf13/cobra"
)

// Server returns a command that runs the server.
func Service() *cobra.Command {
	command := &cobra.Command{
		Use:   "server",
		Short: "Run the blackhole server.",
		RunE: func(command *cobra.Command, args []string) error {
			host, _ := command.Flags().GetString("host")
			port, _ := command.Flags().GetInt("port")
			return server.Start(host, port)
		},
	}
	command.Flags().StringP("host", "", "127.0.0.1", "the server host")
	command.Flags().IntP("port", "", 80, "the server port")
	return command
}
