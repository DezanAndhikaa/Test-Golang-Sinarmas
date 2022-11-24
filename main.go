package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"test-golang/application/product/delivery/rest"
	"test-golang/application/product/repository"
	"test-golang/application/product/usecase"
	infradb "test-golang/infra/db"
)

func main() {
	e := echo.New()

	db := infradb.InitGorm()
	repo := repository.New(db)
	uc := usecase.New(repo)

	productHandler := rest.NewProductHandler(uc)

	productGroup := e.Group("/api/v1/product")
	rest.MapProductRoutes(productGroup, productHandler)

	if err := e.Start(":5000"); err != nil {
		log.Fatal("Error when start server")
	}

}
