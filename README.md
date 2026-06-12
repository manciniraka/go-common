# go-common

Reusable Go packages for backend projects.

This repository contains common packages that are frequently used across Golang backend projects, especially for REST APIs built with Echo, PostgreSQL, GORM, JWT Authentication, and Validation.

## Installation

```bash
go get github.com/manciniraka/go-common
```

---

## Packages

### password

Helper package for password hashing and verification using bcrypt.

#### Hash Password

```go
hashedPassword, err := password.Hash("secret123")
```

#### Compare Password

```go
err := password.Compare(
	hashedPassword,
	"secret123",
)
```

---

### validator

Custom validator implementation for Echo Framework using go-playground/validator.

#### Setup

```go
e.Validator = validator.New()
```

#### Example Request Struct

```go
type RegisterRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
```

#### Validate

```go
if err := c.Validate(&payload); err != nil {
	return err
}
```

---

### logger

Structured JSON logger built on top of slog.

#### Setup

```go
appLogger := logger.NewJSON()
```

#### Example

```go
appLogger.Error(
	"failed create user",
	slog.Any("error", err),
)
```

---

### database

PostgreSQL database connection helper using GORM.

Required environment variables:

```env
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
DB_SSLMODE=
```

#### Connect

```go
db, err := database.ConnectPostgres()
if err != nil {
	log.Fatal(err)
}
```

#### Auto Migrate

Auto migration should be handled inside each project.

```go
err = db.AutoMigrate(
	&entity.User{},
	&entity.Book{},
)
```

---

### jwt

JWT helper package built using golang-jwt/jwt/v5.

#### Generate Token

```go
claims := jwtgo.MapClaims{
	"user_id": 1,
	"email":   "user@mail.com",
	"role":    "admin",
	"exp":     time.Now().Add(24 * time.Hour).Unix(),
}

token, err := jwt.GenerateToken(
	claims,
	os.Getenv("JWT_SECRET"),
)
```

#### Parse Token

```go
claims, err := jwt.ParseToken(
	tokenString,
	os.Getenv("JWT_SECRET"),
)
```

---

## Current Version

### v1.0.0

Included packages:

* password
* validator
* logger
* database
* jwt

---
