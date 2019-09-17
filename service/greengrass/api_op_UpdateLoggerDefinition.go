// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateLoggerDefinitionInput struct {
	_ struct{} `type:"structure"`

	// LoggerDefinitionId is a required field
	LoggerDefinitionId *string `location:"uri" locationName:"LoggerDefinitionId" type:"string" required:"true"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateLoggerDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLoggerDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLoggerDefinitionInput"}

	if s.LoggerDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoggerDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateLoggerDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LoggerDefinitionId != nil {
		v := *s.LoggerDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LoggerDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateLoggerDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateLoggerDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateLoggerDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateLoggerDefinition = "UpdateLoggerDefinition"

// UpdateLoggerDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Updates a logger definition.
//
//    // Example sending a request using UpdateLoggerDefinitionRequest.
//    req := client.UpdateLoggerDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateLoggerDefinition
func (c *Client) UpdateLoggerDefinitionRequest(input *UpdateLoggerDefinitionInput) UpdateLoggerDefinitionRequest {
	op := &aws.Operation{
		Name:       opUpdateLoggerDefinition,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/definition/loggers/{LoggerDefinitionId}",
	}

	if input == nil {
		input = &UpdateLoggerDefinitionInput{}
	}

	req := c.newRequest(op, input, &UpdateLoggerDefinitionOutput{})
	return UpdateLoggerDefinitionRequest{Request: req, Input: input, Copy: c.UpdateLoggerDefinitionRequest}
}

// UpdateLoggerDefinitionRequest is the request type for the
// UpdateLoggerDefinition API operation.
type UpdateLoggerDefinitionRequest struct {
	*aws.Request
	Input *UpdateLoggerDefinitionInput
	Copy  func(*UpdateLoggerDefinitionInput) UpdateLoggerDefinitionRequest
}

// Send marshals and sends the UpdateLoggerDefinition API request.
func (r UpdateLoggerDefinitionRequest) Send(ctx context.Context) (*UpdateLoggerDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLoggerDefinitionResponse{
		UpdateLoggerDefinitionOutput: r.Request.Data.(*UpdateLoggerDefinitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLoggerDefinitionResponse is the response type for the
// UpdateLoggerDefinition API operation.
type UpdateLoggerDefinitionResponse struct {
	*UpdateLoggerDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLoggerDefinition request.
func (r *UpdateLoggerDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
