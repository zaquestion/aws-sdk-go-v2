// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateApplicationSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of an existing application
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// An identifier for the application snapshot.
	//
	// SnapshotName is a required field
	SnapshotName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateApplicationSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApplicationSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApplicationSnapshotInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.SnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotName"))
	}
	if s.SnapshotName != nil && len(*s.SnapshotName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnapshotName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateApplicationSnapshotOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateApplicationSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateApplicationSnapshot = "CreateApplicationSnapshot"

// CreateApplicationSnapshotRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Creates a snapshot of the application's state data.
//
//    // Example sending a request using CreateApplicationSnapshotRequest.
//    req := client.CreateApplicationSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/CreateApplicationSnapshot
func (c *Client) CreateApplicationSnapshotRequest(input *CreateApplicationSnapshotInput) CreateApplicationSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateApplicationSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateApplicationSnapshotInput{}
	}

	req := c.newRequest(op, input, &CreateApplicationSnapshotOutput{})
	return CreateApplicationSnapshotRequest{Request: req, Input: input, Copy: c.CreateApplicationSnapshotRequest}
}

// CreateApplicationSnapshotRequest is the request type for the
// CreateApplicationSnapshot API operation.
type CreateApplicationSnapshotRequest struct {
	*aws.Request
	Input *CreateApplicationSnapshotInput
	Copy  func(*CreateApplicationSnapshotInput) CreateApplicationSnapshotRequest
}

// Send marshals and sends the CreateApplicationSnapshot API request.
func (r CreateApplicationSnapshotRequest) Send(ctx context.Context) (*CreateApplicationSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationSnapshotResponse{
		CreateApplicationSnapshotOutput: r.Request.Data.(*CreateApplicationSnapshotOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationSnapshotResponse is the response type for the
// CreateApplicationSnapshot API operation.
type CreateApplicationSnapshotResponse struct {
	*CreateApplicationSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplicationSnapshot request.
func (r *CreateApplicationSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
