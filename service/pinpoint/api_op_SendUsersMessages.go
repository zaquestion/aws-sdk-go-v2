// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type SendUsersMessagesInput struct {
	_ struct{} `type:"structure" payload:"SendUsersMessageRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the configuration and other settings for a message to send to all
	// the endpoints that are associated with a list of users.
	//
	// SendUsersMessageRequest is a required field
	SendUsersMessageRequest *SendUsersMessageRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s SendUsersMessagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendUsersMessagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendUsersMessagesInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.SendUsersMessageRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("SendUsersMessageRequest"))
	}
	if s.SendUsersMessageRequest != nil {
		if err := s.SendUsersMessageRequest.Validate(); err != nil {
			invalidParams.AddNested("SendUsersMessageRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendUsersMessagesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SendUsersMessageRequest != nil {
		v := s.SendUsersMessageRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SendUsersMessageRequest", v, metadata)
	}
	return nil
}

type SendUsersMessagesOutput struct {
	_ struct{} `type:"structure" payload:"SendUsersMessageResponse"`

	// Provides information about which users and endpoints a message was sent to.
	//
	// SendUsersMessageResponse is a required field
	SendUsersMessageResponse *SendUsersMessageResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s SendUsersMessagesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendUsersMessagesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SendUsersMessageResponse != nil {
		v := s.SendUsersMessageResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SendUsersMessageResponse", v, metadata)
	}
	return nil
}

const opSendUsersMessages = "SendUsersMessages"

// SendUsersMessagesRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates and sends a message to a list of users.
//
//    // Example sending a request using SendUsersMessagesRequest.
//    req := client.SendUsersMessagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/SendUsersMessages
func (c *Client) SendUsersMessagesRequest(input *SendUsersMessagesInput) SendUsersMessagesRequest {
	op := &aws.Operation{
		Name:       opSendUsersMessages,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apps/{application-id}/users-messages",
	}

	if input == nil {
		input = &SendUsersMessagesInput{}
	}

	req := c.newRequest(op, input, &SendUsersMessagesOutput{})
	return SendUsersMessagesRequest{Request: req, Input: input, Copy: c.SendUsersMessagesRequest}
}

// SendUsersMessagesRequest is the request type for the
// SendUsersMessages API operation.
type SendUsersMessagesRequest struct {
	*aws.Request
	Input *SendUsersMessagesInput
	Copy  func(*SendUsersMessagesInput) SendUsersMessagesRequest
}

// Send marshals and sends the SendUsersMessages API request.
func (r SendUsersMessagesRequest) Send(ctx context.Context) (*SendUsersMessagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendUsersMessagesResponse{
		SendUsersMessagesOutput: r.Request.Data.(*SendUsersMessagesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendUsersMessagesResponse is the response type for the
// SendUsersMessages API operation.
type SendUsersMessagesResponse struct {
	*SendUsersMessagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendUsersMessages request.
func (r *SendUsersMessagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
