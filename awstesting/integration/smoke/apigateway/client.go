// +build integration

//Package apigateway provides gucumber integration tests support.
package apigateway

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/apigateway"
)

func init() {
	gucumber.Before("@apigateway", func() {
		gucumber.World["client"] = apigateway.New(smoke.Session)
	})
}
