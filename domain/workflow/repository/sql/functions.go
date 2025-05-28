package sql

import (
	"context"

	modelworkflow "github.com/ideagate/server-controller/domain/workflow/model"
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

func (r *repository) GetListWorkflow(ctx context.Context, req *GetListWorkflowRequest) ([]*modelworkflow.Workflow, error) {
	session := r.db.WithContext(ctx).
		Where("project_id = ?", req.ProjectID).
		Where("application_id = ?", req.ApplicationID).
		Where("entrypoint_id = ?", req.EntrypointID).
		Order("version desc")

	if req.Limit != nil {
		session = session.Limit(*req.Limit)
	}

	var workflows []*modelworkflow.Workflow
	if err := session.Find(&workflows).Error; err != nil {
		return nil, err
	}

	return workflows, nil
}

func (r *repository) GetWorkflow(ctx context.Context, req *GetWorkflowRequest) (*modelworkflow.Workflow, error) {
	session := r.db.WithContext(ctx).
		Where("project_id = ?", req.ProjectID).
		Where("application_id = ?", req.ApplicationID).
		Where("entrypoint_id = ?", req.EntrypointID).
		Where("version = ?", req.Version)

	var workflow modelworkflow.Workflow
	if err := session.Take(&workflow).Error; err != nil {
		return nil, err
	}

	return &workflow, nil
}

func (r *repository) CreateWorkflow(ctx context.Context, req *CreateWorkflowRequest) error {
	session := r.db.WithContext(ctx)

	if err := session.Create(req.Workflow).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateWorkflow(ctx context.Context, req *UpdateWorkflowRequest) error {
	session := r.db.WithContext(ctx).
		Where("project_id = ?", req.Workflow.ProjectID).
		Where("application_id = ?", req.Workflow.ApplicationID).
		Where("entrypoint_id = ?", req.Workflow.EntrypointID).
		Where("version = ?", req.Workflow.Version)

	if err := session.Select("DataBytes").Updates(req.Workflow).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteWorkflow(ctx context.Context, req *DeleteWorkflowRequest) error {
	session := r.db.WithContext(ctx).
		Where("project_id = ?", req.ProjectID).
		Where("application_id = ?", req.ApplicationID).
		Where("entrypoint_id = ?", req.EntrypointID).
		Where("version = ?", req.Version)

	if err := session.Delete(&modelworkflow.Workflow{}).Error; err != nil {
		return err
	}

	return nil
}
