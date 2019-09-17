// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The settings for the updated source of the flow.
type UpdateFlowSourceInput struct {
	_ struct{} `type:"structure"`

	// The type of encryption used on the content ingested from this source.
	Decryption *UpdateEncryption `locationName:"decryption" type:"structure"`

	// A description for the source. This value is not used or seen outside of the
	// current AWS Elemental MediaConnect account.
	Description *string `locationName:"description" type:"string"`

	// The ARN of the entitlement that allows you to subscribe to this flow. The
	// entitlement is set by the flow originator, and the ARN is generated as part
	// of the originator's flow.
	EntitlementArn *string `locationName:"entitlementArn" type:"string"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// The port that the flow will be listening on for incoming content.
	IngestPort *int64 `locationName:"ingestPort" type:"integer"`

	// The smoothing max bitrate for RTP and RTP-FEC streams.
	MaxBitrate *int64 `locationName:"maxBitrate" type:"integer"`

	// The maximum latency in milliseconds for Zixi-based streams.
	MaxLatency *int64 `locationName:"maxLatency" type:"integer"`

	// The protocol that is used by the source.
	Protocol Protocol `locationName:"protocol" type:"string" enum:"true"`

	// SourceArn is a required field
	SourceArn *string `location:"uri" locationName:"sourceArn" type:"string" required:"true"`

	// The stream ID that you want to use for this transport. This parameter applies
	// only to Zixi-based streams.
	StreamId *string `locationName:"streamId" type:"string"`

	// The range of IP addresses that should be allowed to contribute content to
	// your source. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string `locationName:"whitelistCidr" type:"string"`
}

// String returns the string representation
func (s UpdateFlowSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFlowSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFlowSourceInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if s.SourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFlowSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Decryption != nil {
		v := s.Decryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "decryption", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EntitlementArn != nil {
		v := *s.EntitlementArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "entitlementArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IngestPort != nil {
		v := *s.IngestPort

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ingestPort", protocol.Int64Value(v), metadata)
	}
	if s.MaxBitrate != nil {
		v := *s.MaxBitrate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxBitrate", protocol.Int64Value(v), metadata)
	}
	if s.MaxLatency != nil {
		v := *s.MaxLatency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxLatency", protocol.Int64Value(v), metadata)
	}
	if len(s.Protocol) > 0 {
		v := s.Protocol

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "protocol", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StreamId != nil {
		v := *s.StreamId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.WhitelistCidr != nil {
		v := *s.WhitelistCidr

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "whitelistCidr", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "sourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful UpdateFlowSource request. The response includes
// the ARN of the flow that was updated and the updated source configuration.
type UpdateFlowSourceOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that you want to update.
	FlowArn *string `locationName:"flowArn" type:"string"`

	// The settings for the source of the flow.
	Source *Source `locationName:"source" type:"structure"`
}

// String returns the string representation
func (s UpdateFlowSourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFlowSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Source != nil {
		v := s.Source

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "source", v, metadata)
	}
	return nil
}

const opUpdateFlowSource = "UpdateFlowSource"

// UpdateFlowSourceRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Updates the source of a flow.
//
//    // Example sending a request using UpdateFlowSourceRequest.
//    req := client.UpdateFlowSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/UpdateFlowSource
func (c *Client) UpdateFlowSourceRequest(input *UpdateFlowSourceInput) UpdateFlowSourceRequest {
	op := &aws.Operation{
		Name:       opUpdateFlowSource,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/flows/{flowArn}/source/{sourceArn}",
	}

	if input == nil {
		input = &UpdateFlowSourceInput{}
	}

	req := c.newRequest(op, input, &UpdateFlowSourceOutput{})
	return UpdateFlowSourceRequest{Request: req, Input: input, Copy: c.UpdateFlowSourceRequest}
}

// UpdateFlowSourceRequest is the request type for the
// UpdateFlowSource API operation.
type UpdateFlowSourceRequest struct {
	*aws.Request
	Input *UpdateFlowSourceInput
	Copy  func(*UpdateFlowSourceInput) UpdateFlowSourceRequest
}

// Send marshals and sends the UpdateFlowSource API request.
func (r UpdateFlowSourceRequest) Send(ctx context.Context) (*UpdateFlowSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFlowSourceResponse{
		UpdateFlowSourceOutput: r.Request.Data.(*UpdateFlowSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFlowSourceResponse is the response type for the
// UpdateFlowSource API operation.
type UpdateFlowSourceResponse struct {
	*UpdateFlowSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFlowSource request.
func (r *UpdateFlowSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
