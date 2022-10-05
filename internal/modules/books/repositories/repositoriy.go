package repositories

import (
	"goapi/infrastructure/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// BookRepositoryImpl struct for queries from Book model.
type BookRepositoryImpl struct {
	BooksCollection *mongo.Collection
}

func NewBookRepository(ctx *fiber.Ctx) (*BookRepositoryImpl, error) {
	DB_TYPE := os.Getenv("DB_TYPE")
	db, err := database.OpenDBConnection(ctx, DB_TYPE)
	if err != nil {
		// Return status 500 and database connection error.
		return nil, err
	}

	database := db.Mongo.Database("app")
	booksCollection := database.Collection("books")

	return &BookRepositoryImpl{booksCollection}, nil
}
