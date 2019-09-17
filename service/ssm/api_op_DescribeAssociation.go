// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAssociationInput struct {
	_ struct{} `type:"structure"`

	// The association ID for which you want information.
	AssociationId *string `type:"string"`

	// Specify the association version to retrieve. To view the latest version,
	// either specify $LATEST for this parameter, or omit this parameter. To view
	// a list of all associations for an instance, use ListAssociations. To get
	// a list of versions for a specific association, use ListAssociationVersions.
	AssociationVersion *string `type:"string"`

	// The instance ID.
	InstanceId *string `type:"string"`

	// The name of the Systems Manager document.
	Name *string `type:"string"`
}

// String returns the string representation
func (s DescribeAssociationInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the association.
	AssociationDescription *AssociationDescription `type:"structure"`
}

// String returns the string representation
func (s DescribeAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAssociation = "DescribeAssociation"

// DescribeAssociationRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Describes the association for the specified target or instance. If you created
// the association by using the Targets parameter, then you must retrieve the
// association by using the association ID. If you created the association by
// specifying an instance ID and a Systems Manager document, then you retrieve
// the association by specifying the document name and the instance ID.
//
//    // Example sending a request using DescribeAssociationRequest.
//    req := client.DescribeAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeAssociation
func (c *Client) DescribeAssociationRequest(input *DescribeAssociationInput) DescribeAssociationRequest {
	op := &aws.Operation{
		Name:       opDescribeAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAssociationInput{}
	}

	req := c.newRequest(op, input, &DescribeAssociationOutput{})
	return DescribeAssociationRequest{Request: req, Input: input, Copy: c.DescribeAssociationRequest}
}

// DescribeAssociationRequest is the request type for the
// DescribeAssociation API operation.
type DescribeAssociationRequest struct {
	*aws.Request
	Input *DescribeAssociationInput
	Copy  func(*DescribeAssociationInput) DescribeAssociationRequest
}

// Send marshals and sends the DescribeAssociation API request.
func (r DescribeAssociationRequest) Send(ctx context.Context) (*DescribeAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAssociationResponse{
		DescribeAssociationOutput: r.Request.Data.(*DescribeAssociationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAssociationResponse is the response type for the
// DescribeAssociation API operation.
type DescribeAssociationResponse struct {
	*DescribeAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAssociation request.
func (r *DescribeAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
