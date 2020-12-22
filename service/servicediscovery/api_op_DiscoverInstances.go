// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Discovers registered instances for a specified namespace and service. You can
// use DiscoverInstances to discover instances for any type of namespace. For
// public and private DNS namespaces, you can also use DNS queries to discover
// instances.
func (c *Client) DiscoverInstances(ctx context.Context, params *DiscoverInstancesInput, optFns ...func(*Options)) (*DiscoverInstancesOutput, error) {
	if params == nil {
		params = &DiscoverInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DiscoverInstances", params, optFns, addOperationDiscoverInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DiscoverInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DiscoverInstancesInput struct {

	// The name of the namespace that you specified when you registered the instance.
	//
	// This member is required.
	NamespaceName *string

	// The name of the service that you specified when you registered the instance.
	//
	// This member is required.
	ServiceName *string

	// The health status of the instances that you want to discover.
	HealthStatus types.HealthStatusFilter

	// The maximum number of instances that you want AWS Cloud Map to return in the
	// response to a DiscoverInstances request. If you don't specify a value for
	// MaxResults, AWS Cloud Map returns up to 100 instances.
	MaxResults *int32

	// Opportunistic filters to scope the results based on custom attributes. If there
	// are instances that match both the filters specified in both the QueryParameters
	// parameter and this parameter, they are returned. Otherwise, these filters are
	// ignored and only instances that match the filters specified in the
	// QueryParameters parameter are returned.
	OptionalParameters map[string]string

	// Filters to scope the results based on custom attributes for the instance. For
	// example, {version=v1, az=1a}. Only instances that match all the specified
	// key-value pairs will be returned.
	QueryParameters map[string]string
}

type DiscoverInstancesOutput struct {

	// A complex type that contains one HttpInstanceSummary for each registered
	// instance.
	Instances []types.HttpInstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDiscoverInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDiscoverInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDiscoverInstances{}, middleware.After)
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
	if err = addEndpointPrefix_opDiscoverInstancesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDiscoverInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDiscoverInstances(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDiscoverInstancesMiddleware struct {
}

func (*endpointPrefix_opDiscoverInstancesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDiscoverInstancesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDiscoverInstancesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDiscoverInstancesMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDiscoverInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "DiscoverInstances",
	}
}