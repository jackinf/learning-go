package main

import (
	"awesomeProject/tasksRestServer/common"
	"awesomeProject/tasksRestServer/routers"
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
)

func main() {
	common.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	appConfig := common.AppConfig
	addr := appConfig.Server
	server := &http.Server{
		Addr: 	addr,
		Handler: 	n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
