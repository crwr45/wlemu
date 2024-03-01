package cmd

import (
	"github.com/crwr45/wlemu/pkg/resource"
	"github.com/crwr45/wlemu/pkg/rule"
	"github.com/spf13/cobra"

	_ "github.com/crwr45/wlemu/pkg/updaters/stressng"
)

var (
	resourceFile string
	ruleFile     string
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a workload",
	Long:  "Build a workload from one or more manifests and an accompanying rule file",
	Run: func(cmd *cobra.Command, args []string) {
		rule.LoadRulesFromFile(ruleFile)
		resource.ConvertK8sResourceFile(resourceFile)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&ruleFile,
		"rules",
		"c",
		"wlemu.yaml",
		"Rule file containing rules on which image to replace things with",
	)

	buildCmd.Flags().StringVarP(&resourceFile,
		"resource",
		"r",
		"",
		"k8s manifest to convert",
	)
}
