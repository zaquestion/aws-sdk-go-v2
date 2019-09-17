// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CancelMaintenanceWindowExecutionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window execution to stop.
	//
	// WindowExecutionId is a required field
	WindowExecutionId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelMaintenanceWindowExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelMaintenanceWindowExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelMaintenanceWindowExecutionInput"}

	if s.WindowExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowExecutionId"))
	}
	if s.WindowExecutionId != nil && len(*s.WindowExecutionId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowExecutionId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelMaintenanceWindowExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window execution that has been stopped.
	WindowExecutionId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s CancelMaintenanceWindowExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelMaintenanceWindowExecution = "CancelMaintenanceWindowExecution"

// CancelMaintenanceWindowExecutionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Stops a maintenance window execution that is already in progress and cancels
// any tasks in the window that have not already starting running. (Tasks already
// in progress will continue to completion.)
//
//    // Example sending a request using CancelMaintenanceWindowExecutionRequest.
//    req := client.CancelMaintenanceWindowExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CancelMaintenanceWindowExecution
func (c *Client) CancelMaintenanceWindowExecutionRequest(input *CancelMaintenanceWindowExecutionInput) CancelMaintenanceWindowExecutionRequest {
	op := &aws.Operation{
		Name:       opCancelMaintenanceWindowExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelMaintenanceWindowExecutionInput{}
	}

	req := c.newRequest(op, input, &CancelMaintenanceWindowExecutionOutput{})
	return CancelMaintenanceWindowExecutionRequest{Request: req, Input: input, Copy: c.CancelMaintenanceWindowExecutionRequest}
}

// CancelMaintenanceWindowExecutionRequest is the request type for the
// CancelMaintenanceWindowExecution API operation.
type CancelMaintenanceWindowExecutionRequest struct {
	*aws.Request
	Input *CancelMaintenanceWindowExecutionInput
	Copy  func(*CancelMaintenanceWindowExecutionInput) CancelMaintenanceWindowExecutionRequest
}

// Send marshals and sends the CancelMaintenanceWindowExecution API request.
func (r CancelMaintenanceWindowExecutionRequest) Send(ctx context.Context) (*CancelMaintenanceWindowExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelMaintenanceWindowExecutionResponse{
		CancelMaintenanceWindowExecutionOutput: r.Request.Data.(*CancelMaintenanceWindowExecutionOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelMaintenanceWindowExecutionResponse is the response type for the
// CancelMaintenanceWindowExecution API operation.
type CancelMaintenanceWindowExecutionResponse struct {
	*CancelMaintenanceWindowExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelMaintenanceWindowExecution request.
func (r *CancelMaintenanceWindowExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
