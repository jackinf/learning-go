package routers

import (
	"awesomeProject/tasksRestServer/common"
	"awesomeProject/tasksRestServer/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetNoteRouter(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()
	noteRouter.HandleFunc("/Notes", controllers.CreateNote).Methods("POST")
	noteRouter.HandleFunc("/Notes/{id}", controllers.UpdateNote).Methods("UPDATE")
	noteRouter.HandleFunc("/Notes/{id}", controllers.DeleteNote).Methods("DELETE")
	noteRouter.HandleFunc("/Notes", controllers.GetNotes).Methods("GET")
	noteRouter.HandleFunc("/Notes/user/{id}", controllers.GetNotesByTask).Methods("GET")
	noteRouter.HandleFunc("/Notes/{id}", controllers.GetNoteByID).Methods("GET")
	router.PathPrefix("/Notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRouter),
	))
	return router
}
