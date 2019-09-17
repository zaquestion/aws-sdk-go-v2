// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInventoryDeletionsInput struct {
	_ struct{} `type:"structure"`

	// Specify the delete inventory ID for which you want information. This ID was
	// returned by the DeleteInventory action.
	DeletionId *string `type:"string"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInventoryDeletionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInventoryDeletionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInventoryDeletionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInventoryDeletionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of status items for deleted inventory.
	InventoryDeletions []InventoryDeletionStatusItem `type:"list"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInventoryDeletionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInventoryDeletions = "DescribeInventoryDeletions"

// DescribeInventoryDeletionsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Describes a specific delete inventory operation.
//
//    // Example sending a request using DescribeInventoryDeletionsRequest.
//    req := client.DescribeInventoryDeletionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInventoryDeletions
func (c *Client) DescribeInventoryDeletionsRequest(input *DescribeInventoryDeletionsInput) DescribeInventoryDeletionsRequest {
	op := &aws.Operation{
		Name:       opDescribeInventoryDeletions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInventoryDeletionsInput{}
	}

	req := c.newRequest(op, input, &DescribeInventoryDeletionsOutput{})
	return DescribeInventoryDeletionsRequest{Request: req, Input: input, Copy: c.DescribeInventoryDeletionsRequest}
}

// DescribeInventoryDeletionsRequest is the request type for the
// DescribeInventoryDeletions API operation.
type DescribeInventoryDeletionsRequest struct {
	*aws.Request
	Input *DescribeInventoryDeletionsInput
	Copy  func(*DescribeInventoryDeletionsInput) DescribeInventoryDeletionsRequest
}

// Send marshals and sends the DescribeInventoryDeletions API request.
func (r DescribeInventoryDeletionsRequest) Send(ctx context.Context) (*DescribeInventoryDeletionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInventoryDeletionsResponse{
		DescribeInventoryDeletionsOutput: r.Request.Data.(*DescribeInventoryDeletionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInventoryDeletionsResponse is the response type for the
// DescribeInventoryDeletions API operation.
type DescribeInventoryDeletionsResponse struct {
	*DescribeInventoryDeletionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInventoryDeletions request.
func (r *DescribeInventoryDeletionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
