// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListFunctionDefinitionsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListFunctionDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFunctionDefinitionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListFunctionDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	Definitions []DefinitionInformation `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListFunctionDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFunctionDefinitionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Definitions != nil {
		v := s.Definitions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Definitions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFunctionDefinitions = "ListFunctionDefinitions"

// ListFunctionDefinitionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves a list of Lambda function definitions.
//
//    // Example sending a request using ListFunctionDefinitionsRequest.
//    req := client.ListFunctionDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListFunctionDefinitions
func (c *Client) ListFunctionDefinitionsRequest(input *ListFunctionDefinitionsInput) ListFunctionDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListFunctionDefinitions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/functions",
	}

	if input == nil {
		input = &ListFunctionDefinitionsInput{}
	}

	req := c.newRequest(op, input, &ListFunctionDefinitionsOutput{})
	return ListFunctionDefinitionsRequest{Request: req, Input: input, Copy: c.ListFunctionDefinitionsRequest}
}

// ListFunctionDefinitionsRequest is the request type for the
// ListFunctionDefinitions API operation.
type ListFunctionDefinitionsRequest struct {
	*aws.Request
	Input *ListFunctionDefinitionsInput
	Copy  func(*ListFunctionDefinitionsInput) ListFunctionDefinitionsRequest
}

// Send marshals and sends the ListFunctionDefinitions API request.
func (r ListFunctionDefinitionsRequest) Send(ctx context.Context) (*ListFunctionDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFunctionDefinitionsResponse{
		ListFunctionDefinitionsOutput: r.Request.Data.(*ListFunctionDefinitionsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListFunctionDefinitionsResponse is the response type for the
// ListFunctionDefinitions API operation.
type ListFunctionDefinitionsResponse struct {
	*ListFunctionDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFunctionDefinitions request.
func (r *ListFunctionDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
