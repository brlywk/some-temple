package main

import (
	"some-temple/views/pages"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

// NOTE: For now it's fine to have all routes in here. If at some point routes
// will get too much, it can be split up and moved into a routes/ directory.

func NewAppRouter(title string) chi.Router {
	r := chi.NewRouter()

	r.Get("/", templ.Handler(pages.Index(title)).ServeHTTP)

	return r
}
