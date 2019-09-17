// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The API ID for the GraphQL API for the DataSource.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// A description of the DataSource.
	Description *string `locationName:"description" type:"string"`

	// Amazon DynamoDB settings.
	DynamodbConfig *DynamodbDataSourceConfig `locationName:"dynamodbConfig" type:"structure"`

	// Amazon Elasticsearch Service settings.
	ElasticsearchConfig *ElasticsearchDataSourceConfig `locationName:"elasticsearchConfig" type:"structure"`

	// HTTP endpoint settings.
	HttpConfig *HttpDataSourceConfig `locationName:"httpConfig" type:"structure"`

	// AWS Lambda settings.
	LambdaConfig *LambdaDataSourceConfig `locationName:"lambdaConfig" type:"structure"`

	// A user-supplied name for the DataSource.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// Relational database settings.
	RelationalDatabaseConfig *RelationalDatabaseDataSourceConfig `locationName:"relationalDatabaseConfig" type:"structure"`

	// The AWS IAM service role ARN for the data source. The system assumes this
	// role when accessing the data source.
	ServiceRoleArn *string `locationName:"serviceRoleArn" type:"string"`

	// The type of the DataSource.
	//
	// Type is a required field
	Type DataSourceType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataSourceInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}
	if s.DynamodbConfig != nil {
		if err := s.DynamodbConfig.Validate(); err != nil {
			invalidParams.AddNested("DynamodbConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.ElasticsearchConfig != nil {
		if err := s.ElasticsearchConfig.Validate(); err != nil {
			invalidParams.AddNested("ElasticsearchConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.HttpConfig != nil {
		if err := s.HttpConfig.Validate(); err != nil {
			invalidParams.AddNested("HttpConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.LambdaConfig != nil {
		if err := s.LambdaConfig.Validate(); err != nil {
			invalidParams.AddNested("LambdaConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDataSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DynamodbConfig != nil {
		v := s.DynamodbConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dynamodbConfig", v, metadata)
	}
	if s.ElasticsearchConfig != nil {
		v := s.ElasticsearchConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "elasticsearchConfig", v, metadata)
	}
	if s.HttpConfig != nil {
		v := s.HttpConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "httpConfig", v, metadata)
	}
	if s.LambdaConfig != nil {
		v := s.LambdaConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "lambdaConfig", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RelationalDatabaseConfig != nil {
		v := s.RelationalDatabaseConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "relationalDatabaseConfig", v, metadata)
	}
	if s.ServiceRoleArn != nil {
		v := *s.ServiceRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "serviceRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The DataSource object.
	DataSource *DataSource `locationName:"dataSource" type:"structure"`
}

// String returns the string representation
func (s CreateDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDataSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataSource != nil {
		v := s.DataSource

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dataSource", v, metadata)
	}
	return nil
}

const opCreateDataSource = "CreateDataSource"

// CreateDataSourceRequest returns a request value for making API operation for
// AWS AppSync.
//
// Creates a DataSource object.
//
//    // Example sending a request using CreateDataSourceRequest.
//    req := client.CreateDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateDataSource
func (c *Client) CreateDataSourceRequest(input *CreateDataSourceInput) CreateDataSourceRequest {
	op := &aws.Operation{
		Name:       opCreateDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/datasources",
	}

	if input == nil {
		input = &CreateDataSourceInput{}
	}

	req := c.newRequest(op, input, &CreateDataSourceOutput{})
	return CreateDataSourceRequest{Request: req, Input: input, Copy: c.CreateDataSourceRequest}
}

// CreateDataSourceRequest is the request type for the
// CreateDataSource API operation.
type CreateDataSourceRequest struct {
	*aws.Request
	Input *CreateDataSourceInput
	Copy  func(*CreateDataSourceInput) CreateDataSourceRequest
}

// Send marshals and sends the CreateDataSource API request.
func (r CreateDataSourceRequest) Send(ctx context.Context) (*CreateDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDataSourceResponse{
		CreateDataSourceOutput: r.Request.Data.(*CreateDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDataSourceResponse is the response type for the
// CreateDataSource API operation.
type CreateDataSourceResponse struct {
	*CreateDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataSource request.
func (r *CreateDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
