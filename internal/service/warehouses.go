package service

import "github.com/jmoiron/sqlx"

type WarehouseService struct {
	Postgres *sqlx.DB
}

func (ws *WarehouseService) Find() {
}

func (ws *WarehouseService) FindByProductId() {
}

func (ws *WarehouseService) Create() {
}

func (ws *WarehouseService) Update() {
}

func (ws *WarehouseService) Delete() {
}
