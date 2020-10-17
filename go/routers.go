/*
 * velocity limits API server
 * One API enabled:
 * /api/fund - which redirects to FundAccountApi method
 *
 * Other than this a handler to log all requests
 *
 * API version: 1.0.0
 * Contact: daniel.maruani@gmail.com
 */
package _go

import (
	"net/http"
	"strings"

	FundApi "./rest"
	"github.com/gorilla/mux"
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
	Route{
		"FundAccount",
		strings.ToUpper("Post"),
		"/api/fund",
		FundApi.FundAccountApi,
	},
}
