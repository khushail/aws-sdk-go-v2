// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified schema in detail.
func (c *Client) GetSchema(ctx context.Context, params *GetSchemaInput, optFns ...func(*Options)) (*GetSchemaOutput, error) {
	if params == nil {
		params = &GetSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSchema", params, optFns, c.addOperationGetSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSchemaInput struct {

	// This is a wrapper structure to contain schema identity fields. The structure
	// contains:
	//
	// * SchemaId$SchemaArn: The Amazon Resource Name (ARN) of the schema.
	// Either SchemaArn or SchemaName and RegistryName has to be provided.
	//
	// *
	// SchemaId$SchemaName: The name of the schema. Either SchemaArn or SchemaName and
	// RegistryName has to be provided.
	//
	// This member is required.
	SchemaId *types.SchemaId

	noSmithyDocumentSerde
}

type GetSchemaOutput struct {

	// The compatibility mode of the schema.
	Compatibility types.Compatibility

	// The date and time the schema was created.
	CreatedTime *string

	// The data format of the schema definition. Currently AVRO, JSON and PROTOBUF are
	// supported.
	DataFormat types.DataFormat

	// A description of schema if specified when created
	Description *string

	// The latest version of the schema associated with the returned schema definition.
	LatestSchemaVersion int64

	// The next version of the schema associated with the returned schema definition.
	NextSchemaVersion int64

	// The Amazon Resource Name (ARN) of the registry.
	RegistryArn *string

	// The name of the registry.
	RegistryName *string

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string

	// The version number of the checkpoint (the last time the compatibility mode was
	// changed).
	SchemaCheckpoint int64

	// The name of the schema.
	SchemaName *string

	// The status of the schema.
	SchemaStatus types.SchemaStatus

	// The date and time the schema was updated.
	UpdatedTime *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSchema{}, middleware.After)
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
	if err = addOpGetSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetSchema",
	}
}
