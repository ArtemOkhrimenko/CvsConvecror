package main

import (
	"CvsConvecror/internal/app"
	"CvsConvecror/internal/repo"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func main() {

	ctx := context.Background()

	fmt.Println("start")

	db, err := sqlx.ConnectContext(ctx, "pgx", "postgres://postgres:postgres@localhost:5432/monolith_example?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("sqlx.Connect: %s", err.Error()))
	}

	repository := repo.New(db)
	application := app.New(repository)
}
