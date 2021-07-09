// Package cmd implements commands
package cmd

import (
	"github.com/kuritka/doge-action/internal/common"
	"github.com/kuritka/doge-action/internal/common/log"

	"os"

	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var (
	// Verbose output
	Verbose bool
)

var logger = log.Log

var rootCmd = &cobra.Command{
	Short: common.Action,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Info().Msgf("%s %s 🛌 🤺", BrightMagenta("DOGE ACTION"), BrightYellow("started"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logger.Info().Msgf("No parameters included")
			_ = cmd.Help()
			os.Exit(0)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		logger.Info().Msgf("Not sure what to do next? %s %s", BrightWhite("see:"), common.HomeURL)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// Execute runs concrete command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
