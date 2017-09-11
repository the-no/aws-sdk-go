// +build integration

//Package route53domains provides gucumber integration tests support.
package route53domains

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/route53domains"
)

func init() {
	gucumber.Before("@route53domains", func() {
		gucumber.World["client"] = route53domains.New(smoke.Session)
	})
}
