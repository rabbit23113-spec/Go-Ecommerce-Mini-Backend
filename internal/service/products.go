package service

import (
	"main/package/dto"

	"github.com/jmoiron/sqlx"
)

type ProductService struct {
	Postgres *sqlx.DB
}

func (ps *ProductService) Find() ([]dto.ProductDto, error) {
	var products []dto.ProductDto

	if err := ps.Postgres.Get(&products, "SELECT * FROM products"); err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) Create(req dto.CreateProductDto) (dto.ProductDto, error) {
	var product dto.ProductDto
	if err := ps.Postgres.Get("INSERT INTO products (name, price) VALUES ($1, $2)", req.Name, req.Price); err != nil {
		return dto.ProductDto{}, err
	}
	return product, nil
}

func (ps *ProductService) Update() {

}

func (ps *ProductService) Delete(id string) error {
	if _, err := ps.Postgres.Exec("DELETE FROM products WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
