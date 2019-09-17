// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Contains information about an AWS account that is a member of an organization.
type Account struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the account.
	//
	// For more information about ARNs in Organizations, see ARN Formats Supported
	// by Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_permissions.html#orgs-permissions-arns)
	// in the AWS Organizations User Guide.
	Arn *string `type:"string"`

	// The email address associated with the AWS account.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for this parameter is
	// a string of characters that represents a standard Internet email address.
	Email *string `min:"6" type:"string"`

	// The unique identifier (ID) of the account.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an account ID string
	// requires exactly 12 digits.
	Id *string `type:"string"`

	// The method by which the account joined the organization.
	JoinedMethod AccountJoinedMethod `type:"string" enum:"true"`

	// The date the account became a part of the organization.
	JoinedTimestamp *time.Time `type:"timestamp"`

	// The friendly name of the account.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of any of the characters in the ASCII character
	// range.
	Name *string `min:"1" type:"string"`

	// The status of the account in the organization.
	Status AccountStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s Account) String() string {
	return awsutil.Prettify(s)
}

// Contains a list of child entities, either OUs or accounts.
type Child struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of this child entity.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a child ID string
	// requires one of the following:
	//
	//    * Account: a string that consists of exactly 12 digits.
	//
	//    * Organizational unit (OU): a string that begins with "ou-" followed by
	//    from 4 to 32 lower-case letters or digits (the ID of the root that contains
	//    the OU) followed by a second "-" dash and from 8 to 32 additional lower-case
	//    letters or digits.
	Id *string `type:"string"`

	// The type of this child entity.
	Type ChildType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Child) String() string {
	return awsutil.Prettify(s)
}

// Contains the status about a CreateAccount or CreateGovCloudAccount request
// to create an AWS account or an AWS GovCloud (US) account in an organization.
type CreateAccountStatus struct {
	_ struct{} `type:"structure"`

	// If the account was created successfully, the unique identifier (ID) of the
	// new account.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an account ID string
	// requires exactly 12 digits.
	AccountId *string `type:"string"`

	// The account name given to the account when it was created.
	AccountName *string `min:"1" type:"string"`

	// The date and time that the account was created and the request completed.
	CompletedTimestamp *time.Time `type:"timestamp"`

	// If the request failed, a description of the reason for the failure.
	//
	//    * ACCOUNT_LIMIT_EXCEEDED: The account could not be created because you
	//    have reached the limit on the number of accounts in your organization.
	//
	//    * EMAIL_ALREADY_EXISTS: The account could not be created because another
	//    AWS account with that email address already exists.
	//
	//    * INVALID_ADDRESS: The account could not be created because the address
	//    you provided is not valid.
	//
	//    * INVALID_EMAIL: The account could not be created because the email address
	//    you provided is not valid.
	//
	//    * INTERNAL_FAILURE: The account could not be created because of an internal
	//    failure. Try again later. If the problem persists, contact Customer Support.
	FailureReason CreateAccountFailureReason `type:"string" enum:"true"`

	// If the account was created successfully, the unique identifier (ID) of the
	// new account in the AWS GovCloud (US) Region.
	GovCloudAccountId *string `type:"string"`

	// The unique identifier (ID) that references this request. You get this value
	// from the response of the initial CreateAccount request to create the account.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an create account
	// request ID string requires "car-" followed by from 8 to 32 lower-case letters
	// or digits.
	Id *string `type:"string"`

	// The date and time that the request was made for the account creation.
	RequestedTimestamp *time.Time `type:"timestamp"`

	// The status of the request.
	State CreateAccountState `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateAccountStatus) String() string {
	return awsutil.Prettify(s)
}

// A structure that contains details of a service principal that is enabled
// to integrate with AWS Organizations.
type EnabledServicePrincipal struct {
	_ struct{} `type:"structure"`

	// The date that the service principal was enabled for integration with AWS
	// Organizations.
	DateEnabled *time.Time `type:"timestamp"`

	// The name of the service principal. This is typically in the form of a URL,
	// such as: servicename.amazonaws.com.
	ServicePrincipal *string `min:"1" type:"string"`
}

// String returns the string representation
func (s EnabledServicePrincipal) String() string {
	return awsutil.Prettify(s)
}

// Contains information that must be exchanged to securely establish a relationship
// between two accounts (an originator and a recipient). For example, when a
// master account (the originator) invites another account (the recipient) to
// join its organization, the two accounts exchange information as a series
// of handshake requests and responses.
//
// Note: Handshakes that are CANCELED, ACCEPTED, or DECLINED show up in lists
// for only 30 days after entering that state After that they are deleted.
type Handshake struct {
	_ struct{} `type:"structure"`

	// The type of handshake, indicating what action occurs when the recipient accepts
	// the handshake. The following handshake types are supported:
	//
	//    * INVITE: This type of handshake represents a request to join an organization.
	//    It is always sent from the master account to only non-member accounts.
	//
	//    * ENABLE_ALL_FEATURES: This type of handshake represents a request to
	//    enable all features in an organization. It is always sent from the master
	//    account to only invited member accounts. Created accounts do not receive
	//    this because those accounts were created by the organization's master
	//    account and approval is inferred.
	//
	//    * APPROVE_ALL_FEATURES: This type of handshake is sent from the Organizations
	//    service when all member accounts have approved the ENABLE_ALL_FEATURES
	//    invitation. It is sent only to the master account and signals the master
	//    that it can finalize the process to enable all features.
	Action ActionType `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of a handshake.
	//
	// For more information about ARNs in Organizations, see ARN Formats Supported
	// by Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_permissions.html#orgs-permissions-arns)
	// in the AWS Organizations User Guide.
	Arn *string `type:"string"`

	// The date and time that the handshake expires. If the recipient of the handshake
	// request fails to respond before the specified date and time, the handshake
	// becomes inactive and is no longer valid.
	ExpirationTimestamp *time.Time `type:"timestamp"`

	// The unique identifier (ID) of a handshake. The originating account creates
	// the ID when it initiates the handshake.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for handshake ID string
	// requires "h-" followed by from 8 to 32 lower-case letters or digits.
	Id *string `type:"string"`

	// Information about the two accounts that are participating in the handshake.
	Parties []HandshakeParty `type:"list"`

	// The date and time that the handshake request was made.
	RequestedTimestamp *time.Time `type:"timestamp"`

	// Additional information that is needed to process the handshake.
	Resources []HandshakeResource `type:"list"`

	// The current state of the handshake. Use the state to trace the flow of the
	// handshake through the process from its creation to its acceptance. The meaning
	// of each of the valid values is as follows:
	//
	//    * REQUESTED: This handshake was sent to multiple recipients (applicable
	//    to only some handshake types) and not all recipients have responded yet.
	//    The request stays in this state until all recipients respond.
	//
	//    * OPEN: This handshake was sent to multiple recipients (applicable to
	//    only some policy types) and all recipients have responded, allowing the
	//    originator to complete the handshake action.
	//
	//    * CANCELED: This handshake is no longer active because it was canceled
	//    by the originating account.
	//
	//    * ACCEPTED: This handshake is complete because it has been accepted by
	//    the recipient.
	//
	//    * DECLINED: This handshake is no longer active because it was declined
	//    by the recipient account.
	//
	//    * EXPIRED: This handshake is no longer active because the originator did
	//    not receive a response of any kind from the recipient before the expiration
	//    time (15 days).
	State HandshakeState `type:"string" enum:"true"`
}

// String returns the string representation
func (s Handshake) String() string {
	return awsutil.Prettify(s)
}

// Specifies the criteria that are used to select the handshakes for the operation.
type HandshakeFilter struct {
	_ struct{} `type:"structure"`

	// Specifies the type of handshake action.
	//
	// If you specify ActionType, you cannot also specify ParentHandshakeId.
	ActionType ActionType `type:"string" enum:"true"`

	// Specifies the parent handshake. Only used for handshake types that are a
	// child of another type.
	//
	// If you specify ParentHandshakeId, you cannot also specify ActionType.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for handshake ID string
	// requires "h-" followed by from 8 to 32 lower-case letters or digits.
	ParentHandshakeId *string `type:"string"`
}

// String returns the string representation
func (s HandshakeFilter) String() string {
	return awsutil.Prettify(s)
}

// Identifies a participant in a handshake.
type HandshakeParty struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) for the party.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for handshake ID string
	// requires "h-" followed by from 8 to 32 lower-case letters or digits.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The type of party.
	//
	// Type is a required field
	Type HandshakePartyType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s HandshakeParty) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HandshakeParty) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HandshakeParty"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains additional data that is needed to process a handshake.
type HandshakeResource struct {
	_ struct{} `type:"structure"`

	// When needed, contains an additional array of HandshakeResource objects.
	Resources []HandshakeResource `type:"list"`

	// The type of information being passed, specifying how the value is to be interpreted
	// by the other party:
	//
	//    * ACCOUNT - Specifies an AWS account ID number.
	//
	//    * ORGANIZATION - Specifies an organization ID number.
	//
	//    * EMAIL - Specifies the email address that is associated with the account
	//    that receives the handshake.
	//
	//    * OWNER_EMAIL - Specifies the email address associated with the master
	//    account. Included as information about an organization.
	//
	//    * OWNER_NAME - Specifies the name associated with the master account.
	//    Included as information about an organization.
	//
	//    * NOTES - Additional text provided by the handshake initiator and intended
	//    for the recipient to read.
	Type HandshakeResourceType `type:"string" enum:"true"`

	// The information that is passed to the other party in the handshake. The format
	// of the value string must match the requirements of the specified type.
	Value *string `type:"string"`
}

// String returns the string representation
func (s HandshakeResource) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an organization. An organization is a collection of
// accounts that are centrally managed together using consolidated billing,
// organized hierarchically with organizational units (OUs), and controlled
// with policies .
type Organization struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of an organization.
	//
	// For more information about ARNs in Organizations, see ARN Formats Supported
	// by Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_permissions.html#orgs-permissions-arns)
	// in the AWS Organizations User Guide.
	Arn *string `type:"string"`

	// A list of policy types that are enabled for this organization. For example,
	// if your organization has all features enabled, then service control policies
	// (SCPs) are included in the list.
	//
	// Even if a policy type is shown as available in the organization, you can
	// separately enable and disable them at the root level by using EnablePolicyType
	// and DisablePolicyType. Use ListRoots to see the status of a policy type in
	// that root.
	AvailablePolicyTypes []PolicyTypeSummary `type:"list"`

	// Specifies the functionality that currently is available to the organization.
	// If set to "ALL", then all features are enabled and policies can be applied
	// to accounts in the organization. If set to "CONSOLIDATED_BILLING", then only
	// consolidated billing functionality is available. For more information, see
	// Enabling All Features in Your Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
	// in the AWS Organizations User Guide.
	FeatureSet OrganizationFeatureSet `type:"string" enum:"true"`

	// The unique identifier (ID) of an organization.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an organization ID
	// string requires "o-" followed by from 10 to 32 lower-case letters or digits.
	Id *string `type:"string"`

	// The Amazon Resource Name (ARN) of the account that is designated as the master
	// account for the organization.
	//
	// For more information about ARNs in Organizations, see ARN Formats Supported
	// by Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_permissions.html#orgs-permissions-arns)
	// in the AWS Organizations User Guide.
	MasterAccountArn *string `type:"string"`

	// The email address that is associated with the AWS account that is designated
	// as the master account for the organization.
	MasterAccountEmail *string `min:"6" type:"string"`

	// The unique identifier (ID) of the master account of an organization.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an account ID string
	// requires exactly 12 digits.
	MasterAccountId *string `type:"string"`
}

// String returns the string representation
func (s Organization) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an organizational unit (OU). An OU is a container
// of AWS accounts within a root of an organization. Policies that are attached
// to an OU apply to all accounts contained in that OU and in any child OUs.
type OrganizationalUnit struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of this OU.
	//
	// For more information about ARNs in Organizations, see ARN Formats Supported
	// by Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_permissions.html#orgs-permissions-arns)
	// in the AWS Organizations User Guide.
	Arn *string `type:"string"`

	// The unique identifier (ID) associated with this OU.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an organizational
	// unit ID string requires "ou-" followed by from 4 to 32 lower-case letters
	// or digits (the ID of the root that contains the OU) followed by a second
	// "-" dash and from 8 to 32 additional lower-case letters or digits.
	Id *string `type:"string"`

	// The friendly name of this OU.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of any of the characters in the ASCII character
	// range.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s OrganizationalUnit) String() string {
	return awsutil.Prettify(s)
}

// Contains information about either a root or an organizational unit (OU) that
// can contain OUs or accounts in an organization.
type Parent struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the parent entity.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	//    * Root: a string that begins with "r-" followed by from 4 to 32 lower-case
	//    letters or digits.
	//
	//    * Organizational unit (OU): a string that begins with "ou-" followed by
	//    from 4 to 32 lower-case letters or digits (the ID of the root that the
	//    OU is in) followed by a second "-" dash and from 8 to 32 additional lower-case
	//    letters or digits.
	Id *string `type:"string"`

	// The type of the parent entity.
	Type ParentType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Parent) String() string {
	return awsutil.Prettify(s)
}

// Contains rules to be applied to the affected accounts. Policies can be attached
// directly to accounts, or to roots and OUs to affect all accounts in those
// hierarchies.
type Policy struct {
	_ struct{} `type:"structure"`

	// The text content of the policy.
	Content *string `min:"1" type:"string"`

	// A structure that contains additional details about the policy.
	PolicySummary *PolicySummary `type:"structure"`
}

// String returns the string representation
func (s Policy) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a policy, but does not include the content. To
// see the content of a policy, see DescribePolicy.
type PolicySummary struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the policy.
	//
	// For more information about ARNs in Organizations, see ARN Formats Supported
	// by Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_permissions.html#orgs-permissions-arns)
	// in the AWS Organizations User Guide.
	Arn *string `type:"string"`

	// A boolean value that indicates whether the specified policy is an AWS managed
	// policy. If true, then you can attach the policy to roots, OUs, or accounts,
	// but you cannot edit it.
	AwsManaged *bool `type:"boolean"`

	// The description of the policy.
	Description *string `type:"string"`

	// The unique identifier (ID) of the policy.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string
	// requires "p-" followed by from 8 to 128 lower-case letters or digits.
	Id *string `type:"string"`

	// The friendly name of the policy.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of any of the characters in the ASCII character
	// range.
	Name *string `min:"1" type:"string"`

	// The type of policy.
	Type PolicyType `type:"string" enum:"true"`
}

// String returns the string representation
func (s PolicySummary) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a root, OU, or account that a policy is attached
// to.
type PolicyTargetSummary struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the policy target.
	//
	// For more information about ARNs in Organizations, see ARN Formats Supported
	// by Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_permissions.html#orgs-permissions-arns)
	// in the AWS Organizations User Guide.
	Arn *string `type:"string"`

	// The friendly name of the policy target.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of any of the characters in the ASCII character
	// range.
	Name *string `min:"1" type:"string"`

	// The unique identifier (ID) of the policy target.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a target ID string
	// requires one of the following:
	//
	//    * Root: a string that begins with "r-" followed by from 4 to 32 lower-case
	//    letters or digits.
	//
	//    * Account: a string that consists of exactly 12 digits.
	//
	//    * Organizational unit (OU): a string that begins with "ou-" followed by
	//    from 4 to 32 lower-case letters or digits (the ID of the root that the
	//    OU is in) followed by a second "-" dash and from 8 to 32 additional lower-case
	//    letters or digits.
	TargetId *string `type:"string"`

	// The type of the policy target.
	Type TargetType `type:"string" enum:"true"`
}

// String returns the string representation
func (s PolicyTargetSummary) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a policy type and its status in the associated
// root.
type PolicyTypeSummary struct {
	_ struct{} `type:"structure"`

	// The status of the policy type as it relates to the associated root. To attach
	// a policy of the specified type to a root or to an OU or account in that root,
	// it must be available in the organization and enabled for that root.
	Status PolicyTypeStatus `type:"string" enum:"true"`

	// The name of the policy type.
	Type PolicyType `type:"string" enum:"true"`
}

// String returns the string representation
func (s PolicyTypeSummary) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a root. A root is a top-level parent node in the hierarchy
// of an organization that can contain organizational units (OUs) and accounts.
// Every root contains every AWS account in the organization. Each root enables
// the accounts to be organized in a different way and to have different policy
// types enabled for use in that root.
type Root struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the root.
	//
	// For more information about ARNs in Organizations, see ARN Formats Supported
	// by Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_permissions.html#orgs-permissions-arns)
	// in the AWS Organizations User Guide.
	Arn *string `type:"string"`

	// The unique identifier (ID) for the root.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a root ID string
	// requires "r-" followed by from 4 to 32 lower-case letters or digits.
	Id *string `type:"string"`

	// The friendly name of the root.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of any of the characters in the ASCII character
	// range.
	Name *string `min:"1" type:"string"`

	// The types of policies that are currently enabled for the root and therefore
	// can be attached to the root or to its OUs or accounts.
	//
	// Even if a policy type is shown as available in the organization, you can
	// separately enable and disable them at the root level by using EnablePolicyType
	// and DisablePolicyType. Use DescribeOrganization to see the availability of
	// the policy types in that organization.
	PolicyTypes []PolicyTypeSummary `type:"list"`
}

// String returns the string representation
func (s Root) String() string {
	return awsutil.Prettify(s)
}

// A custom key-value pair associated with a resource such as an account within
// your organization.
type Tag struct {
	_ struct{} `type:"structure"`

	// The key identifier, or name, of the tag.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The string value that's associated with the key of the tag. You can set the
	// value of a tag to an empty string, but you can't set the value of a tag to
	// null.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
