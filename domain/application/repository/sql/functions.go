package sql

import (
	"context"

	"github.com/bayu-aditya/ideagate/backend/server/controller/domain/application/model"
	"gorm.io/gorm"
)

func NewApplicationRepository(db *gorm.DB) ApplicationSQLRepository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetListApplication(ctx context.Context, req *GetListApplicationRequest) ([]*model.Application, error) {
	session := r.db.WithContext(ctx)

	if req.ProjectID != nil {
		session = session.Where("project_id = ?", *req.ProjectID)
	}

	if req.ApplicationID != nil {
		session = session.Where("id = ?", *req.ApplicationID)
	}

	if req.Limit == 0 {
		req.Limit = 1
	}

	var applications []*model.Application

	err := session.
		Find(&applications).
		Limit(req.Limit).
		Error

	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (r *repository) CreateApplication(ctx context.Context, req *CreateApplicationRequest) error {
	session := r.db.WithContext(ctx)

	data := model.Application{
		ProjectID: req.ProjectID,
		ID:        req.ApplicationID,
		Name:      req.Name,
	}

	if err := session.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateApplication(ctx context.Context, req *UpdateApplicationRequest) error {
	session := r.db.WithContext(ctx)

	err := session.
		Model(&model.Application{}).
		Select(req.Fields).
		Where("id = ? AND project_id = ?", req.ApplicationID, req.ProjectID).
		Updates(req.Data).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteApplication(ctx context.Context, req *DeleteApplicationRequest) error {
	session := r.db.WithContext(ctx)

	err := session.
		Where("id = ? AND project_id = ?", req.ApplicationID, req.ProjectID).
		Delete(&model.Application{}).
		Error

	if err != nil {
		return err
	}

	return nil
}
