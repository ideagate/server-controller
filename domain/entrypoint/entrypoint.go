package entrypoint

import (
	"context"

	pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"
	"github.com/ideagate/server-controller/domain/entrypoint/repository/sql"
	"github.com/ideagate/server-controller/domain/entrypoint/usecase"
	"gorm.io/gorm"
)

type Domain interface {
	GetListEntrypoint(ctx context.Context, req *GetListEntrypointRequest) (*GetListEntrypointResponse, error)
	GetEntrypoint(ctx context.Context, req *GetEntrypointRequest) (*GetEntrypointResponse, error)
	CreateEntrypoint(ctx context.Context, req *CreateEntrypointRequest) error
	DeleteEntrypoint(ctx context.Context, req *DeleteEntrypointRequest) error
}

func New(db *gorm.DB) Domain {
	repoSql := sql.New(db)
	return usecase.New(repoSql)
}

type GetListEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	Type          pbendpoint.EndpointType
}

type GetListEntrypointResponse struct {
	Entrypoints []*pbendpoint.Endpoint
}

type GetEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
}

type GetEntrypointResponse struct {
	Entrypoint *pbendpoint.Endpoint
}

type CreateEntrypointRequest struct {
	Entrypoint *pbendpoint.Endpoint
}

type DeleteEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
}
