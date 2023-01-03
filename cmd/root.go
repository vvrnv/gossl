package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vvrnv/gossl/internal/log"
)

var rootCmd = &cobra.Command{
	Use:   "gossl",
	Short: "gossl is a simple SSL Certificate Checker",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
