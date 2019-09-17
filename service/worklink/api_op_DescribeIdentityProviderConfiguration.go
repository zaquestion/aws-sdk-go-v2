// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeIdentityProviderConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeIdentityProviderConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIdentityProviderConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeIdentityProviderConfigurationInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeIdentityProviderConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeIdentityProviderConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The SAML metadata document provided by the user’s identity provider.
	IdentityProviderSamlMetadata *string `min:"1" type:"string"`

	// The type of identity provider.
	IdentityProviderType IdentityProviderType `type:"string" enum:"true"`

	// The SAML metadata document uploaded to the user’s identity provider.
	ServiceProviderSamlMetadata *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeIdentityProviderConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeIdentityProviderConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IdentityProviderSamlMetadata != nil {
		v := *s.IdentityProviderSamlMetadata

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityProviderSamlMetadata", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.IdentityProviderType) > 0 {
		v := s.IdentityProviderType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityProviderType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ServiceProviderSamlMetadata != nil {
		v := *s.ServiceProviderSamlMetadata

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ServiceProviderSamlMetadata", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeIdentityProviderConfiguration = "DescribeIdentityProviderConfiguration"

// DescribeIdentityProviderConfigurationRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Describes the identity provider configuration of the specified fleet.
//
//    // Example sending a request using DescribeIdentityProviderConfigurationRequest.
//    req := client.DescribeIdentityProviderConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeIdentityProviderConfiguration
func (c *Client) DescribeIdentityProviderConfigurationRequest(input *DescribeIdentityProviderConfigurationInput) DescribeIdentityProviderConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeIdentityProviderConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/describeIdentityProviderConfiguration",
	}

	if input == nil {
		input = &DescribeIdentityProviderConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeIdentityProviderConfigurationOutput{})
	return DescribeIdentityProviderConfigurationRequest{Request: req, Input: input, Copy: c.DescribeIdentityProviderConfigurationRequest}
}

// DescribeIdentityProviderConfigurationRequest is the request type for the
// DescribeIdentityProviderConfiguration API operation.
type DescribeIdentityProviderConfigurationRequest struct {
	*aws.Request
	Input *DescribeIdentityProviderConfigurationInput
	Copy  func(*DescribeIdentityProviderConfigurationInput) DescribeIdentityProviderConfigurationRequest
}

// Send marshals and sends the DescribeIdentityProviderConfiguration API request.
func (r DescribeIdentityProviderConfigurationRequest) Send(ctx context.Context) (*DescribeIdentityProviderConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIdentityProviderConfigurationResponse{
		DescribeIdentityProviderConfigurationOutput: r.Request.Data.(*DescribeIdentityProviderConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeIdentityProviderConfigurationResponse is the response type for the
// DescribeIdentityProviderConfiguration API operation.
type DescribeIdentityProviderConfigurationResponse struct {
	*DescribeIdentityProviderConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIdentityProviderConfiguration request.
func (r *DescribeIdentityProviderConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
