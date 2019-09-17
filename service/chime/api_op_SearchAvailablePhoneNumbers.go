// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type SearchAvailablePhoneNumbersInput struct {
	_ struct{} `type:"structure"`

	// The area code used to filter results.
	AreaCode *string `location:"querystring" locationName:"area-code" type:"string"`

	// The city used to filter results.
	City *string `location:"querystring" locationName:"city" type:"string"`

	// The country used to filter results.
	Country *string `location:"querystring" locationName:"country" type:"string"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The token to use to retrieve the next page of results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	// The state used to filter results.
	State *string `location:"querystring" locationName:"state" type:"string"`

	// The toll-free prefix that you use to filter results.
	TollFreePrefix *string `location:"querystring" locationName:"toll-free-prefix" min:"3" type:"string"`
}

// String returns the string representation
func (s SearchAvailablePhoneNumbersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchAvailablePhoneNumbersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchAvailablePhoneNumbersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.TollFreePrefix != nil && len(*s.TollFreePrefix) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TollFreePrefix", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SearchAvailablePhoneNumbersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AreaCode != nil {
		v := *s.AreaCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "area-code", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.City != nil {
		v := *s.City

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "city", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Country != nil {
		v := *s.Country

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "country", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.State != nil {
		v := *s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "state", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TollFreePrefix != nil {
		v := *s.TollFreePrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "toll-free-prefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type SearchAvailablePhoneNumbersOutput struct {
	_ struct{} `type:"structure"`

	// List of phone numbers, in E.164 format.
	E164PhoneNumbers []string `type:"list"`
}

// String returns the string representation
func (s SearchAvailablePhoneNumbersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SearchAvailablePhoneNumbersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.E164PhoneNumbers != nil {
		v := s.E164PhoneNumbers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "E164PhoneNumbers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opSearchAvailablePhoneNumbers = "SearchAvailablePhoneNumbers"

// SearchAvailablePhoneNumbersRequest returns a request value for making API operation for
// Amazon Chime.
//
// Searches phone numbers that can be ordered.
//
//    // Example sending a request using SearchAvailablePhoneNumbersRequest.
//    req := client.SearchAvailablePhoneNumbersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/SearchAvailablePhoneNumbers
func (c *Client) SearchAvailablePhoneNumbersRequest(input *SearchAvailablePhoneNumbersInput) SearchAvailablePhoneNumbersRequest {
	op := &aws.Operation{
		Name:       opSearchAvailablePhoneNumbers,
		HTTPMethod: "GET",
		HTTPPath:   "/search?type=phone-numbers",
	}

	if input == nil {
		input = &SearchAvailablePhoneNumbersInput{}
	}

	req := c.newRequest(op, input, &SearchAvailablePhoneNumbersOutput{})
	return SearchAvailablePhoneNumbersRequest{Request: req, Input: input, Copy: c.SearchAvailablePhoneNumbersRequest}
}

// SearchAvailablePhoneNumbersRequest is the request type for the
// SearchAvailablePhoneNumbers API operation.
type SearchAvailablePhoneNumbersRequest struct {
	*aws.Request
	Input *SearchAvailablePhoneNumbersInput
	Copy  func(*SearchAvailablePhoneNumbersInput) SearchAvailablePhoneNumbersRequest
}

// Send marshals and sends the SearchAvailablePhoneNumbers API request.
func (r SearchAvailablePhoneNumbersRequest) Send(ctx context.Context) (*SearchAvailablePhoneNumbersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchAvailablePhoneNumbersResponse{
		SearchAvailablePhoneNumbersOutput: r.Request.Data.(*SearchAvailablePhoneNumbersOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SearchAvailablePhoneNumbersResponse is the response type for the
// SearchAvailablePhoneNumbers API operation.
type SearchAvailablePhoneNumbersResponse struct {
	*SearchAvailablePhoneNumbersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchAvailablePhoneNumbers request.
func (r *SearchAvailablePhoneNumbersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
