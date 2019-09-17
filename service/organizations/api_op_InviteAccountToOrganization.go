// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type InviteAccountToOrganizationInput struct {
	_ struct{} `type:"structure"`

	// Additional information that you want to include in the generated email to
	// the recipient account owner.
	Notes *string `type:"string"`

	// The identifier (ID) of the AWS account that you want to invite to join your
	// organization. This is a JSON object that contains the following elements:
	//
	// { "Type": "ACCOUNT", "Id": "< account id number >" }
	//
	// If you use the AWS CLI, you can submit this as a single string, similar to
	// the following example:
	//
	// --target Id=123456789012,Type=ACCOUNT
	//
	// If you specify "Type": "ACCOUNT", you must provide the AWS account ID number
	// as the Id. If you specify "Type": "EMAIL", you must specify the email address
	// that is associated with the account.
	//
	// --target Id=diego@example.com,Type=EMAIL
	//
	// Target is a required field
	Target *HandshakeParty `type:"structure" required:"true"`
}

// String returns the string representation
func (s InviteAccountToOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InviteAccountToOrganizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InviteAccountToOrganizationInput"}

	if s.Target == nil {
		invalidParams.Add(aws.NewErrParamRequired("Target"))
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			invalidParams.AddNested("Target", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type InviteAccountToOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the handshake that is created to
	// support this invitation request.
	Handshake *Handshake `type:"structure"`
}

// String returns the string representation
func (s InviteAccountToOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opInviteAccountToOrganization = "InviteAccountToOrganization"

// InviteAccountToOrganizationRequest returns a request value for making API operation for
// AWS Organizations.
//
// Sends an invitation to another account to join your organization as a member
// account. AWS Organizations sends email on your behalf to the email address
// that is associated with the other account's owner. The invitation is implemented
// as a Handshake whose details are in the response.
//
//    * You can invite AWS accounts only from the same seller as the master
//    account. For example, if your organization's master account was created
//    by Amazon Internet Services Pvt. Ltd (AISPL), an AWS seller in India,
//    you can invite only other AISPL accounts to your organization. You can't
//    combine accounts from AISPL and AWS or from any other AWS seller. For
//    more information, see Consolidated Billing in India (http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/useconsolidatedbilliing-India.html).
//
//    * If you receive an exception that indicates that you exceeded your account
//    limits for the organization or that the operation failed because your
//    organization is still initializing, wait one hour and then try again.
//    If the error persists after an hour, contact AWS Support (https://console.aws.amazon.com/support/home#/).
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using InviteAccountToOrganizationRequest.
//    req := client.InviteAccountToOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/InviteAccountToOrganization
func (c *Client) InviteAccountToOrganizationRequest(input *InviteAccountToOrganizationInput) InviteAccountToOrganizationRequest {
	op := &aws.Operation{
		Name:       opInviteAccountToOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InviteAccountToOrganizationInput{}
	}

	req := c.newRequest(op, input, &InviteAccountToOrganizationOutput{})
	return InviteAccountToOrganizationRequest{Request: req, Input: input, Copy: c.InviteAccountToOrganizationRequest}
}

// InviteAccountToOrganizationRequest is the request type for the
// InviteAccountToOrganization API operation.
type InviteAccountToOrganizationRequest struct {
	*aws.Request
	Input *InviteAccountToOrganizationInput
	Copy  func(*InviteAccountToOrganizationInput) InviteAccountToOrganizationRequest
}

// Send marshals and sends the InviteAccountToOrganization API request.
func (r InviteAccountToOrganizationRequest) Send(ctx context.Context) (*InviteAccountToOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InviteAccountToOrganizationResponse{
		InviteAccountToOrganizationOutput: r.Request.Data.(*InviteAccountToOrganizationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InviteAccountToOrganizationResponse is the response type for the
// InviteAccountToOrganization API operation.
type InviteAccountToOrganizationResponse struct {
	*InviteAccountToOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InviteAccountToOrganization request.
func (r *InviteAccountToOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
