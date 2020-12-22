// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a bulk thing provisioning task.
func (c *Client) StartThingRegistrationTask(ctx context.Context, params *StartThingRegistrationTaskInput, optFns ...func(*Options)) (*StartThingRegistrationTaskOutput, error) {
	if params == nil {
		params = &StartThingRegistrationTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartThingRegistrationTask", params, optFns, addOperationStartThingRegistrationTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartThingRegistrationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartThingRegistrationTaskInput struct {

	// The S3 bucket that contains the input file.
	//
	// This member is required.
	InputFileBucket *string

	// The name of input file within the S3 bucket. This file contains a newline
	// delimited JSON file. Each line contains the parameter values to provision one
	// device (thing).
	//
	// This member is required.
	InputFileKey *string

	// The IAM role ARN that grants permission the input file.
	//
	// This member is required.
	RoleArn *string

	// The provisioning template.
	//
	// This member is required.
	TemplateBody *string
}

type StartThingRegistrationTaskOutput struct {

	// The bulk thing provisioning task ID.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartThingRegistrationTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartThingRegistrationTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartThingRegistrationTask{}, middleware.After)
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
	if err = addOpStartThingRegistrationTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartThingRegistrationTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartThingRegistrationTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "StartThingRegistrationTask",
	}
}