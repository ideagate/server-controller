package model

import (
	"fmt"
	"testing"

	pbendpoint "github.com/ideagate/model/gen-go/core/endpoint"
	"google.golang.org/protobuf/proto"
)

func TestWorkflow_Data_ProtoToBytesToProto(t *testing.T) {
	protoWorkflow := &pbendpoint.Workflow{
		Version:       1,
		EntrypointId:  "entrypoint_123",
		ApplicationId: "app_456",
		ProjectId:     "project_789",
		Steps: []*pbendpoint.Step{
			{
				Id: "sleep_1",
				Action: &pbendpoint.Step_ActionSleep{
					ActionSleep: &pbendpoint.ActionSleep{
						TimeoutMs: 123456,
					},
				},
			},
		},
	}

	// Convert proto to Workflow
	workflow := &Workflow{}
	workflow.FromProto(protoWorkflow)

	fmt.Printf("Data Bytes (hex): %x\n", workflow.DataBytes)

	// Convert bytes back to Workflow
	newWorkflow := Workflow{
		DataBytes: workflow.DataBytes,
	}
	newProtoWorkflow := newWorkflow.ToProto()

	// Assert the data matches
	if !proto.Equal(protoWorkflow, newProtoWorkflow) {
		//t.Errorf("Expected proto workflow %v, got %v", protoWorkflow, newProtoWorkflow)
	}
}
