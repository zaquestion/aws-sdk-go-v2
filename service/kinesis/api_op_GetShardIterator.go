// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for GetShardIterator.
type GetShardIteratorInput struct {
	_ struct{} `type:"structure"`

	// The shard ID of the Kinesis Data Streams shard to get the iterator for.
	//
	// ShardId is a required field
	ShardId *string `min:"1" type:"string" required:"true"`

	// Determines how the shard iterator is used to start reading data records from
	// the shard.
	//
	// The following are the valid Amazon Kinesis shard iterator types:
	//
	//    * AT_SEQUENCE_NUMBER - Start reading from the position denoted by a specific
	//    sequence number, provided in the value StartingSequenceNumber.
	//
	//    * AFTER_SEQUENCE_NUMBER - Start reading right after the position denoted
	//    by a specific sequence number, provided in the value StartingSequenceNumber.
	//
	//    * AT_TIMESTAMP - Start reading from the position denoted by a specific
	//    time stamp, provided in the value Timestamp.
	//
	//    * TRIM_HORIZON - Start reading at the last untrimmed record in the shard
	//    in the system, which is the oldest data record in the shard.
	//
	//    * LATEST - Start reading just after the most recent record in the shard,
	//    so that you always read the most recent data in the shard.
	//
	// ShardIteratorType is a required field
	ShardIteratorType ShardIteratorType `type:"string" required:"true" enum:"true"`

	// The sequence number of the data record in the shard from which to start reading.
	// Used with shard iterator type AT_SEQUENCE_NUMBER and AFTER_SEQUENCE_NUMBER.
	StartingSequenceNumber *string `type:"string"`

	// The name of the Amazon Kinesis data stream.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`

	// The time stamp of the data record from which to start reading. Used with
	// shard iterator type AT_TIMESTAMP. A time stamp is the Unix epoch date with
	// precision in milliseconds. For example, 2016-04-04T19:58:46.480-00:00 or
	// 1459799926.480. If a record with this exact time stamp does not exist, the
	// iterator returned is for the next (later) record. If the time stamp is older
	// than the current trim horizon, the iterator returned is for the oldest untrimmed
	// data record (TRIM_HORIZON).
	Timestamp *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s GetShardIteratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetShardIteratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetShardIteratorInput"}

	if s.ShardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardId"))
	}
	if s.ShardId != nil && len(*s.ShardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardId", 1))
	}
	if len(s.ShardIteratorType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ShardIteratorType"))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output for GetShardIterator.
type GetShardIteratorOutput struct {
	_ struct{} `type:"structure"`

	// The position in the shard from which to start reading data records sequentially.
	// A shard iterator specifies this position using the sequence number of a data
	// record in a shard.
	ShardIterator *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetShardIteratorOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetShardIterator = "GetShardIterator"

// GetShardIteratorRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Gets an Amazon Kinesis shard iterator. A shard iterator expires 5 minutes
// after it is returned to the requester.
//
// A shard iterator specifies the shard position from which to start reading
// data records sequentially. The position is specified using the sequence number
// of a data record in a shard. A sequence number is the identifier associated
// with every record ingested in the stream, and is assigned when a record is
// put into the stream. Each stream has one or more shards.
//
// You must specify the shard iterator type. For example, you can set the ShardIteratorType
// parameter to read exactly from the position denoted by a specific sequence
// number by using the AT_SEQUENCE_NUMBER shard iterator type. Alternatively,
// the parameter can read right after the sequence number by using the AFTER_SEQUENCE_NUMBER
// shard iterator type, using sequence numbers returned by earlier calls to
// PutRecord, PutRecords, GetRecords, or DescribeStream. In the request, you
// can specify the shard iterator type AT_TIMESTAMP to read records from an
// arbitrary point in time, TRIM_HORIZON to cause ShardIterator to point to
// the last untrimmed record in the shard in the system (the oldest data record
// in the shard), or LATEST so that you always read the most recent data in
// the shard.
//
// When you read repeatedly from a stream, use a GetShardIterator request to
// get the first shard iterator for use in your first GetRecords request and
// for subsequent reads use the shard iterator returned by the GetRecords request
// in NextShardIterator. A new shard iterator is returned by every GetRecords
// request in NextShardIterator, which you use in the ShardIterator parameter
// of the next GetRecords request.
//
// If a GetShardIterator request is made too often, you receive a ProvisionedThroughputExceededException.
// For more information about throughput limits, see GetRecords, and Streams
// Limits (http://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// If the shard is closed, GetShardIterator returns a valid iterator for the
// last sequence number of the shard. A shard can be closed as a result of using
// SplitShard or MergeShards.
//
// GetShardIterator has a limit of five transactions per second per account
// per open shard.
//
//    // Example sending a request using GetShardIteratorRequest.
//    req := client.GetShardIteratorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/GetShardIterator
func (c *Client) GetShardIteratorRequest(input *GetShardIteratorInput) GetShardIteratorRequest {
	op := &aws.Operation{
		Name:       opGetShardIterator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetShardIteratorInput{}
	}

	req := c.newRequest(op, input, &GetShardIteratorOutput{})
	return GetShardIteratorRequest{Request: req, Input: input, Copy: c.GetShardIteratorRequest}
}

// GetShardIteratorRequest is the request type for the
// GetShardIterator API operation.
type GetShardIteratorRequest struct {
	*aws.Request
	Input *GetShardIteratorInput
	Copy  func(*GetShardIteratorInput) GetShardIteratorRequest
}

// Send marshals and sends the GetShardIterator API request.
func (r GetShardIteratorRequest) Send(ctx context.Context) (*GetShardIteratorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetShardIteratorResponse{
		GetShardIteratorOutput: r.Request.Data.(*GetShardIteratorOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetShardIteratorResponse is the response type for the
// GetShardIterator API operation.
type GetShardIteratorResponse struct {
	*GetShardIteratorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetShardIterator request.
func (r *GetShardIteratorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
