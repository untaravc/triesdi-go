# Gin App Base

## Instalation

1. Install Go
   Check if instalation success.

```
go version
```

2. Create New Project

```
mkdir gin-gorm-project
cd gin-gorm-project
go mod init gin-gorm-project
```

3. Install Gin

```
go get -u github.com/gin-gonic/gin
```

4. Install Gorm & Drivers

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u gorm.io/driver/postgres
go get -u gorm.io/driver/sqlite
```

5. Run App
   Set the .env. Copy from .env.example

```
go run main.go
```

6. Run Dev mode

```
air serve
```

## Features

- Gorm config mysql, posgres
- Gin routing
- Auth & Middelware
- Redis cache

---

## Explanation of Folders

- `app/`  
  Main function goes here.

- `app/config/`  
  Contains configuration files, such as database connections.

- `app/controllers/`  
  Holds the logic for handling HTTP requests.

- `app/services/`  
  Contain the logic function.

- `app/request/`  
  Contain struct of request. HTTP POST should bind to these struct

- `app/response/`  
  Contain struct of response. HTTP should return to these struct

- `app/models/`  
  Includes the schema definitions for the database.

- `app/utils/`  
  Utility functions that can be reused across the project.

- `bootstrap/`  
  Tailoring app function.

- `routes/`  
  Defines the API routes.

- `public`  
  Contain files exposed to public

- `main.go`  
  The main entry point of the application.

- `go.mod`  
  Specifies the Go module and dependencies.

## App Flow

- `main.go`  
  Entry point of the application

- `bootstrap/index.bootstrap.go`  
  Tailoring app function.

- routes/
  API end point goes here

---

## Update

Consider updates every Monday in each week

`[24.12.30]`
- Create readme.md, folder structure, .gitignore
- App config, cors config, db connection, redis cache, log config
- Routing, Controller

## Upcoming Feature
- DB migration
- Posgres config
- Mailer
- Firebase config
