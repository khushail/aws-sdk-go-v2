// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The TestRole operation tests the IAM role used to create the pipeline. The
// TestRole action lets you determine whether the IAM role you are using has
// sufficient permissions to let Elastic Transcoder perform tasks associated with
// the transcoding process. The action attempts to assume the specified IAM role,
// checks read access to the input and output buckets, and tries to send a test
// notification to Amazon SNS topics that you specify.
//
// Deprecated: This operation has been deprecated.
func (c *Client) TestRole(ctx context.Context, params *TestRoleInput, optFns ...func(*Options)) (*TestRoleOutput, error) {
	if params == nil {
		params = &TestRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestRole", params, optFns, addOperationTestRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The TestRoleRequest structure.
type TestRoleInput struct {

	// The Amazon S3 bucket that contains media files to be transcoded. The action
	// attempts to read from this bucket.
	//
	// This member is required.
	InputBucket *string

	// The Amazon S3 bucket that Elastic Transcoder writes transcoded media files to.
	// The action attempts to read from this bucket.
	//
	// This member is required.
	OutputBucket *string

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder
	// to test.
	//
	// This member is required.
	Role *string

	// The ARNs of one or more Amazon Simple Notification Service (Amazon SNS) topics
	// that you want the action to send a test notification to.
	//
	// This member is required.
	Topics []string
}

// The TestRoleResponse structure.
type TestRoleOutput struct {

	// If the Success element contains false, this value is an array of one or more
	// error messages that were generated during the test process.
	Messages []string

	// If the operation is successful, this value is true; otherwise, the value is
	// false.
	Success *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTestRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestRole{}, middleware.After)
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
	if err = addOpTestRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "TestRole",
	}
}