// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListLabelingJobsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only labeling jobs created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only labeling jobs created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns only labeling jobs modified after the specified time
	// (timestamp).
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only labeling jobs modified before the specified time
	// (timestamp).
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of labeling jobs to return in each page of the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the labeling job name. This filter returns only labeling jobs
	// whose name contains the specified string.
	NameContains *string `type:"string"`

	// If the result of the previous ListLabelingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of labeling jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`

	// The field to sort results by. The default is CreationTime.
	SortBy SortBy `type:"string" enum:"true"`

	// The sort order for results. The default is Ascending.
	SortOrder SortOrder `type:"string" enum:"true"`

	// A filter that retrieves only labeling jobs with a specific status.
	StatusEquals LabelingJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListLabelingJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListLabelingJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListLabelingJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListLabelingJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of LabelingJobSummary objects, each describing a labeling job.
	LabelingJobSummaryList []LabelingJobSummary `type:"list"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of labeling jobs, use it in the subsequent request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListLabelingJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListLabelingJobs = "ListLabelingJobs"

// ListLabelingJobsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Gets a list of labeling jobs.
//
//    // Example sending a request using ListLabelingJobsRequest.
//    req := client.ListLabelingJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListLabelingJobs
func (c *Client) ListLabelingJobsRequest(input *ListLabelingJobsInput) ListLabelingJobsRequest {
	op := &aws.Operation{
		Name:       opListLabelingJobs,
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
		input = &ListLabelingJobsInput{}
	}

	req := c.newRequest(op, input, &ListLabelingJobsOutput{})
	return ListLabelingJobsRequest{Request: req, Input: input, Copy: c.ListLabelingJobsRequest}
}

// ListLabelingJobsRequest is the request type for the
// ListLabelingJobs API operation.
type ListLabelingJobsRequest struct {
	*aws.Request
	Input *ListLabelingJobsInput
	Copy  func(*ListLabelingJobsInput) ListLabelingJobsRequest
}

// Send marshals and sends the ListLabelingJobs API request.
func (r ListLabelingJobsRequest) Send(ctx context.Context) (*ListLabelingJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListLabelingJobsResponse{
		ListLabelingJobsOutput: r.Request.Data.(*ListLabelingJobsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListLabelingJobsRequestPaginator returns a paginator for ListLabelingJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListLabelingJobsRequest(input)
//   p := sagemaker.NewListLabelingJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListLabelingJobsPaginator(req ListLabelingJobsRequest) ListLabelingJobsPaginator {
	return ListLabelingJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListLabelingJobsInput
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

// ListLabelingJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListLabelingJobsPaginator struct {
	aws.Pager
}

func (p *ListLabelingJobsPaginator) CurrentPage() *ListLabelingJobsOutput {
	return p.Pager.CurrentPage().(*ListLabelingJobsOutput)
}

// ListLabelingJobsResponse is the response type for the
// ListLabelingJobs API operation.
type ListLabelingJobsResponse struct {
	*ListLabelingJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListLabelingJobs request.
func (r *ListLabelingJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
