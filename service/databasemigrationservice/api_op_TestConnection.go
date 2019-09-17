// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TestConnectionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the replication instance.
	//
	// ReplicationInstanceArn is a required field
	ReplicationInstanceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TestConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestConnectionInput"}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if s.ReplicationInstanceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationInstanceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TestConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The connection tested.
	Connection *Connection `type:"structure"`
}

// String returns the string representation
func (s TestConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opTestConnection = "TestConnection"

// TestConnectionRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Tests the connection between the replication instance and the endpoint.
//
//    // Example sending a request using TestConnectionRequest.
//    req := client.TestConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/TestConnection
func (c *Client) TestConnectionRequest(input *TestConnectionInput) TestConnectionRequest {
	op := &aws.Operation{
		Name:       opTestConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestConnectionInput{}
	}

	req := c.newRequest(op, input, &TestConnectionOutput{})
	return TestConnectionRequest{Request: req, Input: input, Copy: c.TestConnectionRequest}
}

// TestConnectionRequest is the request type for the
// TestConnection API operation.
type TestConnectionRequest struct {
	*aws.Request
	Input *TestConnectionInput
	Copy  func(*TestConnectionInput) TestConnectionRequest
}

// Send marshals and sends the TestConnection API request.
func (r TestConnectionRequest) Send(ctx context.Context) (*TestConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestConnectionResponse{
		TestConnectionOutput: r.Request.Data.(*TestConnectionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestConnectionResponse is the response type for the
// TestConnection API operation.
type TestConnectionResponse struct {
	*TestConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestConnection request.
func (r *TestConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
