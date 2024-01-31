package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

//type Shelter struct {
//	Name    string
//	Address string
//}
//
//type Pet struct {
//	Name        string
//	Age         int
//	Weight      float64
//	Height      float64
//	SocialMedia map[string]string
//	Hobbies     []string
//	Shelter     Shelter
//}

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (c Config) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}

func main() {
	cfg := Config{
		User:     "user",
		Password: "password",
		Host:     "localhost",
		Port:     "1111",
		DBName:   "phogo",
	}
	conn, err := pgx.Connect(context.Background(), cfg.String())
	defer conn.Close(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
    		id SERIAL PRIMARY KEY,
    		name TEXT,
    		email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
    		id SERIAL PRIMARY KEY,
    		user_id INT NOT NULL,
    		amount INT,
    		description TEXT
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected")
}
