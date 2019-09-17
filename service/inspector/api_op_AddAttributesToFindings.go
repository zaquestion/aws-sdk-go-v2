// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddAttributesToFindingsInput struct {
	_ struct{} `type:"structure"`

	// The array of attributes that you want to assign to specified findings.
	//
	// Attributes is a required field
	Attributes []Attribute `locationName:"attributes" type:"list" required:"true"`

	// The ARNs that specify the findings that you want to assign attributes to.
	//
	// FindingArns is a required field
	FindingArns []string `locationName:"findingArns" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s AddAttributesToFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddAttributesToFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddAttributesToFindingsInput"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.FindingArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingArns"))
	}
	if s.FindingArns != nil && len(s.FindingArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FindingArns", 1))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddAttributesToFindingsOutput struct {
	_ struct{} `type:"structure"`

	// Attribute details that cannot be described. An error code is provided for
	// each failed item.
	//
	// FailedItems is a required field
	FailedItems map[string]FailedItemDetails `locationName:"failedItems" type:"map" required:"true"`
}

// String returns the string representation
func (s AddAttributesToFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddAttributesToFindings = "AddAttributesToFindings"

// AddAttributesToFindingsRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Assigns attributes (key and value pairs) to the findings that are specified
// by the ARNs of the findings.
//
//    // Example sending a request using AddAttributesToFindingsRequest.
//    req := client.AddAttributesToFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/AddAttributesToFindings
func (c *Client) AddAttributesToFindingsRequest(input *AddAttributesToFindingsInput) AddAttributesToFindingsRequest {
	op := &aws.Operation{
		Name:       opAddAttributesToFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddAttributesToFindingsInput{}
	}

	req := c.newRequest(op, input, &AddAttributesToFindingsOutput{})
	return AddAttributesToFindingsRequest{Request: req, Input: input, Copy: c.AddAttributesToFindingsRequest}
}

// AddAttributesToFindingsRequest is the request type for the
// AddAttributesToFindings API operation.
type AddAttributesToFindingsRequest struct {
	*aws.Request
	Input *AddAttributesToFindingsInput
	Copy  func(*AddAttributesToFindingsInput) AddAttributesToFindingsRequest
}

// Send marshals and sends the AddAttributesToFindings API request.
func (r AddAttributesToFindingsRequest) Send(ctx context.Context) (*AddAttributesToFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddAttributesToFindingsResponse{
		AddAttributesToFindingsOutput: r.Request.Data.(*AddAttributesToFindingsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddAttributesToFindingsResponse is the response type for the
// AddAttributesToFindings API operation.
type AddAttributesToFindingsResponse struct {
	*AddAttributesToFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddAttributesToFindings request.
func (r *AddAttributesToFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
