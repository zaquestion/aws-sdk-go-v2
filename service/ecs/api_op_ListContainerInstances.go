// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListContainerInstancesInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instances to list. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// You can filter the results of a ListContainerInstances operation with cluster
	// query language statements. For more information, see Cluster Query Language
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html)
	// in the Amazon Elastic Container Service Developer Guide.
	Filter *string `locationName:"filter" type:"string"`

	// The maximum number of container instance results returned by ListContainerInstances
	// in paginated output. When this parameter is used, ListContainerInstances
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListContainerInstances request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListContainerInstances returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListContainerInstances
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Filters the container instances by status. For example, if you specify the
	// DRAINING status, the results include only container instances that have been
	// set to DRAINING using UpdateContainerInstancesState. If you do not specify
	// this parameter, the default is to include container instances set to all
	// states other than INACTIVE.
	Status ContainerInstanceStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListContainerInstancesInput) String() string {
	return awsutil.Prettify(s)
}

type ListContainerInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The list of container instances with full ARN entries for each container
	// instance associated with the specified cluster.
	ContainerInstanceArns []string `locationName:"containerInstanceArns" type:"list"`

	// The nextToken value to include in a future ListContainerInstances request.
	// When the results of a ListContainerInstances request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListContainerInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListContainerInstances = "ListContainerInstances"

// ListContainerInstancesRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Returns a list of container instances in a specified cluster. You can filter
// the results of a ListContainerInstances operation with cluster query language
// statements inside the filter parameter. For more information, see Cluster
// Query Language (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using ListContainerInstancesRequest.
//    req := client.ListContainerInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/ListContainerInstances
func (c *Client) ListContainerInstancesRequest(input *ListContainerInstancesInput) ListContainerInstancesRequest {
	op := &aws.Operation{
		Name:       opListContainerInstances,
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
		input = &ListContainerInstancesInput{}
	}

	req := c.newRequest(op, input, &ListContainerInstancesOutput{})
	return ListContainerInstancesRequest{Request: req, Input: input, Copy: c.ListContainerInstancesRequest}
}

// ListContainerInstancesRequest is the request type for the
// ListContainerInstances API operation.
type ListContainerInstancesRequest struct {
	*aws.Request
	Input *ListContainerInstancesInput
	Copy  func(*ListContainerInstancesInput) ListContainerInstancesRequest
}

// Send marshals and sends the ListContainerInstances API request.
func (r ListContainerInstancesRequest) Send(ctx context.Context) (*ListContainerInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListContainerInstancesResponse{
		ListContainerInstancesOutput: r.Request.Data.(*ListContainerInstancesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListContainerInstancesRequestPaginator returns a paginator for ListContainerInstances.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListContainerInstancesRequest(input)
//   p := ecs.NewListContainerInstancesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListContainerInstancesPaginator(req ListContainerInstancesRequest) ListContainerInstancesPaginator {
	return ListContainerInstancesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListContainerInstancesInput
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

// ListContainerInstancesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListContainerInstancesPaginator struct {
	aws.Pager
}

func (p *ListContainerInstancesPaginator) CurrentPage() *ListContainerInstancesOutput {
	return p.Pager.CurrentPage().(*ListContainerInstancesOutput)
}

// ListContainerInstancesResponse is the response type for the
// ListContainerInstances API operation.
type ListContainerInstancesResponse struct {
	*ListContainerInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListContainerInstances request.
func (r *ListContainerInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
