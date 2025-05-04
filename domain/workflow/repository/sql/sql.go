package sql

import (
	"context"

	"github.com/ideagate/server-controller/domain/workflow/model"
)

type Repository interface {
	GetListWorkflow(ctx context.Context, req *GetListWorkflowRequest) ([]*model.Workflow, error)
	GetWorkflow(ctx context.Context, req *GetWorkflowRequest) (*model.Workflow, error)
	CreateWorkflow(ctx context.Context, req *CreateWorkflowRequest) error
	UpdateWorkflow(ctx context.Context, req *UpdateWorkflowRequest) error
	DeleteWorkflow(ctx context.Context, req *DeleteWorkflowRequest) error
}

type GetListWorkflowRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
	Limit         *int
}

type GetWorkflowRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
	Version       int64
}

type CreateWorkflowRequest struct {
	Workflow *model.Workflow
}

type UpdateWorkflowRequest struct {
	Workflow *model.Workflow
}

type DeleteWorkflowRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
	Version       int64
}
