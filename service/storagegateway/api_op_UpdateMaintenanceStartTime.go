// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the following fields:
//
//    * UpdateMaintenanceStartTimeInput$DayOfMonth
//
//    * UpdateMaintenanceStartTimeInput$DayOfWeek
//
//    * UpdateMaintenanceStartTimeInput$HourOfDay
//
//    * UpdateMaintenanceStartTimeInput$MinuteOfHour
type UpdateMaintenanceStartTimeInput struct {
	_ struct{} `type:"structure"`

	// The day of the month component of the maintenance start time represented
	// as an ordinal number from 1 to 28, where 1 represents the first day of the
	// month and 28 represents the last day of the month.
	//
	// This value is only available for tape and volume gateways.
	DayOfMonth *int64 `min:"1" type:"integer"`

	// The day of the week component of the maintenance start time week represented
	// as an ordinal number from 0 to 6, where 0 represents Sunday and 6 Saturday.
	DayOfWeek *int64 `type:"integer"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// The hour component of the maintenance start time represented as hh, where
	// hh is the hour (00 to 23). The hour of the day is in the time zone of the
	// gateway.
	//
	// HourOfDay is a required field
	HourOfDay *int64 `type:"integer" required:"true"`

	// The minute component of the maintenance start time represented as mm, where
	// mm is the minute (00 to 59). The minute of the hour is in the time zone of
	// the gateway.
	//
	// MinuteOfHour is a required field
	MinuteOfHour *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s UpdateMaintenanceStartTimeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMaintenanceStartTimeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMaintenanceStartTimeInput"}
	if s.DayOfMonth != nil && *s.DayOfMonth < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("DayOfMonth", 1))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if s.HourOfDay == nil {
		invalidParams.Add(aws.NewErrParamRequired("HourOfDay"))
	}

	if s.MinuteOfHour == nil {
		invalidParams.Add(aws.NewErrParamRequired("MinuteOfHour"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the of the gateway whose maintenance start time
// is updated.
type UpdateMaintenanceStartTimeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateMaintenanceStartTimeOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateMaintenanceStartTime = "UpdateMaintenanceStartTime"

// UpdateMaintenanceStartTimeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates a gateway's weekly maintenance start time information, including
// day and time of the week. The maintenance time is the time in your gateway's
// time zone.
//
//    // Example sending a request using UpdateMaintenanceStartTimeRequest.
//    req := client.UpdateMaintenanceStartTimeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateMaintenanceStartTime
func (c *Client) UpdateMaintenanceStartTimeRequest(input *UpdateMaintenanceStartTimeInput) UpdateMaintenanceStartTimeRequest {
	op := &aws.Operation{
		Name:       opUpdateMaintenanceStartTime,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateMaintenanceStartTimeInput{}
	}

	req := c.newRequest(op, input, &UpdateMaintenanceStartTimeOutput{})
	return UpdateMaintenanceStartTimeRequest{Request: req, Input: input, Copy: c.UpdateMaintenanceStartTimeRequest}
}

// UpdateMaintenanceStartTimeRequest is the request type for the
// UpdateMaintenanceStartTime API operation.
type UpdateMaintenanceStartTimeRequest struct {
	*aws.Request
	Input *UpdateMaintenanceStartTimeInput
	Copy  func(*UpdateMaintenanceStartTimeInput) UpdateMaintenanceStartTimeRequest
}

// Send marshals and sends the UpdateMaintenanceStartTime API request.
func (r UpdateMaintenanceStartTimeRequest) Send(ctx context.Context) (*UpdateMaintenanceStartTimeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMaintenanceStartTimeResponse{
		UpdateMaintenanceStartTimeOutput: r.Request.Data.(*UpdateMaintenanceStartTimeOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMaintenanceStartTimeResponse is the response type for the
// UpdateMaintenanceStartTime API operation.
type UpdateMaintenanceStartTimeResponse struct {
	*UpdateMaintenanceStartTimeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMaintenanceStartTime request.
func (r *UpdateMaintenanceStartTimeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
