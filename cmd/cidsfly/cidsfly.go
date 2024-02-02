/*
Copyright © 2024 Gino Massei <ginomassei@icloud.com>
*/
package cidsfly

import (
	"github.com/spf13/cobra"
)

// cidsflyCmd represents the cidsfly command
var CidsflyCmd = &cobra.Command{
	Use:   "cidsfly",
	Short: "pallet of commands to manage your cidsfly instance",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cidsflyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cidsflyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
