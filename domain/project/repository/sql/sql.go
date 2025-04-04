package sql

import (
	"context"

	"github.com/bayu-aditya/ideagate/backend/server/controller/domain/project/model"
)

type ProjectSQLRepository interface {
	GetListProject(ctx context.Context, req *GetListProjectRequest) ([]*model.Project, error)
	CreateProject(ctx context.Context, req *CreateProjectRequest) error
	UpdateProject(ctx context.Context, req *UpdateProjectRequest) error
	DeleteProject(ctx context.Context, req *DeleteProjectRequest) error

	CreateProjectUser(ctx context.Context, req *CreateProjectUserRequest) error
}

type GetListProjectRequest struct{}

type CreateProjectRequest struct {
	ProjectID string
	Name      string
}

type UpdateProjectRequest struct {
	ProjectID string
	Data      model.Project
	Fields    []string
}

type DeleteProjectRequest struct {
	ProjectID string
}
type CreateProjectUserRequest struct {
	ProjectID string
	UserID    int64
}
