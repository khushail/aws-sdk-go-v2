// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new, empty pipeline. Use PutPipelineDefinition to populate the
// pipeline.
func (c *Client) CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) {
	if params == nil {
		params = &CreatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePipeline", params, optFns, addOperationCreatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreatePipeline.
type CreatePipelineInput struct {

	// The name for the pipeline. You can use the same name for multiple pipelines
	// associated with your AWS account, because AWS Data Pipeline assigns each
	// pipeline a unique pipeline identifier.
	//
	// This member is required.
	Name *string

	// A unique identifier. This identifier is not the same as the pipeline identifier
	// assigned by AWS Data Pipeline. You are responsible for defining the format and
	// ensuring the uniqueness of this identifier. You use this parameter to ensure
	// idempotency during repeated calls to CreatePipeline. For example, if the first
	// call to CreatePipeline does not succeed, you can pass in the same unique
	// identifier and pipeline name combination on a subsequent call to CreatePipeline.
	// CreatePipeline ensures that if a pipeline already exists with the same name and
	// unique identifier, a new pipeline is not created. Instead, you'll receive the
	// pipeline identifier from the previous attempt. The uniqueness of the name and
	// unique identifier combination is scoped to the AWS account or IAM user
	// credentials.
	//
	// This member is required.
	UniqueId *string

	// The description for the pipeline.
	Description *string

	// A list of tags to associate with the pipeline at creation. Tags let you control
	// access to pipelines. For more information, see Controlling User Access to
	// Pipelines
	// (https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
	// in the AWS Data Pipeline Developer Guide.
	Tags []types.Tag
}

// Contains the output of CreatePipeline.
type CreatePipelineOutput struct {

	// The ID that AWS Data Pipeline assigns the newly created pipeline. For example,
	// df-06372391ZG65EXAMPLE.
	//
	// This member is required.
	PipelineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePipeline{}, middleware.After)
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
	if err = addOpCreatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "CreatePipeline",
	}
}