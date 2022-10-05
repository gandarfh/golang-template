package database

import (
	"context"
	"goapi/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/tryvium-travels/memongo"
	"github.com/tryvium-travels/memongo/memongolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongodbConnection func for connection to mongodb database.
func MemoryMongodbConnection(ctx *fiber.Ctx) (*mongo.Client, error) {
	server, err := memongo.StartWithOptions(&memongo.Options{
		MongoVersion: "4.0.5",
		LogLevel:     memongolog.LogLevelSilent,
	})

	if err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "Failed to start memory mongodb server",
			"error": true,
		})
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(server.URI()))
	if err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "Error when try connect to mongodb",
			"error": true,
		})

	}

	if err = client.Connect(context.Background()); err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "timout when connect to client",
			"error": true,
		})
	}

	// Try to ping database.
	if err := client.Ping(context.Background(), nil); err != nil {
		// close database connection
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "Couldn't connect to client! Server timeout connection",
			"error": true,
		})
	}

	return client, nil
}
