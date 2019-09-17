// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateInputSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	// InputSecurityGroupId is a required field
	InputSecurityGroupId *string `location:"uri" locationName:"inputSecurityGroupId" type:"string" required:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`

	WhitelistRules []InputWhitelistRuleCidr `locationName:"whitelistRules" type:"list"`
}

// String returns the string representation
func (s UpdateInputSecurityGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateInputSecurityGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateInputSecurityGroupInput"}

	if s.InputSecurityGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputSecurityGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateInputSecurityGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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
	if s.WhitelistRules != nil {
		v := s.WhitelistRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "whitelistRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.InputSecurityGroupId != nil {
		v := *s.InputSecurityGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "inputSecurityGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateInputSecurityGroupOutput struct {
	_ struct{} `type:"structure"`

	// An Input Security Group
	SecurityGroup *InputSecurityGroup `locationName:"securityGroup" type:"structure"`
}

// String returns the string representation
func (s UpdateInputSecurityGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateInputSecurityGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SecurityGroup != nil {
		v := s.SecurityGroup

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "securityGroup", v, metadata)
	}
	return nil
}

const opUpdateInputSecurityGroup = "UpdateInputSecurityGroup"

// UpdateInputSecurityGroupRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Update an Input Security Group's Whilelists.
//
//    // Example sending a request using UpdateInputSecurityGroupRequest.
//    req := client.UpdateInputSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/UpdateInputSecurityGroup
func (c *Client) UpdateInputSecurityGroupRequest(input *UpdateInputSecurityGroupInput) UpdateInputSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateInputSecurityGroup,
		HTTPMethod: "PUT",
		HTTPPath:   "/prod/inputSecurityGroups/{inputSecurityGroupId}",
	}

	if input == nil {
		input = &UpdateInputSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateInputSecurityGroupOutput{})
	return UpdateInputSecurityGroupRequest{Request: req, Input: input, Copy: c.UpdateInputSecurityGroupRequest}
}

// UpdateInputSecurityGroupRequest is the request type for the
// UpdateInputSecurityGroup API operation.
type UpdateInputSecurityGroupRequest struct {
	*aws.Request
	Input *UpdateInputSecurityGroupInput
	Copy  func(*UpdateInputSecurityGroupInput) UpdateInputSecurityGroupRequest
}

// Send marshals and sends the UpdateInputSecurityGroup API request.
func (r UpdateInputSecurityGroupRequest) Send(ctx context.Context) (*UpdateInputSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateInputSecurityGroupResponse{
		UpdateInputSecurityGroupOutput: r.Request.Data.(*UpdateInputSecurityGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateInputSecurityGroupResponse is the response type for the
// UpdateInputSecurityGroup API operation.
type UpdateInputSecurityGroupResponse struct {
	*UpdateInputSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateInputSecurityGroup request.
func (r *UpdateInputSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
