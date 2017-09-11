// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"time"

	"github.com/the-no/aws-sdk-go/aws"
	"github.com/the-no/aws-sdk-go/aws/request"
)

// WaitUntilResourceRecordSetsChanged uses the Route 53 API operation
// GetChange to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Route53) WaitUntilResourceRecordSetsChanged(input *GetChangeInput) error {
	return c.WaitUntilResourceRecordSetsChangedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilResourceRecordSetsChangedWithContext is an extended version of WaitUntilResourceRecordSetsChanged.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Route53) WaitUntilResourceRecordSetsChangedWithContext(ctx aws.Context, input *GetChangeInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilResourceRecordSetsChanged",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "ChangeInfo.Status",
				Expected: "INSYNC",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetChangeInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetChangeRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
