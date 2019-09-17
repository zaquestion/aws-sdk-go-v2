// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to get the status for a health check.
type GetHealthCheckStatusInput struct {
	_ struct{} `type:"structure"`

	// The ID for the health check that you want the current status for. When you
	// created the health check, CreateHealthCheck returned the ID in the response,
	// in the HealthCheckId element.
	//
	// If you want to check the status of a calculated health check, you must use
	// the Amazon Route 53 console or the CloudWatch console. You can't use GetHealthCheckStatus
	// to get the status of a calculated health check.
	//
	// HealthCheckId is a required field
	HealthCheckId *string `location:"uri" locationName:"HealthCheckId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetHealthCheckStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetHealthCheckStatusInput"}

	if s.HealthCheckId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HealthCheckId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHealthCheckStatusInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.HealthCheckId != nil {
		v := *s.HealthCheckId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "HealthCheckId", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response to a GetHealthCheck request.
type GetHealthCheckStatusOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains one HealthCheckObservation element for each Amazon Route
	// 53 health checker that is reporting a status about the health check endpoint.
	//
	// HealthCheckObservations is a required field
	HealthCheckObservations []HealthCheckObservation `locationNameList:"HealthCheckObservation" type:"list" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckStatusOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHealthCheckStatusOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HealthCheckObservations != nil {
		v := s.HealthCheckObservations

		metadata := protocol.Metadata{ListLocationName: "HealthCheckObservation"}
		ls0 := e.List(protocol.BodyTarget, "HealthCheckObservations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetHealthCheckStatus = "GetHealthCheckStatus"

// GetHealthCheckStatusRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets status of a specified health check.
//
//    // Example sending a request using GetHealthCheckStatusRequest.
//    req := client.GetHealthCheckStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetHealthCheckStatus
func (c *Client) GetHealthCheckStatusRequest(input *GetHealthCheckStatusInput) GetHealthCheckStatusRequest {
	op := &aws.Operation{
		Name:       opGetHealthCheckStatus,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/healthcheck/{HealthCheckId}/status",
	}

	if input == nil {
		input = &GetHealthCheckStatusInput{}
	}

	req := c.newRequest(op, input, &GetHealthCheckStatusOutput{})
	return GetHealthCheckStatusRequest{Request: req, Input: input, Copy: c.GetHealthCheckStatusRequest}
}

// GetHealthCheckStatusRequest is the request type for the
// GetHealthCheckStatus API operation.
type GetHealthCheckStatusRequest struct {
	*aws.Request
	Input *GetHealthCheckStatusInput
	Copy  func(*GetHealthCheckStatusInput) GetHealthCheckStatusRequest
}

// Send marshals and sends the GetHealthCheckStatus API request.
func (r GetHealthCheckStatusRequest) Send(ctx context.Context) (*GetHealthCheckStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHealthCheckStatusResponse{
		GetHealthCheckStatusOutput: r.Request.Data.(*GetHealthCheckStatusOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHealthCheckStatusResponse is the response type for the
// GetHealthCheckStatus API operation.
type GetHealthCheckStatusResponse struct {
	*GetHealthCheckStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHealthCheckStatus request.
func (r *GetHealthCheckStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
