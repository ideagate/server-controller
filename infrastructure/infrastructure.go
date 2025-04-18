package infrastructure

import (
	"github.com/ideagate/core/config"
	"gorm.io/gorm"
)

type Infrastructure struct {
	Postgres *gorm.DB
}

func NewInfrastructure(cfg *config.Config) (*Infrastructure, error) {
	postgresConn, err := initializePostgres(cfg)
	if err != nil {
		return nil, err
	}

	return &Infrastructure{
		Postgres: postgresConn,
	}, nil
}
