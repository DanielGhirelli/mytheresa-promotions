package application

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mytheresa-promotions/pkg/domain/product"

	"github.com/gorilla/mux"
)

type handler struct {
	service product.Service
}

func NewHandler(service product.Service) product.Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.getProducts).Methods("GET")
}

func (h *handler) getProducts(response http.ResponseWriter, request *http.Request) {
	var priceLessThanFilter *int

	// Params
	priceLessThan := request.URL.Query().Get("priceLessThan")
	category := request.URL.Query().Get("category")

	if priceLessThan != "" {
		price, err := strconv.Atoi(priceLessThan)
		if err != nil {
			http.Error(response, "Invalid priceLessThan parameter: "+priceLessThan, http.StatusBadRequest)
			return
		}
		priceLessThanFilter = &price
	}

	// Get products from service
	products, err := h.service.GetProducts(category, priceLessThanFilter)
	if err != nil {
		http.Error(response, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Return response
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(map[string]interface{}{"products": products})
}
