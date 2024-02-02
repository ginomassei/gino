/*
Copyright © 2024 Gino Massei <ginomassei@icloud.com>
*/
package db

import (
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var DbCmd = &cobra.Command{
	Use:   "db",
	Short: "db is a pallet of commands to manage your database",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
