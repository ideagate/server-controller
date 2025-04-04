package grpc

import (
	"context"

	"github.com/bayu-aditya/ideagate/backend/core/utils"
	"github.com/bayu-aditya/ideagate/backend/model/gen-go/dashboard"
	"github.com/bayu-aditya/ideagate/backend/server/controller/domain/application/usecase"
)

func (s *DashboardServiceServer) GetListApplication(ctx context.Context, req *dashboard.GetListApplicationRequest) (*dashboard.GetListApplicationResponse, error) {
	reqListApp := &usecase.GetListApplicationRequest{
		ProjectID: req.GetProjectId(),
	}

	if req.GetApplicationId() != "" {
		reqListApp.ApplicationID = utils.ToPtr(req.GetApplicationId())
	}

	resListApp, err := s.usecaseApplication.GetListApplication(ctx, reqListApp)
	if err != nil {
		return nil, err
	}

	return &dashboard.GetListApplicationResponse{
		Applications: resListApp.Applications,
	}, nil
}

func (s *DashboardServiceServer) CreateApplication(ctx context.Context, req *dashboard.CreateApplicationRequest) (*dashboard.CreateApplicationResponse, error) {
	reqCreateApp := &usecase.CreateApplicationRequest{
		ProjectID:     req.GetProjectId(),
		ApplicationID: req.GetApplicationId(),
		Name:          req.GetName(),
	}

	if err := s.usecaseApplication.CreateApplication(ctx, reqCreateApp); err != nil {
		return nil, err
	}

	return &dashboard.CreateApplicationResponse{}, nil
}

func (s *DashboardServiceServer) UpdateApplication(ctx context.Context, req *dashboard.UpdateApplicationRequest) (*dashboard.UpdateApplicationResponse, error) {
	reqUpdateApp := &usecase.UpdateApplicationRequest{
		ProjectID:     req.GetProjectId(),
		ApplicationID: req.GetApplicationId(),
		Values:        make(map[string]any),
	}

	for key, val := range req.GetValues().GetFields() {
		switch key {
		case "name":
			reqUpdateApp.Values["name"] = val.GetStringValue()

		case "description":
			reqUpdateApp.Values["description"] = val.GetStringValue()
		}
	}

	if err := s.usecaseApplication.UpdateApplication(ctx, reqUpdateApp); err != nil {
		return nil, err
	}

	return &dashboard.UpdateApplicationResponse{}, nil
}

func (s *DashboardServiceServer) DeleteApplication(ctx context.Context, req *dashboard.DeleteApplicationRequest) (*dashboard.DeleteApplicationResponse, error) {
	reqDeleteApp := &usecase.DeleteApplicationRequest{
		ProjectID:     req.GetProjectId(),
		ApplicationID: req.GetApplicationId(),
	}

	if err := s.usecaseApplication.DeleteApplication(ctx, reqDeleteApp); err != nil {
		return nil, err
	}

	return &dashboard.DeleteApplicationResponse{}, nil
}
