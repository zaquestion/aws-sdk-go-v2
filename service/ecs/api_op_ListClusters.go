// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListClustersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of cluster results returned by ListClusters in paginated
	// output. When this parameter is used, ListClusters only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListClusters
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then ListClusters returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListClusters request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClustersInput) String() string {
	return awsutil.Prettify(s)
}

type ListClustersOutput struct {
	_ struct{} `type:"structure"`

	// The list of full Amazon Resource Name (ARN) entries for each cluster associated
	// with your account.
	ClusterArns []string `locationName:"clusterArns" type:"list"`

	// The nextToken value to include in a future ListClusters request. When the
	// results of a ListClusters request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClustersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListClusters = "ListClusters"

// ListClustersRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Returns a list of existing clusters.
//
//    // Example sending a request using ListClustersRequest.
//    req := client.ListClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/ListClusters
func (c *Client) ListClustersRequest(input *ListClustersInput) ListClustersRequest {
	op := &aws.Operation{
		Name:       opListClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListClustersInput{}
	}

	req := c.newRequest(op, input, &ListClustersOutput{})
	return ListClustersRequest{Request: req, Input: input, Copy: c.ListClustersRequest}
}

// ListClustersRequest is the request type for the
// ListClusters API operation.
type ListClustersRequest struct {
	*aws.Request
	Input *ListClustersInput
	Copy  func(*ListClustersInput) ListClustersRequest
}

// Send marshals and sends the ListClusters API request.
func (r ListClustersRequest) Send(ctx context.Context) (*ListClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListClustersResponse{
		ListClustersOutput: r.Request.Data.(*ListClustersOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListClustersRequestPaginator returns a paginator for ListClusters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListClustersRequest(input)
//   p := ecs.NewListClustersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListClustersPaginator(req ListClustersRequest) ListClustersPaginator {
	return ListClustersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListClustersInput
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

// ListClustersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListClustersPaginator struct {
	aws.Pager
}

func (p *ListClustersPaginator) CurrentPage() *ListClustersOutput {
	return p.Pager.CurrentPage().(*ListClustersOutput)
}

// ListClustersResponse is the response type for the
// ListClusters API operation.
type ListClustersResponse struct {
	*ListClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListClusters request.
func (r *ListClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
