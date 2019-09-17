// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyWorkspaceStateInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	//
	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" required:"true"`

	// The WorkSpace state.
	//
	// WorkspaceState is a required field
	WorkspaceState TargetWorkspaceState `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ModifyWorkspaceStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyWorkspaceStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyWorkspaceStateInput"}

	if s.WorkspaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceId"))
	}
	if len(s.WorkspaceState) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceState"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyWorkspaceStateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyWorkspaceStateOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyWorkspaceState = "ModifyWorkspaceState"

// ModifyWorkspaceStateRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Sets the state of the specified WorkSpace.
//
// To maintain a WorkSpace without being interrupted, set the WorkSpace state
// to ADMIN_MAINTENANCE. WorkSpaces in this state do not respond to requests
// to reboot, stop, start, rebuild, or restore. An AutoStop WorkSpace in this
// state is not stopped. Users cannot log into a WorkSpace in the ADMIN_MAINTENANCE
// state.
//
//    // Example sending a request using ModifyWorkspaceStateRequest.
//    req := client.ModifyWorkspaceStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/ModifyWorkspaceState
func (c *Client) ModifyWorkspaceStateRequest(input *ModifyWorkspaceStateInput) ModifyWorkspaceStateRequest {
	op := &aws.Operation{
		Name:       opModifyWorkspaceState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyWorkspaceStateInput{}
	}

	req := c.newRequest(op, input, &ModifyWorkspaceStateOutput{})
	return ModifyWorkspaceStateRequest{Request: req, Input: input, Copy: c.ModifyWorkspaceStateRequest}
}

// ModifyWorkspaceStateRequest is the request type for the
// ModifyWorkspaceState API operation.
type ModifyWorkspaceStateRequest struct {
	*aws.Request
	Input *ModifyWorkspaceStateInput
	Copy  func(*ModifyWorkspaceStateInput) ModifyWorkspaceStateRequest
}

// Send marshals and sends the ModifyWorkspaceState API request.
func (r ModifyWorkspaceStateRequest) Send(ctx context.Context) (*ModifyWorkspaceStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyWorkspaceStateResponse{
		ModifyWorkspaceStateOutput: r.Request.Data.(*ModifyWorkspaceStateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyWorkspaceStateResponse is the response type for the
// ModifyWorkspaceState API operation.
type ModifyWorkspaceStateResponse struct {
	*ModifyWorkspaceStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyWorkspaceState request.
func (r *ModifyWorkspaceStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
