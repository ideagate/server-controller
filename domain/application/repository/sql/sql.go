package sql

import (
	"context"

	"github.com/ideagate/server-controller/domain/application/model"
)

type ApplicationSQLRepository interface {
	GetListApplication(ctx context.Context, req *GetListApplicationRequest) ([]*model.Application, error)
	CreateApplication(ctx context.Context, req *CreateApplicationRequest) error
	UpdateApplication(ctx context.Context, req *UpdateApplicationRequest) error
	DeleteApplication(ctx context.Context, req *DeleteApplicationRequest) error
}

type GetListApplicationRequest struct {
	ProjectID     *string
	ApplicationID *string
	Limit         int
}

type CreateApplicationRequest struct {
	ProjectID     string
	ApplicationID string
	Name          string
}

type UpdateApplicationRequest struct {
	ProjectID     string
	ApplicationID string
	Data          model.Application
	Fields        []string
}

type DeleteApplicationRequest struct {
	ProjectID     string
	ApplicationID string
}
