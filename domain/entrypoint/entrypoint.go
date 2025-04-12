package entrypoint

import (
	"context"

	"github.com/ideagate/server-controller/domain/entrypoint/repository/sql"
	"github.com/ideagate/server-controller/domain/entrypoint/usecase"
	"gorm.io/gorm"
)

type Domain interface {
	GetListEntrypoint(ctx context.Context, req *usecase.GetListEntrypointRequest) (*usecase.GetListEntrypointResponse, error)
	GetEntrypoint(ctx context.Context, req *usecase.GetEntrypointRequest) (*usecase.GetEntrypointResponse, error)
	CreateEntrypoint(ctx context.Context, req *usecase.CreateEntrypointRequest) error
	DeleteEntrypoint(ctx context.Context, req *usecase.DeleteEntrypointRequest) error
}

var _ Domain = &usecase.Usecase{}

func New(db *gorm.DB) Domain {
	repoSql := sql.New(db)
	return usecase.New(repoSql)
}
