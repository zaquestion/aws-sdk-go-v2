// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetRouteResponsesInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"maxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// RouteId is a required field
	RouteId *string `location:"uri" locationName:"routeId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRouteResponsesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRouteResponsesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRouteResponsesInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.RouteId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRouteResponsesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteId != nil {
		v := *s.RouteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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

type GetRouteResponsesOutput struct {
	_ struct{} `type:"structure"`

	Items []RouteResponse `locationName:"items" type:"list"`

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetRouteResponsesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRouteResponsesOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opGetRouteResponses = "GetRouteResponses"

// GetRouteResponsesRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets the RouteResponses for a Route.
//
//    // Example sending a request using GetRouteResponsesRequest.
//    req := client.GetRouteResponsesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetRouteResponses
func (c *Client) GetRouteResponsesRequest(input *GetRouteResponsesInput) GetRouteResponsesRequest {
	op := &aws.Operation{
		Name:       opGetRouteResponses,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}/routeresponses",
	}

	if input == nil {
		input = &GetRouteResponsesInput{}
	}

	req := c.newRequest(op, input, &GetRouteResponsesOutput{})
	return GetRouteResponsesRequest{Request: req, Input: input, Copy: c.GetRouteResponsesRequest}
}

// GetRouteResponsesRequest is the request type for the
// GetRouteResponses API operation.
type GetRouteResponsesRequest struct {
	*aws.Request
	Input *GetRouteResponsesInput
	Copy  func(*GetRouteResponsesInput) GetRouteResponsesRequest
}

// Send marshals and sends the GetRouteResponses API request.
func (r GetRouteResponsesRequest) Send(ctx context.Context) (*GetRouteResponsesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRouteResponsesResponse{
		GetRouteResponsesOutput: r.Request.Data.(*GetRouteResponsesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRouteResponsesResponse is the response type for the
// GetRouteResponses API operation.
type GetRouteResponsesResponse struct {
	*GetRouteResponsesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRouteResponses request.
func (r *GetRouteResponsesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
