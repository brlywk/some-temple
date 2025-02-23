package main

import (
	"context"
	"fmt"
	"net/http"
	"some-temple/appctx"
	"some-temple/views/pages"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

// withCurrentPath adds the current route path to every request.
func withCurrentPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[withCurrentPath]", "request url", r.URL.Path)
		ctx := context.WithValue(r.Context(), appctx.KeyCurrentPath, r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

// NOTE: For now it's fine to have all routes in here. If at some point routes
// will get too much, it can be split up and moved into a routes/ directory.

// newAppRouter handles all app requests.
func newAppRouter(title string) chi.Router {
	r := chi.NewRouter()

	// add the current path shown to every request
	r.Use(withCurrentPath)

	// root
	r.Get("/", templ.Handler(pages.Index(title)).ServeHTTP)

	// pages/
	r.Mount("/pages", newPagesRouter())

	return r
}

// newPagesRouter handles all /pages routes.
func newPagesRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/dashboard", templ.Handler(pages.Dashboard()).ServeHTTP)
	r.Get("/test", templ.Handler(pages.TestPage()).ServeHTTP)

	return r
}
