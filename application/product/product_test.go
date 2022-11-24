package product_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"test-golang/application/product"
	"test-golang/application/product/dto"
	mock_repository "test-golang/application/product/mocks/repository"
	"test-golang/application/product/usecase"
	"test-golang/domain"
)

var someError = errors.New("some error")

var _ = Describe("Product", func() {
	var (
		mockController *gomock.Controller
		productRepo    *mock_repository.MockRepository
		productUsecase product.Usecase
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		productRepo = mock_repository.NewMockRepository(mockController)
		productUsecase = usecase.New(productRepo)
	})

	Describe("Create Product", func() {
		It("success create product", func() {
			request := dto.CreateProductRequest{}

			productRepo.EXPECT().Create(gomock.Any()).Return(&domain.Product{
				ID:         1,
				Name:       "Mock",
				Qty:        1,
				AuditTable: domain.AuditTable{},
			}, nil)

			_, err := productUsecase.Create(request)
			Expect(err).Should(Succeed())
		})

		It("fail create product", func() {
			request := dto.CreateProductRequest{}

			productRepo.EXPECT().Create(gomock.Any()).Return(nil, someError)

			_, err := productUsecase.Create(request)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("Get Product By ID", func() {
		It("success get product", func() {
			mockData := domain.Product{
				ID:   1,
				Name: "test",
				Qty:  1,
			}

			productRepo.EXPECT().GetByID(gomock.Any()).Return(&mockData, nil)

			_, err := productUsecase.GetByID(1)
			Expect(err).Should(Succeed())
		})

		It("fail get product", func() {
			productRepo.EXPECT().GetByID(gomock.Any()).Return(nil, someError)

			_, err := productUsecase.GetByID(1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("Get All Product", func() {
		It("success get all product", func() {
			mockData := []domain.Product{
				{
					ID:   1,
					Name: "test",
					Qty:  1,
				},
			}
			request := dto.PaginationRequest{}

			productRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockData, nil)

			_, err := productUsecase.Get(request)
			Expect(err).Should(Succeed())
		})

		It("fail get all product", func() {
			request := dto.PaginationRequest{}

			productRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, someError)

			_, err := productUsecase.Get(request)
			Expect(err).Should(HaveOccurred())
		})
	})

})
