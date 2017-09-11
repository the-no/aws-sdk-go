// +build integration

//Package swf provides gucumber integration tests support.
package swf

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/swf"
)

func init() {
	gucumber.Before("@swf", func() {
		gucumber.World["client"] = swf.New(smoke.Session)
	})
}
