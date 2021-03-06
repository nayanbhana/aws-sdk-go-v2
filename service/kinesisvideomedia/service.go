// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideomedia

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// KinesisVideoMedia provides the API operation methods for making requests to
// Amazon Kinesis Video Streams Media. See this package's package overview docs
// for details on the service.
//
// KinesisVideoMedia methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type KinesisVideoMedia struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*KinesisVideoMedia)

// Used for custom request initialization logic
var initRequest func(*KinesisVideoMedia, *aws.Request)

// Service information constants
const (
	ServiceName = "kinesisvideo" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName    // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the KinesisVideoMedia client with a config.
//
// Example:
//     // Create a KinesisVideoMedia client from just a config.
//     svc := kinesisvideomedia.New(myConfig)
func New(config aws.Config) *KinesisVideoMedia {
	var signingName string
	signingRegion := config.Region

	svc := &KinesisVideoMedia{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-09-30",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a KinesisVideoMedia operation and runs any
// custom request initialization.
func (c *KinesisVideoMedia) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
