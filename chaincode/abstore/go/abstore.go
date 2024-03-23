package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// User 구조체는 사용자 정보를 나타냅니다.
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Property 구조체는 부동산 정보를 나타냅니다.
type Property struct {
	ID          string `json:"id"`
	OwnerID     string `json:"ownerId"`
	Description string `json:"description"`
}

// ABstore 체인코드의 구조체 정의를 확장합니다.
type ABstore struct {
	contractapi.Contract
}

// RegisterUser는 새로운 사용자를 등록합니다.
func (t *ABstore) RegisterUser(ctx contractapi.TransactionContextInterface, userID, name string) error {
	user := User{
		ID:   userID,
		Name: name,
	}
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(userID, userBytes)
}

// RegisterProperty는 새로운 부동산을 등록합니다.
func (t *ABstore) RegisterProperty(ctx contractapi.TransactionContextInterface, propertyID, ownerID, description string) error {
	property := Property{
		ID:          propertyID,
		OwnerID:     ownerID,
		Description: description,
	}
	propertyBytes, err := json.Marshal(property)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(propertyID, propertyBytes)
}

// UpdatePropertyDescription은 부동산 설명을 업데이트합니다.
func (t *ABstore) UpdatePropertyDescription(ctx contractapi.TransactionContextInterface, propertyID, newDescription string) error {
	propertyBytes, err := ctx.GetStub().GetState(propertyID)
	if err != nil {
		return err
	}

	if propertyBytes == nil {
		return fmt.Errorf("property %s does not exist", propertyID)
	}

	property := Property{}
	err = json.Unmarshal(propertyBytes, &property)
	if err != nil {
		return err
	}

	property.Description = newDescription
	updatedPropertyBytes, err := json.Marshal(property)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(propertyID, updatedPropertyBytes)
}

// QueryProperty는 부동산 정보를 조회합니다.
func (t *ABstore) QueryProperty(ctx contractapi.TransactionContextInterface, propertyID string) (*Property, error) {
	propertyBytes, err := ctx.GetStub().GetState(propertyID)
	if err != nil {
		return nil, err
	}

	if propertyBytes == nil {
		return nil, fmt.Errorf("property %s does not exist", propertyID)
	}

	property := Property{}
	err = json.Unmarshal(propertyBytes, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

// TransferPropertyOwnership은 부동산 소유권을 이전합니다.
func (t *ABstore) TransferPropertyOwnership(ctx contractapi.TransactionContextInterface, propertyID, newOwnerID string) error {
	propertyBytes, err := ctx.GetStub().GetState(propertyID)
	if err != nil {
		return err
	}

	if propertyBytes == nil {
		return fmt.Errorf("property %s does not exist", propertyID)
	}

	property := Property{}
	err = json.Unmarshal(propertyBytes, &property)
	if err != nil {
		return err
	}

	property.OwnerID = newOwnerID
