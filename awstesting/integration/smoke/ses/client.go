// +build integration

//Package ses provides gucumber integration tests support.
package ses

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/ses"
)

func init() {
	gucumber.Before("@ses", func() {
		gucumber.World["client"] = ses.New(smoke.Session)
	})
}
