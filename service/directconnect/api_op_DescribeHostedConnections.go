// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeHostedConnectionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the interconnect or LAG.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeHostedConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeHostedConnectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeHostedConnectionsInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeHostedConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// The connections.
	Connections []Connection `locationName:"connections" type:"list"`
}

// String returns the string representation
func (s DescribeHostedConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeHostedConnections = "DescribeHostedConnections"

// DescribeHostedConnectionsRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Lists the hosted connections that have been provisioned on the specified
// interconnect or link aggregation group (LAG).
//
// Intended for use by AWS Direct Connect Partners only.
//
//    // Example sending a request using DescribeHostedConnectionsRequest.
//    req := client.DescribeHostedConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeHostedConnections
func (c *Client) DescribeHostedConnectionsRequest(input *DescribeHostedConnectionsInput) DescribeHostedConnectionsRequest {
	op := &aws.Operation{
		Name:       opDescribeHostedConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeHostedConnectionsInput{}
	}

	req := c.newRequest(op, input, &DescribeHostedConnectionsOutput{})
	return DescribeHostedConnectionsRequest{Request: req, Input: input, Copy: c.DescribeHostedConnectionsRequest}
}

// DescribeHostedConnectionsRequest is the request type for the
// DescribeHostedConnections API operation.
type DescribeHostedConnectionsRequest struct {
	*aws.Request
	Input *DescribeHostedConnectionsInput
	Copy  func(*DescribeHostedConnectionsInput) DescribeHostedConnectionsRequest
}

// Send marshals and sends the DescribeHostedConnections API request.
func (r DescribeHostedConnectionsRequest) Send(ctx context.Context) (*DescribeHostedConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHostedConnectionsResponse{
		DescribeHostedConnectionsOutput: r.Request.Data.(*DescribeHostedConnectionsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeHostedConnectionsResponse is the response type for the
// DescribeHostedConnections API operation.
type DescribeHostedConnectionsResponse struct {
	*DescribeHostedConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHostedConnections request.
func (r *DescribeHostedConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
