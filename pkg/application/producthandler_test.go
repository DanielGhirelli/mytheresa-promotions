package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"mytheresa-promotions/pkg/domain/product"
	"mytheresa-promotions/pkg/service"

	"github.com/gorilla/mux"
)

type mockRepository struct {
	products []product.RawProduct
	err      error
}

func (m *mockRepository) GetAll() ([]product.RawProduct, error) {
	return m.products, m.err
}

func TestProductHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		mockProducts   []product.RawProduct
		expectedStatus int
	}{
		{"success", "/products", []product.RawProduct{{SKU: "1", Price: 10000}}, http.StatusOK},
		{"filter by category", "/products?category=boots", []product.RawProduct{{Category: "boots", Price: 10000}}, http.StatusOK},
		{"filter by price", "/products?priceLessThan=5000", []product.RawProduct{{Price: 4000}}, http.StatusOK},
		{"invalid price filter", "/products?priceLessThan=abc", nil, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			repo := &mockRepository{products: tt.mockProducts}
			handler := NewHandler(service.NewService(repo))
			handler.RegisterRoutes(router)

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}
