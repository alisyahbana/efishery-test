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
	SetRoute(router)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	fmt.Println(fmt.Sprintf("Starting Efishery Commodity API HTTP Server on %d", app.GetConfig().Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", app.GetConfig().Port), n)
	if err != nil {
		panic(err)
	}
}

func SetRoute(router *httprouter.Router) {
	router.GET("/", handler.InfoHandler)
	router.POST("/register", handler.RegisterHandler)
	router.POST("/login", handler.LoginHandler)
	router.GET("/auth", handler.AuthTokenHandler)
}
