// +build integration

//Package kms provides gucumber integration tests support.
package kms

import (
	"github.com/gucumber/gucumber"
	"github.com/the-no/aws-sdk-go/awstesting/integration/smoke"
	"github.com/the-no/aws-sdk-go/service/kms"
)

func init() {
	gucumber.Before("@kms", func() {
		gucumber.World["client"] = kms.New(smoke.Session)
	})
}
