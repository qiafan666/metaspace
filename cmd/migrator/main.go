package main

import (
	"context"
	"embed"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

//go:embed "migrations"
var migrations embed.FS

var log = logrus.New()

func main() {
	var err error

	log.Infoln("Application Initializing")
	defer log.Infoln("Completed")

	// Config
	// export DB_URL='mysql://root:!devpass123456@tcp(3.20.122.137:3306)/metaspacetest?parseTime=true&charset=utf8mb4'
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalf("Missing database connection string. Please set ENV DB_URL")
	}
	fmt.Println(dbUrl)

	if !strings.Contains(dbUrl, "x-migrations-table") {
		fmt.Println("Appending x-migrations-table to DB URL")
		dbUrl += "&x-migrations-table=gameserver_schema_migrations"
	}

	if err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	// Main context with stopping capabilities
	mainCtx, mainCtxCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer mainCtxCancel()

	if !strings.Contains(dbUrl, "multiStatements=true") {
		log.Info("Appending multiStatements=true to DB URL")
		dbUrl += "&multiStatements=true"
	}

	d, err := iofs.New(migrations, "migrations")
	m, err := migrate.NewWithSourceInstance("iofs", d, dbUrl)
	if err != nil {
		log.Fatalf("Error creating migrator: %v", err)
	}

	version, _, _ := m.Version()
	log.Infof("CURRENT MIGRATION VERSION: %d", version)

	done := make(chan error)
	go func() {
		//done <- m.Force(int(version + 1))
		done <- m.Up()
	}()

	select {
	case <-mainCtx.Done():
		log.Infof("Stopping migration")
		m.GracefulStop <- true

		err := <-done
		if err != nil {
			log.Fatalf("Error migrating database: %v", err)
		}
	case err := <-done:
		switch err {
		case nil:
			log.Infof("MIGRATION COMPLETED")
		case migrate.ErrNoChange:
			log.Infof("MIGRATION NOT NEEDED")
		default:
			log.Fatalf("Error migrating database: %v", err)
		}
	}

	version, _, _ = m.Version()
	log.Infof("CURRENT MIGRATION VERSION: %d", version)

	sErr, dbErr := m.Close()
	if sErr != nil {
		log.Errorf("Error closing source: %v", sErr)
	}
	if dbErr != nil {
		log.Errorf("Error closing database: %v", dbErr)
	}
}
