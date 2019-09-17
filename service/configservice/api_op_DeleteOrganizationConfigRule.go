// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteOrganizationConfigRuleInput struct {
	_ struct{} `type:"structure"`

	// OrganizationConfigRuleName is a required field
	OrganizationConfigRuleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOrganizationConfigRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOrganizationConfigRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOrganizationConfigRuleInput"}

	if s.OrganizationConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationConfigRuleName"))
	}
	if s.OrganizationConfigRuleName != nil && len(*s.OrganizationConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationConfigRuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOrganizationConfigRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOrganizationConfigRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteOrganizationConfigRule = "DeleteOrganizationConfigRule"

// DeleteOrganizationConfigRuleRequest returns a request value for making API operation for
// AWS Config.
//
//    // Example sending a request using DeleteOrganizationConfigRuleRequest.
//    req := client.DeleteOrganizationConfigRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteOrganizationConfigRule
func (c *Client) DeleteOrganizationConfigRuleRequest(input *DeleteOrganizationConfigRuleInput) DeleteOrganizationConfigRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteOrganizationConfigRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteOrganizationConfigRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteOrganizationConfigRuleOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteOrganizationConfigRuleRequest{Request: req, Input: input, Copy: c.DeleteOrganizationConfigRuleRequest}
}

// DeleteOrganizationConfigRuleRequest is the request type for the
// DeleteOrganizationConfigRule API operation.
type DeleteOrganizationConfigRuleRequest struct {
	*aws.Request
	Input *DeleteOrganizationConfigRuleInput
	Copy  func(*DeleteOrganizationConfigRuleInput) DeleteOrganizationConfigRuleRequest
}

// Send marshals and sends the DeleteOrganizationConfigRule API request.
func (r DeleteOrganizationConfigRuleRequest) Send(ctx context.Context) (*DeleteOrganizationConfigRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOrganizationConfigRuleResponse{
		DeleteOrganizationConfigRuleOutput: r.Request.Data.(*DeleteOrganizationConfigRuleOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOrganizationConfigRuleResponse is the response type for the
// DeleteOrganizationConfigRule API operation.
type DeleteOrganizationConfigRuleResponse struct {
	*DeleteOrganizationConfigRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOrganizationConfigRule request.
func (r *DeleteOrganizationConfigRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
