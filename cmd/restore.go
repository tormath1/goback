package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	pb "github.com/tormath1/goback/server/proto"
)

func init() {
	rootCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{
	Use:   "restore ",
	Short: "restore a volume to his latest version",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		restore(cmd, args...)
	},
}

func restore(cmd *cobra.Command, args ...string) {
	_, err := manager.Restore(context.Background(), &pb.RestoreVolumeRequest{
		VolumeName: args[0],
	})
	if err != nil {
		log.Fatalf("unable to restore volume: %v", err)
	}
}
