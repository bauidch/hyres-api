package routers

import (
	"encoding/json"
	. "github.com/bauidch/hyres-api/logger"
	"github.com/bauidch/hyrt-api/dao"
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
	Route{"Index", "GET", "/v1/", Index },
	Route{"Series", "GET", "/v1/series", dao.GetAllSeries },
	Route{"Add Serie", "POST", "/v1/series", dao.CreateSerie },
	Route{"Get One Serie", "GET", "/v1/series/{id}", dao.GetOneSeries },
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Index(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
