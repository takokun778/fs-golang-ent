package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"fs-golang-ent/ent"
	"fs-golang-ent/ent/user"

	"github.com/google/uuid"
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

	ctx := context.Background()

	uuid, _ := uuid.NewRandom()

	birthday, _ := time.Parse("2006-01-02", "2000-04-18")

	cU, err := client.User.
		Create().
		SetID(uuid).
		SetFirstName("玲").
		SetLastName("大園").
		SetBirthday(birthday).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	fmt.Println(cU)

	qU, err := client.User.
		Query().
		Where(user.IDEQ(cU.ID)).
		Only(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	fmt.Println(qU)

	uU, err := client.User.
		UpdateOneID(qU.ID).
		SetFirstName("れい").
		SetLastName("おおぞの").
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	fmt.Println(uU)

	err = client.User.
		DeleteOneID(uU.ID).
		Exec(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
}
