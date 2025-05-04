package usecase

import pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"

type GetWorkflowsRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
}

type GetWorkflowsResponse struct {
	Workflows []*pbendpoint.Workflow
}

type GetWorkflowRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
	Version       int64
}

type GetWorkflowResponse struct {
	Workflow *pbendpoint.Workflow
}

type CreateWorkflowRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string

	// New workflow will be copied from this version.
	// If not specified, then new workflow will be created from scratch.
	FromVersion *int64
}

type CreateWorkflowResponse struct {
	Version int64
}

type UpdateWorkflowRequest struct {
	Workflow *pbendpoint.Workflow
}

type UpdateWorkflowResponse struct{}

type DeleteWorkflowRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
	Version       int64
}
