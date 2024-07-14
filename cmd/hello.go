package cmd

import (
	"github.com/gkwa/enoughparse/core"
	"github.com/spf13/cobra"
)

var outputFormat string

var helloCmd = &cobra.Command{
	Use:   "hello <image_file>",
	Short: "Generate a Google Maps link from an image's GPS data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		err := core.ProcessImage(args[0], outputFormat, logger)
		if err != nil {
			logger.Error(err, "Failed to process image")
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringVar(&outputFormat, "format", "text", "Output format (json or text)")
}
