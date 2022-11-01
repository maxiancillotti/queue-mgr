package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func setApiRoutes() {

	httpRouter.HandleFunc("/api/jobs",
		prMDW.PanicRecover(
			authMDW.AuthorizationAllow(
				c.POST))).
		Methods(http.MethodPost)

	httpRouter.HandleFunc("/api/jobs",
		prMDW.PanicRecover(
			authMDW.AuthorizationAllow(
				c.GETCollectionByStatus))).
		Methods(http.MethodGet).
		MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
			return r.ContentLength > 0
		})

	httpRouter.HandleFunc("/api/jobs",
		prMDW.PanicRecover(
			authMDW.AuthorizationAllow(
				c.GETCollection))).
		Methods(http.MethodGet).
		MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
			return r.ContentLength <= 0
		})
}

func setRoutesBase() {
	// Health
	httpRouter.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	}).Methods(http.MethodGet)
}

// Wrap httpRouter before starting the server
func CaselessMatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
