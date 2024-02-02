package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBConf struct {
	Uri string `mapstructure:"uri"`
}

type Mongo struct {
	ctx context.Context
}

func NewMongo(ctx context.Context) *Mongo {
	return &Mongo{ctx: ctx}
}

// func (m *Mongo) Connect(host, port, mongoUser, mongoPass, mongoDb string) (*mongo.Client, error) {
func (m *Mongo) Connect(uri string) (*mongo.Client, error) {
	// url := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", mongoUser, mongoPass, host, port, mongoDb)

	client, err := mongo.Connect(m.ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Error connecting to MongoDB! %v\n", err)
		return nil, err
	}

	if err := client.Ping(m.ctx, readpref.Primary()); err != nil {
		fmt.Printf("Error pinging to MongoDB! %v\n", err)
		return nil, err
	}

	fmt.Printf("Connected to MongoDB!\n")
	return client, nil
}
