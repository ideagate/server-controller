package usecase

import pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"

type GetListEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	Type          pbendpoint.EndpointType
}

type GetListEntrypointResponse struct {
	Entrypoints []*pbendpoint.Endpoint
}

type GetEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
}

type GetEntrypointResponse struct {
	Entrypoint *pbendpoint.Endpoint
}

type CreateEntrypointRequest struct {
	Entrypoint *pbendpoint.Endpoint
}

type DeleteEntrypointRequest struct {
	ProjectID     string
	ApplicationID string
	EntrypointID  string
}
