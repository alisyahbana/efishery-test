package handler

import (
	"fmt"
	"github.com/alisyahbana/efishery-test/pkg/common/helper"
	"github.com/alisyahbana/efishery-test/pkg/service/product"
	"github.com/alisyahbana/efishery-test/pkg/service/user"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func FetchProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bearerToken := request.Header.Get("Authorization")
	tempString := strings.Split(bearerToken, " ")
	if bearerToken == "" || len(tempString) == 1 {
		helper.ErrorReturn(writer, http.StatusUnauthorized, fmt.Errorf("token required"))
		return
	}
	tokenString := tempString[1]
	claims, err := user.New().ValidateToken(tokenString)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusInternalServerError, err)
		return
	}

	if claims == nil {
		helper.ErrorReturn(writer, http.StatusUnauthorized, fmt.Errorf("token invalid"))
		return
	}

	if claims.Role != "admin" {
		helper.ErrorReturn(writer, http.StatusUnauthorized, fmt.Errorf("role invalid"))
		return
	}

	products, err := product.New().GetAllProducts()
	if err != nil {
		helper.ErrorReturn(writer, http.StatusInternalServerError, err)
		return
	}
	helper.SuccessReturn(writer, products)
	return
}

func FetchProductCompiled(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bearerToken := request.Header.Get("Authorization")
	tempString := strings.Split(bearerToken, " ")
	if bearerToken == "" || len(tempString) == 1 {
		helper.ErrorReturn(writer, http.StatusUnauthorized, fmt.Errorf("token required"))
		return
	}
	tokenString := tempString[1]
	claims, err := user.New().ValidateToken(tokenString)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusInternalServerError, err)
		return
	}

	if claims == nil {
		helper.ErrorReturn(writer, http.StatusUnauthorized, fmt.Errorf("token invalid"))
		return
	}

	if claims.Role != "admin" {
		helper.ErrorReturn(writer, http.StatusUnauthorized, fmt.Errorf("role invalid"))
		return
	}

	products, err := product.New().GetCompiledProduct()
	if err != nil {
		helper.ErrorReturn(writer, http.StatusInternalServerError, err)
		return
	}
	helper.SuccessReturn(writer, products)
	return
}
