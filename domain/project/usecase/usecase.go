package usecase

import (
	"context"

	pbproject "github.com/bayu-aditya/ideagate/backend/model/gen-go/core/project"
)

type ProjectUsecase interface {
	GetListProject(ctx context.Context, req *GetListProjectRequest) (*GetListProjectResponse, error)
	CreateProject(ctx context.Context, req *CreateProjectRequest) error
	UpdateProject(ctx context.Context, req *UpdateProjectRequest) error
	DeleteProject(ctx context.Context, req *DeleteProjectRequest) error
}

type GetListProjectRequest struct{}

type GetListProjectResponse struct {
	Projects []*pbproject.Project
}

type CreateProjectRequest struct {
	UserID    int64
	ProjectID string
	Name      string
}

type UpdateProjectRequest struct {
	ProjectID string
	Values    map[string]any
}

type DeleteProjectRequest struct {
	ProjectID string
}
