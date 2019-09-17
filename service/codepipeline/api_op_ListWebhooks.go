// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListWebhooksInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token that was returned from the previous ListWebhooks call, which can
	// be used to return the next set of webhooks in the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListWebhooksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListWebhooksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListWebhooksInput"}
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

type ListWebhooksOutput struct {
	_ struct{} `type:"structure"`

	// If the amount of returned information is significantly large, an identifier
	// is also returned and can be used in a subsequent ListWebhooks call to return
	// the next set of webhooks in the list.
	NextToken *string `min:"1" type:"string"`

	// The JSON detail returned for each webhook in the list output for the ListWebhooks
	// call.
	Webhooks []ListWebhookItem `locationName:"webhooks" type:"list"`
}

// String returns the string representation
func (s ListWebhooksOutput) String() string {
	return awsutil.Prettify(s)
}

const opListWebhooks = "ListWebhooks"

// ListWebhooksRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Gets a listing of all the webhooks in this region for this account. The output
// lists all webhooks and includes the webhook URL and ARN, as well the configuration
// for each webhook.
//
//    // Example sending a request using ListWebhooksRequest.
//    req := client.ListWebhooksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/ListWebhooks
func (c *Client) ListWebhooksRequest(input *ListWebhooksInput) ListWebhooksRequest {
	op := &aws.Operation{
		Name:       opListWebhooks,
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
		input = &ListWebhooksInput{}
	}

	req := c.newRequest(op, input, &ListWebhooksOutput{})
	return ListWebhooksRequest{Request: req, Input: input, Copy: c.ListWebhooksRequest}
}

// ListWebhooksRequest is the request type for the
// ListWebhooks API operation.
type ListWebhooksRequest struct {
	*aws.Request
	Input *ListWebhooksInput
	Copy  func(*ListWebhooksInput) ListWebhooksRequest
}

// Send marshals and sends the ListWebhooks API request.
func (r ListWebhooksRequest) Send(ctx context.Context) (*ListWebhooksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListWebhooksResponse{
		ListWebhooksOutput: r.Request.Data.(*ListWebhooksOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListWebhooksRequestPaginator returns a paginator for ListWebhooks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListWebhooksRequest(input)
//   p := codepipeline.NewListWebhooksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListWebhooksPaginator(req ListWebhooksRequest) ListWebhooksPaginator {
	return ListWebhooksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListWebhooksInput
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

// ListWebhooksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListWebhooksPaginator struct {
	aws.Pager
}

func (p *ListWebhooksPaginator) CurrentPage() *ListWebhooksOutput {
	return p.Pager.CurrentPage().(*ListWebhooksOutput)
}

// ListWebhooksResponse is the response type for the
// ListWebhooks API operation.
type ListWebhooksResponse struct {
	*ListWebhooksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListWebhooks request.
func (r *ListWebhooksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
