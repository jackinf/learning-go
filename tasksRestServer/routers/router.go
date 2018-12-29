package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRoutes(router)
	router = SetTaskRouter(router)
	router = SetNoteRouter(router)
	return router
}