package cmd

import (
	"github.com/crwr45/wlemu/pkg/config"
	"github.com/crwr45/wlemu/pkg/resource"
	"github.com/spf13/cobra"

	_ "github.com/crwr45/wlemu/pkg/replace/stressng"
)

var (
	resourceFile string
	configFile   string
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a workload",
	Long:  "Build a workload from one or more manifests and an accompanying config file",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfigFromFile(configFile)
		resource.LoadK8sResourceFile(resourceFile)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&configFile,
		"config",
		"c",
		"wlemu.yaml",
		"Config file containing rules on which image to replace things with",
	)

	buildCmd.Flags().StringVarP(&resourceFile,
		"resource",
		"r",
		"",
		"k8s manifest to convert",
	)
}
