// +build integration

//Package emr provides gucumber integration tests support.
package emr

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/emr"
)

func init() {
	gucumber.Before("@emr", func() {
		gucumber.World["client"] = emr.New(smoke.Session)
	})
}
