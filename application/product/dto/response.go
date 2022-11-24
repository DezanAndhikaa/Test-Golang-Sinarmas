package dto

import "test-golang/domain"

type GetAllResponse struct {
	Page int `json:"page"`
	Size int `json:"size"`

	Data []domain.Product `json:"data"`
}
