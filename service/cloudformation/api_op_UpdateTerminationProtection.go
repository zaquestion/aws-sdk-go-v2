// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateTerminationProtectionInput struct {
	_ struct{} `type:"structure"`

	// Whether to enable termination protection on the specified stack.
	//
	// EnableTerminationProtection is a required field
	EnableTerminationProtection *bool `type:"boolean" required:"true"`

	// The name or unique ID of the stack for which you want to set termination
	// protection.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTerminationProtectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTerminationProtectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTerminationProtectionInput"}

	if s.EnableTerminationProtection == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnableTerminationProtection"))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateTerminationProtectionOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the stack.
	StackId *string `type:"string"`
}

// String returns the string representation
func (s UpdateTerminationProtectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTerminationProtection = "UpdateTerminationProtection"

// UpdateTerminationProtectionRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Updates termination protection for the specified stack. If a user attempts
// to delete a stack with termination protection enabled, the operation fails
// and the stack remains unchanged. For more information, see Protecting a Stack
// From Being Deleted (AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html)
// in the AWS CloudFormation User Guide.
//
// For nested stacks (AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html),
// termination protection is set on the root stack and cannot be changed directly
// on the nested stack.
//
//    // Example sending a request using UpdateTerminationProtectionRequest.
//    req := client.UpdateTerminationProtectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/UpdateTerminationProtection
func (c *Client) UpdateTerminationProtectionRequest(input *UpdateTerminationProtectionInput) UpdateTerminationProtectionRequest {
	op := &aws.Operation{
		Name:       opUpdateTerminationProtection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTerminationProtectionInput{}
	}

	req := c.newRequest(op, input, &UpdateTerminationProtectionOutput{})
	return UpdateTerminationProtectionRequest{Request: req, Input: input, Copy: c.UpdateTerminationProtectionRequest}
}

// UpdateTerminationProtectionRequest is the request type for the
// UpdateTerminationProtection API operation.
type UpdateTerminationProtectionRequest struct {
	*aws.Request
	Input *UpdateTerminationProtectionInput
	Copy  func(*UpdateTerminationProtectionInput) UpdateTerminationProtectionRequest
}

// Send marshals and sends the UpdateTerminationProtection API request.
func (r UpdateTerminationProtectionRequest) Send(ctx context.Context) (*UpdateTerminationProtectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTerminationProtectionResponse{
		UpdateTerminationProtectionOutput: r.Request.Data.(*UpdateTerminationProtectionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTerminationProtectionResponse is the response type for the
// UpdateTerminationProtection API operation.
type UpdateTerminationProtectionResponse struct {
	*UpdateTerminationProtectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTerminationProtection request.
func (r *UpdateTerminationProtectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
