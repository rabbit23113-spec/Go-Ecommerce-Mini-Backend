package service

import (
	"main/package/dto"

	"github.com/jmoiron/sqlx"
)

type WarehouseService struct {
	Postgres *sqlx.DB
}

func (ws *WarehouseService) Find() ([]dto.WarehouseDto, error) {
	var warehouses []dto.WarehouseDto
	if err := ws.Postgres.Get(&warehouses, "SELECT * FROM warehouses"); err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (ws *WarehouseService) FindByProductId(productId string) ([]dto.WarehouseDto, error) {
	var warehouses []dto.WarehouseDto
	if err := ws.Postgres.Get(&warehouses, "SELECT * FROM warehouses WHERE $1 IN product_ids", productId); err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (ws *WarehouseService) Create(req dto.CreateWarehouseDto) (dto.WarehouseDto, error) {
	var warehouse dto.WarehouseDto
	if err := ws.Postgres.Get(&warehouse, "INSERT INTO warehouses (location, status) VALUES ($1, $2) RETURNING *", req.Location, req.Status); err != nil {
		return dto.WarehouseDto{}, err
	}
	return warehouse, nil
}

func (ws *WarehouseService) Update() {
}

func (ws *WarehouseService) Delete(id string) error {
	_, err := ws.Postgres.Exec("DELETE FROM warehouses WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
