package main

import (
	"github.com/arbeeorlar/jwt-example/initializers"
	"github.com/arbeeorlar/jwt-example/migrations"
	"github.com/arbeeorlar/jwt-example/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
}

func main() {
	migrations.AutoMigrate()
	r := gin.Default()
	routes.AuthRoutes(r)
	r.Run()
}

////
/***
go
gorm
postgre driver
env
jwt


go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt



*/
