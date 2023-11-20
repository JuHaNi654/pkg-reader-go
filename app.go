package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JuHaNi654/pkg-reader/pkg/handlers"
	"github.com/JuHaNi654/pkg-reader/pkg/sqlite"
	"github.com/JuHaNi654/pkg-reader/pkg/system"
	_ "github.com/mattn/go-sqlite3"
)

const dbFile = "sqlite.db"

func main() {
	log.Println("Remove existing sqlite file")
	os.Remove(dbFile)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating new sqlite file")
	dbClient := sqlite.NewSQLRepository(db)
	if err := dbClient.Migrate(); err != nil {
		log.Fatal(err)
	}

	log.Println("Adding pkg info to the database")
	pkgs, _ := system.GetPkgs()
	for _, pkg := range pkgs {
		err := dbClient.Insert(pkg)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer stop()

	srv := &http.Server{
		Addr:    ":8888",
		Handler: handlers.GetRouter(dbClient),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully, press Ctrl+c again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
