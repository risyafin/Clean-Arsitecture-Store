package products

import "kasir/modules/categories"

type Product struct {
	Id          int                 `json:"id"`
	Nama        string              `json:"nama"`
	Price       int                 `json:"price"`
	Category_Id int                 `json:"category_id"`
	Category    categories.Category `json:"category"`
}
