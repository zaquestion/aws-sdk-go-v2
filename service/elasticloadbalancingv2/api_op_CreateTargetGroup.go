// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTargetGroupInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether health checks are enabled. If the target type is lambda,
	// health checks are disabled by default but can be enabled. If the target type
	// is instance or ip, health checks are always enabled and cannot be disabled.
	HealthCheckEnabled *bool `type:"boolean"`

	// The approximate amount of time, in seconds, between health checks of an individual
	// target. For HTTP and HTTPS health checks, the range is 5–300 seconds. For
	// TCP health checks, the supported values are 10 and 30 seconds. If the target
	// type is instance or ip, the default is 30 seconds. If the target type is
	// lambda, the default is 35 seconds.
	HealthCheckIntervalSeconds *int64 `min:"5" type:"integer"`

	// [HTTP/HTTPS health checks] The ping path that is the destination on the targets
	// for health checks. The default is /.
	HealthCheckPath *string `min:"1" type:"string"`

	// The port the load balancer uses when performing health checks on targets.
	// The default is traffic-port, which is the port on which each target receives
	// traffic from the load balancer.
	HealthCheckPort *string `type:"string"`

	// The protocol the load balancer uses when performing health checks on targets.
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers,
	// the default is TCP. The TCP protocol is supported for health checks only
	// if the protocol of the target group is TCP, TLS, UDP, or TCP_UDP. The TLS,
	// UDP, and TCP_UDP protocols are not supported for health checks.
	HealthCheckProtocol ProtocolEnum `type:"string" enum:"true"`

	// The amount of time, in seconds, during which no response from a target means
	// a failed health check. For target groups with a protocol of HTTP or HTTPS,
	// the default is 5 seconds. For target groups with a protocol of TCP or TLS,
	// this value must be 6 seconds for HTTP health checks and 10 seconds for TCP
	// and HTTPS health checks. If the target type is lambda, the default is 30
	// seconds.
	HealthCheckTimeoutSeconds *int64 `min:"2" type:"integer"`

	// The number of consecutive health checks successes required before considering
	// an unhealthy target healthy. For target groups with a protocol of HTTP or
	// HTTPS, the default is 5. For target groups with a protocol of TCP or TLS,
	// the default is 3. If the target type is lambda, the default is 5.
	HealthyThresholdCount *int64 `min:"2" type:"integer"`

	// [HTTP/HTTPS health checks] The HTTP codes to use when checking for a successful
	// response from a target.
	Matcher *Matcher `type:"structure"`

	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of 32
	// characters, must contain only alphanumeric characters or hyphens, and must
	// not begin or end with a hyphen.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The port on which the targets receive traffic. This port is used unless you
	// specify a port override when registering the target. If the target is a Lambda
	// function, this parameter does not apply.
	Port *int64 `min:"1" type:"integer"`

	// The protocol to use for routing traffic to the targets. For Application Load
	// Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers,
	// the supported protocols are TCP, TLS, UDP, or TCP_UDP. A TCP_UDP listener
	// must be associated with a TCP_UDP target group. If the target is a Lambda
	// function, this parameter does not apply.
	Protocol ProtocolEnum `type:"string" enum:"true"`

	// The type of target that you must specify when registering targets with this
	// target group. You can't specify targets for a target group using more than
	// one target type.
	//
	//    * instance - Targets are specified by instance ID. This is the default
	//    value. If the target group protocol is UDP or TCP_UDP, the target type
	//    must be instance.
	//
	//    * ip - Targets are specified by IP address. You can specify IP addresses
	//    from the subnets of the virtual private cloud (VPC) for the target group,
	//    the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and
	//    the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable
	//    IP addresses.
	//
	//    * lambda - The target groups contains a single Lambda function.
	TargetType TargetTypeEnum `type:"string" enum:"true"`

	// The number of consecutive health check failures required before considering
	// a target unhealthy. For target groups with a protocol of HTTP or HTTPS, the
	// default is 2. For target groups with a protocol of TCP or TLS, this value
	// must be the same as the healthy threshold count. If the target type is lambda,
	// the default is 2.
	UnhealthyThresholdCount *int64 `min:"2" type:"integer"`

	// The identifier of the virtual private cloud (VPC). If the target is a Lambda
	// function, this parameter does not apply.
	VpcId *string `type:"string"`
}

// String returns the string representation
func (s CreateTargetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTargetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTargetGroupInput"}
	if s.HealthCheckIntervalSeconds != nil && *s.HealthCheckIntervalSeconds < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckIntervalSeconds", 5))
	}
	if s.HealthCheckPath != nil && len(*s.HealthCheckPath) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HealthCheckPath", 1))
	}
	if s.HealthCheckTimeoutSeconds != nil && *s.HealthCheckTimeoutSeconds < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckTimeoutSeconds", 2))
	}
	if s.HealthyThresholdCount != nil && *s.HealthyThresholdCount < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthyThresholdCount", 2))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Port != nil && *s.Port < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Port", 1))
	}
	if s.UnhealthyThresholdCount != nil && *s.UnhealthyThresholdCount < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("UnhealthyThresholdCount", 2))
	}
	if s.Matcher != nil {
		if err := s.Matcher.Validate(); err != nil {
			invalidParams.AddNested("Matcher", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateTargetGroupOutput struct {
	_ struct{} `type:"structure"`

	// Information about the target group.
	TargetGroups []TargetGroup `type:"list"`
}

// String returns the string representation
func (s CreateTargetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTargetGroup = "CreateTargetGroup"

// CreateTargetGroupRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Creates a target group.
//
// To register targets with the target group, use RegisterTargets. To update
// the health check settings for the target group, use ModifyTargetGroup. To
// monitor the health of targets in the target group, use DescribeTargetHealth.
//
// To route traffic to the targets in a target group, specify the target group
// in an action using CreateListener or CreateRule.
//
// To delete a target group, use DeleteTargetGroup.
//
// This operation is idempotent, which means that it completes at most one time.
// If you attempt to create multiple target groups with the same settings, each
// call succeeds.
//
// For more information, see Target Groups for Your Application Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html)
// in the Application Load Balancers Guide or Target Groups for Your Network
// Load Balancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html)
// in the Network Load Balancers Guide.
//
//    // Example sending a request using CreateTargetGroupRequest.
//    req := client.CreateTargetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/CreateTargetGroup
func (c *Client) CreateTargetGroupRequest(input *CreateTargetGroupInput) CreateTargetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateTargetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTargetGroupInput{}
	}

	req := c.newRequest(op, input, &CreateTargetGroupOutput{})
	return CreateTargetGroupRequest{Request: req, Input: input, Copy: c.CreateTargetGroupRequest}
}

// CreateTargetGroupRequest is the request type for the
// CreateTargetGroup API operation.
type CreateTargetGroupRequest struct {
	*aws.Request
	Input *CreateTargetGroupInput
	Copy  func(*CreateTargetGroupInput) CreateTargetGroupRequest
}

// Send marshals and sends the CreateTargetGroup API request.
func (r CreateTargetGroupRequest) Send(ctx context.Context) (*CreateTargetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTargetGroupResponse{
		CreateTargetGroupOutput: r.Request.Data.(*CreateTargetGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTargetGroupResponse is the response type for the
// CreateTargetGroup API operation.
type CreateTargetGroupResponse struct {
	*CreateTargetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTargetGroup request.
func (r *CreateTargetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
