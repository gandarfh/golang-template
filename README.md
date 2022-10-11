# Boilerplate - Go API

## Architecture

```bash
├── internal
│   ├── middleware
│   ├── routes
│   ├── modules
│   │   ├── [module]
│   │   │   ├── [module]_controller.go
│   │   │   ├── [module]_routes.go
│   │   │   ├── dto
│   │   │   ├── entities
│   │   │   ├── repositories
│   │   │   └── services
│   │   └── ...
├── docs
├── infrastructure
│   ├── cache
│   ├── database
│   └── migrations
├── pkg
│   ├── permissions
│   ├── configs
│   ├── errors
│   ├── jwt
│   ├── pagination
├── └── utils
└── main.go
```

## Quick start

1. Copy `.env.example` to `.env` and fill it with your environment values.

2. Run with docker:

```bash
make docker.run
```

3. Go to API Docs page (Swagger): [localhost:5000/docs/index.html](http://localhost:5000/docs/index.html)
   If dont exist swagger page, generate then with this command:

> To install swag: `go install github.com/swaggo/swag/cmd/swag@latest`

```bash
make swag
```

4. To use live reload, download air package.

> To install: `go install github.com/cosmtrek/air@latest`

5. To execute live reload:

```bash
air run .
```

## 📦 Features Roadmap

- [ ] Deployment to aws
- [x] Http request implementation
- [x] Logs implementation
- [ ] Implementation AWS stuffs, like: SNS, SKS
- [ ] Documentation - how create new modules
- [ ] Documentation - how connect to database
- [ ] Documentation - how private route by credentials
- [ ] Documentation - how create a new token
- [ ] Documentation - how generate swagger documentation
- [x] JWT Token generator
- [x] Middleware to valid credentials or roles (provided by JWT token)
- [x] Support to MongoDb, MySql, Postgres, Redis
- [x] Pagination with MongoDb
- [x] Swagger generator
- [x] Struct validator who handle path and validation tag on response
- [x] Multi routes implementation, not found, swagger, public, private
- [x] Dockerfile to execute binary file
