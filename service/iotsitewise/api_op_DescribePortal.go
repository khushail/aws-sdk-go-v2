// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a portal.
func (c *Client) DescribePortal(ctx context.Context, params *DescribePortalInput, optFns ...func(*Options)) (*DescribePortalOutput, error) {
	if params == nil {
		params = &DescribePortalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePortal", params, optFns, addOperationDescribePortalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePortalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePortalInput struct {

	// The ID of the portal.
	//
	// This member is required.
	PortalId *string
}

type DescribePortalOutput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the portal, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:portal/${PortalId}
	//
	// This member is required.
	PortalArn *string

	// The AWS SSO application generated client ID (used with AWS SSO APIs). AWS IoT
	// SiteWise includes portalClientId for only portals that use AWS SSO to
	// authenticate users.
	//
	// This member is required.
	PortalClientId *string

	// The AWS administrator's contact email address.
	//
	// This member is required.
	PortalContactEmail *string

	// The date the portal was created, in Unix epoch time.
	//
	// This member is required.
	PortalCreationDate *time.Time

	// The ID of the portal.
	//
	// This member is required.
	PortalId *string

	// The date the portal was last updated, in Unix epoch time.
	//
	// This member is required.
	PortalLastUpdateDate *time.Time

	// The name of the portal.
	//
	// This member is required.
	PortalName *string

	// The URL for the AWS IoT SiteWise Monitor portal. You can use this URL to access
	// portals that use AWS SSO for authentication. For portals that use IAM for
	// authentication, you must use the CreatePresignedPortalUrl
	// (https://docs.aws.amazon.com/AWS IoT SiteWise API
	// ReferenceAPI_CreatePresignedPortalUrl.html) operation to create a URL that you
	// can use to access the portal.
	//
	// This member is required.
	PortalStartUrl *string

	// The current status of the portal, which contains a state and any error message.
	//
	// This member is required.
	PortalStatus *types.PortalStatus

	// The service to use to authenticate users to the portal.
	PortalAuthMode types.AuthMode

	// The portal's description.
	PortalDescription *string

	// The portal's logo image, which is available at a URL.
	PortalLogoImageLocation *types.ImageLocation

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the service role that allows the portal's users to access your AWS IoT SiteWise
	// resources on your behalf. For more information, see Using service roles for AWS
	// IoT SiteWise Monitor
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html)
	// in the AWS IoT SiteWise User Guide.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePortalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePortal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePortal{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribePortalMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribePortalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePortal(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribePortalMiddleware struct {
}

func (*endpointPrefix_opDescribePortalMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribePortalMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addEndpointPrefix_opDescribePortalMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribePortalMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePortal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribePortal",
	}
}