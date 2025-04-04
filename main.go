package main

import (
	"os"

	"github.com/bayu-aditya/ideagate/backend/core/utils/log"
	"github.com/bayu-aditya/ideagate/backend/server/controller/app/grpc"
	"github.com/bayu-aditya/ideagate/backend/server/controller/app/migration"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name: "Server Controller",
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "start grpc server",
				Action: grpc.Action,
			},
			{
				Name:  "migrate",
				Usage: "migrate the database schema",
				Subcommands: []*cli.Command{
					{
						Name:   "create",
						Usage:  migration.ActionCreateUsage,
						Flags:  migration.ActionCreateFlags,
						Action: migration.ActionCreate,
					},
					{
						Name:   "up",
						Usage:  "run the migration files",
						Action: migration.ActionUp,
					},
					{
						Name:   "down",
						Usage:  "rollback the migration",
						Action: migration.ActionDown,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
