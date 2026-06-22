package dto

import (
	"time"

	"github.com/google/uuid"
)

// users

type UserDto struct {
	Id           uuid.UUID `json:"id" db:"id"`
	Role         string    `json:"role" db:"role"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"passwordHash" db:"password_hash"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

type CreateUserDto struct {
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"passwordHash" db:"password_hash"`
}

// products

type ProductDto struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Price     uint64    `json:"price" db:"price"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type CreateProductDto struct {
	Name  string `json:"name" db:"name"`
	Price uint64 `json:"price" db:"price"`
}

// warehouses

type WarehouseDto struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Location  string    `json:"location" db:"location"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type CreateWarehouseDto struct {
	Location string `json:"location" db:"location"`
	Status   string `json:"status" db:"status"`
}

type AddProductsToWarehouseDto struct {
	ProductIds []uuid.UUID `json:"productIds"`
}

// orders

type OrderDto struct {
	Id         uuid.UUID   `json:"id" db:"id"`
	UserId     uuid.UUID   `json:"userId" db:"user_id"`
	ProductIds []uuid.UUID `json:"productIds" db:"product_ids"`
	Status     string      `json:"status" db:"status"`
	CreatedAt  time.Time   `json:"createdAt" db:"created_at"`
}

type CreateOrderDto struct {
	UserId     uuid.UUID   `json:"userId" db:"user_id"`
	ProductIds []uuid.UUID `json:"productIds" db:"product_ids"`
}

// auth session

type AuthSessionDto struct {
	Id        uuid.UUID `json:"id" db:"id"`
	UserId    uuid.UUID `json:"userId" db:"user_id"`
	TokenHash string    `json:"tokenHash" db:"token_hash"`
	RevokedAt time.Time `json:"revokedAt" db:"revoked_at"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type CreateAuthSessionResponse struct {
	AccessToken string `json:"refreshToken"`
}

type SignInDto struct {
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"passwordHash" db:"password_hash"`
}
