package main

import (
	"github.com/cunla/velocity-limiter/logic"
	FundApi "github.com/cunla/velocity-limiter/rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
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
		handler = logic.Logger(handler, route.Name)

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

func main() {

	log.Printf("Server started")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
