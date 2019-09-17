// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeProvisioningParametersInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The path identifier of the product. This value is optional if the product
	// has a default path, and required if the product has more than one path. To
	// list the paths for a product, use ListLaunchPaths.
	PathId *string `min:"1" type:"string"`

	// The product identifier.
	//
	// ProductId is a required field
	ProductId *string `min:"1" type:"string" required:"true"`

	// The identifier of the provisioning artifact.
	//
	// ProvisioningArtifactId is a required field
	ProvisioningArtifactId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProvisioningParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProvisioningParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProvisioningParametersInput"}
	if s.PathId != nil && len(*s.PathId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PathId", 1))
	}

	if s.ProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductId"))
	}
	if s.ProductId != nil && len(*s.ProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductId", 1))
	}

	if s.ProvisioningArtifactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisioningArtifactId"))
	}
	if s.ProvisioningArtifactId != nil && len(*s.ProvisioningArtifactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisioningArtifactId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeProvisioningParametersOutput struct {
	_ struct{} `type:"structure"`

	// Information about the constraints used to provision the product.
	ConstraintSummaries []ConstraintSummary `type:"list"`

	// Information about the parameters used to provision the product.
	ProvisioningArtifactParameters []ProvisioningArtifactParameter `type:"list"`

	// An object that contains information about preferences, such as regions and
	// accounts, for the provisioning artifact.
	ProvisioningArtifactPreferences *ProvisioningArtifactPreferences `type:"structure"`

	// Information about the TagOptions associated with the resource.
	TagOptions []TagOptionSummary `type:"list"`

	// Any additional metadata specifically related to the provisioning of the product.
	// For example, see the Version field of the CloudFormation template.
	UsageInstructions []UsageInstruction `type:"list"`
}

// String returns the string representation
func (s DescribeProvisioningParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeProvisioningParameters = "DescribeProvisioningParameters"

// DescribeProvisioningParametersRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the configuration required to provision the specified
// product using the specified provisioning artifact.
//
// If the output contains a TagOption key with an empty list of values, there
// is a TagOption conflict for that key. The end user cannot take action to
// fix the conflict, and launch is not blocked. In subsequent calls to ProvisionProduct,
// do not include conflicted TagOption keys as tags, or this causes the error
// "Parameter validation failed: Missing required parameter in Tags[N]:Value".
// Tag the provisioned product with the value sc-tagoption-conflict-portfolioId-productId.
//
//    // Example sending a request using DescribeProvisioningParametersRequest.
//    req := client.DescribeProvisioningParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProvisioningParameters
func (c *Client) DescribeProvisioningParametersRequest(input *DescribeProvisioningParametersInput) DescribeProvisioningParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeProvisioningParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeProvisioningParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeProvisioningParametersOutput{})
	return DescribeProvisioningParametersRequest{Request: req, Input: input, Copy: c.DescribeProvisioningParametersRequest}
}

// DescribeProvisioningParametersRequest is the request type for the
// DescribeProvisioningParameters API operation.
type DescribeProvisioningParametersRequest struct {
	*aws.Request
	Input *DescribeProvisioningParametersInput
	Copy  func(*DescribeProvisioningParametersInput) DescribeProvisioningParametersRequest
}

// Send marshals and sends the DescribeProvisioningParameters API request.
func (r DescribeProvisioningParametersRequest) Send(ctx context.Context) (*DescribeProvisioningParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProvisioningParametersResponse{
		DescribeProvisioningParametersOutput: r.Request.Data.(*DescribeProvisioningParametersOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProvisioningParametersResponse is the response type for the
// DescribeProvisioningParameters API operation.
type DescribeProvisioningParametersResponse struct {
	*DescribeProvisioningParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProvisioningParameters request.
func (r *DescribeProvisioningParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
