// +build integration

//Package route53 provides gucumber integration tests support.
package route53

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/service/route53"
	"github.com/the-nothe-no/aws-sdk-go/awstesting/integration/smoke"
)

func init() {
	gucumber.Before("@route53", func() {
		gucumber.World["client"] = route53.New(smoke.Session)
	})
}
