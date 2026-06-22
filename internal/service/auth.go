package service

import (
	"crypto/sha256"
	"main/package/dto"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Postgres *sqlx.DB
}

func (as *AuthService) SignUp(req dto.CreateUserDto) (dto.CreateAuthSessionResponse, error) {
	var id uuid.UUID
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return dto.CreateAuthSessionResponse{}, err
	}
	if err := as.Postgres.Get(&id, "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING *", req.Email, string(passwordHash)); err != nil {
		return dto.CreateAuthSessionResponse{}, err
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{})
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{})
	accessToken, err := access.SignedString("123")
	if err != nil {
		return dto.CreateAuthSessionResponse{}, err
	}
	refreshToken, err := refresh.SignedString("123")
	refreshTokenHash := sha256.New().Sum([]byte(refreshToken))
	as.Postgres.Exec("INSERT INTO auth_sessions (user_id, token_hash) VALUES ($1, $2)", id, refreshTokenHash)
	return dto.CreateAuthSessionResponse{
		AccessToken: accessToken,
	}, nil
}

func (as *AuthService) SignIn(req dto.SignInDto) (dto.CreateAuthSessionResponse, error) {
	var id dto.UserDto
	if err := as.Postgres.Get(&id, "SELECT id FROM users WHERE id = $1", req.Email); err != nil {
		return dto.CreateAuthSessionResponse{}, err
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{})
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{})
	accessToken, err := access.SignedString("123")
	if err != nil {
		return dto.CreateAuthSessionResponse{}, err
	}
	refreshToken, err := refresh.SignedString("123")
	refreshTokenHash := sha256.New().Sum([]byte(refreshToken))
	as.Postgres.Exec("INSERT INTO auth_sessions (user_id, token_hash) VALUES ($1, $2)", id, refreshTokenHash)
	return dto.CreateAuthSessionResponse{
		AccessToken: accessToken,
	}, nil
}

func (as *AuthService) LogOut(userId uuid.UUID) {
}

func (as *AuthService) Revoke(tokenHash string) {
}
