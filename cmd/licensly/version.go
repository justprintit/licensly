package main

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/justprintit/licensly"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version prints " + CmdName + "'s version",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(CmdName, licensly.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
