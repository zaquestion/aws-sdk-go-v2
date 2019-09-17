// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type AcceptInvitationInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector of the GuardDuty member account.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// This value is used to validate the master account to the member account.
	//
	// InvitationId is a required field
	InvitationId *string `locationName:"invitationId" type:"string" required:"true"`

	// The account ID of the master GuardDuty account whose invitation you're accepting.
	//
	// MasterId is a required field
	MasterId *string `locationName:"masterId" type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptInvitationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptInvitationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptInvitationInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.InvitationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InvitationId"))
	}

	if s.MasterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MasterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptInvitationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InvitationId != nil {
		v := *s.InvitationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "invitationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MasterId != nil {
		v := *s.MasterId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "masterId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AcceptInvitationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AcceptInvitationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptInvitationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAcceptInvitation = "AcceptInvitation"

// AcceptInvitationRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Accepts the invitation to be monitored by a master GuardDuty account.
//
//    // Example sending a request using AcceptInvitationRequest.
//    req := client.AcceptInvitationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/AcceptInvitation
func (c *Client) AcceptInvitationRequest(input *AcceptInvitationInput) AcceptInvitationRequest {
	op := &aws.Operation{
		Name:       opAcceptInvitation,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/master",
	}

	if input == nil {
		input = &AcceptInvitationInput{}
	}

	req := c.newRequest(op, input, &AcceptInvitationOutput{})
	return AcceptInvitationRequest{Request: req, Input: input, Copy: c.AcceptInvitationRequest}
}

// AcceptInvitationRequest is the request type for the
// AcceptInvitation API operation.
type AcceptInvitationRequest struct {
	*aws.Request
	Input *AcceptInvitationInput
	Copy  func(*AcceptInvitationInput) AcceptInvitationRequest
}

// Send marshals and sends the AcceptInvitation API request.
func (r AcceptInvitationRequest) Send(ctx context.Context) (*AcceptInvitationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptInvitationResponse{
		AcceptInvitationOutput: r.Request.Data.(*AcceptInvitationOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptInvitationResponse is the response type for the
// AcceptInvitation API operation.
type AcceptInvitationResponse struct {
	*AcceptInvitationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptInvitation request.
func (r *AcceptInvitationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
