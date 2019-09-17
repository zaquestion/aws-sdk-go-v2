// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the of the gateway.
type DescribeWorkingStorageInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeWorkingStorageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkingStorageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkingStorageInput"}

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

// A JSON object containing the following fields:
type DescribeWorkingStorageOutput struct {
	_ struct{} `type:"structure"`

	// An array of the gateway's local disk IDs that are configured as working storage.
	// Each local disk ID is specified as a string (minimum length of 1 and maximum
	// length of 300). If no local disks are configured as working storage, then
	// the DiskIds array is empty.
	DiskIds []string `type:"list"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	GatewayARN *string `min:"50" type:"string"`

	// The total working storage in bytes allocated for the gateway. If no working
	// storage is configured for the gateway, this field returns 0.
	WorkingStorageAllocatedInBytes *int64 `type:"long"`

	// The total working storage in bytes in use by the gateway. If no working storage
	// is configured for the gateway, this field returns 0.
	WorkingStorageUsedInBytes *int64 `type:"long"`
}

// String returns the string representation
func (s DescribeWorkingStorageOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeWorkingStorage = "DescribeWorkingStorage"

// DescribeWorkingStorageRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns information about the working storage of a gateway. This operation
// is only supported in the stored volumes gateway type. This operation is deprecated
// in cached volumes API version (20120630). Use DescribeUploadBuffer instead.
//
// Working storage is also referred to as upload buffer. You can also use the
// DescribeUploadBuffer operation to add upload buffer to a stored volume gateway.
//
// The response includes disk IDs that are configured as working storage, and
// it includes the amount of working storage allocated and used.
//
//    // Example sending a request using DescribeWorkingStorageRequest.
//    req := client.DescribeWorkingStorageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeWorkingStorage
func (c *Client) DescribeWorkingStorageRequest(input *DescribeWorkingStorageInput) DescribeWorkingStorageRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkingStorage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeWorkingStorageInput{}
	}

	req := c.newRequest(op, input, &DescribeWorkingStorageOutput{})
	return DescribeWorkingStorageRequest{Request: req, Input: input, Copy: c.DescribeWorkingStorageRequest}
}

// DescribeWorkingStorageRequest is the request type for the
// DescribeWorkingStorage API operation.
type DescribeWorkingStorageRequest struct {
	*aws.Request
	Input *DescribeWorkingStorageInput
	Copy  func(*DescribeWorkingStorageInput) DescribeWorkingStorageRequest
}

// Send marshals and sends the DescribeWorkingStorage API request.
func (r DescribeWorkingStorageRequest) Send(ctx context.Context) (*DescribeWorkingStorageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkingStorageResponse{
		DescribeWorkingStorageOutput: r.Request.Data.(*DescribeWorkingStorageOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeWorkingStorageResponse is the response type for the
// DescribeWorkingStorage API operation.
type DescribeWorkingStorageResponse struct {
	*DescribeWorkingStorageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkingStorage request.
func (r *DescribeWorkingStorageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
