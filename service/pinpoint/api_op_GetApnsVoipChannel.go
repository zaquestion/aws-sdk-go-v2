// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetApnsVoipChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetApnsVoipChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApnsVoipChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApnsVoipChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApnsVoipChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetApnsVoipChannelOutput struct {
	_ struct{} `type:"structure" payload:"APNSVoipChannelResponse"`

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP channel for an application.
	//
	// APNSVoipChannelResponse is a required field
	APNSVoipChannelResponse *APNSVoipChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetApnsVoipChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApnsVoipChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.APNSVoipChannelResponse != nil {
		v := s.APNSVoipChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSVoipChannelResponse", v, metadata)
	}
	return nil
}

const opGetApnsVoipChannel = "GetApnsVoipChannel"

// GetApnsVoipChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of the APNs VoIP channel
// for an application.
//
//    // Example sending a request using GetApnsVoipChannelRequest.
//    req := client.GetApnsVoipChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetApnsVoipChannel
func (c *Client) GetApnsVoipChannelRequest(input *GetApnsVoipChannelInput) GetApnsVoipChannelRequest {
	op := &aws.Operation{
		Name:       opGetApnsVoipChannel,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/channels/apns_voip",
	}

	if input == nil {
		input = &GetApnsVoipChannelInput{}
	}

	req := c.newRequest(op, input, &GetApnsVoipChannelOutput{})
	return GetApnsVoipChannelRequest{Request: req, Input: input, Copy: c.GetApnsVoipChannelRequest}
}

// GetApnsVoipChannelRequest is the request type for the
// GetApnsVoipChannel API operation.
type GetApnsVoipChannelRequest struct {
	*aws.Request
	Input *GetApnsVoipChannelInput
	Copy  func(*GetApnsVoipChannelInput) GetApnsVoipChannelRequest
}

// Send marshals and sends the GetApnsVoipChannel API request.
func (r GetApnsVoipChannelRequest) Send(ctx context.Context) (*GetApnsVoipChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApnsVoipChannelResponse{
		GetApnsVoipChannelOutput: r.Request.Data.(*GetApnsVoipChannelOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApnsVoipChannelResponse is the response type for the
// GetApnsVoipChannel API operation.
type GetApnsVoipChannelResponse struct {
	*GetApnsVoipChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApnsVoipChannel request.
func (r *GetApnsVoipChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
