// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateOrganizationalUnitInput struct {
	_ struct{} `type:"structure"`

	// The new name that you want to assign to the OU.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of any of the characters in the ASCII character
	// range.
	Name *string `min:"1" type:"string"`

	// The unique identifier (ID) of the OU that you want to rename. You can get
	// the ID from the ListOrganizationalUnitsForParent operation.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an organizational
	// unit ID string requires "ou-" followed by from 4 to 32 lower-case letters
	// or digits (the ID of the root that contains the OU) followed by a second
	// "-" dash and from 8 to 32 additional lower-case letters or digits.
	//
	// OrganizationalUnitId is a required field
	OrganizationalUnitId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateOrganizationalUnitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateOrganizationalUnitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateOrganizationalUnitInput"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.OrganizationalUnitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationalUnitId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateOrganizationalUnitOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains the details about the specified OU, including its
	// new name.
	OrganizationalUnit *OrganizationalUnit `type:"structure"`
}

// String returns the string representation
func (s UpdateOrganizationalUnitOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateOrganizationalUnit = "UpdateOrganizationalUnit"

// UpdateOrganizationalUnitRequest returns a request value for making API operation for
// AWS Organizations.
//
// Renames the specified organizational unit (OU). The ID and ARN don't change.
// The child OUs and accounts remain in place, and any attached policies of
// the OU remain attached.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using UpdateOrganizationalUnitRequest.
//    req := client.UpdateOrganizationalUnitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/UpdateOrganizationalUnit
func (c *Client) UpdateOrganizationalUnitRequest(input *UpdateOrganizationalUnitInput) UpdateOrganizationalUnitRequest {
	op := &aws.Operation{
		Name:       opUpdateOrganizationalUnit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateOrganizationalUnitInput{}
	}

	req := c.newRequest(op, input, &UpdateOrganizationalUnitOutput{})
	return UpdateOrganizationalUnitRequest{Request: req, Input: input, Copy: c.UpdateOrganizationalUnitRequest}
}

// UpdateOrganizationalUnitRequest is the request type for the
// UpdateOrganizationalUnit API operation.
type UpdateOrganizationalUnitRequest struct {
	*aws.Request
	Input *UpdateOrganizationalUnitInput
	Copy  func(*UpdateOrganizationalUnitInput) UpdateOrganizationalUnitRequest
}

// Send marshals and sends the UpdateOrganizationalUnit API request.
func (r UpdateOrganizationalUnitRequest) Send(ctx context.Context) (*UpdateOrganizationalUnitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateOrganizationalUnitResponse{
		UpdateOrganizationalUnitOutput: r.Request.Data.(*UpdateOrganizationalUnitOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateOrganizationalUnitResponse is the response type for the
// UpdateOrganizationalUnit API operation.
type UpdateOrganizationalUnitResponse struct {
	*UpdateOrganizationalUnitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateOrganizationalUnit request.
func (r *UpdateOrganizationalUnitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
