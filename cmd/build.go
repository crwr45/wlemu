package cmd

import (
	"github.com/crwr45/wlemu/pkg/resource"
	"github.com/crwr45/wlemu/pkg/rule"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	_ "github.com/crwr45/wlemu/pkg/updaters/stressng"
)

const (
	rulesFileFlag    = "rules"
	resourceFileFlag = "resource"
	outputDirFlag    = "output_dir"
)

var (
	resourceFile string
	outDir       string
	ruleFile     string
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a workload",
	Long:  "Build a workload from one or more manifests and an accompanying rule file",
	Run: func(cmd *cobra.Command, args []string) {
		rule.LoadRulesFromFile(ruleFile)
		resource.ConvertK8sResourceFile(resourceFile, outDir)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&ruleFile,
		rulesFileFlag,
		"c",
		"wlemu.yaml",
		"Rule file containing rules on which image to replace things with",
	)
	checkFlagMarkErr(rulesFileFlag, buildCmd.MarkFlagFilename(rulesFileFlag))

	buildCmd.Flags().StringVarP(&resourceFile,
		resourceFileFlag,
		"r",
		"",
		"k8s manifest to convert",
	)
	checkFlagMarkErr(resourceFile, buildCmd.MarkFlagFilename(resourceFileFlag))

	buildCmd.Flags().StringVarP(&outDir,
		outputDirFlag,
		"o",
		"",
		"Directory where outputs will be written",
	)
	checkFlagMarkErr(outputDirFlag, buildCmd.MarkFlagRequired(outputDirFlag))
	checkFlagMarkErr(outputDirFlag, buildCmd.MarkFlagDirname(outputDirFlag))
}

func checkFlagMarkErr(flagName string, err error) {
	if err != nil {
		log.Fatal().Err(err).Str("flagName", flagName).Msg("failed to mark flag")
	}
}
