package chaincode

import (
	"encoding/json"
	"fmt"
	"tractionx/models"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PropertyContract struct {
	contractapi.Contract
}

// PropertyExists returns true when property with given ID exists in world state
func (c *PropertyContract) PropertyExists(
	ctx contractapi.TransactionContextInterface,
	propertyId string) (bool, error) {

	// GetStub: should provide a way to access the stub set by Init/Invoke
	// GetState: returns the value of the specified `key` from the ledger.
	data, err := ctx.GetStub().GetState(propertyId)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

func (c *PropertyContract) ReadProperty(
	ctx contractapi.TransactionContextInterface,
	id string,
) (*models.Property, error) {

	propertyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}

	if propertyJSON == nil {
		return nil, fmt.Errorf("the property %s does not exists", id)
	}

	var property models.Property
	err = json.Unmarshal([]byte(propertyJSON), &property)
	if err != nil {
		return nil, err
	}

	return &property, nil

}

func (c *PropertyContract) TransferProperty(
	ctx contractapi.TransactionContextInterface,
	id string,
	newOwner string) (string, error) {

	property, err := c.ReadProperty(ctx, id)
	if err != nil {
		return "", err
	}

	oldOwner := property.OwnedBy
	property.OwnedBy = newOwner

	propertyJSON, err := json.Marshal(property)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(id, propertyJSON)
	if err != nil {
		return "", err
	}
	return oldOwner, nil
}
