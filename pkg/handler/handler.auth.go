package handler

import (
	"encoding/json"
	"github.com/alisyahbana/efishery-test/pkg/common/helper"
	"github.com/alisyahbana/efishery-test/pkg/service/user"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strings"
)

func RegisterHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jsonBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusInternalServerError, err)
		return
	}
	var payload user.RegisterPayload
	err = json.Unmarshal(jsonBody, &payload)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusBadRequest, err)
		return
	}

	resp, err := user.New().Register(payload)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusBadRequest, err)
		return
	}

	helper.SuccessReturn(writer, resp)
	return
}

func LoginHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jsonBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusInternalServerError, err)
		return
	}
	var payload user.LoginPayload
	err = json.Unmarshal(jsonBody, &payload)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusBadRequest, err)
		return
	}

	resp, err := user.New().Login(payload)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusBadRequest, err)
		return
	}

	helper.SuccessReturn(writer, resp)
	return
}

func AuthTokenHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bearerToken := request.Header.Get("Authorization")
	tempString := strings.Split(bearerToken, " ")
	tokenString := tempString[1]

	claims, err := user.New().ValidateToken(tokenString)
	if err != nil {
		helper.ErrorReturn(writer, http.StatusInternalServerError, err)
		return
	}

	helper.SuccessReturn(writer, claims)
	return
}
