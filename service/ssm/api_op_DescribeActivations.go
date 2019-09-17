// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeActivationsInput struct {
	_ struct{} `type:"structure"`

	// A filter to view information about your activations.
	Filters []DescribeActivationsFilter `type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeActivationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeActivationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeActivationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeActivationsOutput struct {
	_ struct{} `type:"structure"`

	// A list of activations for your AWS account.
	ActivationList []Activation `type:"list"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeActivationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeActivations = "DescribeActivations"

// DescribeActivationsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Describes details about the activation, such as the date and time the activation
// was created, its expiration date, the IAM role assigned to the instances
// in the activation, and the number of instances registered by using this activation.
//
//    // Example sending a request using DescribeActivationsRequest.
//    req := client.DescribeActivationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeActivations
func (c *Client) DescribeActivationsRequest(input *DescribeActivationsInput) DescribeActivationsRequest {
	op := &aws.Operation{
		Name:       opDescribeActivations,
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
		input = &DescribeActivationsInput{}
	}

	req := c.newRequest(op, input, &DescribeActivationsOutput{})
	return DescribeActivationsRequest{Request: req, Input: input, Copy: c.DescribeActivationsRequest}
}

// DescribeActivationsRequest is the request type for the
// DescribeActivations API operation.
type DescribeActivationsRequest struct {
	*aws.Request
	Input *DescribeActivationsInput
	Copy  func(*DescribeActivationsInput) DescribeActivationsRequest
}

// Send marshals and sends the DescribeActivations API request.
func (r DescribeActivationsRequest) Send(ctx context.Context) (*DescribeActivationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeActivationsResponse{
		DescribeActivationsOutput: r.Request.Data.(*DescribeActivationsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeActivationsRequestPaginator returns a paginator for DescribeActivations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeActivationsRequest(input)
//   p := ssm.NewDescribeActivationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeActivationsPaginator(req DescribeActivationsRequest) DescribeActivationsPaginator {
	return DescribeActivationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeActivationsInput
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

// DescribeActivationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeActivationsPaginator struct {
	aws.Pager
}

func (p *DescribeActivationsPaginator) CurrentPage() *DescribeActivationsOutput {
	return p.Pager.CurrentPage().(*DescribeActivationsOutput)
}

// DescribeActivationsResponse is the response type for the
// DescribeActivations API operation.
type DescribeActivationsResponse struct {
	*DescribeActivationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeActivations request.
func (r *DescribeActivationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
