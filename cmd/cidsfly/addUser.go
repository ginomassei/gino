/*
Copyright © 2024 Gino Massei <ginomassei@icloud.com>
*/
package cidsfly

import (
	"context"
	"fmt"
	"gino/config"
	"gino/database"
	"gino/models"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// addUserCmd represents the addUser command
var addUserCmd = &cobra.Command{
	Use:   "add-user",
	Short: "Add a new user to the CidsFly database",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		var dbConfig database.DBConf
		config.Load("database", &dbConfig)

		// Connect to the MongoDB database
		mongo := database.NewMongo(context.Background())
		mongoManager, err := mongo.Connect(dbConfig.Uri)
		if err != nil {
			fmt.Printf("Error connecting to the database %v", err)
			panic(err)
		}

		createUser(name, mongoManager)
	},
}

func init() {
	addUserCmd.Flags().StringP("name", "n", "", "full name of the user like Gino Massei")

	if err := addUserCmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}

	CidsflyCmd.AddCommand(addUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createUser(name string, mongoManager *mongo.Client) {
	// Create an user with employee on the mongo database
	// Create an user
	rolIdFromHex, _ := primitive.ObjectIDFromHex("653a549283299a58b532f0a5")

	splitName := strings.Split(name, " ")
	firstName := splitName[0]
	lastName := splitName[1]

	username := strings.Replace(strings.ToLower(name), " ", ".", -1)

	user := models.User{
		Username: username,
		Password: "9dcd687d19fd260c844a14554115f2848ebb0be3b66b",
		RolId:    rolIdFromHex,
		Changed:  false,
	}

	insertResult, err := mongoManager.Database("cids").Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("Error inserting user %v", err)
		panic(err)
	}

	fmt.Printf("Inserted user with ID %v\n", insertResult.InsertedID)

	// Get first and last name from the user
	areaIdFromHex, _ := primitive.ObjectIDFromHex("653a549283299a58b532f0af")
	seniorityIdFromHex, _ := primitive.ObjectIDFromHex("653a549283299a58b532f0b5")
	// Create an employee
	employee := models.Employee{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       "",
		BirthDate:   time.Now(),
		Cellphone:   "00000000000",
		JoinDate:    time.Now(),
		GitlabId:    0,
		DiscordUser: "",
		AreaId:      areaIdFromHex,
		SeniorityId: seniorityIdFromHex,
		UserId:      insertResult.InsertedID.(primitive.ObjectID),
	}

	insertResult, err = mongoManager.Database("cids").Collection("employees").InsertOne(context.Background(), employee)
	if err != nil {
		fmt.Printf("Error inserting employee %v", err)
		panic(err)
	}

	fmt.Printf("Inserted employee with ID %v\n", insertResult.InsertedID)
}
