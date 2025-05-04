package model

import "github.com/ideagate/core/utils/errors"

var (
	ErrProjectIDRequired       = errors.New("project_id is required")
	ErrApplicationIDRequired   = errors.New("application_id is required")
	ErrEntrypointIDRequired    = errors.New("entrypoint_id is required")
	ErrEntrypointNotFound      = errors.New("entrypoint not found")
	ErrWorkflowVersionRequired = errors.New("workflow version is required")
)
