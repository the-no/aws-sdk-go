// +build integration

//Package devicefarm provides gucumber integration tests support.
package devicefarm

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/aws"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/devicefarm"
)

func init() {
	gucumber.Before("@devicefarm", func() {
		// FIXME remove custom region
		gucumber.World["client"] = devicefarm.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
