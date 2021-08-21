package main

import (
	"context"
	"fmt"
	"log"

	"fs-golang-ent/ent"

	_ "github.com/lib/pq"
)

func main() {
	entOptions := []ent.Option{}

	entOptions = append(entOptions, ent.Debug())

	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "ent"
	dbName := "ent"
	dbPass := "ent"
	datasource := fmt.Sprintf("host=%s port=%s user=%s dbname= %s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	client, err := ent.Open("postgres", datasource, entOptions...)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.TODO()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
