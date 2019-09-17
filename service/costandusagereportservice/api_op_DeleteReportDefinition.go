// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costandusagereportservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Deletes the specified report.
type DeleteReportDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The name of the report that you want to create. The name must be unique,
	// is case sensitive, and can't include spaces.
	ReportName *string `type:"string"`
}

// String returns the string representation
func (s DeleteReportDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// If the action is successful, the service sends back an HTTP 200 response.
type DeleteReportDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// Whether the deletion was successful or not.
	ResponseMessage *string `type:"string"`
}

// String returns the string representation
func (s DeleteReportDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteReportDefinition = "DeleteReportDefinition"

// DeleteReportDefinitionRequest returns a request value for making API operation for
// AWS Cost and Usage Report Service.
//
// Deletes the specified report.
//
//    // Example sending a request using DeleteReportDefinitionRequest.
//    req := client.DeleteReportDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cur-2017-01-06/DeleteReportDefinition
func (c *Client) DeleteReportDefinitionRequest(input *DeleteReportDefinitionInput) DeleteReportDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteReportDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteReportDefinitionInput{}
	}

	req := c.newRequest(op, input, &DeleteReportDefinitionOutput{})
	return DeleteReportDefinitionRequest{Request: req, Input: input, Copy: c.DeleteReportDefinitionRequest}
}

// DeleteReportDefinitionRequest is the request type for the
// DeleteReportDefinition API operation.
type DeleteReportDefinitionRequest struct {
	*aws.Request
	Input *DeleteReportDefinitionInput
	Copy  func(*DeleteReportDefinitionInput) DeleteReportDefinitionRequest
}

// Send marshals and sends the DeleteReportDefinition API request.
func (r DeleteReportDefinitionRequest) Send(ctx context.Context) (*DeleteReportDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReportDefinitionResponse{
		DeleteReportDefinitionOutput: r.Request.Data.(*DeleteReportDefinitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReportDefinitionResponse is the response type for the
// DeleteReportDefinition API operation.
type DeleteReportDefinitionResponse struct {
	*DeleteReportDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReportDefinition request.
func (r *DeleteReportDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
