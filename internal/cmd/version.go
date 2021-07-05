package cmd

import (
	"github.com/kuritka/doge-action/internal/common"

	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "version",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Info().Msgf("%s %s", BrightWhite("version:"), BrightWhite(common.Version))
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
