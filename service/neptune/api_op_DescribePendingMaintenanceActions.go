// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePendingMaintenanceActionsInput struct {
	_ struct{} `type:"structure"`

	// A filter that specifies one or more resources to return pending maintenance
	// actions for.
	//
	// Supported filters:
	//
	//    * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	//    Resource Names (ARNs). The results list will only include pending maintenance
	//    actions for the DB clusters identified by these ARNs.
	//
	//    * db-instance-id - Accepts DB instance identifiers and DB instance ARNs.
	//    The results list will only include pending maintenance actions for the
	//    DB instances identified by these ARNs.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribePendingMaintenanceActions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to a number of records specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The ARN of a resource to return pending maintenance actions for.
	ResourceIdentifier *string `type:"string"`
}

// String returns the string representation
func (s DescribePendingMaintenanceActionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePendingMaintenanceActionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePendingMaintenanceActionsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePendingMaintenanceActionsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous DescribePendingMaintenanceActions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to a number of records specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of the pending maintenance actions for the resource.
	PendingMaintenanceActions []ResourcePendingMaintenanceActions `locationNameList:"ResourcePendingMaintenanceActions" type:"list"`
}

// String returns the string representation
func (s DescribePendingMaintenanceActionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePendingMaintenanceActions = "DescribePendingMaintenanceActions"

// DescribePendingMaintenanceActionsRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Returns a list of resources (for example, DB instances) that have at least
// one pending maintenance action.
//
//    // Example sending a request using DescribePendingMaintenanceActionsRequest.
//    req := client.DescribePendingMaintenanceActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribePendingMaintenanceActions
func (c *Client) DescribePendingMaintenanceActionsRequest(input *DescribePendingMaintenanceActionsInput) DescribePendingMaintenanceActionsRequest {
	op := &aws.Operation{
		Name:       opDescribePendingMaintenanceActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePendingMaintenanceActionsInput{}
	}

	req := c.newRequest(op, input, &DescribePendingMaintenanceActionsOutput{})
	return DescribePendingMaintenanceActionsRequest{Request: req, Input: input, Copy: c.DescribePendingMaintenanceActionsRequest}
}

// DescribePendingMaintenanceActionsRequest is the request type for the
// DescribePendingMaintenanceActions API operation.
type DescribePendingMaintenanceActionsRequest struct {
	*aws.Request
	Input *DescribePendingMaintenanceActionsInput
	Copy  func(*DescribePendingMaintenanceActionsInput) DescribePendingMaintenanceActionsRequest
}

// Send marshals and sends the DescribePendingMaintenanceActions API request.
func (r DescribePendingMaintenanceActionsRequest) Send(ctx context.Context) (*DescribePendingMaintenanceActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePendingMaintenanceActionsResponse{
		DescribePendingMaintenanceActionsOutput: r.Request.Data.(*DescribePendingMaintenanceActionsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePendingMaintenanceActionsResponse is the response type for the
// DescribePendingMaintenanceActions API operation.
type DescribePendingMaintenanceActionsResponse struct {
	*DescribePendingMaintenanceActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePendingMaintenanceActions request.
func (r *DescribePendingMaintenanceActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
