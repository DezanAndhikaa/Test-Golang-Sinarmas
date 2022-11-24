package dto

type CreateProductRequest struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

type UpdateProductRequest struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

type PaginationRequest struct {
	Page  int `query:"page"`
	Size  int `query:"size"`
	Limit int `query:"limit"`
}
