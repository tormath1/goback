package cmd

import (
	"context"
	"log"

	pb "github.com/tormath1/goback/server/proto"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(saveCmd)
}

var saveCmd = &cobra.Command{
	Use:   "save ",
	Short: "backup a volume in a location",
	Run: func(cmd *cobra.Command, args []string) {
		saveVolume(manager, args...)
	},
}

func saveVolume(manager pb.ManagerClient, args ...string) {
	_, err := manager.SaveVolume(context.Background(), &pb.SaveVolumeRequest{
		VolumeName:  args[0],
		Destination: args[1],
	})
	if err != nil {
		log.Fatalf("unable to save volume: %v", err)
	}
}
