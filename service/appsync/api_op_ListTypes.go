// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListTypesInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The type format: SDL or JSON.
	//
	// Format is a required field
	Format TypeDefinitionFormat `location:"querystring" locationName:"format" type:"string" required:"true" enum:"true"`

	// The maximum number of results you want the request to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTypesInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTypesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListTypesOutput struct {
	_ struct{} `type:"structure"`

	// An identifier to be passed in the next request to this operation to return
	// the next set of items in the list.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The Type objects.
	Types []Type `locationName:"types" type:"list"`
}

// String returns the string representation
func (s ListTypesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTypesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Types != nil {
		v := s.Types

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "types", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListTypes = "ListTypes"

// ListTypesRequest returns a request value for making API operation for
// AWS AppSync.
//
// Lists the types for a given API.
//
//    // Example sending a request using ListTypesRequest.
//    req := client.ListTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListTypes
func (c *Client) ListTypesRequest(input *ListTypesInput) ListTypesRequest {
	op := &aws.Operation{
		Name:       opListTypes,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/types",
	}

	if input == nil {
		input = &ListTypesInput{}
	}

	req := c.newRequest(op, input, &ListTypesOutput{})
	return ListTypesRequest{Request: req, Input: input, Copy: c.ListTypesRequest}
}

// ListTypesRequest is the request type for the
// ListTypes API operation.
type ListTypesRequest struct {
	*aws.Request
	Input *ListTypesInput
	Copy  func(*ListTypesInput) ListTypesRequest
}

// Send marshals and sends the ListTypes API request.
func (r ListTypesRequest) Send(ctx context.Context) (*ListTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTypesResponse{
		ListTypesOutput: r.Request.Data.(*ListTypesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTypesResponse is the response type for the
// ListTypes API operation.
type ListTypesResponse struct {
	*ListTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTypes request.
func (r *ListTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
