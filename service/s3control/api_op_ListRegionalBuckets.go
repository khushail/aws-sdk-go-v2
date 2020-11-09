// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// Returns a list of all Outposts buckets in an Outposts that are owned by the
// authenticated sender of the request. For more information, see Using Amazon S3
// on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html)
// in the Amazon Simple Storage Service Developer Guide. For an example of the
// request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint
// hostname prefix and outpost-id in your API request, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_ListRegionalBuckets.html#API_control_ListRegionalBuckets_Examples)
// section below.
func (c *Client) ListRegionalBuckets(ctx context.Context, params *ListRegionalBucketsInput, optFns ...func(*Options)) (*ListRegionalBucketsOutput, error) {
	if params == nil {
		params = &ListRegionalBucketsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRegionalBuckets", params, optFns, addOperationListRegionalBucketsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRegionalBucketsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegionalBucketsInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	//
	MaxResults *int32

	//
	NextToken *string

	// The ID of the AWS Outposts. This is required by Amazon S3 on Outposts buckets.
	OutpostId *string
}

type ListRegionalBucketsOutput struct {

	// NextToken is sent when isTruncated is true, which means there are more buckets
	// that can be listed. The next list requests to Amazon S3 can be continued with
	// this NextToken. NextToken is obfuscated and is not a real key.
	NextToken *string

	//
	RegionalBucketList []*types.RegionalBucket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRegionalBucketsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListRegionalBuckets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListRegionalBuckets{}, middleware.After)
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
	if err = addEndpointPrefix_opListRegionalBucketsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListRegionalBucketsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRegionalBuckets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opListRegionalBucketsMiddleware struct {
}

func (*endpointPrefix_opListRegionalBucketsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListRegionalBucketsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*ListRegionalBucketsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListRegionalBucketsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListRegionalBucketsMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opListRegionalBuckets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListRegionalBuckets",
	}
}
