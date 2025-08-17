package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

// This Product was not supposed to be returned. WE should return a DTO instead
// However, we will return it for now to keep the example simple.
func (u *ProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}