package handler

import (
	"github.com/alisyahbana/efishery-test/pkg/common/helper"
	"github.com/alisyahbana/efishery-test/pkg/service/product"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func FetchProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	products, err := product.New().GetAllProducts()
	if err != nil {
		helper.ErrorReturn(writer, http.StatusInternalServerError, err)
		return
	}
	helper.SuccessReturn(writer, products)
	return
}
