// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRulesPackagesInput struct {
	_ struct{} `type:"structure"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 10. The maximum value is 500.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListRulesPackages action. Subsequent
	// calls to the action fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListRulesPackagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRulesPackagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRulesPackagesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRulesPackagesOutput struct {
	_ struct{} `type:"structure"`

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to
	// be listed, this parameter is set to null.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The list of ARNs that specifies the rules packages returned by the action.
	//
	// RulesPackageArns is a required field
	RulesPackageArns []string `locationName:"rulesPackageArns" type:"list" required:"true"`
}

// String returns the string representation
func (s ListRulesPackagesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRulesPackages = "ListRulesPackages"

// ListRulesPackagesRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Lists all available Amazon Inspector rules packages.
//
//    // Example sending a request using ListRulesPackagesRequest.
//    req := client.ListRulesPackagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/ListRulesPackages
func (c *Client) ListRulesPackagesRequest(input *ListRulesPackagesInput) ListRulesPackagesRequest {
	op := &aws.Operation{
		Name:       opListRulesPackages,
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
		input = &ListRulesPackagesInput{}
	}

	req := c.newRequest(op, input, &ListRulesPackagesOutput{})
	return ListRulesPackagesRequest{Request: req, Input: input, Copy: c.ListRulesPackagesRequest}
}

// ListRulesPackagesRequest is the request type for the
// ListRulesPackages API operation.
type ListRulesPackagesRequest struct {
	*aws.Request
	Input *ListRulesPackagesInput
	Copy  func(*ListRulesPackagesInput) ListRulesPackagesRequest
}

// Send marshals and sends the ListRulesPackages API request.
func (r ListRulesPackagesRequest) Send(ctx context.Context) (*ListRulesPackagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRulesPackagesResponse{
		ListRulesPackagesOutput: r.Request.Data.(*ListRulesPackagesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRulesPackagesRequestPaginator returns a paginator for ListRulesPackages.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRulesPackagesRequest(input)
//   p := inspector.NewListRulesPackagesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRulesPackagesPaginator(req ListRulesPackagesRequest) ListRulesPackagesPaginator {
	return ListRulesPackagesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRulesPackagesInput
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

// ListRulesPackagesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRulesPackagesPaginator struct {
	aws.Pager
}

func (p *ListRulesPackagesPaginator) CurrentPage() *ListRulesPackagesOutput {
	return p.Pager.CurrentPage().(*ListRulesPackagesOutput)
}

// ListRulesPackagesResponse is the response type for the
// ListRulesPackages API operation.
type ListRulesPackagesResponse struct {
	*ListRulesPackagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRulesPackages request.
func (r *ListRulesPackagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
