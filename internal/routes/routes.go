package routes

import (
	"net/http"

	"github.com/go-chi/chi"

	"targetsistemas/internal/handler"
)

func InitializeRoutes() {
	r := chi.NewRouter()

	r.Get("/n1", handler.N1Handler)
	r.Get("/n2", handler.N2Handler)
	r.Get("/n3", handler.N3Handler)
	r.Get("/n4", handler.N4Handler)
	r.Get("/n5", handler.N5Handler)

	http.Handle("/", r)
}
