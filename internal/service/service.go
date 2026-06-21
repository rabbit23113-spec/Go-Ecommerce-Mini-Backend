package service

import (
	"github.com/jmoiron/sqlx"
)

type Service struct {
	AuthService      *AuthService
	OrderService     *OrderService
	ProductService   *ProductService
	WarehouseService *WarehouseService
	UserService      *UserService
	Postgres         *sqlx.DB
}

func (s *Service) InitService() *Service {
	authService := AuthService{s.Postgres}
	orderService := OrderService{s.Postgres}
	productService := ProductService{s.Postgres}
	warehouseService := WarehouseService{s.Postgres}
	userService := UserService{s.Postgres}
	return &Service{
		AuthService:      &authService,
		OrderService:     &orderService,
		ProductService:   &productService,
		WarehouseService: &warehouseService,
		UserService:      &userService,
		Postgres:         s.Postgres,
	}
}
