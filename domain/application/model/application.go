package model

import (
	"time"

	pbapplication "github.com/ideagate/model/gen-go/core/application"
	"github.com/jackc/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Application struct {
	ID          string
	ProjectID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
}

func (a *Application) TableName() string {
	return "application"
}

func (a *Application) BeforeCreate(_ *gorm.DB) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Application) BeforeUpdate(_ *gorm.DB) error {
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Application) ToProtoModel() *pbapplication.Application {
	return &pbapplication.Application{
		Id:          a.ID,
		ProjectId:   a.ProjectID,
		CreatedAt:   timestamppb.New(a.CreatedAt),
		UpdatedAt:   timestamppb.New(a.UpdatedAt),
		Name:        a.Name,
		Description: a.Description,
	}
}

type Endpoint struct {
	ID            string
	ApplicationID string
	ProjectID     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Method        string
	Path          string
	Name          string
	Description   string
	Settings      pgtype.JSONB `gorm:"type:jsonb"`
}

func (e *Endpoint) TableName() string {
	return "endpoint"
}

type Workflow struct {
	Version       int
	EndpointID    string
	ApplicationID string
	ProjectID     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Data          pgtype.JSONB `gorm:"type:jsonb"`
}

func (w *Workflow) TableName() string {
	return "workflow"
}
