// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetAccountInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccountInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccountInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetAccountOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account details.
	Account *Account `type:"structure"`
}

// String returns the string representation
func (s GetAccountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccountOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Account != nil {
		v := s.Account

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Account", v, metadata)
	}
	return nil
}

const opGetAccount = "GetAccount"

// GetAccountRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves details for the specified Amazon Chime account, such as account
// type and supported licenses.
//
//    // Example sending a request using GetAccountRequest.
//    req := client.GetAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetAccount
func (c *Client) GetAccountRequest(input *GetAccountInput) GetAccountRequest {
	op := &aws.Operation{
		Name:       opGetAccount,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{accountId}",
	}

	if input == nil {
		input = &GetAccountInput{}
	}

	req := c.newRequest(op, input, &GetAccountOutput{})
	return GetAccountRequest{Request: req, Input: input, Copy: c.GetAccountRequest}
}

// GetAccountRequest is the request type for the
// GetAccount API operation.
type GetAccountRequest struct {
	*aws.Request
	Input *GetAccountInput
	Copy  func(*GetAccountInput) GetAccountRequest
}

// Send marshals and sends the GetAccount API request.
func (r GetAccountRequest) Send(ctx context.Context) (*GetAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountResponse{
		GetAccountOutput: r.Request.Data.(*GetAccountOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountResponse is the response type for the
// GetAccount API operation.
type GetAccountResponse struct {
	*GetAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccount request.
func (r *GetAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
