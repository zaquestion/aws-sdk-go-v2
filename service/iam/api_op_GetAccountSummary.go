// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccountSummaryInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetAccountSummaryInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a successful GetAccountSummary request.
type GetAccountSummaryOutput struct {
	_ struct{} `type:"structure"`

	// A set of key–value pairs containing information about IAM entity usage
	// and IAM quotas.
	SummaryMap map[string]int64 `type:"map"`
}

// String returns the string representation
func (s GetAccountSummaryOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAccountSummary = "GetAccountSummary"

// GetAccountSummaryRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves information about IAM entity usage and IAM quotas in the AWS account.
//
// For information about limitations on IAM entities, see Limitations on IAM
// Entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html)
// in the IAM User Guide.
//
//    // Example sending a request using GetAccountSummaryRequest.
//    req := client.GetAccountSummaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetAccountSummary
func (c *Client) GetAccountSummaryRequest(input *GetAccountSummaryInput) GetAccountSummaryRequest {
	op := &aws.Operation{
		Name:       opGetAccountSummary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAccountSummaryInput{}
	}

	req := c.newRequest(op, input, &GetAccountSummaryOutput{})
	return GetAccountSummaryRequest{Request: req, Input: input, Copy: c.GetAccountSummaryRequest}
}

// GetAccountSummaryRequest is the request type for the
// GetAccountSummary API operation.
type GetAccountSummaryRequest struct {
	*aws.Request
	Input *GetAccountSummaryInput
	Copy  func(*GetAccountSummaryInput) GetAccountSummaryRequest
}

// Send marshals and sends the GetAccountSummary API request.
func (r GetAccountSummaryRequest) Send(ctx context.Context) (*GetAccountSummaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountSummaryResponse{
		GetAccountSummaryOutput: r.Request.Data.(*GetAccountSummaryOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountSummaryResponse is the response type for the
// GetAccountSummary API operation.
type GetAccountSummaryResponse struct {
	*GetAccountSummaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccountSummary request.
func (r *GetAccountSummaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
