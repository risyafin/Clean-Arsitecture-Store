package products

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetAllProducts() ([]Product, error) {
	products, err := usecase.Repo.GetAllProducts()
	return products, err
}

func (usecase Usecase) GetProductById(id string) (*Product, error) {
	product, err := usecase.Repo.GetProductById(id)
	return product, err
}

func (usecase Usecase) CreateProduct(product Product) error {
	err := usecase.Repo.CreateProduct(product)
	return err
}

func (usecase Usecase) UpdateProductById(id string, product Product) error {
	err := usecase.Repo.UpdateProductByID(id, product)
	return err
}

func (usecase Usecase) DeleteProductById(id string) error {
	err := usecase.Repo.DeleteProductById(id)
	return err
}
