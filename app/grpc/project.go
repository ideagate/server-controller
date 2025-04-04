package grpc

import (
	"context"

	"github.com/bayu-aditya/ideagate/backend/model/gen-go/dashboard"
	"github.com/bayu-aditya/ideagate/backend/server/controller/domain/project/usecase"
)

func (s *DashboardServiceServer) GetListProject(ctx context.Context, req *dashboard.GetListProjectRequest) (*dashboard.GetListProjectResponse, error) {
	reqListProject := &usecase.GetListProjectRequest{}

	resListProject, err := s.usecaseProject.GetListProject(ctx, reqListProject)
	if err != nil {
		return nil, err
	}

	return &dashboard.GetListProjectResponse{
		Projects: resListProject.Projects,
	}, nil
}

func (s *DashboardServiceServer) CreateProject(ctx context.Context, req *dashboard.CreateProjectRequest) (*dashboard.CreateProjectResponse, error) {
	reqCreateProject := &usecase.CreateProjectRequest{
		UserID:    2, // TODO fill this with user id from token
		ProjectID: req.GetProjectId(),
		Name:      req.GetName(),
	}

	if err := s.usecaseProject.CreateProject(ctx, reqCreateProject); err != nil {
		return nil, err
	}

	return &dashboard.CreateProjectResponse{}, nil
}

func (s *DashboardServiceServer) UpdateProject(ctx context.Context, req *dashboard.UpdateProjectRequest) (*dashboard.UpdateProjectResponse, error) {
	reqUpdateProject := &usecase.UpdateProjectRequest{
		ProjectID: req.GetProjectId(),
		Values:    make(map[string]interface{}),
	}

	for key, val := range req.GetValues().GetFields() {
		switch key {
		case "name":
			reqUpdateProject.Values["name"] = val.GetStringValue()

		case "description":
			reqUpdateProject.Values["description"] = val.GetStringValue()
		}
	}

	if err := s.usecaseProject.UpdateProject(ctx, reqUpdateProject); err != nil {
		return nil, err
	}

	return &dashboard.UpdateProjectResponse{}, nil
}

func (s *DashboardServiceServer) DeleteProject(ctx context.Context, req *dashboard.DeleteProjectRequest) (*dashboard.DeleteProjectResponse, error) {
	reqDeleteProject := &usecase.DeleteProjectRequest{
		ProjectID: req.GetProjectId(),
	}

	if err := s.usecaseProject.DeleteProject(ctx, reqDeleteProject); err != nil {
		return nil, err
	}

	return &dashboard.DeleteProjectResponse{}, nil
}
