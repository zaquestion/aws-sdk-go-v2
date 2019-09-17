// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteDatasetInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset to delete.
	//
	// DatasetArn is a required field
	DatasetArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDatasetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDatasetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDatasetInput"}

	if s.DatasetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDatasetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDatasetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDataset = "DeleteDataset"

// DeleteDatasetRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Deletes an Amazon Forecast dataset created using the CreateDataset operation.
// To be deleted, the dataset must have a status of ACTIVE or CREATE_FAILED.
// Use the DescribeDataset operation to get the status.
//
//    // Example sending a request using DeleteDatasetRequest.
//    req := client.DeleteDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/DeleteDataset
func (c *Client) DeleteDatasetRequest(input *DeleteDatasetInput) DeleteDatasetRequest {
	op := &aws.Operation{
		Name:       opDeleteDataset,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDatasetInput{}
	}

	req := c.newRequest(op, input, &DeleteDatasetOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDatasetRequest{Request: req, Input: input, Copy: c.DeleteDatasetRequest}
}

// DeleteDatasetRequest is the request type for the
// DeleteDataset API operation.
type DeleteDatasetRequest struct {
	*aws.Request
	Input *DeleteDatasetInput
	Copy  func(*DeleteDatasetInput) DeleteDatasetRequest
}

// Send marshals and sends the DeleteDataset API request.
func (r DeleteDatasetRequest) Send(ctx context.Context) (*DeleteDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDatasetResponse{
		DeleteDatasetOutput: r.Request.Data.(*DeleteDatasetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDatasetResponse is the response type for the
// DeleteDataset API operation.
type DeleteDatasetResponse struct {
	*DeleteDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDataset request.
func (r *DeleteDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
