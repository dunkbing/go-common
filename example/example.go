package main

import (
	"fmt"
	"github.com/dunkbing/go-common/middlewares"
	"github.com/dunkbing/go-common/mux"
	"log"
	"net/http"
)

func mid(i int) mux.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("mid", i, "start")
			next.ServeHTTP(w, r)
			fmt.Println("mid", i, "done")
		})
	}
}

func someHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[the handler ran here]")
	n, err := fmt.Fprintln(w, "Hello world of", r.URL.Path)
	if err != nil {
		log.Println(err)
	}
	log.Println(n, "bytes written")
}

func main() {
	origins := []string{"http://localhost:8000/"}
	r := mux.NewRouter(middlewares.EnableCORS(origins), mid(0))

	r.Group(func(r *mux.Router) {
		r.Use(mid(2))

		r.Get("/foo", someHandler)
	})

	r.Group(func(r *mux.Router) {
		r.Use(mid(3))

		r.Get("/bar", someHandler, mid(4))
		r.Get("/baz", someHandler, mid(5))
	})

	r.Post("/foobar", someHandler)

	log.Fatal(http.ListenAndServe(":3000", r))
}
