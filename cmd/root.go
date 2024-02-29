package cmd

import (
	"fmt"
	"os"

	"github.com/crwr45/wlemu/pkg/logging"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	logLevel          string
	humanReadableLogs bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wlemu",
	Short: "Create and Manage emulated workloads",
	Long:  "Create and Manage emulated workloads",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { }
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logging.SetupLogging(logLevel, humanReadableLogs)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&humanReadableLogs,
		"human-readable",
		"H",
		false,
		"Disables JSON-formatted logs",
	)
	rootCmd.PersistentFlags().StringVarP(
		&logLevel,
		"verbosity",
		"v",
		zerolog.WarnLevel.String(),
		fmt.Sprintf(
			"Log level (%s, %s, %s, %s, %s, %s)",
			zerolog.LevelDebugValue,
			zerolog.LevelInfoValue,
			zerolog.LevelWarnValue,
			zerolog.LevelErrorValue,
			zerolog.LevelFatalValue,
			zerolog.LevelPanicValue),
	)
}
