// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeUploadBufferInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUploadBufferInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUploadBufferInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUploadBufferInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeUploadBufferOutput struct {
	_ struct{} `type:"structure"`

	// An array of the gateway's local disk IDs that are configured as working storage.
	// Each local disk ID is specified as a string (minimum length of 1 and maximum
	// length of 300). If no local disks are configured as working storage, then
	// the DiskIds array is empty.
	DiskIds []string `type:"list"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	GatewayARN *string `min:"50" type:"string"`

	// The total number of bytes allocated in the gateway's as upload buffer.
	UploadBufferAllocatedInBytes *int64 `type:"long"`

	// The total number of bytes being used in the gateway's upload buffer.
	UploadBufferUsedInBytes *int64 `type:"long"`
}

// String returns the string representation
func (s DescribeUploadBufferOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUploadBuffer = "DescribeUploadBuffer"

// DescribeUploadBufferRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns information about the upload buffer of a gateway. This operation
// is supported for the stored volume, cached volume and tape gateway types.
//
// The response includes disk IDs that are configured as upload buffer space,
// and it includes the amount of upload buffer space allocated and used.
//
//    // Example sending a request using DescribeUploadBufferRequest.
//    req := client.DescribeUploadBufferRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeUploadBuffer
func (c *Client) DescribeUploadBufferRequest(input *DescribeUploadBufferInput) DescribeUploadBufferRequest {
	op := &aws.Operation{
		Name:       opDescribeUploadBuffer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUploadBufferInput{}
	}

	req := c.newRequest(op, input, &DescribeUploadBufferOutput{})
	return DescribeUploadBufferRequest{Request: req, Input: input, Copy: c.DescribeUploadBufferRequest}
}

// DescribeUploadBufferRequest is the request type for the
// DescribeUploadBuffer API operation.
type DescribeUploadBufferRequest struct {
	*aws.Request
	Input *DescribeUploadBufferInput
	Copy  func(*DescribeUploadBufferInput) DescribeUploadBufferRequest
}

// Send marshals and sends the DescribeUploadBuffer API request.
func (r DescribeUploadBufferRequest) Send(ctx context.Context) (*DescribeUploadBufferResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUploadBufferResponse{
		DescribeUploadBufferOutput: r.Request.Data.(*DescribeUploadBufferOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUploadBufferResponse is the response type for the
// DescribeUploadBuffer API operation.
type DescribeUploadBufferResponse struct {
	*DescribeUploadBufferOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUploadBuffer request.
func (r *DescribeUploadBufferResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
