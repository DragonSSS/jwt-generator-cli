package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var privatePem string

var rootCmd = &cobra.Command{
	Use:   "jwt",
	Short: "root cmd of jwt-generator",
	Long:  "root cmd of jwt-generator",
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("this is a test")
	//},
}

// Execute adds all child commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
