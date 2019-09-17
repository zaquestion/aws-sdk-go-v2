// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request that enables the user as an administrator.
type AdminEnableUserInput struct {
	_ struct{} `type:"structure"`

	// The user pool ID for the user pool where you want to enable the user.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user name of the user you wish to enable.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AdminEnableUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminEnableUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminEnableUserInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server for the request to enable a user
// as an administrator.
type AdminEnableUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AdminEnableUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminEnableUser = "AdminEnableUser"

// AdminEnableUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Enables the specified user as an administrator. Works on any user.
//
// Requires developer credentials.
//
//    // Example sending a request using AdminEnableUserRequest.
//    req := client.AdminEnableUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminEnableUser
func (c *Client) AdminEnableUserRequest(input *AdminEnableUserInput) AdminEnableUserRequest {
	op := &aws.Operation{
		Name:       opAdminEnableUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminEnableUserInput{}
	}

	req := c.newRequest(op, input, &AdminEnableUserOutput{})
	return AdminEnableUserRequest{Request: req, Input: input, Copy: c.AdminEnableUserRequest}
}

// AdminEnableUserRequest is the request type for the
// AdminEnableUser API operation.
type AdminEnableUserRequest struct {
	*aws.Request
	Input *AdminEnableUserInput
	Copy  func(*AdminEnableUserInput) AdminEnableUserRequest
}

// Send marshals and sends the AdminEnableUser API request.
func (r AdminEnableUserRequest) Send(ctx context.Context) (*AdminEnableUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminEnableUserResponse{
		AdminEnableUserOutput: r.Request.Data.(*AdminEnableUserOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminEnableUserResponse is the response type for the
// AdminEnableUser API operation.
type AdminEnableUserResponse struct {
	*AdminEnableUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminEnableUser request.
func (r *AdminEnableUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
