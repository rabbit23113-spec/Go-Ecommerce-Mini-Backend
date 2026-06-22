package service

import (
	"main/package/dto"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Postgres *sqlx.DB
}

func (us *UserService) Me(id string) (dto.UserDto, error) {
	var userFromDb dto.UserDto
	if err := us.Postgres.Get(&userFromDb, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return dto.UserDto{}, err
	}
	return userFromDb, nil
}

func (us *UserService) FindAll() ([]dto.UserDto, error) {
	var usersFromDb []dto.UserDto
	if err := us.Postgres.Get(&usersFromDb, "SELECT * FROM users"); err != nil {
		return nil, err
	}
	return usersFromDb, nil
}

func (us *UserService) FindById(id string) (dto.UserDto, error) {
	var userFromDb dto.UserDto
	if err := us.Postgres.Get(&userFromDb, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return dto.UserDto{}, err
	}
	return userFromDb, nil
}

func (us *UserService) FindByEmail(email string) (dto.UserDto, error) {
	var userFromDb dto.UserDto
	if err := us.Postgres.Get(&userFromDb, "SELECT * FROM users WHERE email = $1", email); err != nil {
		return dto.UserDto{}, err
	}
	return userFromDb, nil
}

func (us *UserService) Create(req dto.CreateUserDto) (dto.UserDto, error) {
	var createdUser dto.UserDto
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserDto{}, err
	}
	if err = us.Postgres.Get(&createdUser, "INSERT INTO users (email, password_hash) VALUES ($1, $2)", req.Email, string(passwordHash)); err != nil {
		return dto.UserDto{}, err
	}
	return createdUser, nil
}

func (us *UserService) Update() {
}

func (us *UserService) Delete(id string) error {
	if _, err := us.Postgres.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
