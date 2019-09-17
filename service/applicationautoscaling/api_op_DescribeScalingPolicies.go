// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeScalingPoliciesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of scalable targets. This value can be between 1 and 50.
	// The default value is 50.
	//
	// If this parameter is used, the operation returns up to MaxResults results
	// at a time, along with a NextToken value. To get the next set of results,
	// include the NextToken value in a subsequent call. If this parameter is not
	// used, the operation returns up to 50 results and a NextToken value, if applicable.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of results.
	NextToken *string `type:"string"`

	// The names of the scaling policies to describe.
	PolicyNames []string `type:"list"`

	// The identifier of the resource associated with the scaling policy. This string
	// consists of the resource type and unique identifier. If you specify a scalable
	// dimension, you must also specify a resource ID.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot Fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variants - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	ResourceId *string `min:"1" type:"string"`

	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property. If you specify a scalable dimension, you must
	// also specify a resource ID.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    Fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	ScalableDimension ScalableDimension `type:"string" enum:"true"`

	// The namespace of the AWS service that provides the resource or custom-resource
	// for a resource provided by your own application or service. For more information,
	// see AWS Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the Amazon Web Services General Reference.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DescribeScalingPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScalingPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScalingPoliciesInput"}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}
	if len(s.ServiceNamespace) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceNamespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeScalingPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The token required to get the next set of results. This value is null if
	// there are no more results to return.
	NextToken *string `type:"string"`

	// Information about the scaling policies.
	ScalingPolicies []ScalingPolicy `type:"list"`
}

// String returns the string representation
func (s DescribeScalingPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeScalingPolicies = "DescribeScalingPolicies"

// DescribeScalingPoliciesRequest returns a request value for making API operation for
// Application Auto Scaling.
//
// Describes the Application Auto Scaling scaling policies for the specified
// service namespace.
//
// You can filter the results using ResourceId, ScalableDimension, and PolicyNames.
//
// To create a scaling policy or update an existing one, see PutScalingPolicy.
// If you are no longer using a scaling policy, you can delete it using DeleteScalingPolicy.
//
//    // Example sending a request using DescribeScalingPoliciesRequest.
//    req := client.DescribeScalingPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/DescribeScalingPolicies
func (c *Client) DescribeScalingPoliciesRequest(input *DescribeScalingPoliciesInput) DescribeScalingPoliciesRequest {
	op := &aws.Operation{
		Name:       opDescribeScalingPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeScalingPoliciesInput{}
	}

	req := c.newRequest(op, input, &DescribeScalingPoliciesOutput{})
	return DescribeScalingPoliciesRequest{Request: req, Input: input, Copy: c.DescribeScalingPoliciesRequest}
}

// DescribeScalingPoliciesRequest is the request type for the
// DescribeScalingPolicies API operation.
type DescribeScalingPoliciesRequest struct {
	*aws.Request
	Input *DescribeScalingPoliciesInput
	Copy  func(*DescribeScalingPoliciesInput) DescribeScalingPoliciesRequest
}

// Send marshals and sends the DescribeScalingPolicies API request.
func (r DescribeScalingPoliciesRequest) Send(ctx context.Context) (*DescribeScalingPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScalingPoliciesResponse{
		DescribeScalingPoliciesOutput: r.Request.Data.(*DescribeScalingPoliciesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeScalingPoliciesRequestPaginator returns a paginator for DescribeScalingPolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeScalingPoliciesRequest(input)
//   p := applicationautoscaling.NewDescribeScalingPoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeScalingPoliciesPaginator(req DescribeScalingPoliciesRequest) DescribeScalingPoliciesPaginator {
	return DescribeScalingPoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeScalingPoliciesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeScalingPoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeScalingPoliciesPaginator struct {
	aws.Pager
}

func (p *DescribeScalingPoliciesPaginator) CurrentPage() *DescribeScalingPoliciesOutput {
	return p.Pager.CurrentPage().(*DescribeScalingPoliciesOutput)
}

// DescribeScalingPoliciesResponse is the response type for the
// DescribeScalingPolicies API operation.
type DescribeScalingPoliciesResponse struct {
	*DescribeScalingPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScalingPolicies request.
func (r *DescribeScalingPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
