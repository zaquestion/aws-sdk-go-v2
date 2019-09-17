// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePatchGroupsInput struct {
	_ struct{} `type:"structure"`

	// One or more filters. Use a filter to return a more specific list of results.
	Filters []PatchOrchestratorFilter `type:"list"`

	// The maximum number of patch groups to return (per page).
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribePatchGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePatchGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePatchGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePatchGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Each entry in the array contains:
	//
	// PatchGroup: string (between 1 and 256 characters, Regex: ^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$)
	//
	// PatchBaselineIdentity: A PatchBaselineIdentity element.
	Mappings []PatchGroupPatchBaselineMapping `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribePatchGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePatchGroups = "DescribePatchGroups"

// DescribePatchGroupsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Lists all patch groups that have been registered with patch baselines.
//
//    // Example sending a request using DescribePatchGroupsRequest.
//    req := client.DescribePatchGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribePatchGroups
func (c *Client) DescribePatchGroupsRequest(input *DescribePatchGroupsInput) DescribePatchGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribePatchGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePatchGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribePatchGroupsOutput{})
	return DescribePatchGroupsRequest{Request: req, Input: input, Copy: c.DescribePatchGroupsRequest}
}

// DescribePatchGroupsRequest is the request type for the
// DescribePatchGroups API operation.
type DescribePatchGroupsRequest struct {
	*aws.Request
	Input *DescribePatchGroupsInput
	Copy  func(*DescribePatchGroupsInput) DescribePatchGroupsRequest
}

// Send marshals and sends the DescribePatchGroups API request.
func (r DescribePatchGroupsRequest) Send(ctx context.Context) (*DescribePatchGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePatchGroupsResponse{
		DescribePatchGroupsOutput: r.Request.Data.(*DescribePatchGroupsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePatchGroupsResponse is the response type for the
// DescribePatchGroups API operation.
type DescribePatchGroupsResponse struct {
	*DescribePatchGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePatchGroups request.
func (r *DescribePatchGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
