// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetApisInput struct {
	_ struct{} `type:"structure"`

	MaxResults *string `location:"querystring" locationName:"maxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetApisInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApisInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetApisOutput struct {
	_ struct{} `type:"structure"`

	Items []Api `locationName:"items" type:"list"`

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetApisOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApisOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetApis = "GetApis"

// GetApisRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets a collection of Api resources.
//
//    // Example sending a request using GetApisRequest.
//    req := client.GetApisRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetApis
func (c *Client) GetApisRequest(input *GetApisInput) GetApisRequest {
	op := &aws.Operation{
		Name:       opGetApis,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis",
	}

	if input == nil {
		input = &GetApisInput{}
	}

	req := c.newRequest(op, input, &GetApisOutput{})
	return GetApisRequest{Request: req, Input: input, Copy: c.GetApisRequest}
}

// GetApisRequest is the request type for the
// GetApis API operation.
type GetApisRequest struct {
	*aws.Request
	Input *GetApisInput
	Copy  func(*GetApisInput) GetApisRequest
}

// Send marshals and sends the GetApis API request.
func (r GetApisRequest) Send(ctx context.Context) (*GetApisResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApisResponse{
		GetApisOutput: r.Request.Data.(*GetApisOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApisResponse is the response type for the
// GetApis API operation.
type GetApisResponse struct {
	*GetApisOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApis request.
func (r *GetApisResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
