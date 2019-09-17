// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to update the user pool client.
type UpdateUserPoolClientInput struct {
	_ struct{} `type:"structure"`

	// Set to code to initiate a code grant flow, which provides an authorization
	// code as the response. This code can be exchanged for access tokens with the
	// token endpoint.
	AllowedOAuthFlows []OAuthFlowType `type:"list"`

	// Set to TRUE if the client is allowed to follow the OAuth protocol when interacting
	// with Cognito user pools.
	AllowedOAuthFlowsUserPoolClient *bool `type:"boolean"`

	// A list of allowed OAuth scopes. Currently supported values are "phone", "email",
	// "openid", and "Cognito".
	AllowedOAuthScopes []string `type:"list"`

	// The Amazon Pinpoint analytics configuration for collecting metrics for this
	// user pool.
	AnalyticsConfiguration *AnalyticsConfigurationType `type:"structure"`

	// A list of allowed redirect (callback) URLs for the identity providers.
	//
	// A redirect URI must:
	//
	//    * Be an absolute URI.
	//
	//    * Be registered with the authorization server.
	//
	//    * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2).
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	CallbackURLs []string `type:"list"`

	// The ID of the client associated with the user pool.
	//
	// ClientId is a required field
	ClientId *string `min:"1" type:"string" required:"true"`

	// The client name from the update user pool client request.
	ClientName *string `min:"1" type:"string"`

	// The default redirect URI. Must be in the CallbackURLs list.
	//
	// A redirect URI must:
	//
	//    * Be an absolute URI.
	//
	//    * Be registered with the authorization server.
	//
	//    * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2).
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectURI *string `min:"1" type:"string"`

	// Explicit authentication flows.
	ExplicitAuthFlows []ExplicitAuthFlowsType `type:"list"`

	// A list of allowed logout URLs for the identity providers.
	LogoutURLs []string `type:"list"`

	// The read-only attributes of the user pool.
	ReadAttributes []string `type:"list"`

	// The time limit, in days, after which the refresh token is no longer valid
	// and cannot be used.
	RefreshTokenValidity *int64 `type:"integer"`

	// A list of provider names for the identity providers that are supported on
	// this client.
	SupportedIdentityProviders []string `type:"list"`

	// The user pool ID for the user pool where you want to update the user pool
	// client.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The writeable attributes of the user pool.
	WriteAttributes []string `type:"list"`
}

// String returns the string representation
func (s UpdateUserPoolClientInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserPoolClientInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserPoolClientInput"}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}
	if s.ClientName != nil && len(*s.ClientName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientName", 1))
	}
	if s.DefaultRedirectURI != nil && len(*s.DefaultRedirectURI) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultRedirectURI", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.AnalyticsConfiguration != nil {
		if err := s.AnalyticsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("AnalyticsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server to the request to update the user
// pool client.
type UpdateUserPoolClientOutput struct {
	_ struct{} `type:"structure"`

	// The user pool client value from the response from the server when an update
	// user pool client request is made.
	UserPoolClient *UserPoolClientType `type:"structure"`
}

// String returns the string representation
func (s UpdateUserPoolClientOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateUserPoolClient = "UpdateUserPoolClient"

// UpdateUserPoolClientRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Updates the specified user pool app client with the specified attributes.
// If you don't provide a value for an attribute, it will be set to the default
// value. You can get a list of the current user pool app client settings with .
//
//    // Example sending a request using UpdateUserPoolClientRequest.
//    req := client.UpdateUserPoolClientRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/UpdateUserPoolClient
func (c *Client) UpdateUserPoolClientRequest(input *UpdateUserPoolClientInput) UpdateUserPoolClientRequest {
	op := &aws.Operation{
		Name:       opUpdateUserPoolClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateUserPoolClientInput{}
	}

	req := c.newRequest(op, input, &UpdateUserPoolClientOutput{})
	return UpdateUserPoolClientRequest{Request: req, Input: input, Copy: c.UpdateUserPoolClientRequest}
}

// UpdateUserPoolClientRequest is the request type for the
// UpdateUserPoolClient API operation.
type UpdateUserPoolClientRequest struct {
	*aws.Request
	Input *UpdateUserPoolClientInput
	Copy  func(*UpdateUserPoolClientInput) UpdateUserPoolClientRequest
}

// Send marshals and sends the UpdateUserPoolClient API request.
func (r UpdateUserPoolClientRequest) Send(ctx context.Context) (*UpdateUserPoolClientResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserPoolClientResponse{
		UpdateUserPoolClientOutput: r.Request.Data.(*UpdateUserPoolClientOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserPoolClientResponse is the response type for the
// UpdateUserPoolClient API operation.
type UpdateUserPoolClientResponse struct {
	*UpdateUserPoolClientOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserPoolClient request.
func (r *UpdateUserPoolClientResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
