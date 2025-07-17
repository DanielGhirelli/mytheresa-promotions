package service

import (
	"fmt"
	"mytheresa-promotions/pkg/domain/product"
	"testing"
)

type mockRepository struct {
	products []product.RawProduct
	err      error
}

func (m *mockRepository) GetAll() ([]product.RawProduct, error) {
	return m.products, m.err
}

func TestGetProducts(t *testing.T) {
	testProducts := []product.RawProduct{
		{SKU: "1", Name: "Boots", Category: "boots", Price: 10000},
		{SKU: "2", Name: "Sneakers", Category: "sneakers", Price: 8000},
		{SKU: "000003", Name: "Special", Category: "sandals", Price: 5000},
	}

	tests := []struct {
		name          string
		repoProducts  []product.RawProduct
		repoError     error
		category      string
		priceLessThan *int
		wantCount     int
		wantErr       bool
	}{
		{"all products", testProducts, nil, "", nil, 3, false},
		{"filter by category", testProducts, nil, "boots", nil, 1, false},
		{"filter by price", testProducts, nil, "", intPtr(9000), 2, false},
		{"combined filters", testProducts, nil, "boots", intPtr(15000), 1, false},
		{"repository error", nil, ErrNotFound, "", nil, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockRepository{
				products: tt.repoProducts,
				err:      tt.repoError,
			}
			service := NewService(repo)

			products, err := service.GetProducts(tt.category, tt.priceLessThan)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(products) != tt.wantCount {
				t.Errorf("Expected %d products, got %d", tt.wantCount, len(products))
			}
		})
	}
}

func intPtr(i int) *int {
	return &i
}

var ErrNotFound = fmt.Errorf("not found")

func TestGetProductDiscountedPrice(t *testing.T) {
	tests := []struct {
		name     string
		product  product.Product
		expected int
	}{
		{"boots discount",
			product.Product{SKU: "000002", Category: "boots", Price: product.Price{Original: 10000}},
			7000},
		{"sku discount",
			product.Product{SKU: "000003", Price: product.Price{Original: 10000}},
			8500},
		{"both discounts",
			product.Product{SKU: "000003", Category: "boots", Price: product.Price{Original: 10000}},
			7000},
		{"no discount",
			product.Product{SKU: "000005", Price: product.Price{Original: 10000}},
			0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := product.GetProductDiscountedPrice(tt.product)
			if result.Final != tt.expected {
				t.Errorf("Expected final price %d, got %d", tt.expected, result.Final)
			}
		})
	}
}
