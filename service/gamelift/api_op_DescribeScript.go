// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves properties for a Realtime script. To request a script record, specify
// the script ID. If successful, an object containing the script properties is
// returned. Learn more Amazon Web Services Realtime Servers
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html)
// Related actions CreateScript | ListScripts | DescribeScript | UpdateScript |
// DeleteScript | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DescribeScript(ctx context.Context, params *DescribeScriptInput, optFns ...func(*Options)) (*DescribeScriptOutput, error) {
	if params == nil {
		params = &DescribeScriptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScript", params, optFns, c.addOperationDescribeScriptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScriptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScriptInput struct {

	// A unique identifier for the Realtime script to retrieve properties for. You can
	// use either the script ID or ARN value.
	//
	// This member is required.
	ScriptId *string

	noSmithyDocumentSerde
}

type DescribeScriptOutput struct {

	// A set of properties describing the requested script.
	Script *types.Script

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScriptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeScript{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeScript{}, middleware.After)
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
	if err = addOpDescribeScriptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScript(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeScript(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeScript",
	}
}
