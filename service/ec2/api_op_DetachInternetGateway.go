// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type DetachInternetGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the internet gateway.
	//
	// InternetGatewayId is a required field
	InternetGatewayId *string `locationName:"internetGatewayId" type:"string" required:"true"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `locationName:"vpcId" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachInternetGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachInternetGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachInternetGatewayInput"}

	if s.InternetGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InternetGatewayId"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachInternetGatewayOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DetachInternetGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachInternetGateway = "DetachInternetGateway"

// DetachInternetGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Detaches an internet gateway from a VPC, disabling connectivity between the
// internet and the VPC. The VPC must not contain any running instances with
// Elastic IP addresses or public IPv4 addresses.
//
//    // Example sending a request using DetachInternetGatewayRequest.
//    req := client.DetachInternetGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DetachInternetGateway
func (c *Client) DetachInternetGatewayRequest(input *DetachInternetGatewayInput) DetachInternetGatewayRequest {
	op := &aws.Operation{
		Name:       opDetachInternetGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachInternetGatewayInput{}
	}

	req := c.newRequest(op, input, &DetachInternetGatewayOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DetachInternetGatewayRequest{Request: req, Input: input, Copy: c.DetachInternetGatewayRequest}
}

// DetachInternetGatewayRequest is the request type for the
// DetachInternetGateway API operation.
type DetachInternetGatewayRequest struct {
	*aws.Request
	Input *DetachInternetGatewayInput
	Copy  func(*DetachInternetGatewayInput) DetachInternetGatewayRequest
}

// Send marshals and sends the DetachInternetGateway API request.
func (r DetachInternetGatewayRequest) Send(ctx context.Context) (*DetachInternetGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachInternetGatewayResponse{
		DetachInternetGatewayOutput: r.Request.Data.(*DetachInternetGatewayOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachInternetGatewayResponse is the response type for the
// DetachInternetGateway API operation.
type DetachInternetGatewayResponse struct {
	*DetachInternetGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachInternetGateway request.
func (r *DetachInternetGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
