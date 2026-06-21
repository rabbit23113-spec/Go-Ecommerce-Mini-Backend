package service

import "github.com/jmoiron/sqlx"

type AuthService struct {
	Postgres *sqlx.DB
}

func (as *AuthService) SignUp() {
}

func (as *AuthService) SignIn() {
}

func (as *AuthService) LogOut() {
}

func (as *AuthService) Revoke() {
}
