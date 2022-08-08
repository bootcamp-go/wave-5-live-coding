package products

import (
	"database/sql"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
)

type Repository interface {
	Store(domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetById(id uint64) (domain.Product, error)
	UpdateTotal(domain.Product) (domain.Product, error)
	UpdatePartial(domain.Product) (domain.Product, error)
	Delete(id uint64) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(product domain.Product) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(
		"INSERT INTO PRODUCTS " +
			"(Name, Color, Price, Stock, Code, Published, Created_at)" +
			" VALUES (?, ?, ?, ?, ?, ?, CURDATE())",
	)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	sqlRes, err := stmt.Exec(product.Name, product.Color, product.Price, product.Stock, product.Code, product.Published)
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, err := sqlRes.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}
	product.Id = uint64(insertedId)
	return product, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *repository) GetById(id uint64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) UpdateTotal(domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) UpdatePartial(domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) Delete(id uint64) (domain.Product, error) {
	return domain.Product{}, nil
}

func partialUpdate(oldProduct domain.Product, newProduct domain.Product) domain.Product {
	if newProduct.Name != "" {
		oldProduct.Name = newProduct.Name
	}

	if newProduct.Color != "" {
		oldProduct.Color = newProduct.Color
	}

	if newProduct.Price != 0 {
		oldProduct.Price = newProduct.Price
	}

	oldProduct.Stock = newProduct.Stock

	if newProduct.Code != "" {
		oldProduct.Code = newProduct.Code
	}
	return oldProduct
}
