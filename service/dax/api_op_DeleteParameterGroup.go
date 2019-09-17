// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the parameter group to delete.
	//
	// ParameterGroupName is a required field
	ParameterGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteParameterGroupInput"}

	if s.ParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// A user-specified message for this action (i.e., a reason for deleting the
	// parameter group).
	DeletionMessage *string `type:"string"`
}

// String returns the string representation
func (s DeleteParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteParameterGroup = "DeleteParameterGroup"

// DeleteParameterGroupRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Deletes the specified parameter group. You cannot delete a parameter group
// if it is associated with any DAX clusters.
//
//    // Example sending a request using DeleteParameterGroupRequest.
//    req := client.DeleteParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DeleteParameterGroup
func (c *Client) DeleteParameterGroupRequest(input *DeleteParameterGroupInput) DeleteParameterGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteParameterGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteParameterGroupOutput{})
	return DeleteParameterGroupRequest{Request: req, Input: input, Copy: c.DeleteParameterGroupRequest}
}

// DeleteParameterGroupRequest is the request type for the
// DeleteParameterGroup API operation.
type DeleteParameterGroupRequest struct {
	*aws.Request
	Input *DeleteParameterGroupInput
	Copy  func(*DeleteParameterGroupInput) DeleteParameterGroupRequest
}

// Send marshals and sends the DeleteParameterGroup API request.
func (r DeleteParameterGroupRequest) Send(ctx context.Context) (*DeleteParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteParameterGroupResponse{
		DeleteParameterGroupOutput: r.Request.Data.(*DeleteParameterGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteParameterGroupResponse is the response type for the
// DeleteParameterGroup API operation.
type DeleteParameterGroupResponse struct {
	*DeleteParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteParameterGroup request.
func (r *DeleteParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
