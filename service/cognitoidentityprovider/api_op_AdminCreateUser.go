// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to create a user in the specified user pool.
type AdminCreateUserInput struct {
	_ struct{} `type:"structure"`

	// Specify "EMAIL" if email will be used to send the welcome message. Specify
	// "SMS" if the phone number will be used. The default value is "SMS". More
	// than one value can be specified.
	DesiredDeliveryMediums []DeliveryMediumType `type:"list"`

	// This parameter is only used if the phone_number_verified or email_verified
	// attribute is set to True. Otherwise, it is ignored.
	//
	// If this parameter is set to True and the phone number or email address specified
	// in the UserAttributes parameter already exists as an alias with a different
	// user, the API call will migrate the alias from the previous user to the newly
	// created user. The previous user will no longer be able to log in using that
	// alias.
	//
	// If this parameter is set to False, the API throws an AliasExistsException
	// error if the alias already exists. The default value is False.
	ForceAliasCreation *bool `type:"boolean"`

	// Set to "RESEND" to resend the invitation message to a user that already exists
	// and reset the expiration limit on the user's account. Set to "SUPPRESS" to
	// suppress sending the message. Only one value can be specified.
	MessageAction MessageActionType `type:"string" enum:"true"`

	// The user's temporary password. This password must conform to the password
	// policy that you specified when you created the user pool.
	//
	// The temporary password is valid only once. To complete the Admin Create User
	// flow, the user must enter the temporary password in the sign-in page along
	// with a new password to be used in all future sign-ins.
	//
	// This parameter is not required. If you do not specify a value, Amazon Cognito
	// generates one for you.
	//
	// The temporary password can only be used until the user account expiration
	// limit that you specified when you created the user pool. To reset the account
	// after that time limit, you must call AdminCreateUser again, specifying "RESEND"
	// for the MessageAction parameter.
	TemporaryPassword *string `min:"6" type:"string"`

	// An array of name-value pairs that contain user attributes and attribute values
	// to be set for the user to be created. You can create a user without specifying
	// any attributes other than Username. However, any attributes that you specify
	// as required (in or in the Attributes tab of the console) must be supplied
	// either by you (in your call to AdminCreateUser) or by the user (when he or
	// she signs up in response to your welcome message).
	//
	// For custom attributes, you must prepend the custom: prefix to the attribute
	// name.
	//
	// To send a message inviting the user to sign up, you must specify the user's
	// email address or phone number. This can be done in your call to AdminCreateUser
	// or in the Users tab of the Amazon Cognito console for managing your user
	// pools.
	//
	// In your call to AdminCreateUser, you can set the email_verified attribute
	// to True, and you can set the phone_number_verified attribute to True. (You
	// can also do this by calling .)
	//
	//    * email: The email address of the user to whom the message that contains
	//    the code and username will be sent. Required if the email_verified attribute
	//    is set to True, or if "EMAIL" is specified in the DesiredDeliveryMediums
	//    parameter.
	//
	//    * phone_number: The phone number of the user to whom the message that
	//    contains the code and username will be sent. Required if the phone_number_verified
	//    attribute is set to True, or if "SMS" is specified in the DesiredDeliveryMediums
	//    parameter.
	UserAttributes []AttributeType `type:"list"`

	// The user pool ID for the user pool where the user will be created.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The username for the user. Must be unique within the user pool. Must be a
	// UTF-8 string between 1 and 128 characters. After the user is created, the
	// username cannot be changed.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true"`

	// The user's validation data. This is an array of name-value pairs that contain
	// user attributes and attribute values that you can use for custom validation,
	// such as restricting the types of user accounts that can be registered. For
	// example, you might choose to allow or disallow user sign-up based on the
	// user's domain.
	//
	// To configure custom validation, you must create a Pre Sign-up Lambda trigger
	// for the user pool as described in the Amazon Cognito Developer Guide. The
	// Lambda trigger receives the validation data and uses it in the validation
	// process.
	//
	// The user's validation data is not persisted.
	ValidationData []AttributeType `type:"list"`
}

// String returns the string representation
func (s AdminCreateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminCreateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminCreateUserInput"}
	if s.TemporaryPassword != nil && len(*s.TemporaryPassword) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("TemporaryPassword", 6))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}
	if s.UserAttributes != nil {
		for i, v := range s.UserAttributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UserAttributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ValidationData != nil {
		for i, v := range s.ValidationData {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ValidationData", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server to the request to create the user.
type AdminCreateUserOutput struct {
	_ struct{} `type:"structure"`

	// The newly created user.
	User *UserType `type:"structure"`
}

// String returns the string representation
func (s AdminCreateUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminCreateUser = "AdminCreateUser"

// AdminCreateUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Creates a new user in the specified user pool.
//
// If MessageAction is not set, the default is to send a welcome message via
// email or phone (SMS).
//
// This message is based on a template that you configured in your call to or
// . This template includes your custom sign-up instructions and placeholders
// for user name and temporary password.
//
// Alternatively, you can call AdminCreateUser with “SUPPRESS” for the MessageAction
// parameter, and Amazon Cognito will not send any email.
//
// In either case, the user will be in the FORCE_CHANGE_PASSWORD state until
// they sign in and change their password.
//
// AdminCreateUser requires developer credentials.
//
//    // Example sending a request using AdminCreateUserRequest.
//    req := client.AdminCreateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminCreateUser
func (c *Client) AdminCreateUserRequest(input *AdminCreateUserInput) AdminCreateUserRequest {
	op := &aws.Operation{
		Name:       opAdminCreateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminCreateUserInput{}
	}

	req := c.newRequest(op, input, &AdminCreateUserOutput{})
	return AdminCreateUserRequest{Request: req, Input: input, Copy: c.AdminCreateUserRequest}
}

// AdminCreateUserRequest is the request type for the
// AdminCreateUser API operation.
type AdminCreateUserRequest struct {
	*aws.Request
	Input *AdminCreateUserInput
	Copy  func(*AdminCreateUserInput) AdminCreateUserRequest
}

// Send marshals and sends the AdminCreateUser API request.
func (r AdminCreateUserRequest) Send(ctx context.Context) (*AdminCreateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminCreateUserResponse{
		AdminCreateUserOutput: r.Request.Data.(*AdminCreateUserOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminCreateUserResponse is the response type for the
// AdminCreateUser API operation.
type AdminCreateUserResponse struct {
	*AdminCreateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminCreateUser request.
func (r *AdminCreateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
