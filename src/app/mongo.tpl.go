package app

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetupMongoDB driver
func SetupMongoDB() {
	conf := Config()
	if client, err := mongo.NewClient(options.Client().ApplyURI(conf.String("mongo.conStr", ""))); err != nil {
		panic(err)
	} else {
		ctx, cancel := MongoOperationCtx()
		defer cancel()
		if err := client.Connect(ctx); err != nil {
			panic(err)
		} else {
			db := client.Database(conf.String("database.name", ""))
			_container.Register("--APP-MONGO-CLIENT", client)
			_container.Register("--APP-MONGO-DB", db)
		}
	}
}

// MongoClient get mongo client
// leave name empty to resolve default
func MongoClient(names ...string) *mongo.Client {
	name := "--APP-MONGO-CLIENT"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(*mongo.Client); ok {
			return res
		}
	}
	return nil
}

// MongoDB get mongo database
// leave name empty to resolve default
func MongoDB(names ...string) *mongo.Database {
	name := "--APP-MONGO-DB"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(*mongo.Database); ok {
			return res
		}
	}
	return nil
}

// MongoOperationCtx create context for mongo db operations
func MongoOperationCtx() (context.Context, context.CancelFunc) {
	ttl := Config().Int("database.ttl", 10)
	return context.WithTimeout(context.Background(), time.Duration(ttl)*time.Second)
}
