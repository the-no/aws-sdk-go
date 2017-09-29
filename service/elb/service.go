// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elb

import (
	"errors"

	"github.com/the-no/aws-sdk-go/aws"
	"github.com/the-no/aws-sdk-go/aws/client"
	"github.com/the-no/aws-sdk-go/aws/client/metadata"
	"github.com/the-no/aws-sdk-go/aws/request"
	"github.com/the-no/aws-sdk-go/aws/signer/v4"
	"github.com/the-no/aws-sdk-go/private/protocol/query"
)

// ELB provides the API operation methods for making requests to
// Elastic Load Balancing. See this package's package overview docs
// for details on the service.
//
// ELB methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ELB struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "elasticloadbalancing" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName            // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ELB client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ELB client from just a session.
//     svc := elb.New(mySession)
//
//     // Create a ELB client with additional configuration
//     svc := elb.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ELB {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ELB {
	svc := &ELB{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-06-01",
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

// newRequest creates a new request for a ELB operation and runs any
// custom request initialization.
func (c *ELB) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

func (c *ELB) CreateResource(typ string, data []byte) (r aws.Referencer, attr aws.Attrabuter, err error) {
	return nil, nil, errors.New("Invail Resource Type!")
}

func (c *ELB) DeleteResource(typ string, r aws.Referencer) (err error) {
	return errors.New("Invail Resource Type!")
}
