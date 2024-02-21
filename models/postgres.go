package models

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"io/fs"
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

func Open(cfg PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.String())

	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	return db, err
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect(string(goose.DialectPostgres))

	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	err = goose.Up(db, dir)

	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	return nil
}

func MigrateFS(db *sql.DB, fs fs.FS, dir string) error {
	goose.SetBaseFS(fs)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, dir)
}
