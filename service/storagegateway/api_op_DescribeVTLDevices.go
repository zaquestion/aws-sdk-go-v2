// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DescribeVTLDevicesInput
type DescribeVTLDevicesInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// Specifies that the number of VTL devices described be limited to the specified
	// number.
	Limit *int64 `min:"1" type:"integer"`

	// An opaque string that indicates the position at which to begin describing
	// the VTL devices.
	Marker *string `min:"1" type:"string"`

	// An array of strings, where each string represents the Amazon Resource Name
	// (ARN) of a VTL device.
	//
	// All of the specified VTL devices must be from the same gateway. If no VTL
	// devices are specified, the result will contain all devices on the specified
	// gateway.
	VTLDeviceARNs []string `type:"list"`
}

// String returns the string representation
func (s DescribeVTLDevicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVTLDevicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVTLDevicesInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeVTLDevicesOutput
type DescribeVTLDevicesOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	GatewayARN *string `min:"50" type:"string"`

	// An opaque string that indicates the position at which the VTL devices that
	// were fetched for description ended. Use the marker in your next request to
	// fetch the next set of VTL devices in the list. If there are no more VTL devices
	// to describe, this field does not appear in the response.
	Marker *string `min:"1" type:"string"`

	// An array of VTL device objects composed of the Amazon Resource Name(ARN)
	// of the VTL devices.
	VTLDevices []VTLDevice `type:"list"`
}

// String returns the string representation
func (s DescribeVTLDevicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVTLDevices = "DescribeVTLDevices"

// DescribeVTLDevicesRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns a description of virtual tape library (VTL) devices for the specified
// tape gateway. In the response, AWS Storage Gateway returns VTL device information.
//
// This operation is only supported in the tape gateway type.
//
//    // Example sending a request using DescribeVTLDevicesRequest.
//    req := client.DescribeVTLDevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeVTLDevices
func (c *Client) DescribeVTLDevicesRequest(input *DescribeVTLDevicesInput) DescribeVTLDevicesRequest {
	op := &aws.Operation{
		Name:       opDescribeVTLDevices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeVTLDevicesInput{}
	}

	req := c.newRequest(op, input, &DescribeVTLDevicesOutput{})
	return DescribeVTLDevicesRequest{Request: req, Input: input, Copy: c.DescribeVTLDevicesRequest}
}

// DescribeVTLDevicesRequest is the request type for the
// DescribeVTLDevices API operation.
type DescribeVTLDevicesRequest struct {
	*aws.Request
	Input *DescribeVTLDevicesInput
	Copy  func(*DescribeVTLDevicesInput) DescribeVTLDevicesRequest
}

// Send marshals and sends the DescribeVTLDevices API request.
func (r DescribeVTLDevicesRequest) Send(ctx context.Context) (*DescribeVTLDevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVTLDevicesResponse{
		DescribeVTLDevicesOutput: r.Request.Data.(*DescribeVTLDevicesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVTLDevicesRequestPaginator returns a paginator for DescribeVTLDevices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVTLDevicesRequest(input)
//   p := storagegateway.NewDescribeVTLDevicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVTLDevicesPaginator(req DescribeVTLDevicesRequest) DescribeVTLDevicesPaginator {
	return DescribeVTLDevicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeVTLDevicesInput
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

// DescribeVTLDevicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVTLDevicesPaginator struct {
	aws.Pager
}

func (p *DescribeVTLDevicesPaginator) CurrentPage() *DescribeVTLDevicesOutput {
	return p.Pager.CurrentPage().(*DescribeVTLDevicesOutput)
}

// DescribeVTLDevicesResponse is the response type for the
// DescribeVTLDevices API operation.
type DescribeVTLDevicesResponse struct {
	*DescribeVTLDevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVTLDevices request.
func (r *DescribeVTLDevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
