package database

import (
	"goapi/pkg/errors"
	"goapi/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongodbConnection func for connection to mongodb database.
func MongodbConnection(ctx *fiber.Ctx) (*mongo.Client, error) {

	url, err := utils.ConnectionURLBuilder("mongodb")

	if err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":     "Failed connect from provided database",
			"address": url,
		})
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err = client.Connect(ctx.Context()); err != nil {
		return nil, errors.InternalServerError(errors.Message{
			"msg":     "timout error on connect with database",
			"address": url,
		})
	}

	// Try to ping database.
	if err := client.Ping(ctx.Context(), readpref.Primary()); err != nil {
		// close database connection
		defer client.Disconnect(ctx.Context())
		return nil, errors.InternalServerError(errors.Message{
			"msg":   "Couldn't connect to database! Server timeout connection",
			"error": true,
		})
	}

	return client, nil
}
