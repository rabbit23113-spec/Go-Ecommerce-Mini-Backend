package service

import "github.com/jmoiron/sqlx"

type UserService struct {
	Postgres *sqlx.DB
}

func (us *UserService) Me() {
}

func (us *UserService) FindAll() {
}

func (us *UserService) FindById() {
}

func (us *UserService) FindByEmail() {
}

func (us *UserService) Create() {
}

func (us *UserService) Update() {
}

func (us *UserService) Delete() {
}
