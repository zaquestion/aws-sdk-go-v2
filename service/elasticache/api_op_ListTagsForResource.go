// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input parameters for the ListTagsForResource operation.
type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource for which you want the list
	// of tags, for example arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster
	// or arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// ResourceName is a required field
	ResourceName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output from the AddTagsToResource, ListTagsForResource, and
// RemoveTagsFromResource operations.
type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// A list of cost allocation tags as key-value pairs.
	TagList []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsForResource = "ListTagsForResource"

// ListTagsForResourceRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Lists all cost allocation tags currently on the named resource. A cost allocation
// tag is a key-value pair where the key is case-sensitive and the value is
// optional. You can use cost allocation tags to categorize and track your AWS
// costs.
//
// If the cluster is not in the available state, ListTagsForResource returns
// an error.
//
// You can have a maximum of 50 cost allocation tags on an ElastiCache resource.
// For more information, see Monitoring Costs with Tags (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Tagging.html).
//
//    // Example sending a request using ListTagsForResourceRequest.
//    req := client.ListTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ListTagsForResource
func (c *Client) ListTagsForResourceRequest(input *ListTagsForResourceInput) ListTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opListTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &ListTagsForResourceOutput{})
	return ListTagsForResourceRequest{Request: req, Input: input, Copy: c.ListTagsForResourceRequest}
}

// ListTagsForResourceRequest is the request type for the
// ListTagsForResource API operation.
type ListTagsForResourceRequest struct {
	*aws.Request
	Input *ListTagsForResourceInput
	Copy  func(*ListTagsForResourceInput) ListTagsForResourceRequest
}

// Send marshals and sends the ListTagsForResource API request.
func (r ListTagsForResourceRequest) Send(ctx context.Context) (*ListTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForResourceResponse{
		ListTagsForResourceOutput: r.Request.Data.(*ListTagsForResourceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForResourceResponse is the response type for the
// ListTagsForResource API operation.
type ListTagsForResourceResponse struct {
	*ListTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForResource request.
func (r *ListTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
