package main

import (
	"go.sancus.dev/config"
	"go.sancus.dev/config/flags"
	"go.sancus.dev/config/flags/cobra"
)

// Command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves licensly web application",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		flags.GetMapper(cmd.Flags()).Parse()

		_, err := config.Validate(cfg)
		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// router
		// serve
		return cfg.Server.ListenAndServe(nil)
	},
}

// Flags
func init() {
	cobra.NewMapper(serveCmd.Flags()).
		VarP(&cfg.Server.Port, "port", 'p', "HTTP port").
		Var(&cfg.Server.PIDFile, "pid", "path to PID file").
		VarP(&cfg.Server.GracefulTimeout, "graceful", 't', "Maximum duration to wait for in-flight requests")

	rootCmd.AddCommand(serveCmd)
}
