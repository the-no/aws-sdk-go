// +build integration

//Package configservice provides gucumber integration tests support.
package configservice

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/configservice"
)

func init() {
	gucumber.Before("@configservice", func() {
		gucumber.World["client"] = configservice.New(smoke.Session)
	})
}
