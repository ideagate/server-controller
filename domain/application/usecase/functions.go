package usecase

import (
	"context"

	"github.com/bayu-aditya/ideagate/backend/core/utils/errors"
	pbapplication "github.com/bayu-aditya/ideagate/backend/model/gen-go/core/application"
	"github.com/bayu-aditya/ideagate/backend/server/controller/domain/application/repository/sql"
)

func NewApplicationUsecase(repoApp sql.ApplicationSQLRepository) ApplicationUsecase {
	return &usecase{
		repoApp: repoApp,
	}
}

type usecase struct {
	repoApp sql.ApplicationSQLRepository
}

func (u *usecase) GetListApplication(ctx context.Context, req *GetListApplicationRequest) (*GetListApplicationResponse, error) {
	if req.ProjectID == "" {
		return nil, errors.New("ProjectID is required")
	}

	requestRepo := &sql.GetListApplicationRequest{
		ProjectID: &req.ProjectID,
	}

	if req.ApplicationID != nil {
		requestRepo.ApplicationID = req.ApplicationID
	}

	resultRepo, err := u.repoApp.GetListApplication(ctx, requestRepo)
	if err != nil {
		return nil, err
	}

	result := &GetListApplicationResponse{
		Applications: make([]*pbapplication.Application, len(resultRepo)),
	}

	for i := 0; i < len(resultRepo); i++ {
		result.Applications[i] = resultRepo[i].ToProtoModel()
	}

	return result, nil
}

func (u *usecase) CreateApplication(ctx context.Context, req *CreateApplicationRequest) error {
	if req.ProjectID == "" {
		return errors.New("ProjectID is required")
	}

	if req.ApplicationID == "" {
		return errors.New("ApplicationID is required")
	}

	if req.Name == "" {
		return errors.New("Name is required")
	}

	err := u.repoApp.CreateApplication(ctx, &sql.CreateApplicationRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
		Name:          req.Name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) UpdateApplication(ctx context.Context, req *UpdateApplicationRequest) error {
	if req.ProjectID == "" {
		return errors.New("ProjectID is required")
	}

	if req.ApplicationID == "" {
		return errors.New("ApplicationID is required")
	}

	reqRepo := &sql.UpdateApplicationRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
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

	if err := u.repoApp.UpdateApplication(ctx, reqRepo); err != nil {
		return err
	}

	return nil
}

func (u *usecase) DeleteApplication(ctx context.Context, req *DeleteApplicationRequest) error {
	if req.ProjectID == "" {
		return errors.New("ProjectID is required")
	}

	if req.ApplicationID == "" {
		return errors.New("ApplicationID is required")
	}

	err := u.repoApp.DeleteApplication(ctx, &sql.DeleteApplicationRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
	})
	if err != nil {
		return err
	}

	return nil
}
