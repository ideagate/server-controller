package sql

import (
	"context"

	"github.com/ideagate/server-controller/domain/entrypoint/model"
)

type Repository interface {
	GetListEntrypoint(ctx context.Context, req *GetListEntrypointRequest) ([]*model.Entrypoint, error)
	GetEntrypoint(ctx context.Context, req *GetEntrypointRequest) (*model.Entrypoint, error)
	CreateEntrypoint(ctx context.Context, req *CreateEntrypointRequest) error
	DeleteEntrypoint(ctx context.Context, req *DeleteEntrypointRequest) error
}

type GetListEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	Type          *string
}

type GetEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
}

type CreateEntrypointRequest struct {
	Entrypoint *model.Entrypoint
}

type DeleteEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
}
