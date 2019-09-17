// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DetachPolicyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the policy you want to detach. You can get
	// the ID from the ListPolicies or ListPoliciesForTarget operations.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string
	// requires "p-" followed by from 8 to 128 lower-case letters or digits.
	//
	// PolicyId is a required field
	PolicyId *string `type:"string" required:"true"`

	// The unique identifier (ID) of the root, OU, or account that you want to detach
	// the policy from. You can get the ID from the ListRoots, ListOrganizationalUnitsForParent,
	// or ListAccounts operations.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a target ID string
	// requires one of the following:
	//
	//    * Root: a string that begins with "r-" followed by from 4 to 32 lower-case
	//    letters or digits.
	//
	//    * Account: a string that consists of exactly 12 digits.
	//
	//    * Organizational unit (OU): a string that begins with "ou-" followed by
	//    from 4 to 32 lower-case letters or digits (the ID of the root that the
	//    OU is in) followed by a second "-" dash and from 8 to 32 additional lower-case
	//    letters or digits.
	//
	// TargetId is a required field
	TargetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachPolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}

	if s.TargetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DetachPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachPolicy = "DetachPolicy"

// DetachPolicyRequest returns a request value for making API operation for
// AWS Organizations.
//
// Detaches a policy from a target root, organizational unit (OU), or account.
// If the policy being detached is a service control policy (SCP), the changes
// to permissions for IAM users and roles in affected accounts are immediate.
//
// Note: Every root, OU, and account must have at least one SCP attached. If
// you want to replace the default FullAWSAccess policy with one that limits
// the permissions that can be delegated, you must attach the replacement policy
// before you can remove the default one. This is the authorization strategy
// of whitelisting (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_about-scps.html#orgs_policies_whitelist).
// If you instead attach a second SCP and leave the FullAWSAccess SCP still
// attached, and specify "Effect": "Deny" in the second SCP to override the
// "Effect": "Allow" in the FullAWSAccess policy (or any other attached SCP),
// you're using the authorization strategy of blacklisting (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_about-scps.html#orgs_policies_blacklist).
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using DetachPolicyRequest.
//    req := client.DetachPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DetachPolicy
func (c *Client) DetachPolicyRequest(input *DetachPolicyInput) DetachPolicyRequest {
	op := &aws.Operation{
		Name:       opDetachPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachPolicyInput{}
	}

	req := c.newRequest(op, input, &DetachPolicyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DetachPolicyRequest{Request: req, Input: input, Copy: c.DetachPolicyRequest}
}

// DetachPolicyRequest is the request type for the
// DetachPolicy API operation.
type DetachPolicyRequest struct {
	*aws.Request
	Input *DetachPolicyInput
	Copy  func(*DetachPolicyInput) DetachPolicyRequest
}

// Send marshals and sends the DetachPolicy API request.
func (r DetachPolicyRequest) Send(ctx context.Context) (*DetachPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachPolicyResponse{
		DetachPolicyOutput: r.Request.Data.(*DetachPolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachPolicyResponse is the response type for the
// DetachPolicy API operation.
type DetachPolicyResponse struct {
	*DetachPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachPolicy request.
func (r *DetachPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
