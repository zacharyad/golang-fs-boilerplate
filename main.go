// file: main.go
package main

import (
	"log"
	"net/http"

	"github.com/zacharyad/portfolio/view"
	"github.com/zacharyad/portfolio/view/layout"
	"github.com/zacharyad/portfolio/view/partial"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	c := layout.Base(view.Index())

	http.Handle("/", templ.Handler(c))

	http.Handle("/foo", templ.Handler(partial.Comp()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
