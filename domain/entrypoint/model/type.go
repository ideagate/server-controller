package model

import pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"

const (
	EntryPointUnknown EntrypointType = "unknown"
	EntryPointRest    EntrypointType = "rest"
	EntryPointCron    EntrypointType = "cron"
)

type EntrypointType string

func (e EntrypointType) String() string {
	return string(e)
}

func (e EntrypointType) ToProto() pbendpoint.EndpointType {
	switch e {
	case EntryPointRest:
		return pbendpoint.EndpointType_ENDPOINT_TYPE_REST
	case EntryPointCron:
		return pbendpoint.EndpointType_ENDPOINT_TYPE_CRON
	default:
		return pbendpoint.EndpointType_ENDPOINT_TYPE_UNSPECIFIED
	}
}

func (e *EntrypointType) FromProto(endpointType pbendpoint.EndpointType) {
	switch endpointType {
	case pbendpoint.EndpointType_ENDPOINT_TYPE_REST:
		*e = EntryPointRest
	case pbendpoint.EndpointType_ENDPOINT_TYPE_CRON:
		*e = EntryPointCron
	default:
		*e = EntryPointUnknown
	}
}
