package repositories

import (
	"goapi/infrastructure/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// BookRepo struct for queries from Book model.
type BookRepo struct {
	Mongo *mongo.Database
}

func NewBookRepository(ctx *fiber.Ctx) (*BookRepo, error) {
	db, err := database.OpenDBConnection(ctx, "mongodb")

	if err != nil {
		// Return status 500 and database connection error.
		return nil, err
	}

	database := db.Mongo.Database("app")
	return &BookRepo{
		Mongo: database,
	}, nil
}
