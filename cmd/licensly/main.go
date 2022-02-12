package main

import (
	"log"

	"github.com/spf13/cobra"
)

const (
	CmdName           = "licensly"
	DefaultConfigFile = CmdName + ".yaml"
)

var (
	cfg          Config
	cfgFile      string
	cfgReadError error
)

var rootCmd = &cobra.Command{
	Use:   CmdName,
	Short: CmdName + " helps creators manage their licensees",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// root level flags
	pflags := rootCmd.PersistentFlags()
	pflags.StringVarP(&cfgFile, "config-file", "f", DefaultConfigFile, "config file (YAML format)")

	// load config-file before cobra commands
	cobra.OnInitialize(func() {

		if cfgFile != "" {
			var c Config

			if err := c.ReadInFile(cfgFile); err == nil {
				// good config ready
				cfg = c
				return
			} else {
				// report bad config file and move on
				cfgReadError = err
				log.Println(err)
			}
		}

		// didn't load, try defaults
		if err := cfg.Prepare(); err != nil {
			// bad defaults
			log.Fatal(err)
		}
	})
}
