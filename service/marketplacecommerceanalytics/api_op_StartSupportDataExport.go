// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the StartSupportDataExport operation.
type StartSupportDataExportInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file.
	CustomerDefinedValues map[string]string `locationName:"customerDefinedValues" min:"1" type:"map"`

	// Specifies the data set type to be written to the output csv file. The data
	// set types customer_support_contacts_data and test_customer_support_contacts_data
	// both result in a csv file containing the following fields: Product Id, Product
	// Code, Customer Guid, Subscription Guid, Subscription Start Date, Organization,
	// AWS Account Id, Given Name, Surname, Telephone Number, Email, Title, Country
	// Code, ZIP Code, Operation Type, and Operation Time.
	//
	//    * customer_support_contacts_data Customer support contact data. The data
	//    set will contain all changes (Creates, Updates, and Deletes) to customer
	//    support contact data from the date specified in the from_date parameter.
	//
	//    * test_customer_support_contacts_data An example data set containing static
	//    test data in the same format as customer_support_contacts_data
	//
	// DataSetType is a required field
	DataSetType SupportDataSetType `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"true"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	//
	// DestinationS3BucketName is a required field
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The start date from which to retrieve the data set in UTC. This parameter
	// only affects the customer_support_contacts_data data set type.
	//
	// FromDate is a required field
	FromDate *time.Time `locationName:"fromDate" type:"timestamp" required:"true"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	//
	// RoleNameArn is a required field
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	//
	// SnsTopicArn is a required field
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartSupportDataExportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSupportDataExportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSupportDataExportInput"}
	if s.CustomerDefinedValues != nil && len(s.CustomerDefinedValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomerDefinedValues", 1))
	}
	if len(s.DataSetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DataSetType"))
	}

	if s.DestinationS3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationS3BucketName"))
	}
	if s.DestinationS3BucketName != nil && len(*s.DestinationS3BucketName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationS3BucketName", 1))
	}

	if s.FromDate == nil {
		invalidParams.Add(aws.NewErrParamRequired("FromDate"))
	}

	if s.RoleNameArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleNameArn"))
	}
	if s.RoleNameArn != nil && len(*s.RoleNameArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleNameArn", 1))
	}

	if s.SnsTopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the result of the StartSupportDataExport operation.
type StartSupportDataExportOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier representing a specific request to the StartSupportDataExport
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation
func (s StartSupportDataExportOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartSupportDataExport = "StartSupportDataExport"

// StartSupportDataExportRequest returns a request value for making API operation for
// AWS Marketplace Commerce Analytics.
//
// Given a data set type and a from date, asynchronously publishes the requested
// customer support data to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD'T'HH-mm-ss'Z'.csv. If a file with the
// same name already exists (e.g. if the same data set is requested twice),
// the original file will be overwritten by the new file. Requires a Role with
// an attached permissions policy providing Allow permissions for the following
// actions: s3:PutObject, s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish,
// iam:GetRolePolicy.
//
//    // Example sending a request using StartSupportDataExportRequest.
//    req := client.StartSupportDataExportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/StartSupportDataExport
func (c *Client) StartSupportDataExportRequest(input *StartSupportDataExportInput) StartSupportDataExportRequest {
	op := &aws.Operation{
		Name:       opStartSupportDataExport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartSupportDataExportInput{}
	}

	req := c.newRequest(op, input, &StartSupportDataExportOutput{})
	return StartSupportDataExportRequest{Request: req, Input: input, Copy: c.StartSupportDataExportRequest}
}

// StartSupportDataExportRequest is the request type for the
// StartSupportDataExport API operation.
type StartSupportDataExportRequest struct {
	*aws.Request
	Input *StartSupportDataExportInput
	Copy  func(*StartSupportDataExportInput) StartSupportDataExportRequest
}

// Send marshals and sends the StartSupportDataExport API request.
func (r StartSupportDataExportRequest) Send(ctx context.Context) (*StartSupportDataExportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartSupportDataExportResponse{
		StartSupportDataExportOutput: r.Request.Data.(*StartSupportDataExportOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartSupportDataExportResponse is the response type for the
// StartSupportDataExport API operation.
type StartSupportDataExportResponse struct {
	*StartSupportDataExportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartSupportDataExport request.
func (r *StartSupportDataExportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
