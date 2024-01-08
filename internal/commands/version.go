package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version = "latest"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show goup version",
		RunE: func(c *cobra.Command, args []string) error {
			_, err := fmt.Printf("goup version %s\n", Version)
			return err
		},
	}
}
