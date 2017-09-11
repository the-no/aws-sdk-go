// +build integration

//Package codepipeline provides gucumber integration tests support.
package codepipeline

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/codepipeline"
)

func init() {
	gucumber.Before("@codepipeline", func() {
		gucumber.World["client"] = codepipeline.New(smoke.Session)
	})
}
