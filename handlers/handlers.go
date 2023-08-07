package handlers

import (
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/labstack/echo/v4"
)

type OrgSetup struct {
	OrgName      string
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	Gateway      client.Gateway
}

type QueryPayload struct {
	chaincodeID string
	channelID   string
	function    string
	args        []string
}

type InvokePayload struct {
	chaincodeID string
	channelID   string
	function    string
	args        []string
}

func BaseRoute(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{
		"message": "Server is Healthy.",
	})
}

func (setup OrgSetup) Query(e echo.Context) error {

	queryPayload := new(QueryPayload)
	if err := e.Bind(queryPayload); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Unable to create Query.",
		})
	}

	network := setup.Gateway.GetNetwork(queryPayload.channelID)
	contract := network.GetContract(queryPayload.chaincodeID)
	evalResponse, err := contract.EvaluateTransaction(queryPayload.function, queryPayload.args...)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Unable to evaluate Response.",
		})
	}

	return e.JSON(http.StatusBadRequest, echo.Map{
		"data": evalResponse,
	})
}

func (setup *OrgSetup) Invoke(e echo.Context) error {
	invokePayload := new(InvokePayload)
	if err := e.Bind(invokePayload); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Unable to create Query.",
		})
	}

	network := setup.Gateway.GetNetwork(invokePayload.channelID)
	contract := network.GetContract(invokePayload.chaincodeID)
	txn_proposal, err := contract.NewProposal(invokePayload.function,
		client.WithArguments(invokePayload.args...))
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Unable to create Txn proposal.",
		})
	}

	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Unable endorse Txn proposal.",
		})
	}
	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Unable commit endorsed Txn.",
		})
	}

	return e.JSON(http.StatusBadRequest, echo.Map{
		"data": txn_committed,
	})
}
