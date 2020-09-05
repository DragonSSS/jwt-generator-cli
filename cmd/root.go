package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jwt",
	Short: "root cmd of jwt-generator",
	Long:  "root cmd of jwt-generator",
}

// Execute adds all child commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
