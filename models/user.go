package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *pgx.Conn
}

func (us *UserService) Create(email string, password string) (*User, error) {
	email = strings.ToLower(email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	row := us.DB.QueryRow(context.Background(), `INSERT INTO users(email, password_hash) VALUES ($1, $2) RETURNING id`, email, hashedPassword)

	user := User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	err = row.Scan(&user.ID)

	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}