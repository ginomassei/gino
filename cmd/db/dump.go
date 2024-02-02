/*
Copyright © 2024 Gino Massei <ginomassei@icloud.com>
*/
package db

import (
	"fmt"
	"gino/aws"
	"gino/config"
	"gino/database"
	"gino/services"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const backupDir = "/tmp/mongodb-backup"

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump your MongoDB database",
	Long:  "Dump your MongoDB database. And upload the backup to an S3 Bucket",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the upload flag
		upload, err := cmd.Flags().GetBool("upload")
		if err != nil {
			log.Fatal("Error getting upload flag:", err)
		}

		// Load the database configuration
		var dbConfig database.DBConf
		config.Load("database", &dbConfig)

		// Load the dump configuration
		dumpConf := services.DumpConfig{
			DbUri:     dbConfig.Uri,
			BackupDir: backupDir,
		}

		// Instantiate the necessary services
		osService := services.NewOsService()
		dumpService := services.NewDumpService(dumpConf)

		// Create a temporary directory to store the MongoDB dump
		err = osService.CreateTempDir(dumpConf.BackupDir)
		if err != nil {
			log.Fatal("Error creating temporary directory:", err)
		}

		dumpFileName, err := dumpService.Dump()
		if err != nil {
			log.Fatal("Error running mongodump:", err)
		}
		fmt.Println("MongoDB dump created successfully:", dumpFileName)

		if upload {
			var awsConfig aws.AwsCredentials
			config.Load("aws", &awsConfig)

			awsClient := aws.NewAwsClient(awsConfig)

			err := awsClient.UploadFile(aws.UploadData{
				Key:          "backups/",
				FileLocation: dumpConf.BackupDir + "/" + dumpFileName,
				FileName:     dumpFileName,
			})
			if err != nil {
				log.Fatal("Error uploading MongoDB dump to S3:", err)
			}

			fmt.Println("MongoDB dump uploaded successfully to S3:", dumpFileName)
		} else {
			homeDirectory, err := os.UserHomeDir()
			if err != nil {
				log.Fatal("Error getting home directory:", err)
			}

			// Move the MongoDB dump to the home directory
			err = os.Rename(dumpConf.BackupDir+"/"+dumpFileName, homeDirectory+"/"+dumpFileName)
			if err != nil {
				log.Fatal("Error moving MongoDB dump to home directory:", err)
			}
		}

		fmt.Println("Removing temporary directory:", dumpConf.BackupDir)
		// Clean up the temporary directory
		err = osService.CleanUp(dumpConf.BackupDir)
		if err != nil {
			log.Fatal("Error removing temporary directory:", err)
		}
	},
}

func init() {
	DbCmd.AddCommand(dumpCmd)

	dumpCmd.Flags().BoolP("upload", "u", false, "Upload the backup to an S3 Bucket")
}
