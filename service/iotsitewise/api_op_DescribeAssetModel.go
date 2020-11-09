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
	"time"
)

// Retrieves information about an asset model.
func (c *Client) DescribeAssetModel(ctx context.Context, params *DescribeAssetModelInput, optFns ...func(*Options)) (*DescribeAssetModelOutput, error) {
	if params == nil {
		params = &DescribeAssetModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssetModel", params, optFns, addOperationDescribeAssetModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssetModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssetModelInput struct {

	// The ID of the asset model.
	//
	// This member is required.
	AssetModelId *string
}

type DescribeAssetModelOutput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the asset model, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset-model/${AssetModelId}
	//
	// This member is required.
	AssetModelArn *string

	// The date the asset model was created, in Unix epoch time.
	//
	// This member is required.
	AssetModelCreationDate *time.Time

	// The asset model's description.
	//
	// This member is required.
	AssetModelDescription *string

	// A list of asset model hierarchies that each contain a childAssetModelId and a
	// hierarchyId (named id). A hierarchy specifies allowed parent/child asset
	// relationships for an asset model.
	//
	// This member is required.
	AssetModelHierarchies []*types.AssetModelHierarchy

	// The ID of the asset model.
	//
	// This member is required.
	AssetModelId *string

	// The date the asset model was last updated, in Unix epoch time.
	//
	// This member is required.
	AssetModelLastUpdateDate *time.Time

	// The name of the asset model.
	//
	// This member is required.
	AssetModelName *string

	// The list of asset properties for the asset model.
	//
	// This member is required.
	AssetModelProperties []*types.AssetModelProperty

	// The current status of the asset model, which contains a state and any error
	// message.
	//
	// This member is required.
	AssetModelStatus *types.AssetModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAssetModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAssetModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAssetModel{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeAssetModelMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeAssetModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssetModel(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeAssetModelMiddleware struct {
}

func (*endpointPrefix_opDescribeAssetModelMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeAssetModelMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "model." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeAssetModelMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeAssetModelMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAssetModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeAssetModel",
	}
}
