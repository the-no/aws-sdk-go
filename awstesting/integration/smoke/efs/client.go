// +build integration

//Package efs provides gucumber integration tests support.
package efs

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/aws"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/efs"
)

func init() {
	gucumber.Before("@efs", func() {
		// FIXME remove custom region
		gucumber.World["client"] = efs.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
