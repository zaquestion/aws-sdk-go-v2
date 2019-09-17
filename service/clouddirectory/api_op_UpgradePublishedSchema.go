// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpgradePublishedSchemaInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the development schema with the changes used for the upgrade.
	//
	// DevelopmentSchemaArn is a required field
	DevelopmentSchemaArn *string `type:"string" required:"true"`

	// Used for testing whether the Development schema provided is backwards compatible,
	// or not, with the publish schema provided by the user to be upgraded. If schema
	// compatibility fails, an exception would be thrown else the call would succeed.
	// This parameter is optional and defaults to false.
	DryRun *bool `type:"boolean"`

	// Identifies the minor version of the published schema that will be created.
	// This parameter is NOT optional.
	//
	// MinorVersion is a required field
	MinorVersion *string `min:"1" type:"string" required:"true"`

	// The ARN of the published schema to be upgraded.
	//
	// PublishedSchemaArn is a required field
	PublishedSchemaArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpgradePublishedSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpgradePublishedSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpgradePublishedSchemaInput"}

	if s.DevelopmentSchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DevelopmentSchemaArn"))
	}

	if s.MinorVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("MinorVersion"))
	}
	if s.MinorVersion != nil && len(*s.MinorVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MinorVersion", 1))
	}

	if s.PublishedSchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublishedSchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpgradePublishedSchemaInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DevelopmentSchemaArn != nil {
		v := *s.DevelopmentSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DevelopmentSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DryRun != nil {
		v := *s.DryRun

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DryRun", protocol.BoolValue(v), metadata)
	}
	if s.MinorVersion != nil {
		v := *s.MinorVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MinorVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PublishedSchemaArn != nil {
		v := *s.PublishedSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PublishedSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpgradePublishedSchemaOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the upgraded schema that is returned as part of the response.
	UpgradedSchemaArn *string `type:"string"`
}

// String returns the string representation
func (s UpgradePublishedSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpgradePublishedSchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UpgradedSchemaArn != nil {
		v := *s.UpgradedSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UpgradedSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpgradePublishedSchema = "UpgradePublishedSchema"

// UpgradePublishedSchemaRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Upgrades a published schema under a new minor version revision using the
// current contents of DevelopmentSchemaArn.
//
//    // Example sending a request using UpgradePublishedSchemaRequest.
//    req := client.UpgradePublishedSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/UpgradePublishedSchema
func (c *Client) UpgradePublishedSchemaRequest(input *UpgradePublishedSchemaInput) UpgradePublishedSchemaRequest {
	op := &aws.Operation{
		Name:       opUpgradePublishedSchema,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/upgradepublished",
	}

	if input == nil {
		input = &UpgradePublishedSchemaInput{}
	}

	req := c.newRequest(op, input, &UpgradePublishedSchemaOutput{})
	return UpgradePublishedSchemaRequest{Request: req, Input: input, Copy: c.UpgradePublishedSchemaRequest}
}

// UpgradePublishedSchemaRequest is the request type for the
// UpgradePublishedSchema API operation.
type UpgradePublishedSchemaRequest struct {
	*aws.Request
	Input *UpgradePublishedSchemaInput
	Copy  func(*UpgradePublishedSchemaInput) UpgradePublishedSchemaRequest
}

// Send marshals and sends the UpgradePublishedSchema API request.
func (r UpgradePublishedSchemaRequest) Send(ctx context.Context) (*UpgradePublishedSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpgradePublishedSchemaResponse{
		UpgradePublishedSchemaOutput: r.Request.Data.(*UpgradePublishedSchemaOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpgradePublishedSchemaResponse is the response type for the
// UpgradePublishedSchema API operation.
type UpgradePublishedSchemaResponse struct {
	*UpgradePublishedSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpgradePublishedSchema request.
func (r *UpgradePublishedSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
