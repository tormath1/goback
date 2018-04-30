package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	pb "github.com/tormath1/goback/server/proto"
)

func init() {
	scheduleCmd.AddCommand(scheduleListCmd)
	rootCmd.AddCommand(scheduleCmd)
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule ",
	Short: "schedule your Docker volumes backups",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		scheduleVolume(manager, args...)
	},
}

var scheduleListCmd = &cobra.Command{
	Use:   "list ",
	Short: "list your scheduled Docker volumes backups",
	Run: func(cmd *cobra.Command, args []string) {
		entriesList(manager)
	},
}

func scheduleVolume(manager pb.ManagerClient, args ...string) {
	_, err := manager.ScheduleSaving(context.Background(), &pb.ScheduleSavingRequest{
		Schedule: args[2],
		Volume: &pb.SaveVolumeRequest{
			VolumeName:  args[0],
			Destination: args[1],
		},
	})
	if err != nil {
		log.Fatalf("unable to schedule volume: %v", err)
	}
}

func entriesList(manager pb.ManagerClient) {
	entries, err := manager.ListEntries(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("unable to get cron entries: %v", err)
	}
	for _, entry := range entries.Entries {
		fmt.Println(entry)
	}
}
