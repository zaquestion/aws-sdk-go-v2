// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDatasetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The number of items to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the result of the previous request was truncated, the response includes
	// a NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDatasetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatasetGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatasetGroupsInput"}
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

type ListDatasetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that summarize each dataset group's properties.
	DatasetGroups []DatasetGroupSummary `type:"list"`

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDatasetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDatasetGroups = "ListDatasetGroups"

// ListDatasetGroupsRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Returns a list of dataset groups created using the CreateDatasetGroup operation.
// For each dataset group, a summary of its properties, including its Amazon
// Resource Name (ARN), is returned. You can retrieve the complete set of properties
// by using the ARN with the DescribeDatasetGroup operation.
//
//    // Example sending a request using ListDatasetGroupsRequest.
//    req := client.ListDatasetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/ListDatasetGroups
func (c *Client) ListDatasetGroupsRequest(input *ListDatasetGroupsInput) ListDatasetGroupsRequest {
	op := &aws.Operation{
		Name:       opListDatasetGroups,
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
		input = &ListDatasetGroupsInput{}
	}

	req := c.newRequest(op, input, &ListDatasetGroupsOutput{})
	return ListDatasetGroupsRequest{Request: req, Input: input, Copy: c.ListDatasetGroupsRequest}
}

// ListDatasetGroupsRequest is the request type for the
// ListDatasetGroups API operation.
type ListDatasetGroupsRequest struct {
	*aws.Request
	Input *ListDatasetGroupsInput
	Copy  func(*ListDatasetGroupsInput) ListDatasetGroupsRequest
}

// Send marshals and sends the ListDatasetGroups API request.
func (r ListDatasetGroupsRequest) Send(ctx context.Context) (*ListDatasetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDatasetGroupsResponse{
		ListDatasetGroupsOutput: r.Request.Data.(*ListDatasetGroupsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDatasetGroupsRequestPaginator returns a paginator for ListDatasetGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDatasetGroupsRequest(input)
//   p := forecast.NewListDatasetGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDatasetGroupsPaginator(req ListDatasetGroupsRequest) ListDatasetGroupsPaginator {
	return ListDatasetGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDatasetGroupsInput
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

// ListDatasetGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDatasetGroupsPaginator struct {
	aws.Pager
}

func (p *ListDatasetGroupsPaginator) CurrentPage() *ListDatasetGroupsOutput {
	return p.Pager.CurrentPage().(*ListDatasetGroupsOutput)
}

// ListDatasetGroupsResponse is the response type for the
// ListDatasetGroups API operation.
type ListDatasetGroupsResponse struct {
	*ListDatasetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDatasetGroups request.
func (r *ListDatasetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
