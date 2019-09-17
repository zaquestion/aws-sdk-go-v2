// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDomainSuggestionsInput struct {
	_ struct{} `type:"structure"`

	// A domain name that you want to use as the basis for a list of possible domain
	// names. The domain name must contain a top-level domain (TLD), such as .com,
	// that Amazon Route 53 supports. For a list of TLDs, see Domains that You Can
	// Register with Amazon Route 53 (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// If OnlyAvailable is true, Amazon Route 53 returns only domain names that
	// are available. If OnlyAvailable is false, Amazon Route 53 returns domain
	// names without checking whether they're available to be registered. To determine
	// whether the domain is available, you can call checkDomainAvailability for
	// each suggestion.
	//
	// OnlyAvailable is a required field
	OnlyAvailable *bool `type:"boolean" required:"true"`

	// The number of suggested domain names that you want Amazon Route 53 to return.
	//
	// SuggestionCount is a required field
	SuggestionCount *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s GetDomainSuggestionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainSuggestionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainSuggestionsInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.OnlyAvailable == nil {
		invalidParams.Add(aws.NewErrParamRequired("OnlyAvailable"))
	}

	if s.SuggestionCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("SuggestionCount"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDomainSuggestionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of possible domain names. If you specified true for OnlyAvailable
	// in the request, the list contains only domains that are available for registration.
	SuggestionsList []DomainSuggestion `type:"list"`
}

// String returns the string representation
func (s GetDomainSuggestionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDomainSuggestions = "GetDomainSuggestions"

// GetDomainSuggestionsRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// The GetDomainSuggestions operation returns a list of suggested domain names
// given a string, which can either be a domain name or simply a word or phrase
// (without spaces).
//
//    // Example sending a request using GetDomainSuggestionsRequest.
//    req := client.GetDomainSuggestionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/GetDomainSuggestions
func (c *Client) GetDomainSuggestionsRequest(input *GetDomainSuggestionsInput) GetDomainSuggestionsRequest {
	op := &aws.Operation{
		Name:       opGetDomainSuggestions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDomainSuggestionsInput{}
	}

	req := c.newRequest(op, input, &GetDomainSuggestionsOutput{})
	return GetDomainSuggestionsRequest{Request: req, Input: input, Copy: c.GetDomainSuggestionsRequest}
}

// GetDomainSuggestionsRequest is the request type for the
// GetDomainSuggestions API operation.
type GetDomainSuggestionsRequest struct {
	*aws.Request
	Input *GetDomainSuggestionsInput
	Copy  func(*GetDomainSuggestionsInput) GetDomainSuggestionsRequest
}

// Send marshals and sends the GetDomainSuggestions API request.
func (r GetDomainSuggestionsRequest) Send(ctx context.Context) (*GetDomainSuggestionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainSuggestionsResponse{
		GetDomainSuggestionsOutput: r.Request.Data.(*GetDomainSuggestionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainSuggestionsResponse is the response type for the
// GetDomainSuggestions API operation.
type GetDomainSuggestionsResponse struct {
	*GetDomainSuggestionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainSuggestions request.
func (r *GetDomainSuggestionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
