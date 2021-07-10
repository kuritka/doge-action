package cmd

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/kuritka/doge-action/internal/common/container"
	"github.com/kuritka/doge-action/internal/common/runner"
	"github.com/kuritka/doge-action/internal/impl/create"

	"github.com/kuritka/doge-action/internal/common/guards"
	"github.com/kuritka/doge-action/internal/common/resolver"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "Create cluster",
	Long:  "With respect to the configuration inside the action, it creates a cluster.  ",

	Run: func(cmd *cobra.Command, args []string) {
		o, err := resolver.NewResolver().Resolve()
		guards.Must(err, "reading environment variables")
		logger.Debug().Msgf("loaded configuration: \n %s", aurora.BrightWhite(o))
		if o.Verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
		c := container.NewDocker(context.Background())
		runner.NewCommonRunner(create.NewCluster(c, o)).MustRun()

	},
}

func init() {
	rootCmd.AddCommand(createCommand)
}
