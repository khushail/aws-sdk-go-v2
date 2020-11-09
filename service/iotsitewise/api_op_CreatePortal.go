// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a portal, which can contain projects and dashboards. AWS IoT SiteWise
// Monitor uses AWS SSO or IAM to authenticate portal users and manage user
// permissions. Before you can sign in to a new portal, you must add at least one
// identity to that portal. For more information, see Adding or removing portal
// administrators
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/administer-portals.html#portal-change-admins)
// in the AWS IoT SiteWise User Guide.
func (c *Client) CreatePortal(ctx context.Context, params *CreatePortalInput, optFns ...func(*Options)) (*CreatePortalOutput, error) {
	if params == nil {
		params = &CreatePortalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePortal", params, optFns, addOperationCreatePortalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePortalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePortalInput struct {

	// The AWS administrator's contact email address.
	//
	// This member is required.
	PortalContactEmail *string

	// A friendly name for the portal.
	//
	// This member is required.
	PortalName *string

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// a service role that allows the portal's users to access your AWS IoT SiteWise
	// resources on your behalf. For more information, see Using service roles for AWS
	// IoT SiteWise Monitor
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// This member is required.
	RoleArn *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	// The service to use to authenticate users to the portal. Choose from the
	// following options:
	//
	// * SSO – The portal uses AWS Single Sign-On to authenticate
	// users and manage user permissions. Before you can create a portal that uses AWS
	// SSO, you must enable AWS SSO. For more information, see Enabling AWS SSO
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-get-started.html#mon-gs-sso)
	// in the AWS IoT SiteWise User Guide. This option is only available in AWS Regions
	// other than the China Regions.
	//
	// * IAM – The portal uses AWS Identity and Access
	// Management (IAM) to authenticate users and manage user permissions. IAM users
	// must have the iotsitewise:CreatePresignedPortalUrl permission to sign in to the
	// portal. This option is only available in the China Regions.
	//
	// You can't change
	// this value after you create a portal. Default: SSO
	PortalAuthMode types.AuthMode

	// A description for the portal.
	PortalDescription *string

	// A logo image to display in the portal. Upload a square, high-resolution image.
	// The image is displayed on a dark background.
	PortalLogoImageFile *types.ImageFile

	// A list of key-value pairs that contain metadata for the portal. For more
	// information, see Tagging your AWS IoT SiteWise resources
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]*string
}

type CreatePortalOutput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the portal, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:portal/${PortalId}
	//
	// This member is required.
	PortalArn *string

	// The ID of the created portal.
	//
	// This member is required.
	PortalId *string

	// The URL for the AWS IoT SiteWise Monitor portal. You can use this URL to access
	// portals that use AWS SSO for authentication. For portals that use IAM for
	// authentication, you must use the CreatePresignedPortalUrl
	// (https://docs.aws.amazon.com/AWS IoT SiteWise API
	// ReferenceAPI_CreatePresignedPortalUrl.html) operation to create a URL that you
	// can use to access the portal.
	//
	// This member is required.
	PortalStartUrl *string

	// The status of the portal, which contains a state (CREATING after successfully
	// calling this operation) and any error message.
	//
	// This member is required.
	PortalStatus *types.PortalStatus

	// The associated AWS SSO application ID, if the portal uses AWS SSO.
	//
	// This member is required.
	SsoApplicationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePortalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePortal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePortal{}, middleware.After)
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
	if err = addEndpointPrefix_opCreatePortalMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreatePortalMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePortalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePortal(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreatePortalMiddleware struct {
}

func (*endpointPrefix_opCreatePortalMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreatePortalMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "monitor." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreatePortalMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreatePortalMiddleware{}, `OperationSerializer`, middleware.After)
}

type idempotencyToken_initializeOpCreatePortal struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePortal) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePortal) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePortalInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePortalInput ")
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
func addIdempotencyToken_opCreatePortalMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePortal{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePortal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "CreatePortal",
	}
}
