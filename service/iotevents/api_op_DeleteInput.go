// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteInputInput struct {
	_ struct{} `type:"structure"`

	// The name of the input to delete.
	//
	// InputName is a required field
	InputName *string `location:"uri" locationName:"inputName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInputInput"}

	if s.InputName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputName"))
	}
	if s.InputName != nil && len(*s.InputName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInputInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InputName != nil {
		v := *s.InputName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "inputName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteInputOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteInputOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInputOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteInput = "DeleteInput"

// DeleteInputRequest returns a request value for making API operation for
// AWS IoT Events.
//
// Deletes an input.
//
//    // Example sending a request using DeleteInputRequest.
//    req := client.DeleteInputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/DeleteInput
func (c *Client) DeleteInputRequest(input *DeleteInputInput) DeleteInputRequest {
	op := &aws.Operation{
		Name:       opDeleteInput,
		HTTPMethod: "DELETE",
		HTTPPath:   "/inputs/{inputName}",
	}

	if input == nil {
		input = &DeleteInputInput{}
	}

	req := c.newRequest(op, input, &DeleteInputOutput{})
	return DeleteInputRequest{Request: req, Input: input, Copy: c.DeleteInputRequest}
}

// DeleteInputRequest is the request type for the
// DeleteInput API operation.
type DeleteInputRequest struct {
	*aws.Request
	Input *DeleteInputInput
	Copy  func(*DeleteInputInput) DeleteInputRequest
}

// Send marshals and sends the DeleteInput API request.
func (r DeleteInputRequest) Send(ctx context.Context) (*DeleteInputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInputResponse{
		DeleteInputOutput: r.Request.Data.(*DeleteInputOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInputResponse is the response type for the
// DeleteInput API operation.
type DeleteInputResponse struct {
	*DeleteInputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInput request.
func (r *DeleteInputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
