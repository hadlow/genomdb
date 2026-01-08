package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/hadlow/genomdb/internal/file"
	"github.com/hadlow/genomdb/internal/server"
)

// startCmd represents the run command
var startCmd = &cobra.Command{
	Use:   "start [FILE]",
	Short: "Start a new node",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		config, err := file.Get(path)
		if err != nil {
			os.Exit(1)
		}

		if err != nil {
			log.Fatalf("Error opening database: %d", err)
		}

		s := server.NewServer(&config)

		err = s.Serve()
		if err != nil {
			log.Fatal(err)
		}

		s.Close()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
