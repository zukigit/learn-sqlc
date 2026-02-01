package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zukigit/learn-sqlc/db"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:zabbix@rocky10:5432/test")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	quries := db.New(conn)

	err = quries.DeleteAuthorWithName(ctx, "zuki")
	if err != nil {
		return err
	}

	_, err = quries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "zuki",
		Bio: pgtype.Text{
			String: "former techinal lead in DAT",
			Valid:  true,
		},
	})
	if err != nil {
		return nil
	}

	authors, err := quries.ListAuthors(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Authors:")
	for _, author := range authors {
		fmt.Println(author)
	}

	count, err := quries.CountAuthors(ctx)
	if err != nil {
		return err
	}
	fmt.Println("authors count", count)

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	quriesTx := quries.WithTx(tx)

	quriesTx.DeleteAuthorWithName(ctx, "wai")

	quriesTx.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "wai",
	})

	return tx.Commit(ctx)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
