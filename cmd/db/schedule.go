/*
Copyright © 2024 Gino Massei <ginomassei@icloud.com>
*/
package db

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:   "schedule-dump",
	Short: "Schedule a dump of a database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("db schedule-dump called")

		// Get the cron expression
		cronExpression, err := cmd.Flags().GetString("cron")
		if err != nil {
			fmt.Println("Error getting cron expression:", err)
		}

		fmt.Println("Cron expression:", cronExpression)

		// Schedule the dump using the cron expression
		// the command to execute by cron is gino db dump
		// run this command with go
		// Generate the crontab entry
		crontabEntry := fmt.Sprintf("%s %s", cronExpression, "gino db dump")

		// Add the crontab entry
		if err := addToCrontab(crontabEntry); err != nil {
			fmt.Println("Error adding crontab entry:", err)
			return
		}

		fmt.Println("Crontab entry added successfully.")
	},
}

func addToCrontab(entry string) error {
	// Get the existing crontab entries
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil && strings.Contains(err.Error(), "no crontab for") {
		output = []byte{}
	} else if err != nil {
		return fmt.Errorf("error getting current crontab: %v", err)
	}

	// Append the new entry to the existing entries
	newCrontab := string(output) + "\n" + entry

	// Create a temporary file to store the new crontab
	tmpFile, err := os.CreateTemp("", "new_crontab_*")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write the new crontab to the temporary file
	if _, err := tmpFile.WriteString(newCrontab); err != nil {
		return fmt.Errorf("error writing to temporary file: %v", err)
	}
	tmpFile.Close()

	// Load the new crontab from the temporary file
	cmd = exec.Command("crontab", tmpFile.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error updating crontab: %v", err)
	}

	return nil
}

func init() {
	DbCmd.AddCommand(scheduleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	scheduleCmd.Flags().StringP("cron", "c", "", "The cron expression to schedule the dump")
}
