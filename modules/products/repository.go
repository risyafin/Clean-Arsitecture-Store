package products

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllProducts() ([]Product, error) {
	var products []Product
	result := repo.DB.Preload("Category").Find(&products)
	return products, result.Error
}
func (repo Repository) GetProductById(id string) (*Product, error) {
	var product *Product
	result := repo.DB.Preload("Category").First(&product, id)
	return product, result.Error
}

func (repo Repository) CreateProduct(product Product) error {
	result := repo.DB.Create(&product)
	return result.Error
}

func (repo Repository) UpdateProductByID(id string, product Product) error {
	result := repo.DB.Model(&Product{}).Where("id = ?", id).Updates(map[string]interface{}{
		"nama":        product.Nama,
		"price":       product.Price,
		"category_id": product.Category_Id,
	})
	return result.Error
}

func (repo Repository) DeleteProductById(id string) error {
	result := repo.DB.Delete(&Product{}, id)
	return result.Error
}
