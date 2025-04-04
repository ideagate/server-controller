package usecase

import (
	"context"

	pbapplication "github.com/bayu-aditya/ideagate/backend/model/gen-go/core/application"
)

type ApplicationUsecase interface {
	GetListApplication(ctx context.Context, req *GetListApplicationRequest) (*GetListApplicationResponse, error)
	CreateApplication(ctx context.Context, req *CreateApplicationRequest) error
	UpdateApplication(ctx context.Context, req *UpdateApplicationRequest) error
	DeleteApplication(ctx context.Context, req *DeleteApplicationRequest) error
}

type GetListApplicationRequest struct {
	ProjectID     string
	ApplicationID *string
}

type GetListApplicationResponse struct {
	Applications []*pbapplication.Application
}

type CreateApplicationRequest struct {
	ProjectID     string
	ApplicationID string
	Name          string
}

type UpdateApplicationRequest struct {
	ProjectID     string
	ApplicationID string
	Values        map[string]any
}

type DeleteApplicationRequest struct {
	ProjectID     string
	ApplicationID string
}
