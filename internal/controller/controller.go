package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"lesson30/internal/cases"
	"lesson30/internal/entity"
	"net/http"
	"strconv"
)

type Controller struct {
	usecase cases.Usecase
}

func Build(r *chi.Mux, usecase cases.Usecase) {
	ctr := NewController(usecase)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", ctr.CreateUser)
		r.Get("/", ctr.GetAll)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", ctr.Get)
			r.Patch("/", ctr.NewAge)
			r.Delete("/", ctr.DeleteUser)

			r.Route("/friends", func(r chi.Router) {
				r.Get("/", ctr.GetFriendList)
				r.Put("/", ctr.AddFriend)

				r.Route("/{idt}", func(r chi.Router) {
					r.Delete("/", ctr.DelUserFromFL)
				})
			})
		})
	})
}

func NewController(usecase cases.Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (s *Controller) AddFriend(w http.ResponseWriter, r *http.Request) {
	user := &entity.User{}
	idURL := chi.URLParam(r, "id")
	if idURL == "" {
		w.WriteHeader(404)
		w.Write([]byte("User not found."))
		return
	}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid friends name."))
		return
	}

	id, err := strconv.Atoi(idURL)
	response, err := s.usecase.AddFriend(id, user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User " + response + " friendlist`s was updated."))
	return
}

func (s *Controller) NewAge(w http.ResponseWriter, r *http.Request) {
	user := &entity.User{}
	idURL := chi.URLParam(r, "id")
	if idURL == "" {
		w.WriteHeader(404)
		w.Write([]byte("invalid id"))
		return
	}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid age"))
		return
	}

	id, _ := strconv.Atoi(idURL)
	err = s.usecase.NewAge(id, user.Age)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Controller) DelUserFromFL(w http.ResponseWriter, r *http.Request) {
	idURLf := chi.URLParam(r, "id")
	idURLt := chi.URLParam(r, "idt")
	if idURLf == "" || idURLt == "" {
		w.WriteHeader(404)
		w.Write([]byte("User not found."))
		return
	}
	id, err1 := strconv.Atoi(idURLf)
	idt, err2 := strconv.Atoi(idURLt)
	if err1 != nil || err2 != nil {
		w.WriteHeader(404)
		w.Write([]byte("Invalid id."))
		return
	}
	err := s.usecase.DeleteUserFromFL(id, idt)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (s *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idURL := chi.URLParam(r, "id")
	if idURL == "" {
		w.WriteHeader(400)
		w.Write([]byte("Invalid id."))
		return
	}
	id, _ := strconv.Atoi(idURL)
	responce, err := s.usecase.DeleteUser(id)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("user not found"))
	}
	w.WriteHeader(200)
	w.Write([]byte("User " + responce + " was deleted"))

}

func (s *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &entity.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	id, err := s.usecase.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("id: " + strconv.Itoa(id)))
	return
}

func (s *Controller) GetFriendList(w http.ResponseWriter, r *http.Request) {
	idURL := chi.URLParam(r, "id")
	if idURL == "" {
		w.WriteHeader(400)
		w.Write([]byte("invalid id."))
		return
	}
	id, err := strconv.Atoi(idURL)
	response, err := s.usecase.GetFriendList(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(response))
}

func (s *Controller) Get(w http.ResponseWriter, r *http.Request) {
	idURL := chi.URLParam(r, "id")
	if idURL == "" {
		w.WriteHeader(400)
		w.Write([]byte("Invalid id."))
		return
	}
	id, err := strconv.Atoi(idURL)
	response, err := s.usecase.SpecUser(id)
	if err != nil || response == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	return
}

func (s *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := s.usecase.AllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	return

}
