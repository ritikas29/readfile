package album
import (
	"net/http"
	"proj1/logger"
	"github.com/gorilla/mux"
)
var controller = &Controller{Repository:Repository{}}
type Route struct {
	name string
	method string
	pattern string
	HandlerFunc http.HandleFunc
}
type Routes []Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route {
		"AddAlbum",
		"POST",
		"/",
		controller.AddAlbum,
	},
	Route {
		"UpdateAlbum",
		"PUT",
		"/",
		controller.UpdateAlbum,
	},
	Route {
		"DeleteAlbum",
		"DELETE",
		"/{id}",
		controller.DeleteAlbum,
	}
}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true) 
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler,route.Name)
	router.
	methods(route.method).
	path(route.pattern).
	name(route.name)
	Handler(handler)
	}
	return router
}
