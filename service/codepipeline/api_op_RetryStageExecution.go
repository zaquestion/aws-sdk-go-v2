// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a RetryStageExecution action.
type RetryStageExecutionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the pipeline execution in the failed stage to be retried. Use the
	// GetPipelineState action to retrieve the current pipelineExecutionId of the
	// failed stage
	//
	// PipelineExecutionId is a required field
	PipelineExecutionId *string `locationName:"pipelineExecutionId" type:"string" required:"true"`

	// The name of the pipeline that contains the failed stage.
	//
	// PipelineName is a required field
	PipelineName *string `locationName:"pipelineName" min:"1" type:"string" required:"true"`

	// The scope of the retry attempt. Currently, the only supported value is FAILED_ACTIONS.
	//
	// RetryMode is a required field
	RetryMode StageRetryMode `locationName:"retryMode" type:"string" required:"true" enum:"true"`

	// The name of the failed stage to be retried.
	//
	// StageName is a required field
	StageName *string `locationName:"stageName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RetryStageExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RetryStageExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RetryStageExecutionInput"}

	if s.PipelineExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineExecutionId"))
	}

	if s.PipelineName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineName"))
	}
	if s.PipelineName != nil && len(*s.PipelineName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineName", 1))
	}
	if len(s.RetryMode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RetryMode"))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}
	if s.StageName != nil && len(*s.StageName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StageName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a RetryStageExecution action.
type RetryStageExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the current workflow execution in the failed stage.
	PipelineExecutionId *string `locationName:"pipelineExecutionId" type:"string"`
}

// String returns the string representation
func (s RetryStageExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opRetryStageExecution = "RetryStageExecution"

// RetryStageExecutionRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Resumes the pipeline execution by retrying the last failed actions in a stage.
//
//    // Example sending a request using RetryStageExecutionRequest.
//    req := client.RetryStageExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/RetryStageExecution
func (c *Client) RetryStageExecutionRequest(input *RetryStageExecutionInput) RetryStageExecutionRequest {
	op := &aws.Operation{
		Name:       opRetryStageExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RetryStageExecutionInput{}
	}

	req := c.newRequest(op, input, &RetryStageExecutionOutput{})
	return RetryStageExecutionRequest{Request: req, Input: input, Copy: c.RetryStageExecutionRequest}
}

// RetryStageExecutionRequest is the request type for the
// RetryStageExecution API operation.
type RetryStageExecutionRequest struct {
	*aws.Request
	Input *RetryStageExecutionInput
	Copy  func(*RetryStageExecutionInput) RetryStageExecutionRequest
}

// Send marshals and sends the RetryStageExecution API request.
func (r RetryStageExecutionRequest) Send(ctx context.Context) (*RetryStageExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RetryStageExecutionResponse{
		RetryStageExecutionOutput: r.Request.Data.(*RetryStageExecutionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RetryStageExecutionResponse is the response type for the
// RetryStageExecution API operation.
type RetryStageExecutionResponse struct {
	*RetryStageExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RetryStageExecution request.
func (r *RetryStageExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
