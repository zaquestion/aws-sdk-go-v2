// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the parameter group.
	//
	// ParameterGroupName is a required field
	ParameterGroupName *string `type:"string" required:"true"`

	// An array of name-value pairs for the parameters in the group. Each element
	// in the array represents a single parameter.
	//
	// ParameterNameValues is a required field
	ParameterNameValues []ParameterNameValue `type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateParameterGroupInput"}

	if s.ParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupName"))
	}

	if s.ParameterNameValues == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterNameValues"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// The parameter group that has been modified.
	ParameterGroup *ParameterGroup `type:"structure"`
}

// String returns the string representation
func (s UpdateParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateParameterGroup = "UpdateParameterGroup"

// UpdateParameterGroupRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Modifies the parameters of a parameter group. You can modify up to 20 parameters
// in a single request by submitting a list parameter name and value pairs.
//
//    // Example sending a request using UpdateParameterGroupRequest.
//    req := client.UpdateParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/UpdateParameterGroup
func (c *Client) UpdateParameterGroupRequest(input *UpdateParameterGroupInput) UpdateParameterGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateParameterGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateParameterGroupOutput{})
	return UpdateParameterGroupRequest{Request: req, Input: input, Copy: c.UpdateParameterGroupRequest}
}

// UpdateParameterGroupRequest is the request type for the
// UpdateParameterGroup API operation.
type UpdateParameterGroupRequest struct {
	*aws.Request
	Input *UpdateParameterGroupInput
	Copy  func(*UpdateParameterGroupInput) UpdateParameterGroupRequest
}

// Send marshals and sends the UpdateParameterGroup API request.
func (r UpdateParameterGroupRequest) Send(ctx context.Context) (*UpdateParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateParameterGroupResponse{
		UpdateParameterGroupOutput: r.Request.Data.(*UpdateParameterGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateParameterGroupResponse is the response type for the
// UpdateParameterGroup API operation.
type UpdateParameterGroupResponse struct {
	*UpdateParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateParameterGroup request.
func (r *UpdateParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
