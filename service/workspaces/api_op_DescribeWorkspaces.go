// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkspacesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the bundle. All WorkSpaces that are created from this bundle
	// are retrieved. You cannot combine this parameter with any other filter.
	BundleId *string `type:"string"`

	// The identifier of the directory. In addition, you can optionally specify
	// a specific directory user (see UserName). You cannot combine this parameter
	// with any other filter.
	DirectoryId *string `type:"string"`

	// The maximum number of items to return.
	Limit *int64 `min:"1" type:"integer"`

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The name of the directory user. You must specify this parameter with DirectoryId.
	UserName *string `min:"1" type:"string"`

	// The identifiers of the WorkSpaces. You cannot combine this parameter with
	// any other filter.
	//
	// Because the CreateWorkspaces operation is asynchronous, the identifier it
	// returns is not immediately available. If you immediately call DescribeWorkspaces
	// with this identifier, no information is returned.
	WorkspaceIds []string `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeWorkspacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkspacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkspacesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}
	if s.WorkspaceIds != nil && len(s.WorkspaceIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkspaceIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkspacesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next set of results, or null if no more
	// results are available.
	NextToken *string `min:"1" type:"string"`

	// Information about the WorkSpaces.
	//
	// Because CreateWorkspaces is an asynchronous operation, some of the returned
	// information could be incomplete.
	Workspaces []Workspace `type:"list"`
}

// String returns the string representation
func (s DescribeWorkspacesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeWorkspaces = "DescribeWorkspaces"

// DescribeWorkspacesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Describes the specified WorkSpaces.
//
// You can filter the results by using the bundle identifier, directory identifier,
// or owner, but you can specify only one filter at a time.
//
//    // Example sending a request using DescribeWorkspacesRequest.
//    req := client.DescribeWorkspacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeWorkspaces
func (c *Client) DescribeWorkspacesRequest(input *DescribeWorkspacesInput) DescribeWorkspacesRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeWorkspacesInput{}
	}

	req := c.newRequest(op, input, &DescribeWorkspacesOutput{})
	return DescribeWorkspacesRequest{Request: req, Input: input, Copy: c.DescribeWorkspacesRequest}
}

// DescribeWorkspacesRequest is the request type for the
// DescribeWorkspaces API operation.
type DescribeWorkspacesRequest struct {
	*aws.Request
	Input *DescribeWorkspacesInput
	Copy  func(*DescribeWorkspacesInput) DescribeWorkspacesRequest
}

// Send marshals and sends the DescribeWorkspaces API request.
func (r DescribeWorkspacesRequest) Send(ctx context.Context) (*DescribeWorkspacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkspacesResponse{
		DescribeWorkspacesOutput: r.Request.Data.(*DescribeWorkspacesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeWorkspacesRequestPaginator returns a paginator for DescribeWorkspaces.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeWorkspacesRequest(input)
//   p := workspaces.NewDescribeWorkspacesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeWorkspacesPaginator(req DescribeWorkspacesRequest) DescribeWorkspacesPaginator {
	return DescribeWorkspacesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeWorkspacesInput
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

// DescribeWorkspacesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeWorkspacesPaginator struct {
	aws.Pager
}

func (p *DescribeWorkspacesPaginator) CurrentPage() *DescribeWorkspacesOutput {
	return p.Pager.CurrentPage().(*DescribeWorkspacesOutput)
}

// DescribeWorkspacesResponse is the response type for the
// DescribeWorkspaces API operation.
type DescribeWorkspacesResponse struct {
	*DescribeWorkspacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkspaces request.
func (r *DescribeWorkspacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
