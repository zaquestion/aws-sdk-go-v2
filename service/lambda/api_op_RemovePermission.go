// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type RemovePermissionInput struct {
	_ struct{} `type:"structure"`

	// The name of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//    * Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function
	// name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// Specify a version or alias to remove permissions from a published version
	// of the function.
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string"`

	// Only update the policy if the revision ID matches the ID that's specified.
	// Use this option to avoid modifying a policy that has changed since you last
	// read it.
	RevisionId *string `location:"querystring" locationName:"RevisionId" type:"string"`

	// Statement ID of the permission to remove.
	//
	// StatementId is a required field
	StatementId *string `location:"uri" locationName:"StatementId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RemovePermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemovePermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemovePermissionInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.Qualifier != nil && len(*s.Qualifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Qualifier", 1))
	}

	if s.StatementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatementId"))
	}
	if s.StatementId != nil && len(*s.StatementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatementId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemovePermissionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StatementId != nil {
		v := *s.StatementId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "StatementId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Qualifier != nil {
		v := *s.Qualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Qualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RemovePermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemovePermissionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemovePermissionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRemovePermission = "RemovePermission"

// RemovePermissionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Revokes function-use permission from an AWS service or another account. You
// can get the ID of the statement from the output of GetPolicy.
//
//    // Example sending a request using RemovePermissionRequest.
//    req := client.RemovePermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/RemovePermission
func (c *Client) RemovePermissionRequest(input *RemovePermissionInput) RemovePermissionRequest {
	op := &aws.Operation{
		Name:       opRemovePermission,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/policy/{StatementId}",
	}

	if input == nil {
		input = &RemovePermissionInput{}
	}

	req := c.newRequest(op, input, &RemovePermissionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemovePermissionRequest{Request: req, Input: input, Copy: c.RemovePermissionRequest}
}

// RemovePermissionRequest is the request type for the
// RemovePermission API operation.
type RemovePermissionRequest struct {
	*aws.Request
	Input *RemovePermissionInput
	Copy  func(*RemovePermissionInput) RemovePermissionRequest
}

// Send marshals and sends the RemovePermission API request.
func (r RemovePermissionRequest) Send(ctx context.Context) (*RemovePermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemovePermissionResponse{
		RemovePermissionOutput: r.Request.Data.(*RemovePermissionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemovePermissionResponse is the response type for the
// RemovePermission API operation.
type RemovePermissionResponse struct {
	*RemovePermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemovePermission request.
func (r *RemovePermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
