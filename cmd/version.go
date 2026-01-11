package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "GenomDB version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("GenomDB %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
