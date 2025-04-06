package migration

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ideagate/core/config"
	"github.com/ideagate/core/utils/log"
	"github.com/urfave/cli/v2"
)

var (
	folderName = "migrations"

	ActionCreateUsage = "migrate the database schema"
	ActionCreateFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Usage:    "migration name",
			Required: true,
		},
	}
)

func ActionCreate(c *cli.Context) error {
	name := c.String("name")

	fileNameUp, fileNameDown := generateMigrationFile(folderName, name)

	// Create migration folder if not exist
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		if err := os.Mkdir(folderName, os.ModePerm); err != nil {
			return err
		}
	}

	// Create migration file for up and down
	if _, err := os.Create(fileNameUp); err != nil {
		return err
	}

	if _, err := os.Create(fileNameDown); err != nil {
		_ = os.Remove(fileNameUp)
		return err
	}

	return nil
}

func ActionUp(_ *cli.Context) error {
	migration, err := initMigration()
	if err != nil {
		return err
	}

	return migration.Up()
}

func ActionDown(_ *cli.Context) error {
	migration, err := initMigration()
	if err != nil {
		return err
	}

	return migration.Steps(-1)
}

func initMigration() (*migrate.Migrate, error) {
	if err := config.Load("."); err != nil {
		log.Panic(err.Error())
	}
	cfg := config.Get()
	if cfg.Postgres == nil {
		log.Panic("postgres config is not set")
	}

	dbType := "postgres"
	db, err := sql.Open(dbType, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB))
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()

	sourceFile := fmt.Sprintf("file://%s", folderName)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migration, err := migrate.NewWithDatabaseInstance(sourceFile, dbType, driver)
	if err != nil {
		return nil, err
	}

	return migration, nil
}

func generateMigrationFile(folderName, name string) (string, string) {
	nowSec := time.Now().Unix()

	fileNameUp := fmt.Sprintf("%s/%d_%s.up.sql", folderName, nowSec, name)
	fileNameDown := fmt.Sprintf("%s/%d_%s.down.sql", folderName, nowSec, name)

	return fileNameUp, fileNameDown
}
