package api

import (
	"github.com/google/uuid"
)

type ProductServiceInterface interface {
	CreateProduct(name string) DomainModel
	RemoveProduct(id uuid.UUID)
	FindAllProducts() []DomainModel
}

type ProductService struct {
	repository ProductRepositoryInterface
}

func NewProductService(repository ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{repository: repository}
}

func (s *ProductService) CreateProduct(name string) DomainModel {
	dm := DomainModel{ID: uuid.New(), Name: name}
	s.repository.add(dm)
	return dm
}

func (s *ProductService) RemoveProduct(id uuid.UUID) {
	s.repository.remove(id)
}

func (s *ProductService) FindAllProducts() []DomainModel {
	return s.repository.list()
}
