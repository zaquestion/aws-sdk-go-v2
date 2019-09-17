// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateFolderInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// The name of the new folder.
	Name *string `min:"1" type:"string"`

	// The ID of the parent folder.
	//
	// ParentFolderId is a required field
	ParentFolderId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateFolderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFolderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFolderInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.ParentFolderId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentFolderId"))
	}
	if s.ParentFolderId != nil && len(*s.ParentFolderId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ParentFolderId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFolderInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParentFolderId != nil {
		v := *s.ParentFolderId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ParentFolderId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateFolderOutput struct {
	_ struct{} `type:"structure"`

	// The metadata of the folder.
	Metadata *FolderMetadata `type:"structure"`
}

// String returns the string representation
func (s CreateFolderOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFolderOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Metadata", v, metadata)
	}
	return nil
}

const opCreateFolder = "CreateFolder"

// CreateFolderRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Creates a folder with the specified name and parent folder.
//
//    // Example sending a request using CreateFolderRequest.
//    req := client.CreateFolderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateFolder
func (c *Client) CreateFolderRequest(input *CreateFolderInput) CreateFolderRequest {
	op := &aws.Operation{
		Name:       opCreateFolder,
		HTTPMethod: "POST",
		HTTPPath:   "/api/v1/folders",
	}

	if input == nil {
		input = &CreateFolderInput{}
	}

	req := c.newRequest(op, input, &CreateFolderOutput{})
	return CreateFolderRequest{Request: req, Input: input, Copy: c.CreateFolderRequest}
}

// CreateFolderRequest is the request type for the
// CreateFolder API operation.
type CreateFolderRequest struct {
	*aws.Request
	Input *CreateFolderInput
	Copy  func(*CreateFolderInput) CreateFolderRequest
}

// Send marshals and sends the CreateFolder API request.
func (r CreateFolderRequest) Send(ctx context.Context) (*CreateFolderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFolderResponse{
		CreateFolderOutput: r.Request.Data.(*CreateFolderOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFolderResponse is the response type for the
// CreateFolder API operation.
type CreateFolderResponse struct {
	*CreateFolderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFolder request.
func (r *CreateFolderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
