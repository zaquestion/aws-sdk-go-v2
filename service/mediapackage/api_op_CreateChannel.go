// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateChannelInput struct {
	_ struct{} `type:"structure"`

	Description *string `locationName:"description" type:"string"`

	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateChannelInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateChannelOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	Description *string `locationName:"description" type:"string"`

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *HlsIngest `locationName:"hlsIngest" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HlsIngest != nil {
		v := s.HlsIngest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsIngest", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opCreateChannel = "CreateChannel"

// CreateChannelRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Creates a new Channel.
//
//    // Example sending a request using CreateChannelRequest.
//    req := client.CreateChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/CreateChannel
func (c *Client) CreateChannelRequest(input *CreateChannelInput) CreateChannelRequest {
	op := &aws.Operation{
		Name:       opCreateChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/channels",
	}

	if input == nil {
		input = &CreateChannelInput{}
	}

	req := c.newRequest(op, input, &CreateChannelOutput{})
	return CreateChannelRequest{Request: req, Input: input, Copy: c.CreateChannelRequest}
}

// CreateChannelRequest is the request type for the
// CreateChannel API operation.
type CreateChannelRequest struct {
	*aws.Request
	Input *CreateChannelInput
	Copy  func(*CreateChannelInput) CreateChannelRequest
}

// Send marshals and sends the CreateChannel API request.
func (r CreateChannelRequest) Send(ctx context.Context) (*CreateChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateChannelResponse{
		CreateChannelOutput: r.Request.Data.(*CreateChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateChannelResponse is the response type for the
// CreateChannel API operation.
type CreateChannelResponse struct {
	*CreateChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateChannel request.
func (r *CreateChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
