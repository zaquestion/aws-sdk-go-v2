// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDetectEntitiesInput struct {
	_ struct{} `type:"structure"`

	// The language of the input documents. You can specify any of the primary languages
	// supported by Amazon Comprehend: German ("de"), English ("en"), Spanish ("es"),
	// French ("fr"), Italian ("it"), or Portuguese ("pt"). All documents must be
	// in the same language.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// A list containing the text of the input documents. The list can contain a
	// maximum of 25 documents. Each document must contain fewer than 5,000 bytes
	// of UTF-8 encoded characters.
	//
	// TextList is a required field
	TextList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectEntitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDetectEntitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDetectEntitiesInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.TextList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TextList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDetectEntitiesOutput struct {
	_ struct{} `type:"structure"`

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order
	// of the documents in the input list. If there are no errors in the batch,
	// the ErrorList is empty.
	//
	// ErrorList is a required field
	ErrorList []BatchItemError `type:"list" required:"true"`

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the documents
	// in the input list. If all of the documents contain an error, the ResultList
	// is empty.
	//
	// ResultList is a required field
	ResultList []BatchDetectEntitiesItemResult `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectEntitiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDetectEntities = "BatchDetectEntities"

// BatchDetectEntitiesRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Inspects the text of a batch of documents for named entities and returns
// information about them. For more information about named entities, see how-entities
//
//    // Example sending a request using BatchDetectEntitiesRequest.
//    req := client.BatchDetectEntitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectEntities
func (c *Client) BatchDetectEntitiesRequest(input *BatchDetectEntitiesInput) BatchDetectEntitiesRequest {
	op := &aws.Operation{
		Name:       opBatchDetectEntities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDetectEntitiesInput{}
	}

	req := c.newRequest(op, input, &BatchDetectEntitiesOutput{})
	return BatchDetectEntitiesRequest{Request: req, Input: input, Copy: c.BatchDetectEntitiesRequest}
}

// BatchDetectEntitiesRequest is the request type for the
// BatchDetectEntities API operation.
type BatchDetectEntitiesRequest struct {
	*aws.Request
	Input *BatchDetectEntitiesInput
	Copy  func(*BatchDetectEntitiesInput) BatchDetectEntitiesRequest
}

// Send marshals and sends the BatchDetectEntities API request.
func (r BatchDetectEntitiesRequest) Send(ctx context.Context) (*BatchDetectEntitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDetectEntitiesResponse{
		BatchDetectEntitiesOutput: r.Request.Data.(*BatchDetectEntitiesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDetectEntitiesResponse is the response type for the
// BatchDetectEntities API operation.
type BatchDetectEntitiesResponse struct {
	*BatchDetectEntitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDetectEntities request.
func (r *BatchDetectEntitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
