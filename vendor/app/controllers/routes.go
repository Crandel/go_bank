package controllers

import (
	"app/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

var router *mux.Router

// RouteInit - Create new httprouter for ListenAndServe http loop
func RouteInit() *mux.Router {
	models.Automigrate()
	router = mux.NewRouter()
	// Create base list of middlewares
	baseMidList := []alice.Constructor{LogMiddleware}
	// Create auth list of middlewares, extended from base
	authMidList := make([]alice.Constructor, len(baseMidList))
	copy(authMidList, baseMidList)

	// append from base list
	// authMidList = append(authMidList, UserInContext)

	baseAlice := alice.New(baseMidList...)
	authAlice := alice.New(authMidList...)
	router.Handle("/", authAlice.Then(MainHandler)).Name("home")
	router.PathPrefix("/static/").Handler(baseAlice.Then(http.StripPrefix("/static/", http.FileServer(http.Dir("./public")))))
	// router.NotFoundHandler = http.HandlerFunc(NotFoundHandleFunc)
	router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
	// api := router.PathPrefix("/api").Subrouter()
	// api.Handle("/login", baseAlice.Then(LoginHandler)).Methods("GET", "POST").Name("login")
	// api.Handle("/signin", baseAlice.Then(SignHandler)).Methods("GET", "POST").Name("sign")
	// api.Handle("/signout", authAlice.Then(SignOutHandler)).Methods("GET").Name("signout")
	return router
}

// RedirectFunc - redirect to named router
func RedirectFunc(name string) (string, error) {
	url, err := router.Get(name).URL()
	return url.String(), err
}
