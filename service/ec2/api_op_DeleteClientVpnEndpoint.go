// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteClientVpnEndpointInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN to be deleted.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s DeleteClientVpnEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteClientVpnEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteClientVpnEndpointInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteClientVpnEndpointOutput struct {
	_ struct{} `type:"structure"`

	// The current state of the Client VPN endpoint.
	Status *VpnEndpointStatus `locationName:"status" type:"structure"`
}

// String returns the string representation
func (s DeleteClientVpnEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteClientVpnEndpoint = "DeleteClientVpnEndpoint"

// DeleteClientVpnEndpointRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified Client VPN endpoint. You must disassociate all target
// networks before you can delete a Client VPN endpoint.
//
//    // Example sending a request using DeleteClientVpnEndpointRequest.
//    req := client.DeleteClientVpnEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteClientVpnEndpoint
func (c *Client) DeleteClientVpnEndpointRequest(input *DeleteClientVpnEndpointInput) DeleteClientVpnEndpointRequest {
	op := &aws.Operation{
		Name:       opDeleteClientVpnEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteClientVpnEndpointInput{}
	}

	req := c.newRequest(op, input, &DeleteClientVpnEndpointOutput{})
	return DeleteClientVpnEndpointRequest{Request: req, Input: input, Copy: c.DeleteClientVpnEndpointRequest}
}

// DeleteClientVpnEndpointRequest is the request type for the
// DeleteClientVpnEndpoint API operation.
type DeleteClientVpnEndpointRequest struct {
	*aws.Request
	Input *DeleteClientVpnEndpointInput
	Copy  func(*DeleteClientVpnEndpointInput) DeleteClientVpnEndpointRequest
}

// Send marshals and sends the DeleteClientVpnEndpoint API request.
func (r DeleteClientVpnEndpointRequest) Send(ctx context.Context) (*DeleteClientVpnEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteClientVpnEndpointResponse{
		DeleteClientVpnEndpointOutput: r.Request.Data.(*DeleteClientVpnEndpointOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteClientVpnEndpointResponse is the response type for the
// DeleteClientVpnEndpoint API operation.
type DeleteClientVpnEndpointResponse struct {
	*DeleteClientVpnEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteClientVpnEndpoint request.
func (r *DeleteClientVpnEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
