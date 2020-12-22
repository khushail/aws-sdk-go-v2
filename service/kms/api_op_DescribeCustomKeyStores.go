// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about custom key stores
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// in the account and region. This operation is part of the Custom Key Store
// feature
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in AWS KMS, which combines the convenience and extensive integration of
// AWS KMS with the isolation and control of a single-tenant key store. By default,
// this operation returns information about all custom key stores in the account
// and region. To get only information about a particular custom key store, use
// either the CustomKeyStoreName or CustomKeyStoreId parameter (but not both). To
// determine whether the custom key store is connected to its AWS CloudHSM cluster,
// use the ConnectionState element in the response. If an attempt to connect the
// custom key store failed, the ConnectionState value is FAILED and the
// ConnectionErrorCode element in the response indicates the cause of the failure.
// For help interpreting the ConnectionErrorCode, see CustomKeyStoresListEntry.
// Custom key stores have a DISCONNECTED connection state if the key store has
// never been connected or you use the DisconnectCustomKeyStore operation to
// disconnect it. If your custom key store state is CONNECTED but you are having
// trouble using it, make sure that its associated AWS CloudHSM cluster is active
// and contains the minimum number of HSMs required for the operation, if any. For
// help repairing your custom key store, see the Troubleshooting Custom Key Stores
// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html) topic
// in the AWS Key Management Service Developer Guide.
func (c *Client) DescribeCustomKeyStores(ctx context.Context, params *DescribeCustomKeyStoresInput, optFns ...func(*Options)) (*DescribeCustomKeyStoresOutput, error) {
	if params == nil {
		params = &DescribeCustomKeyStoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCustomKeyStores", params, optFns, addOperationDescribeCustomKeyStoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCustomKeyStoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomKeyStoresInput struct {

	// Gets only information about the specified custom key store. Enter the key store
	// ID. By default, this operation gets information about all custom key stores in
	// the account and region. To limit the output to a particular custom key store,
	// you can use either the CustomKeyStoreId or CustomKeyStoreName parameter, but not
	// both.
	CustomKeyStoreId *string

	// Gets only information about the specified custom key store. Enter the friendly
	// name of the custom key store. By default, this operation gets information about
	// all custom key stores in the account and region. To limit the output to a
	// particular custom key store, you can use either the CustomKeyStoreId or
	// CustomKeyStoreName parameter, but not both.
	CustomKeyStoreName *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, AWS KMS does not return more than the specified number of
	// items, but it might return fewer.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string
}

type DescribeCustomKeyStoresOutput struct {

	// Contains metadata about each custom key store.
	CustomKeyStores []types.CustomKeyStoresListEntry

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request.
	NextMarker *string

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in thisresponse to the Marker parameter in a
	// subsequent request.
	Truncated bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCustomKeyStoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCustomKeyStores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCustomKeyStores{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomKeyStores(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCustomKeyStores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DescribeCustomKeyStores",
	}
}