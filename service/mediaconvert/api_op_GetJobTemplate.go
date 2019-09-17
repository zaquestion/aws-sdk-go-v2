// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Query a job template by sending a request with the job template name.
type GetJobTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the job template.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobTemplateInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Successful get job template requests will return an OK message and the job
// template JSON.
type GetJobTemplateOutput struct {
	_ struct{} `type:"structure"`

	// A job template is a pre-made set of encoding instructions that you can use
	// to quickly create a job.
	JobTemplate *JobTemplate `locationName:"jobTemplate" type:"structure"`
}

// String returns the string representation
func (s GetJobTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobTemplate != nil {
		v := s.JobTemplate

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jobTemplate", v, metadata)
	}
	return nil
}

const opGetJobTemplate = "GetJobTemplate"

// GetJobTemplateRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Retrieve the JSON for a specific job template.
//
//    // Example sending a request using GetJobTemplateRequest.
//    req := client.GetJobTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/GetJobTemplate
func (c *Client) GetJobTemplateRequest(input *GetJobTemplateInput) GetJobTemplateRequest {
	op := &aws.Operation{
		Name:       opGetJobTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/2017-08-29/jobTemplates/{name}",
	}

	if input == nil {
		input = &GetJobTemplateInput{}
	}

	req := c.newRequest(op, input, &GetJobTemplateOutput{})
	return GetJobTemplateRequest{Request: req, Input: input, Copy: c.GetJobTemplateRequest}
}

// GetJobTemplateRequest is the request type for the
// GetJobTemplate API operation.
type GetJobTemplateRequest struct {
	*aws.Request
	Input *GetJobTemplateInput
	Copy  func(*GetJobTemplateInput) GetJobTemplateRequest
}

// Send marshals and sends the GetJobTemplate API request.
func (r GetJobTemplateRequest) Send(ctx context.Context) (*GetJobTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobTemplateResponse{
		GetJobTemplateOutput: r.Request.Data.(*GetJobTemplateOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobTemplateResponse is the response type for the
// GetJobTemplate API operation.
type GetJobTemplateResponse struct {
	*GetJobTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJobTemplate request.
func (r *GetJobTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
