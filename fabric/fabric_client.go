package fabric

import (
	"context"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// type FabClient struct{}

var (
	chaincodeID = "basic"
)

type ChaincodeInvoker struct {
	channelClient *channel.Client
}

func NewChainCodeInvoker(
	ctx context.Context,
	sdk *fabsdk.FabricSDK,
	channelId string) (*ChaincodeInvoker, error) {
	channelProvider := sdk.ChannelContext(channelId, fabsdk.WithUser("tractionx"), fabsdk.WithOrg("org1"))
	channelClient, err := channel.New(channelProvider)
	if err != nil {
		return nil, err
	}

	return &ChaincodeInvoker{
		channelClient: channelClient,
	}, nil
}

func (ci *ChaincodeInvoker) InvokePropertyCreate(propertyId, propertyValue string) error {
	req := channel.Request{
		ChaincodeID: chaincodeID,
		Fcn:         "CreateProperty",
		Args:        [][]byte{[]byte(propertyId), []byte(propertyValue)},
	}

	response, err := ci.channelClient.Execute(req)
	if err != nil {
		return err
	}

	if response.ChaincodeStatus != 200 {
		return fmt.Errorf("chaincode invocation failed with status: %d", response.ChaincodeStatus)
	}

	return nil
}
