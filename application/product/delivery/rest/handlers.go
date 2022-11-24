package rest

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"test-golang/application/product"
	"test-golang/application/product/dto"
)

type handler struct {
	usecase product.Usecase
}

func NewProductHandler(uc product.Usecase) handler {
	return handler{usecase: uc}
}

func (h *handler) CreateProduct(c echo.Context) error {
	request := dto.CreateProductRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	response, err := h.usecase.Create(request)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

func (h *handler) GetByID(c echo.Context) error {
	id := c.Param("id")

	idToInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, map[string]string{
			"status": "id must be int",
		})
	}

	response, err := h.usecase.GetByID(idToInt)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

func (h *handler) GetAll(c echo.Context) error {
	request := dto.PaginationRequest{}
	response, err := h.usecase.Get(request)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}
