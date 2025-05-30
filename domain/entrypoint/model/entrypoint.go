package model

import (
	"encoding/json"
	"time"

	pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Entrypoint struct {
	ID            string               `json:"id"`
	ApplicationID string               `json:"application_id"`
	ProjectID     string               `json:"project_id"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	Type          string               `json:"type"`
	Name          datatypes.NullString `json:"name"`
	Description   datatypes.NullString `json:"description"`
	Settings      datatypes.JSON       `json:"settings"`
}

func (e *Entrypoint) TableName() string {
	return "entrypoint"
}

func (e *Entrypoint) BeforeCreate(_ *gorm.DB) error {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return nil
}

func (e *Entrypoint) BeforeUpdate(_ *gorm.DB) error {
	e.UpdatedAt = time.Now()
	return nil
}

func (e *Entrypoint) ToProto() *pbendpoint.Endpoint {
	entrypointType := EntrypointType(e.Type)

	return &pbendpoint.Endpoint{
		Id:            e.ID,
		ApplicationId: e.ApplicationID,
		ProjectId:     e.ProjectID,
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		Type:          entrypointType.ToProto(),
		Name:          e.Name.V,
		Description:   e.Description.V,
		Settings:      e.toProtoSettings().Settings,
	}
}

func (e *Entrypoint) FromProto(entrypoint *pbendpoint.Endpoint) {
	var entrypointType EntrypointType
	entrypointType.FromProto(entrypoint.GetType())

	e.ID = entrypoint.Id
	e.ApplicationID = entrypoint.ApplicationId
	e.ProjectID = entrypoint.ProjectId
	e.CreatedAt = entrypoint.CreatedAt.AsTime()
	e.UpdatedAt = entrypoint.UpdatedAt.AsTime()
	e.Type = entrypointType.String()
	e.Name = datatypes.NullString{V: entrypoint.Name, Valid: true}
	e.Description = datatypes.NullString{V: entrypoint.Description, Valid: true}
	e.fromProtoSettings(entrypoint)
}

func (e *Entrypoint) toProtoSettings() *pbendpoint.Endpoint {
	result := &pbendpoint.Endpoint{}
	settingsJson, _ := e.Settings.MarshalJSON()

	var settings interface{}

	switch EntrypointType(e.Type) {
	case EntryPointRest:
		settings = &pbendpoint.SettingRest{}
		result.Settings = &pbendpoint.Endpoint_SettingRest{
			SettingRest: settings.(*pbendpoint.SettingRest),
		}

	case EntryPointCron:
		settings = &pbendpoint.SettingCron{}
		result.Settings = &pbendpoint.Endpoint_SettingCron{
			SettingCron: settings.(*pbendpoint.SettingCron),
		}
	}

	if settings != nil {
		_ = json.Unmarshal(settingsJson, settings)
	}

	return result
}

func (e *Entrypoint) fromProtoSettings(entrypoint *pbendpoint.Endpoint) {
	var settingsJson []byte
	switch EntrypointType(e.Type) {
	case EntryPointRest:
		settingsJson, _ = json.Marshal(entrypoint.GetSettingRest())
	case EntryPointCron:
		settingsJson, _ = json.Marshal(entrypoint.GetSettingCron())
	}

	if settingsJson != nil {
		_ = e.Settings.UnmarshalJSON(settingsJson)
	}
}
