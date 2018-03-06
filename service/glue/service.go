// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Glue provides the API operation methods for making requests to
// AWS Glue. See this package's package overview docs
// for details on the service.
//
// Glue methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Glue struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Glue)

// Used for custom request initialization logic
var initRequest func(*Glue, *aws.Request)

// Service information constants
const (
	ServiceName = "glue"      // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Glue client with a config.
//
// Example:
//     // Create a Glue client from just a config.
//     svc := glue.New(myConfig)
func New(config aws.Config) *Glue {
	var signingName string
	signingRegion := config.Region

	svc := &Glue{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-03-31",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSGlue",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a Glue operation and runs any
// custom request initialization.
func (c *Glue) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
