// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetQueryResultsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results (rows) to return in this request.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token that specifies where to start pagination if a previous request
	// was truncated.
	NextToken *string `min:"1" type:"string"`

	// The unique ID of the query execution.
	//
	// QueryExecutionId is a required field
	QueryExecutionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetQueryResultsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueryResultsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetQueryResultsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.QueryExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryExecutionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetQueryResultsOutput struct {
	_ struct{} `type:"structure"`

	// A token to be used by the next request if this request is truncated.
	NextToken *string `min:"1" type:"string"`

	// The results of the query execution.
	ResultSet *ResultSet `type:"structure"`

	// The number of rows inserted with a CREATE TABLE AS SELECT statement.
	UpdateCount *int64 `type:"long"`
}

// String returns the string representation
func (s GetQueryResultsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetQueryResults = "GetQueryResults"

// GetQueryResultsRequest returns a request value for making API operation for
// Amazon Athena.
//
// Returns the results of a single query execution specified by QueryExecutionId
// if you have access to the workgroup in which the query ran. This request
// does not execute the query but returns results. Use StartQueryExecution to
// run a query.
//
//    // Example sending a request using GetQueryResultsRequest.
//    req := client.GetQueryResultsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/GetQueryResults
func (c *Client) GetQueryResultsRequest(input *GetQueryResultsInput) GetQueryResultsRequest {
	op := &aws.Operation{
		Name:       opGetQueryResults,
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
		input = &GetQueryResultsInput{}
	}

	req := c.newRequest(op, input, &GetQueryResultsOutput{})
	return GetQueryResultsRequest{Request: req, Input: input, Copy: c.GetQueryResultsRequest}
}

// GetQueryResultsRequest is the request type for the
// GetQueryResults API operation.
type GetQueryResultsRequest struct {
	*aws.Request
	Input *GetQueryResultsInput
	Copy  func(*GetQueryResultsInput) GetQueryResultsRequest
}

// Send marshals and sends the GetQueryResults API request.
func (r GetQueryResultsRequest) Send(ctx context.Context) (*GetQueryResultsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetQueryResultsResponse{
		GetQueryResultsOutput: r.Request.Data.(*GetQueryResultsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetQueryResultsRequestPaginator returns a paginator for GetQueryResults.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetQueryResultsRequest(input)
//   p := athena.NewGetQueryResultsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetQueryResultsPaginator(req GetQueryResultsRequest) GetQueryResultsPaginator {
	return GetQueryResultsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetQueryResultsInput
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

// GetQueryResultsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetQueryResultsPaginator struct {
	aws.Pager
}

func (p *GetQueryResultsPaginator) CurrentPage() *GetQueryResultsOutput {
	return p.Pager.CurrentPage().(*GetQueryResultsOutput)
}

// GetQueryResultsResponse is the response type for the
// GetQueryResults API operation.
type GetQueryResultsResponse struct {
	*GetQueryResultsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetQueryResults request.
func (r *GetQueryResultsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
