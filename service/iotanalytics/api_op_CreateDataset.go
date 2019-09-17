// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateDatasetInput struct {
	_ struct{} `type:"structure"`

	// A list of actions that create the data set contents.
	//
	// Actions is a required field
	Actions []DatasetAction `locationName:"actions" min:"1" type:"list" required:"true"`

	// When data set contents are created they are delivered to destinations specified
	// here.
	ContentDeliveryRules []DatasetContentDeliveryRule `locationName:"contentDeliveryRules" type:"list"`

	// The name of the data set.
	//
	// DatasetName is a required field
	DatasetName *string `locationName:"datasetName" min:"1" type:"string" required:"true"`

	// [Optional] How long, in days, versions of data set contents are kept for
	// the data set. If not specified or set to null, versions of data set contents
	// are retained for at most 90 days. The number of versions of data set contents
	// retained is determined by the versioningConfiguration parameter. (For more
	// information, see https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions)
	RetentionPeriod *RetentionPeriod `locationName:"retentionPeriod" type:"structure"`

	// Metadata which can be used to manage the data set.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`

	// A list of triggers. A trigger causes data set contents to be populated at
	// a specified time interval or when another data set's contents are created.
	// The list of triggers can be empty or contain up to five DataSetTrigger objects.
	Triggers []DatasetTrigger `locationName:"triggers" type:"list"`

	// [Optional] How many versions of data set contents are kept. If not specified
	// or set to null, only the latest version plus the latest succeeded version
	// (if they are different) are kept for the time period specified by the "retentionPeriod"
	// parameter. (For more information, see https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions)
	VersioningConfiguration *VersioningConfiguration `locationName:"versioningConfiguration" type:"structure"`
}

// String returns the string representation
func (s CreateDatasetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatasetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDatasetInput"}

	if s.Actions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Actions"))
	}
	if s.Actions != nil && len(s.Actions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Actions", 1))
	}

	if s.DatasetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetName"))
	}
	if s.DatasetName != nil && len(*s.DatasetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Actions != nil {
		for i, v := range s.Actions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Actions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ContentDeliveryRules != nil {
		for i, v := range s.ContentDeliveryRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ContentDeliveryRules", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RetentionPeriod != nil {
		if err := s.RetentionPeriod.Validate(); err != nil {
			invalidParams.AddNested("RetentionPeriod", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Triggers != nil {
		for i, v := range s.Triggers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Triggers", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VersioningConfiguration != nil {
		if err := s.VersioningConfiguration.Validate(); err != nil {
			invalidParams.AddNested("VersioningConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDatasetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Actions != nil {
		v := s.Actions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "actions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ContentDeliveryRules != nil {
		v := s.ContentDeliveryRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "contentDeliveryRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DatasetName != nil {
		v := *s.DatasetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "datasetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RetentionPeriod != nil {
		v := s.RetentionPeriod

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "retentionPeriod", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Triggers != nil {
		v := s.Triggers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "triggers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VersioningConfiguration != nil {
		v := s.VersioningConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "versioningConfiguration", v, metadata)
	}
	return nil
}

type CreateDatasetOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the data set.
	DatasetArn *string `locationName:"datasetArn" type:"string"`

	// The name of the data set.
	DatasetName *string `locationName:"datasetName" min:"1" type:"string"`

	// How long, in days, data set contents are kept for the data set.
	RetentionPeriod *RetentionPeriod `locationName:"retentionPeriod" type:"structure"`
}

// String returns the string representation
func (s CreateDatasetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDatasetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DatasetArn != nil {
		v := *s.DatasetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "datasetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DatasetName != nil {
		v := *s.DatasetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "datasetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RetentionPeriod != nil {
		v := s.RetentionPeriod

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "retentionPeriod", v, metadata)
	}
	return nil
}

const opCreateDataset = "CreateDataset"

// CreateDatasetRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Creates a data set. A data set stores data retrieved from a data store by
// applying a "queryAction" (a SQL query) or a "containerAction" (executing
// a containerized application). This operation creates the skeleton of a data
// set. The data set can be populated manually by calling "CreateDatasetContent"
// or automatically according to a "trigger" you specify.
//
//    // Example sending a request using CreateDatasetRequest.
//    req := client.CreateDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/CreateDataset
func (c *Client) CreateDatasetRequest(input *CreateDatasetInput) CreateDatasetRequest {
	op := &aws.Operation{
		Name:       opCreateDataset,
		HTTPMethod: "POST",
		HTTPPath:   "/datasets",
	}

	if input == nil {
		input = &CreateDatasetInput{}
	}

	req := c.newRequest(op, input, &CreateDatasetOutput{})
	return CreateDatasetRequest{Request: req, Input: input, Copy: c.CreateDatasetRequest}
}

// CreateDatasetRequest is the request type for the
// CreateDataset API operation.
type CreateDatasetRequest struct {
	*aws.Request
	Input *CreateDatasetInput
	Copy  func(*CreateDatasetInput) CreateDatasetRequest
}

// Send marshals and sends the CreateDataset API request.
func (r CreateDatasetRequest) Send(ctx context.Context) (*CreateDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatasetResponse{
		CreateDatasetOutput: r.Request.Data.(*CreateDatasetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatasetResponse is the response type for the
// CreateDataset API operation.
type CreateDatasetResponse struct {
	*CreateDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataset request.
func (r *CreateDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
