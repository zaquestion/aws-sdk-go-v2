// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstanceProfileInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of your instance profile.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstanceProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstanceProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstanceProfileInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInstanceProfileOutput struct {
	_ struct{} `type:"structure"`

	// An object containing information about your instance profile.
	InstanceProfile *InstanceProfile `locationName:"instanceProfile" type:"structure"`
}

// String returns the string representation
func (s GetInstanceProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstanceProfile = "GetInstanceProfile"

// GetInstanceProfileRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns information about the specified instance profile.
//
//    // Example sending a request using GetInstanceProfileRequest.
//    req := client.GetInstanceProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetInstanceProfile
func (c *Client) GetInstanceProfileRequest(input *GetInstanceProfileInput) GetInstanceProfileRequest {
	op := &aws.Operation{
		Name:       opGetInstanceProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInstanceProfileInput{}
	}

	req := c.newRequest(op, input, &GetInstanceProfileOutput{})
	return GetInstanceProfileRequest{Request: req, Input: input, Copy: c.GetInstanceProfileRequest}
}

// GetInstanceProfileRequest is the request type for the
// GetInstanceProfile API operation.
type GetInstanceProfileRequest struct {
	*aws.Request
	Input *GetInstanceProfileInput
	Copy  func(*GetInstanceProfileInput) GetInstanceProfileRequest
}

// Send marshals and sends the GetInstanceProfile API request.
func (r GetInstanceProfileRequest) Send(ctx context.Context) (*GetInstanceProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstanceProfileResponse{
		GetInstanceProfileOutput: r.Request.Data.(*GetInstanceProfileOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstanceProfileResponse is the response type for the
// GetInstanceProfile API operation.
type GetInstanceProfileResponse struct {
	*GetInstanceProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstanceProfile request.
func (r *GetInstanceProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
