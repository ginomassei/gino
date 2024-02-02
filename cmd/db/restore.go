/*
Copyright © 2024 Gino Massei <ginomassei@icloud.com>
*/
package db

import (
	"fmt"

	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a dump of a database passed by path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("db restore called")
	},
}

func init() {
	DbCmd.AddCommand(restoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	restoreCmd.Flags().StringP("filePath", "f", "", "The file to restore")
}
