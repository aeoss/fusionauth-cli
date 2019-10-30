package cmd

import (
	"github.com/spf13/cobra"
)

var exportCmdDir string

var exportCmdPath string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports a config from the current FusionAuth instance.",
	Long: `This command will export the running config into the specified
current directory so that the config may be imported into another FusionAuth
instance or be managed via file for a more GitOps-friendly workflow.

# Export to the current directory
$ factl export
Done`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		batch.ExportAll()
	},
}

func init() {
	// --path, -p to specify the sub-directory that the config should be dumped into.
	exportCmd.Flags().StringVarP(&exportCmdPath, "path", "p", "config", "The sub-directory that the exported config should be dumped into.")

	rootCmd.AddCommand(exportCmd)
}
