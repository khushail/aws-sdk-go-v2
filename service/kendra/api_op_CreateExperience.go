// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Kendra experience such as a search application. For more
// information on creating a search application experience, see Building a search
// experience with no code
// (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html).
func (c *Client) CreateExperience(ctx context.Context, params *CreateExperienceInput, optFns ...func(*Options)) (*CreateExperienceOutput, error) {
	if params == nil {
		params = &CreateExperienceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExperience", params, optFns, c.addOperationCreateExperienceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExperienceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExperienceInput struct {

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	// A name for your Amazon Kendra experience.
	//
	// This member is required.
	Name *string

	// A token that you provide to identify the request to create your Amazon Kendra
	// experience. Multiple calls to the CreateExperience API with the same client
	// token creates only one Amazon Kendra experience.
	ClientToken *string

	// Provides the configuration information for your Amazon Kendra experience. This
	// includes ContentSourceConfiguration, which specifies the data source IDs and/or
	// FAQ IDs, and UserIdentityConfiguration, which specifies the user or group
	// information to grant access to your Amazon Kendra experience.
	Configuration *types.ExperienceConfiguration

	// A description for your Amazon Kendra experience.
	Description *string

	// The Amazon Resource Name (ARN) of a role with permission to access Query API,
	// QuerySuggestions API, SubmitFeedback API, and Amazon Web Services SSO that
	// stores your user and group information. For more information, see IAM roles for
	// Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn *string

	noSmithyDocumentSerde
}

type CreateExperienceOutput struct {

	// The identifier for your created Amazon Kendra experience.
	//
	// This member is required.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateExperienceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateExperience{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateExperience{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateExperienceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateExperienceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExperience(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateExperience struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateExperience) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateExperience) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateExperienceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateExperienceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateExperienceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateExperience{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateExperience(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateExperience",
	}
}
