package usecase

import (
	"context"

	"github.com/ideagate/core/utils/errors"
	pbproject "github.com/ideagate/model/gen-go/core/project"
	"github.com/ideagate/server-controller/domain/project/repository/sql"
)

func NewProjectUsecase(repoProject sql.ProjectSQLRepository) ProjectUsecase {
	return &usecase{
		repoProject: repoProject,
	}
}

type usecase struct {
	repoProject sql.ProjectSQLRepository
}

func (u *usecase) GetListProject(ctx context.Context, req *GetListProjectRequest) (*GetListProjectResponse, error) {
	resultRepo, err := u.repoProject.GetListProject(ctx, &sql.GetListProjectRequest{})
	if err != nil {
		return nil, err
	}

	result := &GetListProjectResponse{
		Projects: make([]*pbproject.Project, len(resultRepo)),
	}

	for i := 0; i < len(resultRepo); i++ {
		result.Projects[i] = resultRepo[i].ToProtoModel()
	}

	return result, nil
}

func (u *usecase) CreateProject(ctx context.Context, req *CreateProjectRequest) error {
	if req.ProjectID == "" {
		return errors.New("ProjectID is required")
	}

	if req.Name == "" {
		return errors.New("Name is required")
	}

	err := u.repoProject.CreateProject(ctx, &sql.CreateProjectRequest{
		ProjectID: req.ProjectID,
		Name:      req.Name,
	})
	if err != nil {
		return err
	}

	err = u.repoProject.CreateProjectUser(ctx, &sql.CreateProjectUserRequest{
		ProjectID: req.ProjectID,
		UserID:    req.UserID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) UpdateProject(ctx context.Context, req *UpdateProjectRequest) error {
	if req.ProjectID == "" {
		return errors.New("ProjectID is required")
	}

	reqRepo := &sql.UpdateProjectRequest{
		ProjectID: req.ProjectID,
	}

	for key, value := range req.Values {
		switch key {
		case "name":
			reqRepo.Data.Name = value.(string)
			reqRepo.Fields = append(reqRepo.Fields, "name")

		case "description":
			reqRepo.Data.Description = value.(string)
			reqRepo.Fields = append(reqRepo.Fields, "description")
		}
	}

	if len(reqRepo.Fields) == 0 {
		return nil
	}

	if err := u.repoProject.UpdateProject(ctx, reqRepo); err != nil {
		return err
	}

	return nil
}

func (u *usecase) DeleteProject(ctx context.Context, req *DeleteProjectRequest) error {
	if req.ProjectID == "" {
		return errors.New("ProjectID is required")
	}

	err := u.repoProject.DeleteProject(ctx, &sql.DeleteProjectRequest{
		ProjectID: req.ProjectID,
	})
	if err != nil {
		return err
	}

	return nil
}
