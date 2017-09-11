// +build integration

//Package ecs provides gucumber integration tests support.
package ecs

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/aws"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/ecs"
)

func init() {
	gucumber.Before("@ecs", func() {
		// FIXME remove custom region
		gucumber.World["client"] = ecs.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
