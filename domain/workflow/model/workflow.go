package model

import (
	"time"

	pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Workflow struct {
	Version       int64                            `json:"version"`
	EntrypointID  string                           `json:"entrypoint_id"`
	ApplicationID string                           `json:"application_id"`
	ProjectID     string                           `json:"project_id"`
	CreatedAt     time.Time                        `json:"created_at"`
	UpdatedAt     time.Time                        `json:"updated_at"`
	Data          datatypes.JSONType[WorkflowData] `json:"data"`
}

type WorkflowData struct {
	Steps []*pbendpoint.Step `json:"steps"`
	Edges []*pbendpoint.Edge `json:"edges"`
}

func (w *Workflow) TableName() string {
	return "workflow"
}

func (w *Workflow) BeforeCreate(_ *gorm.DB) error {
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Workflow) BeforeUpdate(_ *gorm.DB) error {
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Workflow) ToProto() *pbendpoint.Workflow {
	data := w.Data.Data()

	return &pbendpoint.Workflow{
		Version:       w.Version,
		EntrypointId:  w.EntrypointID,
		ApplicationId: w.ApplicationID,
		ProjectId:     w.ProjectID,
		CreatedAt:     timestamppb.New(w.CreatedAt),
		UpdatedAt:     timestamppb.New(w.UpdatedAt),
		Steps:         data.Steps,
		Edges:         data.Edges,
	}
}

func (w *Workflow) FromProto(workflow *pbendpoint.Workflow) {
	w.Version = workflow.Version
	w.EntrypointID = workflow.EntrypointId
	w.ApplicationID = workflow.ApplicationId
	w.ProjectID = workflow.ProjectId
	w.CreatedAt = workflow.CreatedAt.AsTime()
	w.UpdatedAt = workflow.UpdatedAt.AsTime()
	w.Data = datatypes.NewJSONType(WorkflowData{
		Steps: workflow.Steps,
		Edges: workflow.Edges,
	})
}
