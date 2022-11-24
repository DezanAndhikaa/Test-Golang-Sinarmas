package usecase

import "test-golang/application/product/dto"

func (i impl) Get(request dto.PaginationRequest) (dto.GetAllResponse, error) {
	response := dto.GetAllResponse{}

	allProduct, err := i.repo.Get(request.Page, request.Limit, request.Size)
	if err != nil {
		return response, err
	}

	for _, val := range allProduct {
		response.Data = append(response.Data, val)
	}

	return response, nil
}
