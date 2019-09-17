// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsForProjectInput struct {
	_ struct{} `type:"structure"`

	// The ID of the project to get tags for.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"2" type:"string" required:"true"`

	// Reserved for future use.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// Reserved for future use.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListTagsForProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForProjectInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 2))
	}
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

type ListTagsForProjectOutput struct {
	_ struct{} `type:"structure"`

	// Reserved for future use.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The tags for the project.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s ListTagsForProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsForProject = "ListTagsForProject"

// ListTagsForProjectRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Gets the tags for a project.
//
//    // Example sending a request using ListTagsForProjectRequest.
//    req := client.ListTagsForProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListTagsForProject
func (c *Client) ListTagsForProjectRequest(input *ListTagsForProjectInput) ListTagsForProjectRequest {
	op := &aws.Operation{
		Name:       opListTagsForProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForProjectInput{}
	}

	req := c.newRequest(op, input, &ListTagsForProjectOutput{})
	return ListTagsForProjectRequest{Request: req, Input: input, Copy: c.ListTagsForProjectRequest}
}

// ListTagsForProjectRequest is the request type for the
// ListTagsForProject API operation.
type ListTagsForProjectRequest struct {
	*aws.Request
	Input *ListTagsForProjectInput
	Copy  func(*ListTagsForProjectInput) ListTagsForProjectRequest
}

// Send marshals and sends the ListTagsForProject API request.
func (r ListTagsForProjectRequest) Send(ctx context.Context) (*ListTagsForProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForProjectResponse{
		ListTagsForProjectOutput: r.Request.Data.(*ListTagsForProjectOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForProjectResponse is the response type for the
// ListTagsForProject API operation.
type ListTagsForProjectResponse struct {
	*ListTagsForProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForProject request.
func (r *ListTagsForProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
