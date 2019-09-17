// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRiskConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The app client ID.
	ClientId *string `min:"1" type:"string"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRiskConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRiskConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRiskConfigurationInput"}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeRiskConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The risk configuration.
	//
	// RiskConfiguration is a required field
	RiskConfiguration *RiskConfigurationType `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeRiskConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRiskConfiguration = "DescribeRiskConfiguration"

// DescribeRiskConfigurationRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Describes the risk configuration.
//
//    // Example sending a request using DescribeRiskConfigurationRequest.
//    req := client.DescribeRiskConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DescribeRiskConfiguration
func (c *Client) DescribeRiskConfigurationRequest(input *DescribeRiskConfigurationInput) DescribeRiskConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeRiskConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRiskConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeRiskConfigurationOutput{})
	return DescribeRiskConfigurationRequest{Request: req, Input: input, Copy: c.DescribeRiskConfigurationRequest}
}

// DescribeRiskConfigurationRequest is the request type for the
// DescribeRiskConfiguration API operation.
type DescribeRiskConfigurationRequest struct {
	*aws.Request
	Input *DescribeRiskConfigurationInput
	Copy  func(*DescribeRiskConfigurationInput) DescribeRiskConfigurationRequest
}

// Send marshals and sends the DescribeRiskConfiguration API request.
func (r DescribeRiskConfigurationRequest) Send(ctx context.Context) (*DescribeRiskConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRiskConfigurationResponse{
		DescribeRiskConfigurationOutput: r.Request.Data.(*DescribeRiskConfigurationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRiskConfigurationResponse is the response type for the
// DescribeRiskConfiguration API operation.
type DescribeRiskConfigurationResponse struct {
	*DescribeRiskConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRiskConfiguration request.
func (r *DescribeRiskConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
