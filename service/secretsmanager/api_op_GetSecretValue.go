// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSecretValueInput struct {
	_ struct{} `type:"structure"`

	// Specifies the secret containing the version that you want to retrieve. You
	// can specify either the Amazon Resource Name (ARN) or the friendly name of
	// the secret.
	//
	// If you specify an ARN, we generally recommend that you specify a complete
	// ARN. You can specify a partial ARN too—for example, if you don’t include
	// the final hyphen and six random characters that Secrets Manager adds at the
	// end of the ARN when you created the secret. A partial ARN match can work
	// as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as
	// a partial ARN, then those characters cause Secrets Manager to assume that
	// you’re specifying a complete ARN. This confusion can cause unexpected results.
	// To avoid this situation, we recommend that you don’t create secret names
	// that end with a hyphen followed by six characters.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`

	// Specifies the unique identifier of the version of the secret that you want
	// to retrieve. If you specify this parameter then don't specify VersionStage.
	// If you don't specify either a VersionStage or VersionId then the default
	// is to perform the operation on the version with the VersionStage value of
	// AWSCURRENT.
	//
	// This value is typically a UUID-type (https://wikipedia.org/wiki/Universally_unique_identifier)
	// value with 32 hexadecimal digits.
	VersionId *string `min:"32" type:"string"`

	// Specifies the secret version that you want to retrieve by the staging label
	// attached to the version.
	//
	// Staging labels are used to keep track of different versions during the rotation
	// process. If you use this parameter then don't specify VersionId. If you don't
	// specify either a VersionStage or VersionId, then the default is to perform
	// the operation on the version with the VersionStage value of AWSCURRENT.
	VersionStage *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetSecretValueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSecretValueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSecretValueInput"}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}
	if s.VersionId != nil && len(*s.VersionId) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 32))
	}
	if s.VersionStage != nil && len(*s.VersionStage) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionStage", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetSecretValueOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret.
	ARN *string `min:"20" type:"string"`

	// The date and time that this version of the secret was created.
	CreatedDate *time.Time `type:"timestamp"`

	// The friendly name of the secret.
	Name *string `min:"1" type:"string"`

	// The decrypted part of the protected secret information that was originally
	// provided as binary data in the form of a byte array. The response parameter
	// represents the binary data as a base64-encoded (https://tools.ietf.org/html/rfc4648#section-4)
	// string.
	//
	// This parameter is not used if the secret is created by the Secrets Manager
	// console.
	//
	// If you store custom information in this field of the secret, then you must
	// code your Lambda rotation function to parse and interpret whatever you store
	// in the SecretString or SecretBinary fields.
	//
	// SecretBinary is automatically base64 encoded/decoded by the SDK.
	SecretBinary []byte `type:"blob"`

	// The decrypted part of the protected secret information that was originally
	// provided as a string.
	//
	// If you create this secret by using the Secrets Manager console then only
	// the SecretString parameter contains data. Secrets Manager stores the information
	// as a JSON structure of key/value pairs that the Lambda rotation function
	// knows how to parse.
	//
	// If you store custom information in the secret by using the CreateSecret,
	// UpdateSecret, or PutSecretValue API operations instead of the Secrets Manager
	// console, or by using the Other secret type in the console, then you must
	// code your Lambda rotation function to parse and interpret those values.
	SecretString *string `type:"string"`

	// The unique identifier of this version of the secret.
	VersionId *string `min:"32" type:"string"`

	// A list of all of the staging labels currently attached to this version of
	// the secret.
	VersionStages []string `min:"1" type:"list"`
}

// String returns the string representation
func (s GetSecretValueOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSecretValue = "GetSecretValue"

// GetSecretValueRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Retrieves the contents of the encrypted fields SecretString or SecretBinary
// from the specified version of a secret, whichever contains content.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:GetSecretValue
//
//    * kms:Decrypt - required only if you use a customer-managed AWS KMS key
//    to encrypt the secret. You do not need this permission to use the account's
//    default AWS managed CMK for Secrets Manager.
//
// Related operations
//
//    * To create a new version of the secret with different encrypted information,
//    use PutSecretValue.
//
//    * To retrieve the non-encrypted details for the secret, use DescribeSecret.
//
//    // Example sending a request using GetSecretValueRequest.
//    req := client.GetSecretValueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/GetSecretValue
func (c *Client) GetSecretValueRequest(input *GetSecretValueInput) GetSecretValueRequest {
	op := &aws.Operation{
		Name:       opGetSecretValue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSecretValueInput{}
	}

	req := c.newRequest(op, input, &GetSecretValueOutput{})
	return GetSecretValueRequest{Request: req, Input: input, Copy: c.GetSecretValueRequest}
}

// GetSecretValueRequest is the request type for the
// GetSecretValue API operation.
type GetSecretValueRequest struct {
	*aws.Request
	Input *GetSecretValueInput
	Copy  func(*GetSecretValueInput) GetSecretValueRequest
}

// Send marshals and sends the GetSecretValue API request.
func (r GetSecretValueRequest) Send(ctx context.Context) (*GetSecretValueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSecretValueResponse{
		GetSecretValueOutput: r.Request.Data.(*GetSecretValueOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSecretValueResponse is the response type for the
// GetSecretValue API operation.
type GetSecretValueResponse struct {
	*GetSecretValueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSecretValue request.
func (r *GetSecretValueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
