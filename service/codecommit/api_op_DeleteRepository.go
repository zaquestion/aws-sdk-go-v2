// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a delete repository operation.
type DeleteRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the repository to delete.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRepositoryInput"}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a delete repository operation.
type DeleteRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the repository that was deleted.
	RepositoryId *string `locationName:"repositoryId" type:"string"`
}

// String returns the string representation
func (s DeleteRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRepository = "DeleteRepository"

// DeleteRepositoryRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Deletes a repository. If a specified repository was already deleted, a null
// repository ID will be returned.
//
// Deleting a repository also deletes all associated objects and metadata. After
// a repository is deleted, all future push calls to the deleted repository
// will fail.
//
//    // Example sending a request using DeleteRepositoryRequest.
//    req := client.DeleteRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DeleteRepository
func (c *Client) DeleteRepositoryRequest(input *DeleteRepositoryInput) DeleteRepositoryRequest {
	op := &aws.Operation{
		Name:       opDeleteRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRepositoryInput{}
	}

	req := c.newRequest(op, input, &DeleteRepositoryOutput{})
	return DeleteRepositoryRequest{Request: req, Input: input, Copy: c.DeleteRepositoryRequest}
}

// DeleteRepositoryRequest is the request type for the
// DeleteRepository API operation.
type DeleteRepositoryRequest struct {
	*aws.Request
	Input *DeleteRepositoryInput
	Copy  func(*DeleteRepositoryInput) DeleteRepositoryRequest
}

// Send marshals and sends the DeleteRepository API request.
func (r DeleteRepositoryRequest) Send(ctx context.Context) (*DeleteRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRepositoryResponse{
		DeleteRepositoryOutput: r.Request.Data.(*DeleteRepositoryOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRepositoryResponse is the response type for the
// DeleteRepository API operation.
type DeleteRepositoryResponse struct {
	*DeleteRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRepository request.
func (r *DeleteRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
