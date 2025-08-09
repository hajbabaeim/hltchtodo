package app

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"path/filepath"
)

func (a *App) initLogger() {
	logger := logrus.New()
	level, err := logrus.ParseLevel(a.config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	// Add custom fields
	logger = logger.WithFields(logrus.Fields{
		"service": a.config.Name,
		"version": a.config.Version,
		"env":     a.config.Environment,
	}).Logger
	logrus.SetLevel(level)
	a.logger = logger
	a.logger.Info("Logger initialized successfully")
}

func (a *App) initPostgres() {
	db, err := gorm.Open(postgres.Open(a.dBConnectionString()))
	if err != nil {
		a.panicOnError(err)
	}
	a.sqlMigrate()
	a.postgres = db
}

func (a *App) dBConnectionString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable&client_encoding=UTF8",
		a.config.Database.Username,
		a.config.Database.Password,
		a.config.Database.Host,
		a.config.Database.Port,
		a.config.Database.DBName,
	)
}

func (a *App) sqlMigrate() {
	path, _ := filepath.Abs("data/postgres/migrations")
	migrationDir := &migrate.FileMigrationSource{
		Dir: path,
	}
	executor := &migrate.MigrationSet{
		TableName: "migrations",
	}
	db, _ := a.postgres.DB()
	n, err := executor.Exec(db, "postgres", migrationDir, migrate.Up)
	a.panicOnError(err)
	fmt.Printf("[POSTGRES] applied %d migrations!\n", n)
}
