// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type DeleteBucketReplicationInput struct {
	_ struct{} `type:"structure"`

	// The bucket name.
	//
	// It can take a while to propagate the deletion of a replication configuration
	// to all Amazon S3 systems.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketReplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketReplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBucketReplicationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteBucketReplicationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketReplicationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeleteBucketReplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketReplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketReplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBucketReplication = "DeleteBucketReplication"

// DeleteBucketReplicationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the replication configuration from the bucket. For information about
// replication configuration, see Cross-Region Replication (CRR) (https://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html)
// in the Amazon S3 Developer Guide.
//
//    // Example sending a request using DeleteBucketReplicationRequest.
//    req := client.DeleteBucketReplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketReplication
func (c *Client) DeleteBucketReplicationRequest(input *DeleteBucketReplicationInput) DeleteBucketReplicationRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketReplication,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?replication",
	}

	if input == nil {
		input = &DeleteBucketReplicationInput{}
	}

	req := c.newRequest(op, input, &DeleteBucketReplicationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketReplicationRequest{Request: req, Input: input, Copy: c.DeleteBucketReplicationRequest}
}

// DeleteBucketReplicationRequest is the request type for the
// DeleteBucketReplication API operation.
type DeleteBucketReplicationRequest struct {
	*aws.Request
	Input *DeleteBucketReplicationInput
	Copy  func(*DeleteBucketReplicationInput) DeleteBucketReplicationRequest
}

// Send marshals and sends the DeleteBucketReplication API request.
func (r DeleteBucketReplicationRequest) Send(ctx context.Context) (*DeleteBucketReplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketReplicationResponse{
		DeleteBucketReplicationOutput: r.Request.Data.(*DeleteBucketReplicationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketReplicationResponse is the response type for the
// DeleteBucketReplication API operation.
type DeleteBucketReplicationResponse struct {
	*DeleteBucketReplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketReplication request.
func (r *DeleteBucketReplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
