// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateSimulationJobInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string" idempotencyToken:"true"`

	// The data sources for the simulation job.
	//
	// There is a limit of 100 files and a combined size of 25GB for all DataSourceConfig
	// objects.
	DataSources []DataSourceConfig `locationName:"dataSources" min:"1" type:"list"`

	// The failure behavior the simulation job.
	//
	// Continue
	//
	// Restart the simulation job in the same host instance.
	//
	// Fail
	//
	// Stop the simulation job and terminate the instance.
	FailureBehavior FailureBehavior `locationName:"failureBehavior" type:"string" enum:"true"`

	// The IAM role name that allows the simulation instance to call the AWS APIs
	// that are specified in its associated policies on your behalf. This is how
	// credentials are passed in to your simulation job.
	//
	// IamRole is a required field
	IamRole *string `locationName:"iamRole" min:"1" type:"string" required:"true"`

	// The logging configuration.
	LoggingConfig *LoggingConfig `locationName:"loggingConfig" type:"structure"`

	// The maximum simulation job duration in seconds (up to 14 days or 1,209,600
	// seconds. When maxJobDurationInSeconds is reached, the simulation job will
	// status will transition to Completed.
	//
	// MaxJobDurationInSeconds is a required field
	MaxJobDurationInSeconds *int64 `locationName:"maxJobDurationInSeconds" type:"long" required:"true"`

	// Location for output files generated by the simulation job.
	OutputLocation *OutputLocation `locationName:"outputLocation" type:"structure"`

	// The robot application to use in the simulation job.
	RobotApplications []RobotApplicationConfig `locationName:"robotApplications" min:"1" type:"list"`

	// The simulation application to use in the simulation job.
	SimulationApplications []SimulationApplicationConfig `locationName:"simulationApplications" min:"1" type:"list"`

	// A map that contains tag keys and tag values that are attached to the simulation
	// job.
	Tags map[string]string `locationName:"tags" type:"map"`

	// If your simulation job accesses resources in a VPC, you provide this parameter
	// identifying the list of security group IDs and subnet IDs. These must belong
	// to the same VPC. You must provide at least one security group and one subnet
	// ID.
	VpcConfig *VPCConfig `locationName:"vpcConfig" type:"structure"`
}

// String returns the string representation
func (s CreateSimulationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSimulationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSimulationJobInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.DataSources != nil && len(s.DataSources) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataSources", 1))
	}

	if s.IamRole == nil {
		invalidParams.Add(aws.NewErrParamRequired("IamRole"))
	}
	if s.IamRole != nil && len(*s.IamRole) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IamRole", 1))
	}

	if s.MaxJobDurationInSeconds == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaxJobDurationInSeconds"))
	}
	if s.RobotApplications != nil && len(s.RobotApplications) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RobotApplications", 1))
	}
	if s.SimulationApplications != nil && len(s.SimulationApplications) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SimulationApplications", 1))
	}
	if s.DataSources != nil {
		for i, v := range s.DataSources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DataSources", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.LoggingConfig != nil {
		if err := s.LoggingConfig.Validate(); err != nil {
			invalidParams.AddNested("LoggingConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputLocation != nil {
		if err := s.OutputLocation.Validate(); err != nil {
			invalidParams.AddNested("OutputLocation", err.(aws.ErrInvalidParams))
		}
	}
	if s.RobotApplications != nil {
		for i, v := range s.RobotApplications {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RobotApplications", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SimulationApplications != nil {
		for i, v := range s.SimulationApplications {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SimulationApplications", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			invalidParams.AddNested("VpcConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSimulationJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSources != nil {
		v := s.DataSources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dataSources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.FailureBehavior) > 0 {
		v := s.FailureBehavior

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureBehavior", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.IamRole != nil {
		v := *s.IamRole

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "iamRole", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LoggingConfig != nil {
		v := s.LoggingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "loggingConfig", v, metadata)
	}
	if s.MaxJobDurationInSeconds != nil {
		v := *s.MaxJobDurationInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxJobDurationInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.OutputLocation != nil {
		v := s.OutputLocation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "outputLocation", v, metadata)
	}
	if s.RobotApplications != nil {
		v := s.RobotApplications

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "robotApplications", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SimulationApplications != nil {
		v := s.SimulationApplications

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "simulationApplications", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

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
	if s.VpcConfig != nil {
		v := s.VpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "vpcConfig", v, metadata)
	}
	return nil
}

type CreateSimulationJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the simulation job.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string"`

	// The data sources for the simulation job.
	DataSources []DataSource `locationName:"dataSources" type:"list"`

	// the failure behavior for the simulation job.
	FailureBehavior FailureBehavior `locationName:"failureBehavior" type:"string" enum:"true"`

	// The failure code of the simulation job if it failed:
	//
	// InternalServiceError
	//
	// Internal service error.
	//
	// RobotApplicationCrash
	//
	// Robot application exited abnormally.
	//
	// SimulationApplicationCrash
	//
	// Simulation application exited abnormally.
	//
	// BadPermissionsRobotApplication
	//
	// Robot application bundle could not be downloaded.
	//
	// BadPermissionsSimulationApplication
	//
	// Simulation application bundle could not be downloaded.
	//
	// BadPermissionsS3Output
	//
	// Unable to publish outputs to customer-provided S3 bucket.
	//
	// BadPermissionsCloudwatchLogs
	//
	// Unable to publish logs to customer-provided CloudWatch Logs resource.
	//
	// SubnetIpLimitExceeded
	//
	// Subnet IP limit exceeded.
	//
	// ENILimitExceeded
	//
	// ENI limit exceeded.
	//
	// BadPermissionsUserCredentials
	//
	// Unable to use the Role provided.
	//
	// InvalidBundleRobotApplication
	//
	// Robot bundle cannot be extracted (invalid format, bundling error, or other
	// issue).
	//
	// InvalidBundleSimulationApplication
	//
	// Simulation bundle cannot be extracted (invalid format, bundling error, or
	// other issue).
	//
	// RobotApplicationVersionMismatchedEtag
	//
	// Etag for RobotApplication does not match value during version creation.
	//
	// SimulationApplicationVersionMismatchedEtag
	//
	// Etag for SimulationApplication does not match value during version creation.
	//
	// WrongRegionS3Output
	//
	// S3 output bucket is in a different region than AWS RoboMaker.
	//
	// WrongRegionRobotApplication
	//
	// RobotApplication bucket is in a different region than AWS RoboMaker.
	//
	// WrongRegionSimulationApplication
	//
	// SimulationApplication bucket is in a different region than AWS RoboMaker.
	FailureCode SimulationJobErrorCode `locationName:"failureCode" type:"string" enum:"true"`

	// The IAM role that allows the simulation job to call the AWS APIs that are
	// specified in its associated policies on your behalf.
	IamRole *string `locationName:"iamRole" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the simulation job was last
	// started.
	LastStartedAt *time.Time `locationName:"lastStartedAt" type:"timestamp"`

	// The time, in milliseconds since the epoch, when the simulation job was last
	// updated.
	LastUpdatedAt *time.Time `locationName:"lastUpdatedAt" type:"timestamp"`

	// The logging configuration.
	LoggingConfig *LoggingConfig `locationName:"loggingConfig" type:"structure"`

	// The maximum simulation job duration in seconds.
	MaxJobDurationInSeconds *int64 `locationName:"maxJobDurationInSeconds" type:"long"`

	// Simulation job output files location.
	OutputLocation *OutputLocation `locationName:"outputLocation" type:"structure"`

	// The robot application used by the simulation job.
	RobotApplications []RobotApplicationConfig `locationName:"robotApplications" min:"1" type:"list"`

	// The simulation application used by the simulation job.
	SimulationApplications []SimulationApplicationConfig `locationName:"simulationApplications" min:"1" type:"list"`

	// The simulation job execution duration in milliseconds.
	SimulationTimeMillis *int64 `locationName:"simulationTimeMillis" type:"long"`

	// The status of the simulation job.
	Status SimulationJobStatus `locationName:"status" type:"string" enum:"true"`

	// The list of all tags added to the simulation job.
	Tags map[string]string `locationName:"tags" type:"map"`

	// Information about the vpc configuration.
	VpcConfig *VPCConfigResponse `locationName:"vpcConfig" type:"structure"`
}

// String returns the string representation
func (s CreateSimulationJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSimulationJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientRequestToken != nil {
		v := *s.ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSources != nil {
		v := s.DataSources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dataSources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.FailureBehavior) > 0 {
		v := s.FailureBehavior

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureBehavior", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.FailureCode) > 0 {
		v := s.FailureCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.IamRole != nil {
		v := *s.IamRole

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "iamRole", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastStartedAt != nil {
		v := *s.LastStartedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastStartedAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LoggingConfig != nil {
		v := s.LoggingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "loggingConfig", v, metadata)
	}
	if s.MaxJobDurationInSeconds != nil {
		v := *s.MaxJobDurationInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxJobDurationInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.OutputLocation != nil {
		v := s.OutputLocation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "outputLocation", v, metadata)
	}
	if s.RobotApplications != nil {
		v := s.RobotApplications

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "robotApplications", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SimulationApplications != nil {
		v := s.SimulationApplications

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "simulationApplications", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SimulationTimeMillis != nil {
		v := *s.SimulationTimeMillis

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "simulationTimeMillis", protocol.Int64Value(v), metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
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
	if s.VpcConfig != nil {
		v := s.VpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "vpcConfig", v, metadata)
	}
	return nil
}

const opCreateSimulationJob = "CreateSimulationJob"

// CreateSimulationJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Creates a simulation job.
//
// After 90 days, simulation jobs expire and will be deleted. They will no longer
// be accessible.
//
//    // Example sending a request using CreateSimulationJobRequest.
//    req := client.CreateSimulationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateSimulationJob
func (c *Client) CreateSimulationJobRequest(input *CreateSimulationJobInput) CreateSimulationJobRequest {
	op := &aws.Operation{
		Name:       opCreateSimulationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/createSimulationJob",
	}

	if input == nil {
		input = &CreateSimulationJobInput{}
	}

	req := c.newRequest(op, input, &CreateSimulationJobOutput{})
	return CreateSimulationJobRequest{Request: req, Input: input, Copy: c.CreateSimulationJobRequest}
}

// CreateSimulationJobRequest is the request type for the
// CreateSimulationJob API operation.
type CreateSimulationJobRequest struct {
	*aws.Request
	Input *CreateSimulationJobInput
	Copy  func(*CreateSimulationJobInput) CreateSimulationJobRequest
}

// Send marshals and sends the CreateSimulationJob API request.
func (r CreateSimulationJobRequest) Send(ctx context.Context) (*CreateSimulationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSimulationJobResponse{
		CreateSimulationJobOutput: r.Request.Data.(*CreateSimulationJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSimulationJobResponse is the response type for the
// CreateSimulationJob API operation.
type CreateSimulationJobResponse struct {
	*CreateSimulationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSimulationJob request.
func (r *CreateSimulationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
