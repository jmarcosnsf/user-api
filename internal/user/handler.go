package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)
type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) http.Handler {
	r := chi.NewMux()
	
	h := &Handler{repo: repo}

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/users", h.FindAll)
	r.Get("/users/{id:[a-f0-9-]{36}}", h.FindById)
	r.Post("/users", h.Insert)
	r.Put("/users/{id:[a-f0-9-]{36}}", h.Update)
	r.Delete("/users/{id:[a-f0-9-]{36}}", h.Delete)

	return r
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) FindById(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) Insert(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request){

}