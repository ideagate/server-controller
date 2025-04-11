package usecase

import (
	"context"

	"github.com/ideagate/core/utils"
	pbEndpoint "github.com/ideagate/model/gen-go/core/endpoint"
	"github.com/ideagate/server-controller/domain/entrypoint"
	"github.com/ideagate/server-controller/domain/entrypoint/model"
	"github.com/ideagate/server-controller/domain/entrypoint/repository/sql"
)

func New(repoSql sql.Repository) entrypoint.Domain {
	return &usecase{
		repoSql: repoSql,
	}
}

type usecase struct {
	repoSql sql.Repository
}

func (u *usecase) GetListEntrypoint(ctx context.Context, req *entrypoint.GetListEntrypointRequest) (*entrypoint.GetListEntrypointResponse, error) {
	if req.ProjectID == "" {
		return nil, model.ErrProjectIDRequired
	}

	if req.ApplicationID == "" {
		return nil, model.ErrApplicationIDRequired
	}

	resultRepo, err := u.repoSql.GetListEntrypoint(ctx, &sql.GetListEntrypointRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
		Type:          utils.ToPtr(model.EntryPointRest.String()),
	})
	if err != nil {
		return nil, err
	}

	result := &entrypoint.GetListEntrypointResponse{
		Entrypoints: make([]*pbEndpoint.Endpoint, len(resultRepo)),
	}

	for i := 0; i < len(resultRepo); i++ {
		result.Entrypoints[i] = resultRepo[i].ToProto()
	}

	return result, nil
}

func (u *usecase) GetEntrypoint(ctx context.Context, req *entrypoint.GetEntrypointRequest) (*entrypoint.GetEntrypointResponse, error) {
	if req.ProjectID == "" {
		return nil, model.ErrProjectIDRequired
	}

	if req.ApplicationID == "" {
		return nil, model.ErrApplicationIDRequired
	}

	if req.EntrypointID == "" {
		return nil, model.ErrEntrypointIDRequired
	}

	dataEntrypoint, err := u.repoSql.GetEntrypoint(ctx, &sql.GetEntrypointRequest{
		ProjectID:     req.ProjectID,
		ApplicationID: req.ApplicationID,
		EntrypointID:  req.EntrypointID,
	})
	if err != nil {
		return nil, err
	}

	if dataEntrypoint == nil {
		return nil, model.ErrEntrypointNotFound
	}

	return &entrypoint.GetEntrypointResponse{
		Entrypoint: dataEntrypoint.ToProto(),
	}, nil
}
