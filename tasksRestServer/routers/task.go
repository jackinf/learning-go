package routers

import (
	"awesomeProject/tasksRestServer/common"
	"awesomeProject/tasksRestServer/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetTaskRouter(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("UPDATE")
	taskRouter.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	taskRouter.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/tasks/user/{id}", controllers.GetTasksByUser).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(taskRouter),
	))
	return router
}
