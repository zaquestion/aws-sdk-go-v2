// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to obtain a list of dedicated IP pools.
type ListDedicatedIpPoolsInput struct {
	_ struct{} `type:"structure"`

	// A token returned from a previous call to ListDedicatedIpPools to indicate
	// the position in the list of dedicated IP pools.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// The number of results to show in a single call to ListDedicatedIpPools. If
	// the number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int64 `location:"querystring" locationName:"PageSize" type:"integer"`
}

// String returns the string representation
func (s ListDedicatedIpPoolsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDedicatedIpPoolsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "PageSize", protocol.Int64Value(v), metadata)
	}
	return nil
}

// A list of dedicated IP pools.
type ListDedicatedIpPoolsOutput struct {
	_ struct{} `type:"structure"`

	// A list of all of the dedicated IP pools that are associated with your Amazon
	// Pinpoint account.
	DedicatedIpPools []string `type:"list"`

	// A token that indicates that there are additional IP pools to list. To view
	// additional IP pools, issue another request to ListDedicatedIpPools, passing
	// this token in the NextToken parameter.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDedicatedIpPoolsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDedicatedIpPoolsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DedicatedIpPools != nil {
		v := s.DedicatedIpPools

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DedicatedIpPools", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
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

const opListDedicatedIpPools = "ListDedicatedIpPools"

// ListDedicatedIpPoolsRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// List all of the dedicated IP pools that exist in your Amazon Pinpoint account
// in the current AWS Region.
//
//    // Example sending a request using ListDedicatedIpPoolsRequest.
//    req := client.ListDedicatedIpPoolsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListDedicatedIpPools
func (c *Client) ListDedicatedIpPoolsRequest(input *ListDedicatedIpPoolsInput) ListDedicatedIpPoolsRequest {
	op := &aws.Operation{
		Name:       opListDedicatedIpPools,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/dedicated-ip-pools",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDedicatedIpPoolsInput{}
	}

	req := c.newRequest(op, input, &ListDedicatedIpPoolsOutput{})
	return ListDedicatedIpPoolsRequest{Request: req, Input: input, Copy: c.ListDedicatedIpPoolsRequest}
}

// ListDedicatedIpPoolsRequest is the request type for the
// ListDedicatedIpPools API operation.
type ListDedicatedIpPoolsRequest struct {
	*aws.Request
	Input *ListDedicatedIpPoolsInput
	Copy  func(*ListDedicatedIpPoolsInput) ListDedicatedIpPoolsRequest
}

// Send marshals and sends the ListDedicatedIpPools API request.
func (r ListDedicatedIpPoolsRequest) Send(ctx context.Context) (*ListDedicatedIpPoolsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDedicatedIpPoolsResponse{
		ListDedicatedIpPoolsOutput: r.Request.Data.(*ListDedicatedIpPoolsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDedicatedIpPoolsRequestPaginator returns a paginator for ListDedicatedIpPools.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDedicatedIpPoolsRequest(input)
//   p := pinpointemail.NewListDedicatedIpPoolsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDedicatedIpPoolsPaginator(req ListDedicatedIpPoolsRequest) ListDedicatedIpPoolsPaginator {
	return ListDedicatedIpPoolsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDedicatedIpPoolsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDedicatedIpPoolsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDedicatedIpPoolsPaginator struct {
	aws.Pager
}

func (p *ListDedicatedIpPoolsPaginator) CurrentPage() *ListDedicatedIpPoolsOutput {
	return p.Pager.CurrentPage().(*ListDedicatedIpPoolsOutput)
}

// ListDedicatedIpPoolsResponse is the response type for the
// ListDedicatedIpPools API operation.
type ListDedicatedIpPoolsResponse struct {
	*ListDedicatedIpPoolsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDedicatedIpPools request.
func (r *ListDedicatedIpPoolsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
