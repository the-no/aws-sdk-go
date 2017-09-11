// +build integration

//Package sns provides gucumber integration tests support.
package sns

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/sns"
)

func init() {
	gucumber.Before("@sns", func() {
		gucumber.World["client"] = sns.New(smoke.Session)
	})
}
