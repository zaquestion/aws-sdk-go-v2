// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type UpdateMyUserProfileInput struct {
	_ struct{} `type:"structure"`

	// The user's SSH public key.
	SshPublicKey *string `type:"string"`
}

// String returns the string representation
func (s UpdateMyUserProfileInput) String() string {
	return awsutil.Prettify(s)
}

type UpdateMyUserProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateMyUserProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateMyUserProfile = "UpdateMyUserProfile"

// UpdateMyUserProfileRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Updates a user's SSH public key.
//
// Required Permissions: To use this action, an IAM user must have self-management
// enabled or an attached policy that explicitly grants permissions. For more
// information about user permissions, see Managing User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using UpdateMyUserProfileRequest.
//    req := client.UpdateMyUserProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateMyUserProfile
func (c *Client) UpdateMyUserProfileRequest(input *UpdateMyUserProfileInput) UpdateMyUserProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateMyUserProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateMyUserProfileInput{}
	}

	req := c.newRequest(op, input, &UpdateMyUserProfileOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateMyUserProfileRequest{Request: req, Input: input, Copy: c.UpdateMyUserProfileRequest}
}

// UpdateMyUserProfileRequest is the request type for the
// UpdateMyUserProfile API operation.
type UpdateMyUserProfileRequest struct {
	*aws.Request
	Input *UpdateMyUserProfileInput
	Copy  func(*UpdateMyUserProfileInput) UpdateMyUserProfileRequest
}

// Send marshals and sends the UpdateMyUserProfile API request.
func (r UpdateMyUserProfileRequest) Send(ctx context.Context) (*UpdateMyUserProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMyUserProfileResponse{
		UpdateMyUserProfileOutput: r.Request.Data.(*UpdateMyUserProfileOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMyUserProfileResponse is the response type for the
// UpdateMyUserProfile API operation.
type UpdateMyUserProfileResponse struct {
	*UpdateMyUserProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMyUserProfile request.
func (r *UpdateMyUserProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
