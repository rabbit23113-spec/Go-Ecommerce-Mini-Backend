package service

import "github.com/jmoiron/sqlx"

type OrderService struct {
	Postgres *sqlx.DB
}

func (ors *OrderService) FindAll() {
}

func (ors *OrderService) FindByUserId() {
}
func (ors *OrderService) FindByProductId() {
}

func (ors *OrderService) Create() {
}

func (ors *OrderService) Update() {
}

func (ors *OrderService) Delete() {
}
