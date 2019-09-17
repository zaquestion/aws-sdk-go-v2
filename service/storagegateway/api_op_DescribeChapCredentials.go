// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the Amazon Resource Name (ARN) of the iSCSI volume
// target.
type DescribeChapCredentialsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the DescribeStorediSCSIVolumes
	// operation to return to retrieve the TargetARN for specified VolumeARN.
	//
	// TargetARN is a required field
	TargetARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeChapCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeChapCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeChapCredentialsInput"}

	if s.TargetARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetARN"))
	}
	if s.TargetARN != nil && len(*s.TargetARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing a .
type DescribeChapCredentialsOutput struct {
	_ struct{} `type:"structure"`

	// An array of ChapInfo objects that represent CHAP credentials. Each object
	// in the array contains CHAP credential information for one target-initiator
	// pair. If no CHAP credentials are set, an empty array is returned. CHAP credential
	// information is provided in a JSON object with the following fields:
	//
	//    * InitiatorName: The iSCSI initiator that connects to the target.
	//
	//    * SecretToAuthenticateInitiator: The secret key that the initiator (for
	//    example, the Windows client) must provide to participate in mutual CHAP
	//    with the target.
	//
	//    * SecretToAuthenticateTarget: The secret key that the target must provide
	//    to participate in mutual CHAP with the initiator (e.g. Windows client).
	//
	//    * TargetARN: The Amazon Resource Name (ARN) of the storage volume.
	ChapCredentials []ChapInfo `type:"list"`
}

// String returns the string representation
func (s DescribeChapCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeChapCredentials = "DescribeChapCredentials"

// DescribeChapCredentialsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns an array of Challenge-Handshake Authentication Protocol (CHAP) credentials
// information for a specified iSCSI target, one for each target-initiator pair.
//
//    // Example sending a request using DescribeChapCredentialsRequest.
//    req := client.DescribeChapCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeChapCredentials
func (c *Client) DescribeChapCredentialsRequest(input *DescribeChapCredentialsInput) DescribeChapCredentialsRequest {
	op := &aws.Operation{
		Name:       opDescribeChapCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeChapCredentialsInput{}
	}

	req := c.newRequest(op, input, &DescribeChapCredentialsOutput{})
	return DescribeChapCredentialsRequest{Request: req, Input: input, Copy: c.DescribeChapCredentialsRequest}
}

// DescribeChapCredentialsRequest is the request type for the
// DescribeChapCredentials API operation.
type DescribeChapCredentialsRequest struct {
	*aws.Request
	Input *DescribeChapCredentialsInput
	Copy  func(*DescribeChapCredentialsInput) DescribeChapCredentialsRequest
}

// Send marshals and sends the DescribeChapCredentials API request.
func (r DescribeChapCredentialsRequest) Send(ctx context.Context) (*DescribeChapCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeChapCredentialsResponse{
		DescribeChapCredentialsOutput: r.Request.Data.(*DescribeChapCredentialsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeChapCredentialsResponse is the response type for the
// DescribeChapCredentials API operation.
type DescribeChapCredentialsResponse struct {
	*DescribeChapCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeChapCredentials request.
func (r *DescribeChapCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
