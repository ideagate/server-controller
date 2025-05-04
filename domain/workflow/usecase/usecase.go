package usecase

import (
	"context"

	"github.com/ideagate/core/utils"
	pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"
	modelworkflow "github.com/ideagate/server-controller/domain/workflow/model"
	"github.com/ideagate/server-controller/domain/workflow/repository/sql"
	"github.com/ideagate/server-controller/model"
)

func New(repoSql sql.Repository) *Usecase {
	return &Usecase{
		repoSql: repoSql,
	}
}

type Usecase struct {
	repoSql sql.Repository
}

func (u *Usecase) GetWorkflows(ctx context.Context, req *GetWorkflowsRequest) (*GetWorkflowsResponse, error) {
	if req.ProjectID == "" {
		return nil, model.ErrProjectIDRequired
	}

	if req.ApplicationID == "" {
		return nil, model.ErrApplicationIDRequired
	}

	if req.EntrypointID == "" {
		return nil, model.ErrEntrypointIDRequired
	}

	resultRepo, err := u.repoSql.GetListWorkflow(ctx, &sql.GetListWorkflowRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
		EntrypointID:  req.EntrypointID,
	})
	if err != nil {
		return nil, err
	}

	result := &GetWorkflowsResponse{
		Workflows: make([]*pbendpoint.Workflow, len(resultRepo)),
	}

	for i := 0; i < len(resultRepo); i++ {
		result.Workflows[i] = resultRepo[i].ToProto()
		result.Workflows[i].Edges = nil
		result.Workflows[i].Steps = nil
	}

	return result, nil
}

func (u *Usecase) GetWorkflow(ctx context.Context, req *GetWorkflowRequest) (*GetWorkflowResponse, error) {
	if req.ProjectID == "" {
		return nil, model.ErrProjectIDRequired
	}

	if req.ApplicationID == "" {
		return nil, model.ErrApplicationIDRequired
	}

	if req.EntrypointID == "" {
		return nil, model.ErrEntrypointIDRequired
	}

	if req.Version == 0 {
		return nil, model.ErrWorkflowVersionRequired
	}

	resultRepo, err := u.repoSql.GetWorkflow(ctx, &sql.GetWorkflowRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
		EntrypointID:  req.EntrypointID,
		Version:       req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &GetWorkflowResponse{
		Workflow: resultRepo.ToProto(),
	}, nil
}

func (u *Usecase) CreateWorkflow(ctx context.Context, req *CreateWorkflowRequest) (*CreateWorkflowResponse, error) {
	if req.ProjectID == "" {
		return nil, model.ErrProjectIDRequired
	}

	if req.ApplicationID == "" {
		return nil, model.ErrApplicationIDRequired
	}

	if req.EntrypointID == "" {
		return nil, model.ErrEntrypointIDRequired
	}

	newWorkflow := &modelworkflow.Workflow{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
		EntrypointID:  req.EntrypointID,
		Version:       1,
	}

	// Copy workflow from existing version if specified
	if req.FromVersion != nil {
		resultRepo, err := u.repoSql.GetWorkflow(ctx, &sql.GetWorkflowRequest{
			ProjectID:     req.ProjectID,
			ApplicationID: req.ApplicationID,
			EntrypointID:  req.EntrypointID,
			Version:       *req.FromVersion,
		})
		if err != nil {
			return nil, err
		}
		newWorkflow = resultRepo
	}

	// Determine the new version
	resultRepoLast, err := u.repoSql.GetListWorkflow(ctx, &sql.GetListWorkflowRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
		EntrypointID:  req.EntrypointID,
		Limit:         utils.ToPtr(1),
	})
	if err != nil {
		return nil, err
	}
	if len(resultRepoLast) > 0 {
		newWorkflow.Version = resultRepoLast[0].Version + 1
	}

	// Create new version workflow
	if err := u.repoSql.CreateWorkflow(ctx, &sql.CreateWorkflowRequest{
		Workflow: newWorkflow,
	}); err != nil {
		return nil, err
	}

	return &CreateWorkflowResponse{
		Version: newWorkflow.Version,
	}, nil
}

func (u *Usecase) UpdateWorkflow(ctx context.Context, req *UpdateWorkflowRequest) (*UpdateWorkflowResponse, error) {
	if req.Workflow.GetProjectId() == "" {
		return nil, model.ErrProjectIDRequired
	}

	if req.Workflow.GetApplicationId() == "" {
		return nil, model.ErrApplicationIDRequired
	}

	if req.Workflow.GetEntrypointId() == "" {
		return nil, model.ErrEntrypointIDRequired
	}

	if req.Workflow.Version == 0 {
		return nil, model.ErrWorkflowVersionRequired
	}

	newWorkflow := &modelworkflow.Workflow{}
	newWorkflow.FromProto(req.Workflow)

	err := u.repoSql.UpdateWorkflow(ctx, &sql.UpdateWorkflowRequest{
		Workflow: newWorkflow,
	})
	if err != nil {
		return nil, err
	}

	return &UpdateWorkflowResponse{}, nil
}

func (u *Usecase) DeleteWorkflow(ctx context.Context, req *DeleteWorkflowRequest) error {
	if req.ProjectID == "" {
		return model.ErrProjectIDRequired
	}

	if req.ApplicationID == "" {
		return model.ErrApplicationIDRequired
	}

	if req.EntrypointID == "" {
		return model.ErrEntrypointIDRequired
	}

	if req.Version == 0 {
		return model.ErrWorkflowVersionRequired
	}

	err := u.repoSql.DeleteWorkflow(ctx, &sql.DeleteWorkflowRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
		EntrypointID:  req.EntrypointID,
		Version:       req.Version,
	})
	if err != nil {
		return err
	}

	return nil
}
