// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RefreshSchemasInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the replication instance.
	//
	// ReplicationInstanceArn is a required field
	ReplicationInstanceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RefreshSchemasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RefreshSchemasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RefreshSchemasInput"}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if s.ReplicationInstanceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationInstanceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RefreshSchemasOutput struct {
	_ struct{} `type:"structure"`

	// The status of the refreshed schema.
	RefreshSchemasStatus *RefreshSchemasStatus `type:"structure"`
}

// String returns the string representation
func (s RefreshSchemasOutput) String() string {
	return awsutil.Prettify(s)
}

const opRefreshSchemas = "RefreshSchemas"

// RefreshSchemasRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Populates the schema for the specified endpoint. This is an asynchronous
// operation and can take several minutes. You can check the status of this
// operation by calling the DescribeRefreshSchemasStatus operation.
//
//    // Example sending a request using RefreshSchemasRequest.
//    req := client.RefreshSchemasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/RefreshSchemas
func (c *Client) RefreshSchemasRequest(input *RefreshSchemasInput) RefreshSchemasRequest {
	op := &aws.Operation{
		Name:       opRefreshSchemas,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RefreshSchemasInput{}
	}

	req := c.newRequest(op, input, &RefreshSchemasOutput{})
	return RefreshSchemasRequest{Request: req, Input: input, Copy: c.RefreshSchemasRequest}
}

// RefreshSchemasRequest is the request type for the
// RefreshSchemas API operation.
type RefreshSchemasRequest struct {
	*aws.Request
	Input *RefreshSchemasInput
	Copy  func(*RefreshSchemasInput) RefreshSchemasRequest
}

// Send marshals and sends the RefreshSchemas API request.
func (r RefreshSchemasRequest) Send(ctx context.Context) (*RefreshSchemasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RefreshSchemasResponse{
		RefreshSchemasOutput: r.Request.Data.(*RefreshSchemasOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RefreshSchemasResponse is the response type for the
// RefreshSchemas API operation.
type RefreshSchemasResponse struct {
	*RefreshSchemasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RefreshSchemas request.
func (r *RefreshSchemasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
