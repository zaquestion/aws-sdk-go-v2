// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateDeployment operation.
type CreateDeploymentInput struct {
	_ struct{} `type:"structure"`

	// The name of an AWS CodeDeploy application associated with the IAM user or
	// AWS account.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`

	// Configuration information for an automatic rollback that is added when a
	// deployment is created.
	AutoRollbackConfiguration *AutoRollbackConfiguration `locationName:"autoRollbackConfiguration" type:"structure"`

	// The name of a deployment configuration associated with the IAM user or AWS
	// account.
	//
	// If not specified, the value configured in the deployment group is used as
	// the default. If the deployment group does not have a deployment configuration
	// associated with it, CodeDeployDefault.OneAtATime is used by default.
	DeploymentConfigName *string `locationName:"deploymentConfigName" min:"1" type:"string"`

	// The name of the deployment group.
	DeploymentGroupName *string `locationName:"deploymentGroupName" min:"1" type:"string"`

	// A comment about the deployment.
	Description *string `locationName:"description" type:"string"`

	// Information about how AWS CodeDeploy handles files that already exist in
	// a deployment target location but weren't part of the previous successful
	// deployment.
	//
	// The fileExistsBehavior parameter takes any of the following values:
	//
	//    * DISALLOW: The deployment fails. This is also the default behavior if
	//    no option is specified.
	//
	//    * OVERWRITE: The version of the file from the application revision currently
	//    being deployed replaces the version already on the instance.
	//
	//    * RETAIN: The version of the file already on the instance is kept and
	//    used as part of the new deployment.
	FileExistsBehavior FileExistsBehavior `locationName:"fileExistsBehavior" type:"string" enum:"true"`

	// If true, then if an ApplicationStop, BeforeBlockTraffic, or AfterBlockTraffic
	// deployment lifecycle event to an instance fails, then the deployment continues
	// to the next deployment lifecycle event. For example, if ApplicationStop fails,
	// the deployment continues with DownloadBundle. If BeforeBlockTraffic fails,
	// the deployment continues with BlockTraffic. If AfterBlockTraffic fails, the
	// deployment continues with ApplicationStop.
	//
	// If false or not specified, then if a lifecycle event fails during a deployment
	// to an instance, that deployment fails. If deployment to that instance is
	// part of an overall deployment and the number of healthy hosts is not less
	// than the minimum number of healthy hosts, then a deployment to the next instance
	// is attempted.
	//
	// During a deployment, the AWS CodeDeploy agent runs the scripts specified
	// for ApplicationStop, BeforeBlockTraffic, and AfterBlockTraffic in the AppSpec
	// file from the previous successful deployment. (All other scripts are run
	// from the AppSpec file in the current deployment.) If one of these scripts
	// contains an error and does not run successfully, the deployment can fail.
	//
	// If the cause of the failure is a script from the last successful deployment
	// that will never run successfully, create a new deployment and use ignoreApplicationStopFailures
	// to specify that the ApplicationStop, BeforeBlockTraffic, and AfterBlockTraffic
	// failures should be ignored.
	IgnoreApplicationStopFailures *bool `locationName:"ignoreApplicationStopFailures" type:"boolean"`

	// The type and location of the revision to deploy.
	Revision *RevisionLocation `locationName:"revision" type:"structure"`

	// Information about the instances that belong to the replacement environment
	// in a blue/green deployment.
	TargetInstances *TargetInstances `locationName:"targetInstances" type:"structure"`

	// Indicates whether to deploy to all instances or only to instances that are
	// not running the latest application revision.
	UpdateOutdatedInstancesOnly *bool `locationName:"updateOutdatedInstancesOnly" type:"boolean"`
}

// String returns the string representation
func (s CreateDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.DeploymentConfigName != nil && len(*s.DeploymentConfigName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentConfigName", 1))
	}
	if s.DeploymentGroupName != nil && len(*s.DeploymentGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateDeployment operation.
type CreateDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a deployment.
	DeploymentId *string `locationName:"deploymentId" type:"string"`
}

// String returns the string representation
func (s CreateDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDeployment = "CreateDeployment"

// CreateDeploymentRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Deploys an application revision through the specified deployment group.
//
//    // Example sending a request using CreateDeploymentRequest.
//    req := client.CreateDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/CreateDeployment
func (c *Client) CreateDeploymentRequest(input *CreateDeploymentInput) CreateDeploymentRequest {
	op := &aws.Operation{
		Name:       opCreateDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDeploymentInput{}
	}

	req := c.newRequest(op, input, &CreateDeploymentOutput{})
	return CreateDeploymentRequest{Request: req, Input: input, Copy: c.CreateDeploymentRequest}
}

// CreateDeploymentRequest is the request type for the
// CreateDeployment API operation.
type CreateDeploymentRequest struct {
	*aws.Request
	Input *CreateDeploymentInput
	Copy  func(*CreateDeploymentInput) CreateDeploymentRequest
}

// Send marshals and sends the CreateDeployment API request.
func (r CreateDeploymentRequest) Send(ctx context.Context) (*CreateDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentResponse{
		CreateDeploymentOutput: r.Request.Data.(*CreateDeploymentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentResponse is the response type for the
// CreateDeployment API operation.
type CreateDeploymentResponse struct {
	*CreateDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeployment request.
func (r *CreateDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
