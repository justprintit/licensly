package main

import (
	"github.com/spf13/cobra"
	"os"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dumps effective config",
	Run: func(_ *cobra.Command, _ []string) {
		cfg.WriteTo(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
