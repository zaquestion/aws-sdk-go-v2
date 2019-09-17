// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRecordHistoryInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The access level to use to obtain results. The default is User.
	AccessLevelFilter *AccessLevelFilter `type:"structure"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The search filter to scope the results.
	SearchFilter *ListRecordHistorySearchFilter `type:"structure"`
}

// String returns the string representation
func (s ListRecordHistoryInput) String() string {
	return awsutil.Prettify(s)
}

type ListRecordHistoryOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`

	// The records, in reverse chronological order.
	RecordDetails []RecordDetail `type:"list"`
}

// String returns the string representation
func (s ListRecordHistoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRecordHistory = "ListRecordHistory"

// ListRecordHistoryRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists the specified requests or all performed requests.
//
//    // Example sending a request using ListRecordHistoryRequest.
//    req := client.ListRecordHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListRecordHistory
func (c *Client) ListRecordHistoryRequest(input *ListRecordHistoryInput) ListRecordHistoryRequest {
	op := &aws.Operation{
		Name:       opListRecordHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRecordHistoryInput{}
	}

	req := c.newRequest(op, input, &ListRecordHistoryOutput{})
	return ListRecordHistoryRequest{Request: req, Input: input, Copy: c.ListRecordHistoryRequest}
}

// ListRecordHistoryRequest is the request type for the
// ListRecordHistory API operation.
type ListRecordHistoryRequest struct {
	*aws.Request
	Input *ListRecordHistoryInput
	Copy  func(*ListRecordHistoryInput) ListRecordHistoryRequest
}

// Send marshals and sends the ListRecordHistory API request.
func (r ListRecordHistoryRequest) Send(ctx context.Context) (*ListRecordHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRecordHistoryResponse{
		ListRecordHistoryOutput: r.Request.Data.(*ListRecordHistoryOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRecordHistoryResponse is the response type for the
// ListRecordHistory API operation.
type ListRecordHistoryResponse struct {
	*ListRecordHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRecordHistory request.
func (r *ListRecordHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
