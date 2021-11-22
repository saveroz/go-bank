package main

import (
	"log"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/saveroz/go-bank/internal/postgres"
	"github.com/saveroz/go-bank/internal/rest"
	"github.com/saveroz/go-bank/services/account"
)

func main() {
	env := godotenv.Load(".env")
	if env != nil {
		log.Fatalf("Error loading .env file")
	}
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	db, err := sql.Open("postgres", POSTGRES_DB)
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	g := e.Group("")

	g.Use(
		sentryecho.New(sentryecho.Options{}),
	)

	accountRepository := postgres.NewAccountRepository(db)
	accountService := account.NewService(accountRepository)
	rest.InitAccountHandler(g, accountService)

	e.Logger.Fatal(e.Start(":3000"))

	if err != nil {
		log.Fatal(err)
	}
}
