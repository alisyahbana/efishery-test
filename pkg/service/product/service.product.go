package product

import (
	"encoding/json"
	"fmt"
	"github.com/alisyahbana/efishery-test/pkg/common/helper"
	"github.com/alisyahbana/efishery-test/pkg/service/product/data"
	"log"
	"net/http"
	"strconv"
)

type ProductService struct {
}

func New() ProductService {
	return ProductService{}
}

func (s ProductService) GetAllProducts() ([]data.Product, error) {
	resp, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var listProduct []data.Product

	err = json.NewDecoder(resp.Body).Decode(&listProduct)
	if err != nil {
		return nil, err
	}

	rateUsd, err := helper.GetRatioUSD()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(listProduct); i++ {
		if listProduct[i].Price != nil {
			priceFloat, err := strconv.ParseFloat(*listProduct[i].Price, 64)
			if err != nil {
				log.Println(err, "Failed to convert price to float for uuid:", listProduct[i].ID)
			}
			usdPriceFloat := priceFloat * rateUsd
			listProduct[i].USDPrice = fmt.Sprintf("$%.2f", usdPriceFloat)
		}
	}

	return listProduct, nil
}
