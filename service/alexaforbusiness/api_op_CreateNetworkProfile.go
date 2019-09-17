// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateNetworkProfileInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the Private Certificate Authority (PCA) created in AWS Certificate
	// Manager (ACM). This is used to issue certificates to the devices.
	CertificateAuthorityArn *string `type:"string"`

	// A unique, user-specified identifier for the request that ensures idempotency.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `min:"10" type:"string" required:"true" idempotencyToken:"true"`

	// The current password of the Wi-Fi network.
	CurrentPassword *string `min:"5" type:"string"`

	// Detailed information about a device's network profile.
	Description *string `type:"string"`

	// The authentication standard that is used in the EAP framework. Currently,
	// EAP_TLS is supported.
	EapMethod NetworkEapMethod `type:"string" enum:"true"`

	// The name of the network profile associated with a device.
	//
	// NetworkProfileName is a required field
	NetworkProfileName *string `min:"1" type:"string" required:"true"`

	// The next, or subsequent, password of the Wi-Fi network. This password is
	// asynchronously transmitted to the device and is used when the password of
	// the network changes to NextPassword.
	NextPassword *string `type:"string"`

	// The security type of the Wi-Fi network. This can be WPA2_ENTERPRISE, WPA2_PSK,
	// WPA_PSK, WEP, or OPEN.
	//
	// SecurityType is a required field
	SecurityType NetworkSecurityType `type:"string" required:"true" enum:"true"`

	// The SSID of the Wi-Fi network.
	//
	// Ssid is a required field
	Ssid *string `min:"1" type:"string" required:"true"`

	// The root certificates of your authentication server that is installed on
	// your devices and used to trust your authentication server during EAP negotiation.
	TrustAnchors []string `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateNetworkProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNetworkProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNetworkProfileInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}
	if s.CurrentPassword != nil && len(*s.CurrentPassword) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentPassword", 5))
	}

	if s.NetworkProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkProfileName"))
	}
	if s.NetworkProfileName != nil && len(*s.NetworkProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkProfileName", 1))
	}
	if len(s.SecurityType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SecurityType"))
	}

	if s.Ssid == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ssid"))
	}
	if s.Ssid != nil && len(*s.Ssid) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Ssid", 1))
	}
	if s.TrustAnchors != nil && len(s.TrustAnchors) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrustAnchors", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateNetworkProfileOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the network profile associated with a device.
	NetworkProfileArn *string `type:"string"`
}

// String returns the string representation
func (s CreateNetworkProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNetworkProfile = "CreateNetworkProfile"

// CreateNetworkProfileRequest returns a request value for making API operation for
// Alexa For Business.
//
// Creates a network profile with the specified details.
//
//    // Example sending a request using CreateNetworkProfileRequest.
//    req := client.CreateNetworkProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateNetworkProfile
func (c *Client) CreateNetworkProfileRequest(input *CreateNetworkProfileInput) CreateNetworkProfileRequest {
	op := &aws.Operation{
		Name:       opCreateNetworkProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNetworkProfileInput{}
	}

	req := c.newRequest(op, input, &CreateNetworkProfileOutput{})
	return CreateNetworkProfileRequest{Request: req, Input: input, Copy: c.CreateNetworkProfileRequest}
}

// CreateNetworkProfileRequest is the request type for the
// CreateNetworkProfile API operation.
type CreateNetworkProfileRequest struct {
	*aws.Request
	Input *CreateNetworkProfileInput
	Copy  func(*CreateNetworkProfileInput) CreateNetworkProfileRequest
}

// Send marshals and sends the CreateNetworkProfile API request.
func (r CreateNetworkProfileRequest) Send(ctx context.Context) (*CreateNetworkProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNetworkProfileResponse{
		CreateNetworkProfileOutput: r.Request.Data.(*CreateNetworkProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNetworkProfileResponse is the response type for the
// CreateNetworkProfile API operation.
type CreateNetworkProfileResponse struct {
	*CreateNetworkProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNetworkProfile request.
func (r *CreateNetworkProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
