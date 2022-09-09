package repositories

import (
	"fmt"
	"goapi/app/modules/books/dto"
	"goapi/app/modules/books/entities"
	"goapi/shared/errors"
	"goapi/shared/pagination"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetBooks method for getting all books.
func (repo *BookRepo) GetBooks(ctx *fiber.Ctx) (*dto.BookListAllResponse, error) {
	// Define books variable.
	books := []entities.Book{}

	// Length of books document.
	count, err := repo.Mongo.Collection("books").CountDocuments(ctx.Context(), bson.D{})

	// Error when cant count items into document
	if err != nil {
		return nil, err
	}

	// Generate pagination to mongodb find method
	paginate := pagination.NewMongoPaginate(ctx, count)

	// Send query to database.
	cur, err := repo.Mongo.Collection("books").Find(ctx.Context(), bson.D{}, paginate.Options())

	// Error when cant find any books into database
	if err != nil {
		// TODO tratar esse erro
	return nil, err
}

	// Error when marshal to books struct
	if err := cur.All(ctx.Context(), &books); err != nil {
	// TODO tratar esse erro
return nil, err
	}

	result := pagination.Paginate[entities.Book]{
	Items:         books,
MongoPaginate: paginate,
	}

// Return query result.
return &dto.BookListAllResponse{Paginate: &result}, nil
}

// GetBook method for getting one book by given ID.
func (repo *BookRepo) GetBook(ctx *fiber.Ctx, id uuid.UUID) (*entities.Book, error) {
// Define book variable.
	book := entities.Book{}

	err := repo.Mongo.Collection("books").FindOne(ctx.Context(), bson.D{{Key: "id", Value: id}}).Decode(&book)
if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, errors.NotFound(errors.Message{
			"error": true,
"msg":   fmt.Sprintf("book with the given ID: %s is not found", id),
		})

		}

return nil, err
	}

// Return query result.
return &book, nil
}

// CreateBook method for creating book by given Book object.
func (repo *BookRepo) CreateBook(ctx *fiber.Ctx, b *entities.Book) (*entities.Book, error) {
// Convert struct to bson structure.
	value := b.Value()

// Insert book to database.
	_, err := repo.Mongo.Collection("books").InsertOne(ctx.Context(), value)

	// If dont create the book, return 400
	if err != nil {
		// Return empty object and error.
		return nil, errors.BadRequest(errors.Message{
		"error": true,
	"msg":   "Can't create this book!",
})
	}

// Return query result.
return b, nil
}

// UpdateBook method for updating book by given Book object.
func (repo *BookRepo) UpdateBook(ctx *fiber.Ctx, book *entities.Book) (*entities.Book, error) {
	newBook := entities.Book{}

	filter := bson.D{{Key: "id", Value: book.ID}}

	value := (*book).Value()

	_, err := repo.Mongo.
Collection("books").
	UpdateOne(ctx.Context(), filter, bson.D{{Key: "$set", Value: value}})

	if err != nil {
return nil, errors.NotFound(errors.Message{"msg": err.Error()})
	}

	fmt.Println(value)

return &newBook, nil
}

// DeleteBook method for delete book by given ID.
func (q *BookRepo) DeleteBook(ctx *fiber.Ctx, id uuid.UUID) error {
// Define query string.
	// query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	// _, err := q.SQL.Exec(query, id)
	// if err != nil {
	// 	// Return only error.
// 	return err
	// }

// This query returns nothing.
return nil
}
