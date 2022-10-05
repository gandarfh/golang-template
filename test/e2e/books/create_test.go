package books

import (
	"goapi/internal/routes"
	"goapi/pkg/jwt"
	"goapi/pkg/permissions"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		panic(err)
	}

	data := `{"title": "Book Tittle", "author": "Tester", "book_attrs": { "rating": 3 } }`

	credentials, _ := permissions.GetCredentialsByRole(permissions.AdminRole)
	token, _ := jwt.GenerateNewTokens(uuid.NewString(), permissions.AdminRole, credentials)

	tests := []struct {
		description  string
		route        string
		method       string
		jwtToken     string
		body         io.Reader
		expectedCode int
	}{
		{
			description:  "Shouldn't create book without JWT and body",
			route:        "/api/v1/books",
			method:       "POST",
			jwtToken:     "",
			body:         nil,
			expectedCode: 401,
		},
		{
			description:  "create book with JWT wrong and body",
			route:        "/api/v1/books",
			method:       "POST",
			jwtToken:     "Bearer " + token.Access,
			body:         strings.NewReader(data),
			expectedCode: 201,
		},
	}

	app := fiber.New()
	routes.PrivateRoutes(app)
	routes.PublicRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Authorization", test.jwtToken)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req, -1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
