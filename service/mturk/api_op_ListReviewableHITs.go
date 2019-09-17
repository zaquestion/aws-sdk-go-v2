// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListReviewableHITsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the HIT type of the HITs to consider for the query. If not specified,
	// all HITs for the Reviewer are considered
	HITTypeId *string `min:"1" type:"string"`

	// Limit the number of results returned.
	MaxResults *int64 `min:"1" type:"integer"`

	// Pagination Token
	NextToken *string `min:"1" type:"string"`

	// Can be either Reviewable or Reviewing. Reviewable is the default value.
	Status ReviewableHITStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListReviewableHITsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListReviewableHITsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListReviewableHITsInput"}
	if s.HITTypeId != nil && len(*s.HITTypeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITTypeId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListReviewableHITsOutput struct {
	_ struct{} `type:"structure"`

	// The list of HIT elements returned by the query.
	HITs []HIT `type:"list"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The number of HITs on this page in the filtered results list, equivalent
	// to the number of HITs being returned by this call.
	NumResults *int64 `type:"integer"`
}

// String returns the string representation
func (s ListReviewableHITsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListReviewableHITs = "ListReviewableHITs"

// ListReviewableHITsRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The ListReviewableHITs operation retrieves the HITs with Status equal to
// Reviewable or Status equal to Reviewing that belong to the Requester calling
// the operation.
//
//    // Example sending a request using ListReviewableHITsRequest.
//    req := client.ListReviewableHITsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/ListReviewableHITs
func (c *Client) ListReviewableHITsRequest(input *ListReviewableHITsInput) ListReviewableHITsRequest {
	op := &aws.Operation{
		Name:       opListReviewableHITs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListReviewableHITsInput{}
	}

	req := c.newRequest(op, input, &ListReviewableHITsOutput{})
	return ListReviewableHITsRequest{Request: req, Input: input, Copy: c.ListReviewableHITsRequest}
}

// ListReviewableHITsRequest is the request type for the
// ListReviewableHITs API operation.
type ListReviewableHITsRequest struct {
	*aws.Request
	Input *ListReviewableHITsInput
	Copy  func(*ListReviewableHITsInput) ListReviewableHITsRequest
}

// Send marshals and sends the ListReviewableHITs API request.
func (r ListReviewableHITsRequest) Send(ctx context.Context) (*ListReviewableHITsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListReviewableHITsResponse{
		ListReviewableHITsOutput: r.Request.Data.(*ListReviewableHITsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListReviewableHITsRequestPaginator returns a paginator for ListReviewableHITs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListReviewableHITsRequest(input)
//   p := mturk.NewListReviewableHITsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListReviewableHITsPaginator(req ListReviewableHITsRequest) ListReviewableHITsPaginator {
	return ListReviewableHITsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListReviewableHITsInput
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

// ListReviewableHITsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListReviewableHITsPaginator struct {
	aws.Pager
}

func (p *ListReviewableHITsPaginator) CurrentPage() *ListReviewableHITsOutput {
	return p.Pager.CurrentPage().(*ListReviewableHITsOutput)
}

// ListReviewableHITsResponse is the response type for the
// ListReviewableHITs API operation.
type ListReviewableHITsResponse struct {
	*ListReviewableHITsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListReviewableHITs request.
func (r *ListReviewableHITsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
