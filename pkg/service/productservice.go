package service

import (
	"mytheresa-promotions/pkg/domain/product"
)

const (
	APIResponseLimit = 5
)

type service struct {
	repo product.Repository
}

func NewService(repo product.Repository) product.Service {
	return &service{repo: repo}
}

func (s *service) GetProducts(categoryFilter string, priceLessThanFilter *int) ([]product.ProductResponse, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	filtered := make([]product.ProductResponse, 0, APIResponseLimit)

	for _, raw := range products {
		p := product.Product{
			SKU:      raw.SKU,
			Name:     raw.Name,
			Category: raw.Category,
			Price: product.Price{
				Original: raw.Price,
				Final:    raw.Price,
			},
		}

		if priceLessThanFilter != nil && p.Price.Original > *priceLessThanFilter {
			continue
		}
		if categoryFilter != "" && p.Category != categoryFilter {
			continue
		}

		priceResponse := product.GetProductDiscountedPrice(p)

		filtered = append(filtered, product.ProductResponse{
			SKU:      p.SKU,
			Name:     p.Name,
			Category: p.Category,
			Price:    priceResponse,
		})

		if len(filtered) == APIResponseLimit {
			break
		}
	}

	return filtered, nil
}
