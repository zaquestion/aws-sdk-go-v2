// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutLexiconInput struct {
	_ struct{} `type:"structure"`

	// Content of the PLS lexicon as string data.
	//
	// Content is a required field
	Content *string `type:"string" required:"true"`

	// Name of the lexicon. The name must follow the regular express format [0-9A-Za-z]{1,20}.
	// That is, the name is a case-sensitive alphanumeric string up to 20 characters
	// long.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"LexiconName" type:"string" required:"true"`
}

// String returns the string representation
func (s PutLexiconInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLexiconInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLexiconInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutLexiconInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Content != nil {
		v := *s.Content

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Content", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LexiconName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutLexiconOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutLexiconOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutLexiconOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutLexicon = "PutLexicon"

// PutLexiconRequest returns a request value for making API operation for
// Amazon Polly.
//
// Stores a pronunciation lexicon in an AWS Region. If a lexicon with the same
// name already exists in the region, it is overwritten by the new lexicon.
// Lexicon operations have eventual consistency, therefore, it might take some
// time before the lexicon is available to the SynthesizeSpeech operation.
//
// For more information, see Managing Lexicons (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
//
//    // Example sending a request using PutLexiconRequest.
//    req := client.PutLexiconRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/PutLexicon
func (c *Client) PutLexiconRequest(input *PutLexiconInput) PutLexiconRequest {
	op := &aws.Operation{
		Name:       opPutLexicon,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/lexicons/{LexiconName}",
	}

	if input == nil {
		input = &PutLexiconInput{}
	}

	req := c.newRequest(op, input, &PutLexiconOutput{})
	return PutLexiconRequest{Request: req, Input: input, Copy: c.PutLexiconRequest}
}

// PutLexiconRequest is the request type for the
// PutLexicon API operation.
type PutLexiconRequest struct {
	*aws.Request
	Input *PutLexiconInput
	Copy  func(*PutLexiconInput) PutLexiconRequest
}

// Send marshals and sends the PutLexicon API request.
func (r PutLexiconRequest) Send(ctx context.Context) (*PutLexiconResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLexiconResponse{
		PutLexiconOutput: r.Request.Data.(*PutLexiconOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLexiconResponse is the response type for the
// PutLexicon API operation.
type PutLexiconResponse struct {
	*PutLexiconOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLexicon request.
func (r *PutLexiconResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
