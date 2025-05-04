package grpc

import (
	"context"

	"github.com/ideagate/core/utils"
	"github.com/ideagate/core/utils/errors"
	pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"
	"github.com/ideagate/model/gen-go/dashboard"
	"github.com/ideagate/server-controller/domain/workflow/usecase"
)

func (s *DashboardServiceServer) GetWorkflows(ctx context.Context, req *dashboard.GetWorkflowsRequest) (*dashboard.GetWorkflowsResponse, error) {
	// Request validation
	if req.GetProjectId() == "" || req.GetApplicationId() == "" || req.GetEntrypointId() == "" {
		return nil, errors.New("project_id, application_id, entrypoint_id are required")
	}

	// Check if getting a specific workflow
	if req.GetVersion() != 0 {
		resultWorkflow, err := s.domainWorkflow.GetWorkflow(ctx, &usecase.GetWorkflowRequest{
			ProjectID:     req.GetProjectId(),
			ApplicationID: req.GetApplicationId(),
			EntrypointID:  req.GetEntrypointId(),
			Version:       req.GetVersion(),
		})
		if err != nil {
			return nil, err
		}

		return &dashboard.GetWorkflowsResponse{
			Workflows: []*pbendpoint.Workflow{resultWorkflow.Workflow},
		}, nil
	}

	// Get list of workflows
	requestWorkflows := &usecase.GetWorkflowsRequest{
		ProjectID:     req.GetProjectId(),
		ApplicationID: req.GetApplicationId(),
		EntrypointID:  req.GetEntrypointId(),
	}

	resultWorkflows, err := s.domainWorkflow.GetWorkflows(ctx, requestWorkflows)
	if err != nil {
		return nil, err
	}

	return &dashboard.GetWorkflowsResponse{
		Workflows: resultWorkflows.Workflows,
	}, nil
}

func (s *DashboardServiceServer) CreateWorkflow(ctx context.Context, req *dashboard.CreateWorkflowRequest) (*dashboard.CreateWorkflowResponse, error) {
	// Request validation
	if req.GetProjectId() == "" || req.GetApplicationId() == "" || req.GetEntrypointId() == "" {
		return nil, errors.New("project_id, application_id, entrypoint_id are required")
	}

	// Create workflow
	requestCreate := &usecase.CreateWorkflowRequest{
		ProjectID:     req.GetProjectId(),
		ApplicationID: req.GetApplicationId(),
		EntrypointID:  req.GetEntrypointId(),
	}

	if req.GetFromVersion() != 0 {
		requestCreate.FromVersion = utils.ToPtr(req.GetFromVersion())
	}

	responseCreate, err := s.domainWorkflow.CreateWorkflow(ctx, requestCreate)
	if err != nil {
		return nil, err
	}

	return &dashboard.CreateWorkflowResponse{
		Version: responseCreate.Version,
	}, nil
}

func (s *DashboardServiceServer) UpdateWorkflow(ctx context.Context, req *dashboard.UpdateWorkflowRequest) (*dashboard.UpdateWorkflowResponse, error) {
	// Request validation
	workflow := req.GetWorkflow()
	if workflow.GetProjectId() == "" || workflow.GetApplicationId() == "" || workflow.GetEntrypointId() == "" {
		return nil, errors.New("project_id, application_id, entrypoint_id are required")
	}

	// Update workflow
	requestUpdate := &usecase.UpdateWorkflowRequest{
		Workflow: workflow,
	}

	if _, err := s.domainWorkflow.UpdateWorkflow(ctx, requestUpdate); err != nil {
		return nil, err
	}

	return &dashboard.UpdateWorkflowResponse{}, nil
}

func (s *DashboardServiceServer) DeleteWorkflow(ctx context.Context, req *dashboard.DeleteWorkflowRequest) (*dashboard.DeleteWorkflowResponse, error) {
	// Request validation
	if req.GetProjectId() == "" || req.GetApplicationId() == "" || req.GetEntrypointId() == "" || req.GetVersion() == 0 {
		return nil, errors.New("project_id, application_id, entrypoint_id and version are required")
	}

	// Delete workflow
	requestDelete := &usecase.DeleteWorkflowRequest{
		ProjectID:     req.GetProjectId(),
		ApplicationID: req.GetApplicationId(),
		EntrypointID:  req.GetEntrypointId(),
		Version:       req.GetVersion(),
	}

	if err := s.domainWorkflow.DeleteWorkflow(ctx, requestDelete); err != nil {
		return nil, err
	}

	return &dashboard.DeleteWorkflowResponse{}, nil
}
