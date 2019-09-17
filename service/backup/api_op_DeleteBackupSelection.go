// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteBackupSelectionInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a backup plan.
	//
	// BackupPlanId is a required field
	BackupPlanId *string `location:"uri" locationName:"backupPlanId" type:"string" required:"true"`

	// Uniquely identifies the body of a request to assign a set of resources to
	// a backup plan.
	//
	// SelectionId is a required field
	SelectionId *string `location:"uri" locationName:"selectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackupSelectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackupSelectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBackupSelectionInput"}

	if s.BackupPlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlanId"))
	}

	if s.SelectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SelectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackupSelectionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SelectionId != nil {
		v := *s.SelectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "selectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteBackupSelectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBackupSelectionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackupSelectionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBackupSelection = "DeleteBackupSelection"

// DeleteBackupSelectionRequest returns a request value for making API operation for
// AWS Backup.
//
// Deletes the resource selection associated with a backup plan that is specified
// by the SelectionId.
//
//    // Example sending a request using DeleteBackupSelectionRequest.
//    req := client.DeleteBackupSelectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupSelection
func (c *Client) DeleteBackupSelectionRequest(input *DeleteBackupSelectionInput) DeleteBackupSelectionRequest {
	op := &aws.Operation{
		Name:       opDeleteBackupSelection,
		HTTPMethod: "DELETE",
		HTTPPath:   "/backup/plans/{backupPlanId}/selections/{selectionId}",
	}

	if input == nil {
		input = &DeleteBackupSelectionInput{}
	}

	req := c.newRequest(op, input, &DeleteBackupSelectionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBackupSelectionRequest{Request: req, Input: input, Copy: c.DeleteBackupSelectionRequest}
}

// DeleteBackupSelectionRequest is the request type for the
// DeleteBackupSelection API operation.
type DeleteBackupSelectionRequest struct {
	*aws.Request
	Input *DeleteBackupSelectionInput
	Copy  func(*DeleteBackupSelectionInput) DeleteBackupSelectionRequest
}

// Send marshals and sends the DeleteBackupSelection API request.
func (r DeleteBackupSelectionRequest) Send(ctx context.Context) (*DeleteBackupSelectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupSelectionResponse{
		DeleteBackupSelectionOutput: r.Request.Data.(*DeleteBackupSelectionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupSelectionResponse is the response type for the
// DeleteBackupSelection API operation.
type DeleteBackupSelectionResponse struct {
	*DeleteBackupSelectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackupSelection request.
func (r *DeleteBackupSelectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
