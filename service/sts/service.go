// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"github.com/the-no/aws-sdk-go/aws"
	"github.com/the-no/aws-sdk-go/aws/client"
	"github.com/the-no/aws-sdk-go/aws/client/metadata"
	"github.com/the-no/aws-sdk-go/aws/request"
	"github.com/the-no/aws-sdk-go/aws/signer/v4"
	"github.com/the-no/aws-sdk-go/private/protocol/query"
)

// STS provides the API operation methods for making requests to
// AWS Security Token Service. See this package's package overview docs
// for details on the service.
//
// STS methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type STS struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "sts"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the STS client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a STS client from just a session.
//     svc := sts.New(mySession)
//
//     // Create a STS client with additional configuration
//     svc := sts.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *STS {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *STS {
	svc := &STS{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2011-06-15",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a STS operation and runs any
// custom request initialization.
func (c *STS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
