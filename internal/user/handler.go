package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"user-api/internal/api"

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
	users := h.repo.FindAll()
	api.SendJSON(w, api.Response{Data: users}, http.StatusOK)

}

func (h *Handler) FindById(w http.ResponseWriter, r *http.Request){
	
}

func (h *Handler) Insert(w http.ResponseWriter, r *http.Request){
	var body CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil{
		slog.Error("failed to decode body", "error", err)
		api.SendJSON(w, api.Response{Error: "invalid request body"}, http.StatusBadRequest)
		return
	}

	newUser, err := h.repo.Insert(body.Name,body.Email)
	if err != nil {
		slog.Error("user already exists", "error", err.Error())
		api.SendJSON(w, api.Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	api.SendJSON(w, api.Response{Data: newUser}, http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request){

}

