package cmd

import (
	"fmt"

	"github.com/gkwa/enoughparse/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello <image_file>",
	Short: "Generate a Google Maps link from an image's GPS data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())

		gpsInfo, err := core.ExtractGPSCoordinates(args[0], logger)
		if err != nil {
			logger.Error(err, "Failed to extract GPS coordinates")
			return
		}

		logger.V(1).Info("Extracted GPS coordinates", "gpsInfo", gpsInfo.String())

		link := core.GenerateGoogleMapsLink(gpsInfo)
		fmt.Println(link)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
