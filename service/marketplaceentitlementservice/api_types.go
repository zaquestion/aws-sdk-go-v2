// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplaceentitlementservice

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// An entitlement represents capacity in a product owned by the customer. For
// example, a customer might own some number of users or seats in an SaaS application
// or some amount of data capacity in a multi-tenant database.
type Entitlement struct {
	_ struct{} `type:"structure"`

	// The customer identifier is a handle to each unique customer in an application.
	// Customer identifiers are obtained through the ResolveCustomer operation in
	// AWS Marketplace Metering Service.
	CustomerIdentifier *string `type:"string"`

	// The dimension for which the given entitlement applies. Dimensions represent
	// categories of capacity in a product and are specified when the product is
	// listed in AWS Marketplace.
	Dimension *string `type:"string"`

	// The expiration date represents the minimum date through which this entitlement
	// is expected to remain valid. For contractual products listed on AWS Marketplace,
	// the expiration date is the date at which the customer will renew or cancel
	// their contract. Customers who are opting to renew their contract will still
	// have entitlements with an expiration date.
	ExpirationDate *time.Time `type:"timestamp"`

	// The product code for which the given entitlement applies. Product codes are
	// provided by AWS Marketplace when the product listing is created.
	ProductCode *string `min:"1" type:"string"`

	// The EntitlementValue represents the amount of capacity that the customer
	// is entitled to for the product.
	Value *EntitlementValue `type:"structure"`
}

// String returns the string representation
func (s Entitlement) String() string {
	return awsutil.Prettify(s)
}

// The EntitlementValue represents the amount of capacity that the customer
// is entitled to for the product.
type EntitlementValue struct {
	_ struct{} `type:"structure"`

	// The BooleanValue field will be populated with a boolean value when the entitlement
	// is a boolean type. Otherwise, the field will not be set.
	BooleanValue *bool `type:"boolean"`

	// The DoubleValue field will be populated with a double value when the entitlement
	// is a double type. Otherwise, the field will not be set.
	DoubleValue *float64 `type:"double"`

	// The IntegerValue field will be populated with an integer value when the entitlement
	// is an integer type. Otherwise, the field will not be set.
	IntegerValue *int64 `type:"integer"`

	// The StringValue field will be populated with a string value when the entitlement
	// is a string type. Otherwise, the field will not be set.
	StringValue *string `type:"string"`
}

// String returns the string representation
func (s EntitlementValue) String() string {
	return awsutil.Prettify(s)
}
