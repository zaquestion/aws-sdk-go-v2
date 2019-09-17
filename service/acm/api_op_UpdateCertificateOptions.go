// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type UpdateCertificateOptionsInput struct {
	_ struct{} `type:"structure"`

	// ARN of the requested certificate to update. This must be of the form:
	//
	// arn:aws:acm:us-east-1:account:certificate/12345678-1234-1234-1234-123456789012
	//
	// CertificateArn is a required field
	CertificateArn *string `min:"20" type:"string" required:"true"`

	// Use to update the options for your certificate. Currently, you can specify
	// whether to add your certificate to a transparency log. Certificate transparency
	// makes it possible to detect SSL/TLS certificates that have been mistakenly
	// or maliciously issued. Certificates that have not been logged typically produce
	// an error message in a browser.
	//
	// Options is a required field
	Options *CertificateOptions `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateCertificateOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCertificateOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCertificateOptionsInput"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}
	if s.CertificateArn != nil && len(*s.CertificateArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateArn", 20))
	}

	if s.Options == nil {
		invalidParams.Add(aws.NewErrParamRequired("Options"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateCertificateOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateCertificateOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateCertificateOptions = "UpdateCertificateOptions"

// UpdateCertificateOptionsRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Updates a certificate. Currently, you can use this function to specify whether
// to opt in to or out of recording your certificate in a certificate transparency
// log. For more information, see Opting Out of Certificate Transparency Logging
// (https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency).
//
//    // Example sending a request using UpdateCertificateOptionsRequest.
//    req := client.UpdateCertificateOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/UpdateCertificateOptions
func (c *Client) UpdateCertificateOptionsRequest(input *UpdateCertificateOptionsInput) UpdateCertificateOptionsRequest {
	op := &aws.Operation{
		Name:       opUpdateCertificateOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateCertificateOptionsInput{}
	}

	req := c.newRequest(op, input, &UpdateCertificateOptionsOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateCertificateOptionsRequest{Request: req, Input: input, Copy: c.UpdateCertificateOptionsRequest}
}

// UpdateCertificateOptionsRequest is the request type for the
// UpdateCertificateOptions API operation.
type UpdateCertificateOptionsRequest struct {
	*aws.Request
	Input *UpdateCertificateOptionsInput
	Copy  func(*UpdateCertificateOptionsInput) UpdateCertificateOptionsRequest
}

// Send marshals and sends the UpdateCertificateOptions API request.
func (r UpdateCertificateOptionsRequest) Send(ctx context.Context) (*UpdateCertificateOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCertificateOptionsResponse{
		UpdateCertificateOptionsOutput: r.Request.Data.(*UpdateCertificateOptionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCertificateOptionsResponse is the response type for the
// UpdateCertificateOptions API operation.
type UpdateCertificateOptionsResponse struct {
	*UpdateCertificateOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCertificateOptions request.
func (r *UpdateCertificateOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
