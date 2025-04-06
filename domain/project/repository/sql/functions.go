package sql

import (
	"context"

	"github.com/ideagate/server-controller/domain/project/model"
	"gorm.io/gorm"
)

func NewProjectRepository(db *gorm.DB) ProjectSQLRepository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetListProject(ctx context.Context, req *GetListProjectRequest) ([]*model.Project, error) {
	session := r.db.WithContext(ctx)

	var projects []*model.Project

	if err := session.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *repository) CreateProject(ctx context.Context, req *CreateProjectRequest) error {
	session := r.db.WithContext(ctx)

	data := model.Project{
		ID:   req.ProjectID,
		Name: req.Name,
	}

	if err := session.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateProject(ctx context.Context, req *UpdateProjectRequest) error {
	session := r.db.WithContext(ctx)

	err := session.
		Model(&model.Project{}).
		Select(req.Fields).
		Where("id = ?", req.ProjectID).
		Updates(req.Data).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteProject(ctx context.Context, req *DeleteProjectRequest) error {
	session := r.db.WithContext(ctx)

	err := session.
		Where("id = ?", req.ProjectID).
		Delete(&model.Project{}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateProjectUser(ctx context.Context, req *CreateProjectUserRequest) error {
	session := r.db.WithContext(ctx)

	data := model.ProjectUser{
		ProjectID: req.ProjectID,
		UserID:    req.UserID,
	}

	if err := session.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
