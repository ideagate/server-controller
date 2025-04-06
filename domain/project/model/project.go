package model

import (
	"time"

	pbproject "github.com/ideagate/model/gen-go/core/project"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Project struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
}

func (p *Project) TableName() string {
	return "project"
}

func (p *Project) BeforeCreate(_ *gorm.DB) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Project) BeforeUpdate(_ *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Project) ToProtoModel() *pbproject.Project {
	return &pbproject.Project{
		Id:          p.ID,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
		Name:        p.Name,
		Description: p.Description,
	}
}

type ProjectUser struct {
	ProjectID string
	UserID    int64
	CreatedAt time.Time
}

func (pu *ProjectUser) TableName() string {
	return "project_user"
}

func (pu *ProjectUser) BeforeCreate(_ *gorm.DB) error {
	pu.CreatedAt = time.Now()
	return nil
}
