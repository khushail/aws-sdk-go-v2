// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes locations from a multi-location fleet. When deleting a location, all
// game server process and all instances that are still active in the location are
// shut down. To delete fleet locations, identify the fleet ID and provide a list
// of the locations to be deleted. If successful, GameLift sets the location status
// to DELETING, and begins to shut down existing server processes and terminate
// instances in each location being deleted. When completed, the location status
// changes to TERMINATED. Learn more Setting up GameLift fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related actions CreateFleetLocations | DescribeFleetLocationAttributes |
// DescribeFleetLocationCapacity | DescribeFleetLocationUtilization |
// DescribeFleetAttributes | DescribeFleetCapacity | DescribeFleetUtilization |
// UpdateFleetCapacity | StopFleetActions | DeleteFleetLocations | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DeleteFleetLocations(ctx context.Context, params *DeleteFleetLocationsInput, optFns ...func(*Options)) (*DeleteFleetLocationsOutput, error) {
	if params == nil {
		params = &DeleteFleetLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFleetLocations", params, optFns, c.addOperationDeleteFleetLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFleetLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteFleetLocationsInput struct {

	// A unique identifier for the fleet to delete locations for. You can use either
	// the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The list of fleet locations to delete. Specify locations in the form of an
	// Amazon Web Services Region code, such as us-west-2.
	//
	// This member is required.
	Locations []string

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type DeleteFleetLocationsOutput struct {

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is
	// assigned to a GameLift fleet resource and uniquely identifies it. ARNs are
	// unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912.
	FleetArn *string

	// A unique identifier for the fleet that location attributes are being deleted
	// for.
	FleetId *string

	// The remote locations that are being deleted, with each location status set to
	// DELETING.
	LocationStates []types.LocationState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFleetLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteFleetLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteFleetLocations{}, middleware.After)
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
	if err = addOpDeleteFleetLocationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFleetLocations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFleetLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteFleetLocations",
	}
}
