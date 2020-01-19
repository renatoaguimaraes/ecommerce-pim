package driver

import (
	"context"
	"log"
	"time"

	"github.com/renatoaguimaraes/ecm-pim/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientConn mongo.Client

//Connect creates a mongo client with a pool size 50
func Connect(min, max uint64) {
	clientOptions := options.Client().ApplyURI(util.GetEnv("MONGO_URL", "mongodb://localhost:27017"))
	clientOptions.SetMinPoolSize(min)
	clientOptions.SetMaxPoolSize(max)
	clientOptions.SetMaxConnIdleTime(time.Second * 60)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	clientConn = *client
}

// GetMongoClient return mongo client.
func GetMongoClient() *mongo.Client {
	return &clientConn
}

// Disconnect close mongo connection
func Disconnect() {
	GetMongoClient().Disconnect(context.TODO())
}
