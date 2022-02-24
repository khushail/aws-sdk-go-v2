// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a specified stack. Once the call completes successfully, stack deletion
// starts. Deleted stacks don't show up in the DescribeStacks operation if the
// deletion has been completed successfully.
func (c *Client) DeleteStack(ctx context.Context, params *DeleteStackInput, optFns ...func(*Options)) (*DeleteStackOutput, error) {
	if params == nil {
		params = &DeleteStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStack", params, optFns, c.addOperationDeleteStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DeleteStack action.
type DeleteStackInput struct {

	// The name or the unique stack ID that's associated with the stack.
	//
	// This member is required.
	StackName *string

	// A unique identifier for this DeleteStack request. Specify this token if you plan
	// to retry requests so that CloudFormation knows that you're not attempting to
	// delete a stack with the same name. You might retry DeleteStack requests to
	// ensure that CloudFormation successfully received them. All events initiated by a
	// given stack operation are assigned the same client request token, which you can
	// use to track operations. For example, if you execute a CreateStack operation
	// with the token token1, then all the StackEvents generated by that operation will
	// have ClientRequestToken set as token1. In the console, stack operations display
	// the client request token on the Events tab. Stack operations that are initiated
	// from the console use the token format Console-StackOperation-ID, which helps you
	// easily identify the stack operation . For example, if you create a stack using
	// the console, each stack event would be assigned the same token in the following
	// format: Console-CreateStack-7f59c3cf-00d2-40c7-b2ff-e75db0987002.
	ClientRequestToken *string

	// For stacks in the DELETE_FAILED state, a list of resource logical IDs that are
	// associated with the resources you want to retain. During deletion,
	// CloudFormation deletes the stack but doesn't delete the retained resources.
	// Retaining resources is useful when you can't delete a resource, such as a
	// non-empty S3 bucket, but you want to delete the stack.
	RetainResources []string

	// The Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role
	// that CloudFormation assumes to delete the stack. CloudFormation uses the role's
	// credentials to make calls on your behalf. If you don't specify a value,
	// CloudFormation uses the role that was previously associated with the stack. If
	// no role is available, CloudFormation uses a temporary session that's generated
	// from your user credentials.
	RoleARN *string

	noSmithyDocumentSerde
}

type DeleteStackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteStack{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpDeleteStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeleteStack",
	}
}
