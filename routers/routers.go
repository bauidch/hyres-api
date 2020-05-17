package routers

import (
	. "github.com/bauidch/hyrt-api/logger"
	"github.com/bauidch/hyrt-api/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{"Ping", "GET", "/ping", handlers.Ping },
	Route{"Index", "GET", "/v1/", handlers.Index },
	Route{"Series", "GET", "/v1/series", handlers.GetAllSeries },
	Route{"Add Serie", "POST", "/v1/series", handlers.CreateSerie },
	Route{"Get One Serie", "GET", "/v1/series/{id}", handlers.GetOneSeries },
	//Route{"Delet One Serie", "DEL", "/v1/series/{id}", handlers.GetOneSeries },
	Route{"Seed", "GET", "/v1/seed", handlers.GetAllSeed },
	Route{"Add Seed", "POST", "/v1/seed", handlers.CreateSeed },
	//Route{"Get One Seed", "GET", "/v1/seed/{id}", handlers.GetOneSeed },
	Route{"Get Seed Journal", "GET", "/v1/journal/seed", handlers.GetAllSeedJornal },
}
