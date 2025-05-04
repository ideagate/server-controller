package workflow

import (
	"context"

	"github.com/ideagate/server-controller/domain/workflow/repository/sql"
	"github.com/ideagate/server-controller/domain/workflow/usecase"
	"gorm.io/gorm"
)

type Domain interface {
	GetWorkflows(ctx context.Context, req *usecase.GetWorkflowsRequest) (*usecase.GetWorkflowsResponse, error)
	GetWorkflow(ctx context.Context, req *usecase.GetWorkflowRequest) (*usecase.GetWorkflowResponse, error)
	CreateWorkflow(ctx context.Context, req *usecase.CreateWorkflowRequest) (*usecase.CreateWorkflowResponse, error)
	UpdateWorkflow(ctx context.Context, req *usecase.UpdateWorkflowRequest) (*usecase.UpdateWorkflowResponse, error)
	DeleteWorkflow(ctx context.Context, req *usecase.DeleteWorkflowRequest) error
}

var _ Domain = &usecase.Usecase{}

func New(db *gorm.DB) Domain {
	repoSql := sql.New(db)
	return usecase.New(repoSql)
}
