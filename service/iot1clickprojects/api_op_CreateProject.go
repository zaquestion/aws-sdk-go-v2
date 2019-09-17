// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateProjectInput struct {
	_ struct{} `type:"structure"`

	// An optional description for the project.
	Description *string `locationName:"description" type:"string"`

	// The schema defining the placement to be created. A placement template defines
	// placement default attributes and device templates. You cannot add or remove
	// device templates after the project has been created. However, you can update
	// callbackOverrides for the device templates using the UpdateProject API.
	PlacementTemplate *PlacementTemplate `locationName:"placementTemplate" type:"structure"`

	// The name of the project to create.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"1" type:"string" required:"true"`

	// Optional tags (metadata key/value pairs) to be associated with the project.
	// For example, { {"key1": "value1", "key2": "value2"} }. For more information,
	// see AWS Tagging Strategies (https://aws.amazon.com/answers/account-management/aws-tagging-strategies/).
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProjectInput"}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlacementTemplate != nil {
		v := s.PlacementTemplate

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "placementTemplate", v, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateProjectOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateProject = "CreateProject"

// CreateProjectRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Creates an empty project with a placement template. A project contains zero
// or more placements that adhere to the placement template defined in the project.
//
//    // Example sending a request using CreateProjectRequest.
//    req := client.CreateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/CreateProject
func (c *Client) CreateProjectRequest(input *CreateProjectInput) CreateProjectRequest {
	op := &aws.Operation{
		Name:       opCreateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/projects",
	}

	if input == nil {
		input = &CreateProjectInput{}
	}

	req := c.newRequest(op, input, &CreateProjectOutput{})
	return CreateProjectRequest{Request: req, Input: input, Copy: c.CreateProjectRequest}
}

// CreateProjectRequest is the request type for the
// CreateProject API operation.
type CreateProjectRequest struct {
	*aws.Request
	Input *CreateProjectInput
	Copy  func(*CreateProjectInput) CreateProjectRequest
}

// Send marshals and sends the CreateProject API request.
func (r CreateProjectRequest) Send(ctx context.Context) (*CreateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProjectResponse{
		CreateProjectOutput: r.Request.Data.(*CreateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProjectResponse is the response type for the
// CreateProject API operation.
type CreateProjectResponse struct {
	*CreateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProject request.
func (r *CreateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
