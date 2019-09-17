// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of an UpdatePipeline action.
type UpdatePipelineInput struct {
	_ struct{} `type:"structure"`

	// The name of the pipeline to be updated.
	//
	// Pipeline is a required field
	Pipeline *PipelineDeclaration `locationName:"pipeline" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePipelineInput"}

	if s.Pipeline == nil {
		invalidParams.Add(aws.NewErrParamRequired("Pipeline"))
	}
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			invalidParams.AddNested("Pipeline", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an UpdatePipeline action.
type UpdatePipelineOutput struct {
	_ struct{} `type:"structure"`

	// The structure of the updated pipeline.
	Pipeline *PipelineDeclaration `locationName:"pipeline" type:"structure"`
}

// String returns the string representation
func (s UpdatePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdatePipeline = "UpdatePipeline"

// UpdatePipelineRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Updates a specified pipeline with edits or changes to its structure. Use
// a JSON file with the pipeline structure in conjunction with UpdatePipeline
// to provide the full structure of the pipeline. Updating the pipeline increases
// the version number of the pipeline by 1.
//
//    // Example sending a request using UpdatePipelineRequest.
//    req := client.UpdatePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/UpdatePipeline
func (c *Client) UpdatePipelineRequest(input *UpdatePipelineInput) UpdatePipelineRequest {
	op := &aws.Operation{
		Name:       opUpdatePipeline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePipelineInput{}
	}

	req := c.newRequest(op, input, &UpdatePipelineOutput{})
	return UpdatePipelineRequest{Request: req, Input: input, Copy: c.UpdatePipelineRequest}
}

// UpdatePipelineRequest is the request type for the
// UpdatePipeline API operation.
type UpdatePipelineRequest struct {
	*aws.Request
	Input *UpdatePipelineInput
	Copy  func(*UpdatePipelineInput) UpdatePipelineRequest
}

// Send marshals and sends the UpdatePipeline API request.
func (r UpdatePipelineRequest) Send(ctx context.Context) (*UpdatePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePipelineResponse{
		UpdatePipelineOutput: r.Request.Data.(*UpdatePipelineOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePipelineResponse is the response type for the
// UpdatePipeline API operation.
type UpdatePipelineResponse struct {
	*UpdatePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePipeline request.
func (r *UpdatePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
