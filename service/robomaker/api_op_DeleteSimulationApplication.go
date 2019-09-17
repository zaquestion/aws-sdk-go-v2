// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteSimulationApplicationInput struct {
	_ struct{} `type:"structure"`

	// The application information for the simulation application to delete.
	//
	// Application is a required field
	Application *string `locationName:"application" min:"1" type:"string" required:"true"`

	// The version of the simulation application to delete.
	ApplicationVersion *string `locationName:"applicationVersion" min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteSimulationApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSimulationApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSimulationApplicationInput"}

	if s.Application == nil {
		invalidParams.Add(aws.NewErrParamRequired("Application"))
	}
	if s.Application != nil && len(*s.Application) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Application", 1))
	}
	if s.ApplicationVersion != nil && len(*s.ApplicationVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSimulationApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Application != nil {
		v := *s.Application

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "application", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApplicationVersion != nil {
		v := *s.ApplicationVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "applicationVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteSimulationApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSimulationApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSimulationApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteSimulationApplication = "DeleteSimulationApplication"

// DeleteSimulationApplicationRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Deletes a simulation application.
//
//    // Example sending a request using DeleteSimulationApplicationRequest.
//    req := client.DeleteSimulationApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DeleteSimulationApplication
func (c *Client) DeleteSimulationApplicationRequest(input *DeleteSimulationApplicationInput) DeleteSimulationApplicationRequest {
	op := &aws.Operation{
		Name:       opDeleteSimulationApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/deleteSimulationApplication",
	}

	if input == nil {
		input = &DeleteSimulationApplicationInput{}
	}

	req := c.newRequest(op, input, &DeleteSimulationApplicationOutput{})
	return DeleteSimulationApplicationRequest{Request: req, Input: input, Copy: c.DeleteSimulationApplicationRequest}
}

// DeleteSimulationApplicationRequest is the request type for the
// DeleteSimulationApplication API operation.
type DeleteSimulationApplicationRequest struct {
	*aws.Request
	Input *DeleteSimulationApplicationInput
	Copy  func(*DeleteSimulationApplicationInput) DeleteSimulationApplicationRequest
}

// Send marshals and sends the DeleteSimulationApplication API request.
func (r DeleteSimulationApplicationRequest) Send(ctx context.Context) (*DeleteSimulationApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSimulationApplicationResponse{
		DeleteSimulationApplicationOutput: r.Request.Data.(*DeleteSimulationApplicationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSimulationApplicationResponse is the response type for the
// DeleteSimulationApplication API operation.
type DeleteSimulationApplicationResponse struct {
	*DeleteSimulationApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSimulationApplication request.
func (r *DeleteSimulationApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
