// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteConferenceProviderInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the conference provider.
	//
	// ConferenceProviderArn is a required field
	ConferenceProviderArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConferenceProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConferenceProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConferenceProviderInput"}

	if s.ConferenceProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConferenceProviderArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteConferenceProviderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConferenceProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConferenceProvider = "DeleteConferenceProvider"

// DeleteConferenceProviderRequest returns a request value for making API operation for
// Alexa For Business.
//
// Deletes a conference provider.
//
//    // Example sending a request using DeleteConferenceProviderRequest.
//    req := client.DeleteConferenceProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteConferenceProvider
func (c *Client) DeleteConferenceProviderRequest(input *DeleteConferenceProviderInput) DeleteConferenceProviderRequest {
	op := &aws.Operation{
		Name:       opDeleteConferenceProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConferenceProviderInput{}
	}

	req := c.newRequest(op, input, &DeleteConferenceProviderOutput{})
	return DeleteConferenceProviderRequest{Request: req, Input: input, Copy: c.DeleteConferenceProviderRequest}
}

// DeleteConferenceProviderRequest is the request type for the
// DeleteConferenceProvider API operation.
type DeleteConferenceProviderRequest struct {
	*aws.Request
	Input *DeleteConferenceProviderInput
	Copy  func(*DeleteConferenceProviderInput) DeleteConferenceProviderRequest
}

// Send marshals and sends the DeleteConferenceProvider API request.
func (r DeleteConferenceProviderRequest) Send(ctx context.Context) (*DeleteConferenceProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConferenceProviderResponse{
		DeleteConferenceProviderOutput: r.Request.Data.(*DeleteConferenceProviderOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConferenceProviderResponse is the response type for the
// DeleteConferenceProvider API operation.
type DeleteConferenceProviderResponse struct {
	*DeleteConferenceProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConferenceProvider request.
func (r *DeleteConferenceProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
