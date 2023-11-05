package handlers

import (
	"net/http"

	"github.com/elue-dev/go-program/pkg/renderer"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "about.page.gohtml")
}


