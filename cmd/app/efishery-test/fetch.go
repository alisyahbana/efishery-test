package main

import (
	"fmt"
	"github.com/alisyahbana/efishery-test/pkg/common/app"
	"github.com/alisyahbana/efishery-test/pkg/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	router := httprouter.New()
	SetRouteFetchApp(router)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	fmt.Println(fmt.Sprintf("Starting Efishery Commodity API HTTP Server on %d", app.GetConfig().PortFetch))
	err := http.ListenAndServe(fmt.Sprintf(":%d", app.GetConfig().PortFetch), n)
	if err != nil {
		panic(err)
	}
}

func SetRouteFetchApp(router *httprouter.Router) {
	router.GET("/fetch-product", handler.FetchProduct)
	router.GET("/fetch-product-compiled", handler.FetchProductCompiled)
	router.GET("/auth", handler.AuthTokenHandler)
}
