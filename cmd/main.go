package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"path/filepath"
	"sqlc-demo/config"
	"sqlc-demo/internal/db"
)

func main() {
	configPath := filepath.Join("config", "config.yml")

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	ctx := context.TODO()
	conn, err := pgx.Connect(ctx, cfg.PostgreSQL.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		_ = conn.Close(ctx)
	}(conn, ctx)

	qry := db.New(conn)

	books, err := qry.GetAllBooks(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("所有图书:", books)

	books, err = qry.GetBooksByCategory(ctx, db.BookCategoryEducation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("图书:", books)

	book, err := qry.GetBook(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID为1的图书:", book)
}
