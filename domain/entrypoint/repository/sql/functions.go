package sql

import (
	"context"

	"github.com/ideagate/server-controller/domain/entrypoint/model"
	"gorm.io/gorm"
)

func New(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetListEntrypoint(ctx context.Context, req *GetListEntrypointRequest) ([]*model.Entrypoint, error) {
	session := r.db.WithContext(ctx).
		Where("project_id = ?", req.ProjectID).
		Where("application_id = ?", req.ApplicationID)

	if req.Type != nil {
		session = session.Where("type = ?", *req.Type)
	}

	var entrypoints []*model.Entrypoint

	err := session.
		Find(&entrypoints).
		Error

	if err != nil {
		return nil, err
	}

	return entrypoints, nil
}

func (r *repository) GetEntrypoint(ctx context.Context, req *GetEntrypointRequest) (*model.Entrypoint, error) {
	session := r.db.WithContext(ctx).
		Where("project_id = ?", req.ProjectID).
		Where("application_id = ?", req.ApplicationID).
		Where("id = ?", req.EntrypointID)

	var entrypoint model.Entrypoint
	if err := session.Take(&entrypoint).Error; err != nil {
		return nil, err
	}

	return &entrypoint, nil
}

func (r *repository) CreateEntrypoint(ctx context.Context, req *CreateEntrypointRequest) error {
	session := r.db.WithContext(ctx)

	if err := session.Create(req.Entrypoint).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteEntrypoint(ctx context.Context, req *DeleteEntrypointRequest) error {
	session := r.db.WithContext(ctx).
		Where("project_id = ?", req.ProjectID).
		Where("application_id = ?", req.ApplicationID).
		Where("id = ?", req.EntrypointID)

	if err := session.Delete(&model.Entrypoint{}).Error; err != nil {
		return err
	}

	return nil
}
