package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router

	// this should be from a repository not on a server object
	repository     ProductRepositoryInterface
	productService ProductServiceInterface
}

func NewServer() *Server {

	repository := new(InMemoryRepository)
	productService := NewProductService(repository)
	s := &Server{

		Router:         mux.NewRouter(),
		repository:     repository,
		productService: productService,
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/", s.listDomainObjs()).Methods("GET")
	s.HandleFunc("/", s.createDomainObjs()).Methods("POST")
	s.HandleFunc("/{id}", s.removeDomainObjs()).Methods("DELETE")
}

func (s *Server) listDomainObjs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Conent-Type", "application/json")
		var objects []DomainModel = s.productService.FindAllProducts()

		if err := json.NewEncoder(w).Encode(objects); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("COTO", err)
			return
		}
	}
}

func (s *Server) createDomainObjs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var schema CreateObjectSchema
		if err := json.NewDecoder(r.Body).Decode(&schema); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var dm DomainModel = s.productService.CreateProduct(schema.Name)
		w.Header().Set("Conent-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dm); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeDomainObjs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		s.repository.remove(id)
	}
}
