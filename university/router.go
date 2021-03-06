package university

import (
	"net/http"

	"github.com/Tardishade/YouCanGetIn/logger"
	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"AddUniversity",
		"POST",
		"/",
		controller.AddUniversity,
	},
	Route{
		"UpdateUniversity",
		"PUT",
		"/{id}",
		controller.UpdateUniversity,
	},
	Route{
		"DeleteUniversity",
		"GET",
		"/{id}",
		controller.DeleteUniversity,
	},
	Route{
		"GetGrades",
		"GET",
		"/grades/{id}",
		controller.GetGrades,
	},
	Route{
		"AddGrade",
		"PUT",
		"/grades/",
		controller.AddGrade,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
