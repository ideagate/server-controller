package grpc

import (
	"context"

	"github.com/ideagate/core/utils/errors"
	pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"
	"github.com/ideagate/model/gen-go/dashboard"
	"github.com/ideagate/server-controller/domain/entrypoint"
)

func (s *DashboardServiceServer) GetListEndpoint(ctx context.Context, req *dashboard.GetListEndpointRequest) (*dashboard.GetListEndpointResponse, error) {
	if req.GetProjectId() == "" || req.GetApplicationId() == "" {
		return nil, errors.New("project_id and application_id are required")
	}

	// Check if getting a specific entrypoint
	if req.GetEndpointId() != "" {
		resultEntrypoint, err := s.domainEntrypoint.GetEntrypoint(ctx, &entrypoint.GetEntrypointRequest{
			ProjectID:     req.GetProjectId(),
			ApplicationID: req.GetApplicationId(),
			EntrypointID:  req.GetEndpointId(),
		})
		if err != nil {
			return nil, err
		}

		return &dashboard.GetListEndpointResponse{
			Endpoints: []*pbendpoint.Endpoint{resultEntrypoint.Entrypoint},
		}, nil
	}

	// Get list of entrypoints
	requestEntrypoints := &entrypoint.GetListEntrypointRequest{
		ProjectID:     req.GetProjectId(),
		ApplicationID: req.GetApplicationId(),
	}

	resultEntrypoints, err := s.domainEntrypoint.GetListEntrypoint(ctx, requestEntrypoints)
	if err != nil {
		return nil, err
	}

	return &dashboard.GetListEndpointResponse{
		Endpoints: resultEntrypoints.Entrypoints,
	}, nil
}

func (s *DashboardServiceServer) CreateEndpoint(ctx context.Context, req *dashboard.CreateEndpointRequest) (*dashboard.CreateEndpointResponse, error) {
	if req.GetEndpoint() == nil {
		return nil, errors.New("endpoint are required")
	}

	// Create entrypoint
	requestCreate := &entrypoint.CreateEntrypointRequest{
		Entrypoint: req.GetEndpoint(),
	}

	if err := s.domainEntrypoint.CreateEntrypoint(ctx, requestCreate); err != nil {
		return nil, err
	}

	return &dashboard.CreateEndpointResponse{}, nil
}

func (s *DashboardServiceServer) DeleteEndpoint(ctx context.Context, req *dashboard.DeleteEndpointRequest) (*dashboard.DeleteEndpointResponse, error) {
	if req.GetProjectId() == "" || req.GetApplicationId() == "" || req.GetEndpointId() == "" {
		return nil, errors.New("project_id, application_id and endpoint_id are required")
	}

	// Delete entrypoint
	requestDelete := &entrypoint.DeleteEntrypointRequest{
		ProjectID:     req.GetProjectId(),
		ApplicationID: req.GetApplicationId(),
		EntrypointID:  req.GetEndpointId(),
	}

	if err := s.domainEntrypoint.DeleteEntrypoint(ctx, requestDelete); err != nil {
		return nil, err
	}

	return &dashboard.DeleteEndpointResponse{}, nil
}
