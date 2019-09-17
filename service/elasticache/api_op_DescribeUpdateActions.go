// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeUpdateActionsInput struct {
	_ struct{} `type:"structure"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response
	MaxRecords *int64 `type:"integer"`

	// The replication group IDs
	ReplicationGroupIds []string `type:"list"`

	// The unique ID of the service update
	ServiceUpdateName *string `type:"string"`

	// The status of the service update
	ServiceUpdateStatus []ServiceUpdateStatus `type:"list"`

	// The range of time specified to search for service updates that are in available
	// status
	ServiceUpdateTimeRange *TimeRangeFilter `type:"structure"`

	// Dictates whether to include node level update status in the response
	ShowNodeLevelUpdateStatus *bool `type:"boolean"`

	// The status of the update action.
	UpdateActionStatus []UpdateActionStatus `type:"list"`
}

// String returns the string representation
func (s DescribeUpdateActionsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeUpdateActionsOutput struct {
	_ struct{} `type:"structure"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// Returns a list of update actions
	UpdateActions []UpdateAction `locationNameList:"UpdateAction" type:"list"`
}

// String returns the string representation
func (s DescribeUpdateActionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUpdateActions = "DescribeUpdateActions"

// DescribeUpdateActionsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns details of the update actions
//
//    // Example sending a request using DescribeUpdateActionsRequest.
//    req := client.DescribeUpdateActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeUpdateActions
func (c *Client) DescribeUpdateActionsRequest(input *DescribeUpdateActionsInput) DescribeUpdateActionsRequest {
	op := &aws.Operation{
		Name:       opDescribeUpdateActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeUpdateActionsInput{}
	}

	req := c.newRequest(op, input, &DescribeUpdateActionsOutput{})
	return DescribeUpdateActionsRequest{Request: req, Input: input, Copy: c.DescribeUpdateActionsRequest}
}

// DescribeUpdateActionsRequest is the request type for the
// DescribeUpdateActions API operation.
type DescribeUpdateActionsRequest struct {
	*aws.Request
	Input *DescribeUpdateActionsInput
	Copy  func(*DescribeUpdateActionsInput) DescribeUpdateActionsRequest
}

// Send marshals and sends the DescribeUpdateActions API request.
func (r DescribeUpdateActionsRequest) Send(ctx context.Context) (*DescribeUpdateActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUpdateActionsResponse{
		DescribeUpdateActionsOutput: r.Request.Data.(*DescribeUpdateActionsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeUpdateActionsRequestPaginator returns a paginator for DescribeUpdateActions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeUpdateActionsRequest(input)
//   p := elasticache.NewDescribeUpdateActionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeUpdateActionsPaginator(req DescribeUpdateActionsRequest) DescribeUpdateActionsPaginator {
	return DescribeUpdateActionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeUpdateActionsInput
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

// DescribeUpdateActionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeUpdateActionsPaginator struct {
	aws.Pager
}

func (p *DescribeUpdateActionsPaginator) CurrentPage() *DescribeUpdateActionsOutput {
	return p.Pager.CurrentPage().(*DescribeUpdateActionsOutput)
}

// DescribeUpdateActionsResponse is the response type for the
// DescribeUpdateActions API operation.
type DescribeUpdateActionsResponse struct {
	*DescribeUpdateActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUpdateActions request.
func (r *DescribeUpdateActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
