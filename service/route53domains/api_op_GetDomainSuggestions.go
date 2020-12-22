// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The GetDomainSuggestions operation returns a list of suggested domain names.
func (c *Client) GetDomainSuggestions(ctx context.Context, params *GetDomainSuggestionsInput, optFns ...func(*Options)) (*GetDomainSuggestionsOutput, error) {
	if params == nil {
		params = &GetDomainSuggestionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainSuggestions", params, optFns, addOperationGetDomainSuggestionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainSuggestionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDomainSuggestionsInput struct {

	// A domain name that you want to use as the basis for a list of possible domain
	// names. The top-level domain (TLD), such as .com, must be a TLD that Route 53
	// supports. For a list of supported TLDs, see Domains that You Can Register with
	// Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide. The domain name can contain only the
	// following characters:
	//
	// * Letters a through z. Domain names are not case
	// sensitive.
	//
	// * Numbers 0 through 9.
	//
	// * Hyphen (-). You can't specify a hyphen at
	// the beginning or end of a label.
	//
	// * Period (.) to separate the labels in the
	// name, such as the . in example.com.
	//
	// Internationalized domain names are not
	// supported for some top-level domains. To determine whether the TLD that you want
	// to use supports internationalized domain names, see Domains that You Can
	// Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html).
	//
	// This member is required.
	DomainName *string

	// If OnlyAvailable is true, Route 53 returns only domain names that are available.
	// If OnlyAvailable is false, Route 53 returns domain names without checking
	// whether they're available to be registered. To determine whether the domain is
	// available, you can call checkDomainAvailability for each suggestion.
	//
	// This member is required.
	OnlyAvailable *bool

	// The number of suggested domain names that you want Route 53 to return. Specify a
	// value between 1 and 50.
	//
	// This member is required.
	SuggestionCount int32
}

type GetDomainSuggestionsOutput struct {

	// A list of possible domain names. If you specified true for OnlyAvailable in the
	// request, the list contains only domains that are available for registration.
	SuggestionsList []types.DomainSuggestion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDomainSuggestionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDomainSuggestions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDomainSuggestions{}, middleware.After)
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
	if err = addOpGetDomainSuggestionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainSuggestions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomainSuggestions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "GetDomainSuggestions",
	}
}