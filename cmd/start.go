package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/hadlow/genomdb/internal/file"
	"github.com/hadlow/genomdb/internal/server"
)

var startCmd = &cobra.Command{
	Use:   "start [FILE]",
	Short: "Start a new node",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		config, err := file.Get(path)
		if err != nil {
			log.Fatalf("Error opening config file: %d", err)
		}

		s := server.NewServer(&config)

		err = s.Serve()
		if err != nil {
			log.Fatalf("Error starting server: %d", err)
		}

		s.Close()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
