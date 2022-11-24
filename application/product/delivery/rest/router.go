package rest

import "github.com/labstack/echo/v4"

func MapProductRoutes(group *echo.Group, h handler) {
	group.GET("/", h.GetAll)
	group.POST("/", h.CreateProduct)
	group.GET("/:id", h.GetByID)
}
