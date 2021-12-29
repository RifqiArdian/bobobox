package config

import (
	"bobobox/exception"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"strconv"
	"time"
)

func NewMongoDatabase(configuration Config) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, _ := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	mongoPoolMax, _ := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	mongoMaxIdleTime, _ := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME"))

	client, err := mongo.NewClient(options.Client().
		ApplyURI(configuration.Get("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second))
	exception.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exception.PanicIfNeeded(err)

	database := client.Database(configuration.Get("MONGO_DATABASE"))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func MongoTransOption() *options.TransactionOptions {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	return options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
}