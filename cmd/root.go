package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goback ",
	Short: "schedule your Docker volumes backups with goback",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("unable to execute command: %v", err)
	}
}
