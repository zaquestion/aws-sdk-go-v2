// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateFunctionInput struct {
	_ struct{} `type:"structure"`

	// The GraphQL API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The Function DataSource name.
	//
	// DataSourceName is a required field
	DataSourceName *string `locationName:"dataSourceName" type:"string" required:"true"`

	// The Function description.
	Description *string `locationName:"description" type:"string"`

	// The function ID.
	//
	// FunctionId is a required field
	FunctionId *string `location:"uri" locationName:"functionId" type:"string" required:"true"`

	// The version of the request mapping template. Currently the supported value
	// is 2018-05-29.
	//
	// FunctionVersion is a required field
	FunctionVersion *string `locationName:"functionVersion" type:"string" required:"true"`

	// The Function name.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The Function request mapping template. Functions support only the 2018-05-29
	// version of the request mapping template.
	//
	// RequestMappingTemplate is a required field
	RequestMappingTemplate *string `locationName:"requestMappingTemplate" min:"1" type:"string" required:"true"`

	// The Function request mapping template.
	ResponseMappingTemplate *string `locationName:"responseMappingTemplate" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateFunctionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFunctionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFunctionInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.DataSourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceName"))
	}

	if s.FunctionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionId"))
	}

	if s.FunctionVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionVersion"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RequestMappingTemplate == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestMappingTemplate"))
	}
	if s.RequestMappingTemplate != nil && len(*s.RequestMappingTemplate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RequestMappingTemplate", 1))
	}
	if s.ResponseMappingTemplate != nil && len(*s.ResponseMappingTemplate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResponseMappingTemplate", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFunctionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DataSourceName != nil {
		v := *s.DataSourceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dataSourceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionVersion != nil {
		v := *s.FunctionVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "functionVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestMappingTemplate != nil {
		v := *s.RequestMappingTemplate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestMappingTemplate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResponseMappingTemplate != nil {
		v := *s.ResponseMappingTemplate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "responseMappingTemplate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionId != nil {
		v := *s.FunctionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "functionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateFunctionOutput struct {
	_ struct{} `type:"structure"`

	// The Function object.
	FunctionConfiguration *FunctionConfiguration `locationName:"functionConfiguration" type:"structure"`
}

// String returns the string representation
func (s UpdateFunctionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFunctionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FunctionConfiguration != nil {
		v := s.FunctionConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "functionConfiguration", v, metadata)
	}
	return nil
}

const opUpdateFunction = "UpdateFunction"

// UpdateFunctionRequest returns a request value for making API operation for
// AWS AppSync.
//
// Updates a Function object.
//
//    // Example sending a request using UpdateFunctionRequest.
//    req := client.UpdateFunctionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateFunction
func (c *Client) UpdateFunctionRequest(input *UpdateFunctionInput) UpdateFunctionRequest {
	op := &aws.Operation{
		Name:       opUpdateFunction,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/functions/{functionId}",
	}

	if input == nil {
		input = &UpdateFunctionInput{}
	}

	req := c.newRequest(op, input, &UpdateFunctionOutput{})
	return UpdateFunctionRequest{Request: req, Input: input, Copy: c.UpdateFunctionRequest}
}

// UpdateFunctionRequest is the request type for the
// UpdateFunction API operation.
type UpdateFunctionRequest struct {
	*aws.Request
	Input *UpdateFunctionInput
	Copy  func(*UpdateFunctionInput) UpdateFunctionRequest
}

// Send marshals and sends the UpdateFunction API request.
func (r UpdateFunctionRequest) Send(ctx context.Context) (*UpdateFunctionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFunctionResponse{
		UpdateFunctionOutput: r.Request.Data.(*UpdateFunctionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFunctionResponse is the response type for the
// UpdateFunction API operation.
type UpdateFunctionResponse struct {
	*UpdateFunctionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFunction request.
func (r *UpdateFunctionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
