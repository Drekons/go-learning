package main

import (
	"dz28/pkg/router"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "specify port to use. Defaults to 8080.")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	router.Init(r)

	log.Printf("Server running on port %d ...", port)

	if err := http.ListenAndServe(":"+strconv.Itoa(port), r); err != nil {
		log.Fatalln(err)
	}
}
