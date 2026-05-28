package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")

	fmt.Println("DB URL:", dbURL)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Connected!")
}
