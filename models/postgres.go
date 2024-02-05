package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}

func DevConfig() PostgresConfig {
	return PostgresConfig{
		Username: "user",
		Password: "password",
		Host:     "localhost",
		Port:     "1111",
		DBName:   "phogo",
	}
}

func Open(cfg PostgresConfig) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cfg.String())

	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	return conn, err
}
