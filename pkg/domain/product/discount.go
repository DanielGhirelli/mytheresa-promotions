package product

import (
	"strconv"
	"strings"
)

func GetProductDiscountedPrice(p Product) PriceResponse {
	var maxDiscount int
	var priceResponse PriceResponse

	priceResponse.Original = p.Price.Original
	priceResponse.Final = p.Price.Final
	priceResponse.DiscountPercentage = nil
	priceResponse.Currency = "EUR"

	if strings.ToLower(p.Category) == "boots" {
		maxDiscount = 30
	}
	if p.SKU == "000003" && maxDiscount < 15 {
		maxDiscount = 15
	}

	if maxDiscount > 0 {
		discountStr := strconv.Itoa(maxDiscount) + "%"
		final := p.Price.Original * (100 - maxDiscount) / 100

		priceResponse.Final = final
		priceResponse.DiscountPercentage = &discountStr
	}

	return priceResponse
}
