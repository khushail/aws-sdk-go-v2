// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the specified activity type. This includes
// configuration settings provided when the type was registered and other general
// information about the type. Access Control You can use IAM policies to control
// this action's access to Amazon SWF resources as follows:
//
// * Use a Resource
// element with the domain name to limit the action to only specified domains.
//
// *
// Use an Action element to allow or deny permission to call this action.
//
// *
// Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// * activityType.name: String constraint. The key is
// swf:activityType.name.
//
// * activityType.version: String constraint. The key is
// swf:activityType.version.
//
// If the caller doesn't have sufficient permissions to
// invoke the action, or the parameter values fall outside the specified
// constraints, the action fails. The associated event attribute's cause parameter
// is set to OPERATION_NOT_PERMITTED. For details and example IAM policies, see
// Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) DescribeActivityType(ctx context.Context, params *DescribeActivityTypeInput, optFns ...func(*Options)) (*DescribeActivityTypeOutput, error) {
	if params == nil {
		params = &DescribeActivityTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeActivityType", params, optFns, addOperationDescribeActivityTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeActivityTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActivityTypeInput struct {

	// The activity type to get information about. Activity types are identified by the
	// name and version that were supplied when the activity was registered.
	//
	// This member is required.
	ActivityType *types.ActivityType

	// The name of the domain in which the activity type is registered.
	//
	// This member is required.
	Domain *string
}

// Detailed information about an activity type.
type DescribeActivityTypeOutput struct {

	// The configuration settings registered with the activity type.
	//
	// This member is required.
	Configuration *types.ActivityTypeConfiguration

	// General information about the activity type. The status of activity type
	// (returned in the ActivityTypeInfo structure) can be one of the following.
	//
	// *
	// REGISTERED – The type is registered and available. Workers supporting this type
	// should be running.
	//
	// * DEPRECATED – The type was deprecated using
	// DeprecateActivityType, but is still in use. You should keep workers supporting
	// this type running. You cannot create new tasks of this type.
	//
	// This member is required.
	TypeInfo *types.ActivityTypeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeActivityTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeActivityType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeActivityType{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeActivityTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActivityType(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeActivityType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "DescribeActivityType",
	}
}