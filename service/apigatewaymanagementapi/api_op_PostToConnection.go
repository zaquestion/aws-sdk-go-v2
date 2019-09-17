// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewaymanagementapi

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type PostToConnectionInput struct {
	_ struct{} `type:"structure" payload:"Data"`

	// ConnectionId is a required field
	ConnectionId *string `location:"uri" locationName:"connectionId" type:"string" required:"true"`

	// The data to be sent to the client specified by its connection id.
	//
	// Data is a required field
	Data []byte `type:"blob" required:"true"`
}

// String returns the string representation
func (s PostToConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostToConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostToConnectionInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.Data == nil {
		invalidParams.Add(aws.NewErrParamRequired("Data"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PostToConnectionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ConnectionId != nil {
		v := *s.ConnectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "connectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Data != nil {
		v := s.Data

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Data", protocol.BytesStream(v), metadata)
	}
	return nil
}

type PostToConnectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PostToConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PostToConnectionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPostToConnection = "PostToConnection"

// PostToConnectionRequest returns a request value for making API operation for
// AmazonApiGatewayManagementApi.
//
// Sends the provided data to the specified connection.
//
//    // Example sending a request using PostToConnectionRequest.
//    req := client.PostToConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewaymanagementapi-2018-11-29/PostToConnection
func (c *Client) PostToConnectionRequest(input *PostToConnectionInput) PostToConnectionRequest {
	op := &aws.Operation{
		Name:       opPostToConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/@connections/{connectionId}",
	}

	if input == nil {
		input = &PostToConnectionInput{}
	}

	req := c.newRequest(op, input, &PostToConnectionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PostToConnectionRequest{Request: req, Input: input, Copy: c.PostToConnectionRequest}
}

// PostToConnectionRequest is the request type for the
// PostToConnection API operation.
type PostToConnectionRequest struct {
	*aws.Request
	Input *PostToConnectionInput
	Copy  func(*PostToConnectionInput) PostToConnectionRequest
}

// Send marshals and sends the PostToConnection API request.
func (r PostToConnectionRequest) Send(ctx context.Context) (*PostToConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PostToConnectionResponse{
		PostToConnectionOutput: r.Request.Data.(*PostToConnectionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PostToConnectionResponse is the response type for the
// PostToConnection API operation.
type PostToConnectionResponse struct {
	*PostToConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PostToConnection request.
func (r *PostToConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
