// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type ListBuildsInput struct {
	_ struct{} `type:"structure"`

	// Maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int64 `min:"1" type:"integer"`

	// Token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value.
	NextToken *string `min:"1" type:"string"`

	// Build status to filter results by. To retrieve all builds, leave this parameter
	// empty.
	//
	// Possible build statuses include the following:
	//
	//    * INITIALIZED -- A new build has been defined, but no files have been
	//    uploaded. You cannot create fleets for builds that are in this status.
	//    When a build is successfully created, the build status is set to this
	//    value.
	//
	//    * READY -- The game build has been successfully uploaded. You can now
	//    create new fleets for this build.
	//
	//    * FAILED -- The game build upload failed. You cannot create new fleets
	//    for this build.
	Status BuildStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListBuildsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBuildsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBuildsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type ListBuildsOutput struct {
	_ struct{} `type:"structure"`

	// Collection of build records that match the request.
	Builds []Build `type:"list"`

	// Token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListBuildsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListBuilds = "ListBuilds"

// ListBuildsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves build records for all builds associated with the AWS account in
// use. You can limit results to builds that are in a specific status by using
// the Status parameter. Use the pagination parameters to retrieve results in
// a set of sequential pages.
//
// Build records are not listed in any particular order.
//
// Learn more
//
//  Working with Builds (https://docs.aws.amazon.com/gamelift/latest/developerguide/build-intro.html)
//
// Related operations
//
//    * CreateBuild
//
//    * ListBuilds
//
//    * DescribeBuild
//
//    * UpdateBuild
//
//    * DeleteBuild
//
//    // Example sending a request using ListBuildsRequest.
//    req := client.ListBuildsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/ListBuilds
func (c *Client) ListBuildsRequest(input *ListBuildsInput) ListBuildsRequest {
	op := &aws.Operation{
		Name:       opListBuilds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListBuildsInput{}
	}

	req := c.newRequest(op, input, &ListBuildsOutput{})
	return ListBuildsRequest{Request: req, Input: input, Copy: c.ListBuildsRequest}
}

// ListBuildsRequest is the request type for the
// ListBuilds API operation.
type ListBuildsRequest struct {
	*aws.Request
	Input *ListBuildsInput
	Copy  func(*ListBuildsInput) ListBuildsRequest
}

// Send marshals and sends the ListBuilds API request.
func (r ListBuildsRequest) Send(ctx context.Context) (*ListBuildsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBuildsResponse{
		ListBuildsOutput: r.Request.Data.(*ListBuildsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBuildsResponse is the response type for the
// ListBuilds API operation.
type ListBuildsResponse struct {
	*ListBuildsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBuilds request.
func (r *ListBuildsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
