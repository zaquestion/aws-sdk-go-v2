// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSharedDirectoriesInput struct {
	_ struct{} `type:"structure"`

	// The number of shared directories to return in the response object.
	Limit *int64 `type:"integer"`

	// The DescribeSharedDirectoriesResult.NextToken value from a previous call
	// to DescribeSharedDirectories. Pass null if this is the first call.
	NextToken *string `type:"string"`

	// Returns the identifier of the directory in the directory owner account.
	//
	// OwnerDirectoryId is a required field
	OwnerDirectoryId *string `type:"string" required:"true"`

	// A list of identifiers of all shared directories in your account.
	SharedDirectoryIds []string `type:"list"`
}

// String returns the string representation
func (s DescribeSharedDirectoriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSharedDirectoriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSharedDirectoriesInput"}

	if s.OwnerDirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OwnerDirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSharedDirectoriesOutput struct {
	_ struct{} `type:"structure"`

	// If not null, token that indicates that more results are available. Pass this
	// value for the NextToken parameter in a subsequent call to DescribeSharedDirectories
	// to retrieve the next set of items.
	NextToken *string `type:"string"`

	// A list of all shared directories in your account.
	SharedDirectories []SharedDirectory `type:"list"`
}

// String returns the string representation
func (s DescribeSharedDirectoriesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSharedDirectories = "DescribeSharedDirectories"

// DescribeSharedDirectoriesRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Returns the shared directories in your account.
//
//    // Example sending a request using DescribeSharedDirectoriesRequest.
//    req := client.DescribeSharedDirectoriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DescribeSharedDirectories
func (c *Client) DescribeSharedDirectoriesRequest(input *DescribeSharedDirectoriesInput) DescribeSharedDirectoriesRequest {
	op := &aws.Operation{
		Name:       opDescribeSharedDirectories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSharedDirectoriesInput{}
	}

	req := c.newRequest(op, input, &DescribeSharedDirectoriesOutput{})
	return DescribeSharedDirectoriesRequest{Request: req, Input: input, Copy: c.DescribeSharedDirectoriesRequest}
}

// DescribeSharedDirectoriesRequest is the request type for the
// DescribeSharedDirectories API operation.
type DescribeSharedDirectoriesRequest struct {
	*aws.Request
	Input *DescribeSharedDirectoriesInput
	Copy  func(*DescribeSharedDirectoriesInput) DescribeSharedDirectoriesRequest
}

// Send marshals and sends the DescribeSharedDirectories API request.
func (r DescribeSharedDirectoriesRequest) Send(ctx context.Context) (*DescribeSharedDirectoriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSharedDirectoriesResponse{
		DescribeSharedDirectoriesOutput: r.Request.Data.(*DescribeSharedDirectoriesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSharedDirectoriesResponse is the response type for the
// DescribeSharedDirectories API operation.
type DescribeSharedDirectoriesResponse struct {
	*DescribeSharedDirectoriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSharedDirectories request.
func (r *DescribeSharedDirectoriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
