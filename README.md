# Recipes

## Introduction

- A practical project to learn go
- A web server for recipes management

## Usage

`go run cmd/main.go`

## Packges Used

- Gin: Web framework
- uuid: unique id generator
- viper: environment variables configuration
- mongo-driver: MongoDB driver
- logur: logging package

## Structure

```
.
├── README.md
├── cmd
│   ├── main.go       // entry point
│   └── server        // Starting server package
├── config            // Environment configuration (Currently using viper)
├── internal          // main source code
│   ├── api           // API layer (contains code for routers and some dependences)
│   ├── controllers   // Handler functions for routers
│   ├── databases     // Database connection
│   ├── interfaces    // Ports and Adapters for Service and Repositories layers
│   ├── lib           // Reusable functions
│   ├── models        // Data models
│   ├── repositories  // Repository layer ()
│   └── services      // Service layer
├── pkg               // External packages, also reusable (like logging, uuid, ..)
│   ├── logger
│   ├── setting
│   └── utils
└── tests             // Testing related code
```

## Features

- [ ] Swagger Docs
- [ ] Logging
- [ ] Exception Handling
- [ ] CRUD Recipes
- [ ] CRUD Recipe Books
- [ ] User Authentications: loging, logout, register
- [ ] User Authorizations: private and shared recipies
- [ ] Tests
