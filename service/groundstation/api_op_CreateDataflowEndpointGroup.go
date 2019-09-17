// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateDataflowEndpointGroupInput struct {
	_ struct{} `type:"structure"`

	// EndpointDetails is a required field
	EndpointDetails []EndpointDetails `locationName:"endpointDetails" type:"list" required:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDataflowEndpointGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataflowEndpointGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataflowEndpointGroupInput"}

	if s.EndpointDetails == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointDetails"))
	}
	if s.EndpointDetails != nil {
		for i, v := range s.EndpointDetails {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EndpointDetails", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDataflowEndpointGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EndpointDetails != nil {
		v := s.EndpointDetails

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "endpointDetails", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

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

type CreateDataflowEndpointGroupOutput struct {
	_ struct{} `type:"structure"`

	DataflowEndpointGroupId *string `locationName:"dataflowEndpointGroupId" type:"string"`
}

// String returns the string representation
func (s CreateDataflowEndpointGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDataflowEndpointGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataflowEndpointGroupId != nil {
		v := *s.DataflowEndpointGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dataflowEndpointGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDataflowEndpointGroup = "CreateDataflowEndpointGroup"

// CreateDataflowEndpointGroupRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Creates a DataflowEndpoint group containing the specified list of DataflowEndpoint
// objects.
//
// The name field in each endpoint is used in your mission profile DataflowEndpointConfig
// to specify which endpoints to use during a contact.
//
// When a contact uses multiple DataflowEndpointConfig objects, each Config
// must match a DataflowEndpoint in the same group.
//
//    // Example sending a request using CreateDataflowEndpointGroupRequest.
//    req := client.CreateDataflowEndpointGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/CreateDataflowEndpointGroup
func (c *Client) CreateDataflowEndpointGroupRequest(input *CreateDataflowEndpointGroupInput) CreateDataflowEndpointGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDataflowEndpointGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/dataflowEndpointGroup",
	}

	if input == nil {
		input = &CreateDataflowEndpointGroupInput{}
	}

	req := c.newRequest(op, input, &CreateDataflowEndpointGroupOutput{})
	return CreateDataflowEndpointGroupRequest{Request: req, Input: input, Copy: c.CreateDataflowEndpointGroupRequest}
}

// CreateDataflowEndpointGroupRequest is the request type for the
// CreateDataflowEndpointGroup API operation.
type CreateDataflowEndpointGroupRequest struct {
	*aws.Request
	Input *CreateDataflowEndpointGroupInput
	Copy  func(*CreateDataflowEndpointGroupInput) CreateDataflowEndpointGroupRequest
}

// Send marshals and sends the CreateDataflowEndpointGroup API request.
func (r CreateDataflowEndpointGroupRequest) Send(ctx context.Context) (*CreateDataflowEndpointGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDataflowEndpointGroupResponse{
		CreateDataflowEndpointGroupOutput: r.Request.Data.(*CreateDataflowEndpointGroupOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDataflowEndpointGroupResponse is the response type for the
// CreateDataflowEndpointGroup API operation.
type CreateDataflowEndpointGroupResponse struct {
	*CreateDataflowEndpointGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataflowEndpointGroup request.
func (r *CreateDataflowEndpointGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
