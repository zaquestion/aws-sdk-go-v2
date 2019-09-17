// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dlm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetLifecyclePolicyInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the lifecycle policy.
	//
	// PolicyId is a required field
	PolicyId *string `location:"uri" locationName:"policyId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLifecyclePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLifecyclePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLifecyclePolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLifecyclePolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PolicyId != nil {
		v := *s.PolicyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "policyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetLifecyclePolicyOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about the lifecycle policy.
	Policy *LifecyclePolicy `type:"structure"`
}

// String returns the string representation
func (s GetLifecyclePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLifecyclePolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Policy != nil {
		v := s.Policy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Policy", v, metadata)
	}
	return nil
}

const opGetLifecyclePolicy = "GetLifecyclePolicy"

// GetLifecyclePolicyRequest returns a request value for making API operation for
// Amazon Data Lifecycle Manager.
//
// Gets detailed information about the specified lifecycle policy.
//
//    // Example sending a request using GetLifecyclePolicyRequest.
//    req := client.GetLifecyclePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dlm-2018-01-12/GetLifecyclePolicy
func (c *Client) GetLifecyclePolicyRequest(input *GetLifecyclePolicyInput) GetLifecyclePolicyRequest {
	op := &aws.Operation{
		Name:       opGetLifecyclePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/policies/{policyId}/",
	}

	if input == nil {
		input = &GetLifecyclePolicyInput{}
	}

	req := c.newRequest(op, input, &GetLifecyclePolicyOutput{})
	return GetLifecyclePolicyRequest{Request: req, Input: input, Copy: c.GetLifecyclePolicyRequest}
}

// GetLifecyclePolicyRequest is the request type for the
// GetLifecyclePolicy API operation.
type GetLifecyclePolicyRequest struct {
	*aws.Request
	Input *GetLifecyclePolicyInput
	Copy  func(*GetLifecyclePolicyInput) GetLifecyclePolicyRequest
}

// Send marshals and sends the GetLifecyclePolicy API request.
func (r GetLifecyclePolicyRequest) Send(ctx context.Context) (*GetLifecyclePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLifecyclePolicyResponse{
		GetLifecyclePolicyOutput: r.Request.Data.(*GetLifecyclePolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLifecyclePolicyResponse is the response type for the
// GetLifecyclePolicy API operation.
type GetLifecyclePolicyResponse struct {
	*GetLifecyclePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLifecyclePolicy request.
func (r *GetLifecyclePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
