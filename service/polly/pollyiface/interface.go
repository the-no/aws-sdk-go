// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pollyiface provides an interface to enable mocking the Amazon Polly service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pollyiface

import (
	"github.com/the-no/aws-sdk-go/aws"
	"github.com/the-no/aws-sdk-go/aws/request"
	"github.com/the-no/aws-sdk-go/service/polly"
)

// PollyAPI provides an interface to enable mocking the
// polly.Polly service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Polly.
//    func myFunc(svc pollyiface.PollyAPI) bool {
//        // Make svc.DeleteLexicon request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := polly.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPollyClient struct {
//        pollyiface.PollyAPI
//    }
//    func (m *mockPollyClient) DeleteLexicon(input *polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPollyClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PollyAPI interface {
	DeleteLexicon(*polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error)
	DeleteLexiconWithContext(aws.Context, *polly.DeleteLexiconInput, ...request.Option) (*polly.DeleteLexiconOutput, error)
	DeleteLexiconRequest(*polly.DeleteLexiconInput) (*request.Request, *polly.DeleteLexiconOutput)

	DescribeVoices(*polly.DescribeVoicesInput) (*polly.DescribeVoicesOutput, error)
	DescribeVoicesWithContext(aws.Context, *polly.DescribeVoicesInput, ...request.Option) (*polly.DescribeVoicesOutput, error)
	DescribeVoicesRequest(*polly.DescribeVoicesInput) (*request.Request, *polly.DescribeVoicesOutput)

	GetLexicon(*polly.GetLexiconInput) (*polly.GetLexiconOutput, error)
	GetLexiconWithContext(aws.Context, *polly.GetLexiconInput, ...request.Option) (*polly.GetLexiconOutput, error)
	GetLexiconRequest(*polly.GetLexiconInput) (*request.Request, *polly.GetLexiconOutput)

	ListLexicons(*polly.ListLexiconsInput) (*polly.ListLexiconsOutput, error)
	ListLexiconsWithContext(aws.Context, *polly.ListLexiconsInput, ...request.Option) (*polly.ListLexiconsOutput, error)
	ListLexiconsRequest(*polly.ListLexiconsInput) (*request.Request, *polly.ListLexiconsOutput)

	PutLexicon(*polly.PutLexiconInput) (*polly.PutLexiconOutput, error)
	PutLexiconWithContext(aws.Context, *polly.PutLexiconInput, ...request.Option) (*polly.PutLexiconOutput, error)
	PutLexiconRequest(*polly.PutLexiconInput) (*request.Request, *polly.PutLexiconOutput)

	SynthesizeSpeech(*polly.SynthesizeSpeechInput) (*polly.SynthesizeSpeechOutput, error)
	SynthesizeSpeechWithContext(aws.Context, *polly.SynthesizeSpeechInput, ...request.Option) (*polly.SynthesizeSpeechOutput, error)
	SynthesizeSpeechRequest(*polly.SynthesizeSpeechInput) (*request.Request, *polly.SynthesizeSpeechOutput)
}

var _ PollyAPI = (*polly.Polly)(nil)
