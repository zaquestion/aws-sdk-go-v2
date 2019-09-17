// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRelationalDatabaseMetricDataInput struct {
	_ struct{} `type:"structure"`

	// The end of the time interval from which to get metric data.
	//
	// Constraints:
	//
	//    * Specified in Universal Coordinated Time (UTC).
	//
	//    * Specified in the Unix time format. For example, if you wish to use an
	//    end time of October 1, 2018, at 8 PM UTC, then you input 1538424000 as
	//    the end time.
	//
	// EndTime is a required field
	EndTime *time.Time `locationName:"endTime" type:"timestamp" required:"true"`

	// The name of the metric data to return.
	//
	// MetricName is a required field
	MetricName RelationalDatabaseMetricName `locationName:"metricName" type:"string" required:"true" enum:"true"`

	// The granularity, in seconds, of the returned data points.
	//
	// Period is a required field
	Period *int64 `locationName:"period" min:"60" type:"integer" required:"true"`

	// The name of your database from which to get metric data.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`

	// The start of the time interval from which to get metric data.
	//
	// Constraints:
	//
	//    * Specified in Universal Coordinated Time (UTC).
	//
	//    * Specified in the Unix time format. For example, if you wish to use a
	//    start time of October 1, 2018, at 8 PM UTC, then you input 1538424000
	//    as the start time.
	//
	// StartTime is a required field
	StartTime *time.Time `locationName:"startTime" type:"timestamp" required:"true"`

	// The array of statistics for your metric data request.
	//
	// Statistics is a required field
	Statistics []MetricStatistic `locationName:"statistics" type:"list" required:"true"`

	// The unit for the metric data request.
	//
	// Unit is a required field
	Unit MetricUnit `locationName:"unit" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseMetricDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRelationalDatabaseMetricDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRelationalDatabaseMetricDataInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}
	if len(s.MetricName) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.Period == nil {
		invalidParams.Add(aws.NewErrParamRequired("Period"))
	}
	if s.Period != nil && *s.Period < 60 {
		invalidParams.Add(aws.NewErrParamMinValue("Period", 60))
	}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if s.Statistics == nil {
		invalidParams.Add(aws.NewErrParamRequired("Statistics"))
	}
	if len(s.Unit) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Unit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRelationalDatabaseMetricDataOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your get relational database metric data
	// request.
	MetricData []MetricDatapoint `locationName:"metricData" type:"list"`

	// The name of the metric.
	MetricName RelationalDatabaseMetricName `locationName:"metricName" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseMetricDataOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRelationalDatabaseMetricData = "GetRelationalDatabaseMetricData"

// GetRelationalDatabaseMetricDataRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the data points of the specified metric for a database in Amazon
// Lightsail.
//
//    // Example sending a request using GetRelationalDatabaseMetricDataRequest.
//    req := client.GetRelationalDatabaseMetricDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseMetricData
func (c *Client) GetRelationalDatabaseMetricDataRequest(input *GetRelationalDatabaseMetricDataInput) GetRelationalDatabaseMetricDataRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseMetricData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRelationalDatabaseMetricDataInput{}
	}

	req := c.newRequest(op, input, &GetRelationalDatabaseMetricDataOutput{})
	return GetRelationalDatabaseMetricDataRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseMetricDataRequest}
}

// GetRelationalDatabaseMetricDataRequest is the request type for the
// GetRelationalDatabaseMetricData API operation.
type GetRelationalDatabaseMetricDataRequest struct {
	*aws.Request
	Input *GetRelationalDatabaseMetricDataInput
	Copy  func(*GetRelationalDatabaseMetricDataInput) GetRelationalDatabaseMetricDataRequest
}

// Send marshals and sends the GetRelationalDatabaseMetricData API request.
func (r GetRelationalDatabaseMetricDataRequest) Send(ctx context.Context) (*GetRelationalDatabaseMetricDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseMetricDataResponse{
		GetRelationalDatabaseMetricDataOutput: r.Request.Data.(*GetRelationalDatabaseMetricDataOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseMetricDataResponse is the response type for the
// GetRelationalDatabaseMetricData API operation.
type GetRelationalDatabaseMetricDataResponse struct {
	*GetRelationalDatabaseMetricDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseMetricData request.
func (r *GetRelationalDatabaseMetricDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
