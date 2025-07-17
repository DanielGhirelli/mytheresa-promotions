package product

import "github.com/gorilla/mux"

type Product struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
}

type ProductResponse struct {
	SKU      string        `json:"sku"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Price    PriceResponse `json:"price"`
}

// JSON file structure
type RawProduct struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

type Repository interface {
	GetAll() ([]RawProduct, error)
}

type Service interface {
	GetProducts(categoryFilter string, priceLessThanFilter *int) ([]ProductResponse, error)
}

type Handler interface {
	RegisterRoutes(router *mux.Router)
}
