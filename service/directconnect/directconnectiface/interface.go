// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package directconnectiface provides an interface to enable mocking the AWS Direct Connect service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package directconnectiface

import (
	"github.com/the-no/aws-sdk-go/aws"
	"github.com/the-no/aws-sdk-go/aws/request"
	"github.com/the-no/aws-sdk-go/service/directconnect"
)

// DirectConnectAPI provides an interface to enable mocking the
// directconnect.DirectConnect service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Direct Connect.
//    func myFunc(svc directconnectiface.DirectConnectAPI) bool {
//        // Make svc.AllocateConnectionOnInterconnect request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := directconnect.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDirectConnectClient struct {
//        directconnectiface.DirectConnectAPI
//    }
//    func (m *mockDirectConnectClient) AllocateConnectionOnInterconnect(input *directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDirectConnectClient{}
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
type DirectConnectAPI interface {
	AllocateConnectionOnInterconnect(*directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error)
	AllocateConnectionOnInterconnectWithContext(aws.Context, *directconnect.AllocateConnectionOnInterconnectInput, ...request.Option) (*directconnect.Connection, error)
	AllocateConnectionOnInterconnectRequest(*directconnect.AllocateConnectionOnInterconnectInput) (*request.Request, *directconnect.Connection)

	AllocateHostedConnection(*directconnect.AllocateHostedConnectionInput) (*directconnect.Connection, error)
	AllocateHostedConnectionWithContext(aws.Context, *directconnect.AllocateHostedConnectionInput, ...request.Option) (*directconnect.Connection, error)
	AllocateHostedConnectionRequest(*directconnect.AllocateHostedConnectionInput) (*request.Request, *directconnect.Connection)

	AllocatePrivateVirtualInterface(*directconnect.AllocatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
	AllocatePrivateVirtualInterfaceWithContext(aws.Context, *directconnect.AllocatePrivateVirtualInterfaceInput, ...request.Option) (*directconnect.VirtualInterface, error)
	AllocatePrivateVirtualInterfaceRequest(*directconnect.AllocatePrivateVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	AllocatePublicVirtualInterface(*directconnect.AllocatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
	AllocatePublicVirtualInterfaceWithContext(aws.Context, *directconnect.AllocatePublicVirtualInterfaceInput, ...request.Option) (*directconnect.VirtualInterface, error)
	AllocatePublicVirtualInterfaceRequest(*directconnect.AllocatePublicVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	AssociateConnectionWithLag(*directconnect.AssociateConnectionWithLagInput) (*directconnect.Connection, error)
	AssociateConnectionWithLagWithContext(aws.Context, *directconnect.AssociateConnectionWithLagInput, ...request.Option) (*directconnect.Connection, error)
	AssociateConnectionWithLagRequest(*directconnect.AssociateConnectionWithLagInput) (*request.Request, *directconnect.Connection)

	AssociateHostedConnection(*directconnect.AssociateHostedConnectionInput) (*directconnect.Connection, error)
	AssociateHostedConnectionWithContext(aws.Context, *directconnect.AssociateHostedConnectionInput, ...request.Option) (*directconnect.Connection, error)
	AssociateHostedConnectionRequest(*directconnect.AssociateHostedConnectionInput) (*request.Request, *directconnect.Connection)

	AssociateVirtualInterface(*directconnect.AssociateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
	AssociateVirtualInterfaceWithContext(aws.Context, *directconnect.AssociateVirtualInterfaceInput, ...request.Option) (*directconnect.VirtualInterface, error)
	AssociateVirtualInterfaceRequest(*directconnect.AssociateVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	ConfirmConnection(*directconnect.ConfirmConnectionInput) (*directconnect.ConfirmConnectionOutput, error)
	ConfirmConnectionWithContext(aws.Context, *directconnect.ConfirmConnectionInput, ...request.Option) (*directconnect.ConfirmConnectionOutput, error)
	ConfirmConnectionRequest(*directconnect.ConfirmConnectionInput) (*request.Request, *directconnect.ConfirmConnectionOutput)

	ConfirmPrivateVirtualInterface(*directconnect.ConfirmPrivateVirtualInterfaceInput) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error)
	ConfirmPrivateVirtualInterfaceWithContext(aws.Context, *directconnect.ConfirmPrivateVirtualInterfaceInput, ...request.Option) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error)
	ConfirmPrivateVirtualInterfaceRequest(*directconnect.ConfirmPrivateVirtualInterfaceInput) (*request.Request, *directconnect.ConfirmPrivateVirtualInterfaceOutput)

	ConfirmPublicVirtualInterface(*directconnect.ConfirmPublicVirtualInterfaceInput) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error)
	ConfirmPublicVirtualInterfaceWithContext(aws.Context, *directconnect.ConfirmPublicVirtualInterfaceInput, ...request.Option) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error)
	ConfirmPublicVirtualInterfaceRequest(*directconnect.ConfirmPublicVirtualInterfaceInput) (*request.Request, *directconnect.ConfirmPublicVirtualInterfaceOutput)

	CreateBGPPeer(*directconnect.CreateBGPPeerInput) (*directconnect.CreateBGPPeerOutput, error)
	CreateBGPPeerWithContext(aws.Context, *directconnect.CreateBGPPeerInput, ...request.Option) (*directconnect.CreateBGPPeerOutput, error)
	CreateBGPPeerRequest(*directconnect.CreateBGPPeerInput) (*request.Request, *directconnect.CreateBGPPeerOutput)

	CreateConnection(*directconnect.CreateConnectionInput) (*directconnect.Connection, error)
	CreateConnectionWithContext(aws.Context, *directconnect.CreateConnectionInput, ...request.Option) (*directconnect.Connection, error)
	CreateConnectionRequest(*directconnect.CreateConnectionInput) (*request.Request, *directconnect.Connection)

	CreateInterconnect(*directconnect.CreateInterconnectInput) (*directconnect.Interconnect, error)
	CreateInterconnectWithContext(aws.Context, *directconnect.CreateInterconnectInput, ...request.Option) (*directconnect.Interconnect, error)
	CreateInterconnectRequest(*directconnect.CreateInterconnectInput) (*request.Request, *directconnect.Interconnect)

	CreateLag(*directconnect.CreateLagInput) (*directconnect.Lag, error)
	CreateLagWithContext(aws.Context, *directconnect.CreateLagInput, ...request.Option) (*directconnect.Lag, error)
	CreateLagRequest(*directconnect.CreateLagInput) (*request.Request, *directconnect.Lag)

	CreatePrivateVirtualInterface(*directconnect.CreatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
	CreatePrivateVirtualInterfaceWithContext(aws.Context, *directconnect.CreatePrivateVirtualInterfaceInput, ...request.Option) (*directconnect.VirtualInterface, error)
	CreatePrivateVirtualInterfaceRequest(*directconnect.CreatePrivateVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	CreatePublicVirtualInterface(*directconnect.CreatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
	CreatePublicVirtualInterfaceWithContext(aws.Context, *directconnect.CreatePublicVirtualInterfaceInput, ...request.Option) (*directconnect.VirtualInterface, error)
	CreatePublicVirtualInterfaceRequest(*directconnect.CreatePublicVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	DeleteBGPPeer(*directconnect.DeleteBGPPeerInput) (*directconnect.DeleteBGPPeerOutput, error)
	DeleteBGPPeerWithContext(aws.Context, *directconnect.DeleteBGPPeerInput, ...request.Option) (*directconnect.DeleteBGPPeerOutput, error)
	DeleteBGPPeerRequest(*directconnect.DeleteBGPPeerInput) (*request.Request, *directconnect.DeleteBGPPeerOutput)

	DeleteConnection(*directconnect.DeleteConnectionInput) (*directconnect.Connection, error)
	DeleteConnectionWithContext(aws.Context, *directconnect.DeleteConnectionInput, ...request.Option) (*directconnect.Connection, error)
	DeleteConnectionRequest(*directconnect.DeleteConnectionInput) (*request.Request, *directconnect.Connection)

	DeleteInterconnect(*directconnect.DeleteInterconnectInput) (*directconnect.DeleteInterconnectOutput, error)
	DeleteInterconnectWithContext(aws.Context, *directconnect.DeleteInterconnectInput, ...request.Option) (*directconnect.DeleteInterconnectOutput, error)
	DeleteInterconnectRequest(*directconnect.DeleteInterconnectInput) (*request.Request, *directconnect.DeleteInterconnectOutput)

	DeleteLag(*directconnect.DeleteLagInput) (*directconnect.Lag, error)
	DeleteLagWithContext(aws.Context, *directconnect.DeleteLagInput, ...request.Option) (*directconnect.Lag, error)
	DeleteLagRequest(*directconnect.DeleteLagInput) (*request.Request, *directconnect.Lag)

	DeleteVirtualInterface(*directconnect.DeleteVirtualInterfaceInput) (*directconnect.DeleteVirtualInterfaceOutput, error)
	DeleteVirtualInterfaceWithContext(aws.Context, *directconnect.DeleteVirtualInterfaceInput, ...request.Option) (*directconnect.DeleteVirtualInterfaceOutput, error)
	DeleteVirtualInterfaceRequest(*directconnect.DeleteVirtualInterfaceInput) (*request.Request, *directconnect.DeleteVirtualInterfaceOutput)

	DescribeConnectionLoa(*directconnect.DescribeConnectionLoaInput) (*directconnect.DescribeConnectionLoaOutput, error)
	DescribeConnectionLoaWithContext(aws.Context, *directconnect.DescribeConnectionLoaInput, ...request.Option) (*directconnect.DescribeConnectionLoaOutput, error)
	DescribeConnectionLoaRequest(*directconnect.DescribeConnectionLoaInput) (*request.Request, *directconnect.DescribeConnectionLoaOutput)

	DescribeConnections(*directconnect.DescribeConnectionsInput) (*directconnect.Connections, error)
	DescribeConnectionsWithContext(aws.Context, *directconnect.DescribeConnectionsInput, ...request.Option) (*directconnect.Connections, error)
	DescribeConnectionsRequest(*directconnect.DescribeConnectionsInput) (*request.Request, *directconnect.Connections)

	DescribeConnectionsOnInterconnect(*directconnect.DescribeConnectionsOnInterconnectInput) (*directconnect.Connections, error)
	DescribeConnectionsOnInterconnectWithContext(aws.Context, *directconnect.DescribeConnectionsOnInterconnectInput, ...request.Option) (*directconnect.Connections, error)
	DescribeConnectionsOnInterconnectRequest(*directconnect.DescribeConnectionsOnInterconnectInput) (*request.Request, *directconnect.Connections)

	DescribeHostedConnections(*directconnect.DescribeHostedConnectionsInput) (*directconnect.Connections, error)
	DescribeHostedConnectionsWithContext(aws.Context, *directconnect.DescribeHostedConnectionsInput, ...request.Option) (*directconnect.Connections, error)
	DescribeHostedConnectionsRequest(*directconnect.DescribeHostedConnectionsInput) (*request.Request, *directconnect.Connections)

	DescribeInterconnectLoa(*directconnect.DescribeInterconnectLoaInput) (*directconnect.DescribeInterconnectLoaOutput, error)
	DescribeInterconnectLoaWithContext(aws.Context, *directconnect.DescribeInterconnectLoaInput, ...request.Option) (*directconnect.DescribeInterconnectLoaOutput, error)
	DescribeInterconnectLoaRequest(*directconnect.DescribeInterconnectLoaInput) (*request.Request, *directconnect.DescribeInterconnectLoaOutput)

	DescribeInterconnects(*directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error)
	DescribeInterconnectsWithContext(aws.Context, *directconnect.DescribeInterconnectsInput, ...request.Option) (*directconnect.DescribeInterconnectsOutput, error)
	DescribeInterconnectsRequest(*directconnect.DescribeInterconnectsInput) (*request.Request, *directconnect.DescribeInterconnectsOutput)

	DescribeLags(*directconnect.DescribeLagsInput) (*directconnect.DescribeLagsOutput, error)
	DescribeLagsWithContext(aws.Context, *directconnect.DescribeLagsInput, ...request.Option) (*directconnect.DescribeLagsOutput, error)
	DescribeLagsRequest(*directconnect.DescribeLagsInput) (*request.Request, *directconnect.DescribeLagsOutput)

	DescribeLoa(*directconnect.DescribeLoaInput) (*directconnect.Loa, error)
	DescribeLoaWithContext(aws.Context, *directconnect.DescribeLoaInput, ...request.Option) (*directconnect.Loa, error)
	DescribeLoaRequest(*directconnect.DescribeLoaInput) (*request.Request, *directconnect.Loa)

	DescribeLocations(*directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error)
	DescribeLocationsWithContext(aws.Context, *directconnect.DescribeLocationsInput, ...request.Option) (*directconnect.DescribeLocationsOutput, error)
	DescribeLocationsRequest(*directconnect.DescribeLocationsInput) (*request.Request, *directconnect.DescribeLocationsOutput)

	DescribeTags(*directconnect.DescribeTagsInput) (*directconnect.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *directconnect.DescribeTagsInput, ...request.Option) (*directconnect.DescribeTagsOutput, error)
	DescribeTagsRequest(*directconnect.DescribeTagsInput) (*request.Request, *directconnect.DescribeTagsOutput)

	DescribeVirtualGateways(*directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error)
	DescribeVirtualGatewaysWithContext(aws.Context, *directconnect.DescribeVirtualGatewaysInput, ...request.Option) (*directconnect.DescribeVirtualGatewaysOutput, error)
	DescribeVirtualGatewaysRequest(*directconnect.DescribeVirtualGatewaysInput) (*request.Request, *directconnect.DescribeVirtualGatewaysOutput)

	DescribeVirtualInterfaces(*directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error)
	DescribeVirtualInterfacesWithContext(aws.Context, *directconnect.DescribeVirtualInterfacesInput, ...request.Option) (*directconnect.DescribeVirtualInterfacesOutput, error)
	DescribeVirtualInterfacesRequest(*directconnect.DescribeVirtualInterfacesInput) (*request.Request, *directconnect.DescribeVirtualInterfacesOutput)

	DisassociateConnectionFromLag(*directconnect.DisassociateConnectionFromLagInput) (*directconnect.Connection, error)
	DisassociateConnectionFromLagWithContext(aws.Context, *directconnect.DisassociateConnectionFromLagInput, ...request.Option) (*directconnect.Connection, error)
	DisassociateConnectionFromLagRequest(*directconnect.DisassociateConnectionFromLagInput) (*request.Request, *directconnect.Connection)

	TagResource(*directconnect.TagResourceInput) (*directconnect.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *directconnect.TagResourceInput, ...request.Option) (*directconnect.TagResourceOutput, error)
	TagResourceRequest(*directconnect.TagResourceInput) (*request.Request, *directconnect.TagResourceOutput)

	UntagResource(*directconnect.UntagResourceInput) (*directconnect.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *directconnect.UntagResourceInput, ...request.Option) (*directconnect.UntagResourceOutput, error)
	UntagResourceRequest(*directconnect.UntagResourceInput) (*request.Request, *directconnect.UntagResourceOutput)

	UpdateLag(*directconnect.UpdateLagInput) (*directconnect.Lag, error)
	UpdateLagWithContext(aws.Context, *directconnect.UpdateLagInput, ...request.Option) (*directconnect.Lag, error)
	UpdateLagRequest(*directconnect.UpdateLagInput) (*request.Request, *directconnect.Lag)
}

var _ DirectConnectAPI = (*directconnect.DirectConnect)(nil)
