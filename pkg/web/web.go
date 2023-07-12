package web

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ServeApi() error {

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(5 * time.Second))

	routes := myRoutes{}
	routes.addToMux(router)

	return http.ListenAndServe(":8080", router)
}

type myRoutes struct{}

func (routes *myRoutes) addToMux(mux *chi.Mux) {
	mux.Route("/api", func(router chi.Router) {
		router.Get("/", routes.helloRoute)
	})
}
func (routes *myRoutes) helloRoute(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello"))
}
