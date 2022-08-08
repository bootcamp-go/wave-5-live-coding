package products

import (
	"products_project/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Delete(id int) (domain.Product, error)
	UpdateFields(id int, nombre string, precio int) (domain.Product, error)
	GetById(id int) (domain.Product, error)
	GetByName(nombre string) ([]domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Product, error) {
	return s.repository.GetAll()
}

func (s *service) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	return s.repository.Store(nombre, color, precio, stock, codigo, publicado, fecha)
}

func (s *service) Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	return s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fecha)
}

func (s *service) Delete(id int) (domain.Product, error) {
	return s.repository.Delete(id)
}

func (s *service) UpdateFields(id int, nombre string, precio int) (domain.Product, error) {
	return s.repository.UpdateFields(id, nombre, precio)
}

func (s *service) GetById(id int) (domain.Product, error) {
	return s.repository.GetById(id)
}

func (s *service) GetByName(nombre string) ([]domain.Product, error) {
	return s.repository.GetByName(nombre)
}
