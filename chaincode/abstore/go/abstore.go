package main

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// RealEstateChaincode 구조체 정의
type RealEstateChaincode struct {
	contractapi.Contract
}

// 소유주 등록 함수
func (t *RealEstateChaincode) CreateOwner(ctx contractapi.TransactionContextInterface, ownerID string) error {
	// TODO: 구현
	return nil
}

// 구매자 등록 함수
func (t *RealEstateChaincode) CreateCustomer(ctx contractapi.TransactionContextInterface, customerID string) error {
	// TODO: 구현
	return nil
}

// 건물 소유권 등록 함수
func (t *RealEstateChaincode) CreateBuildingOwnership(ctx contractapi.TransactionContextInterface, ownerID, buildingID string) error {
	// TODO: 구현
	return nil
}

// 건물 소유권 조회 함수
func (t *RealEstateChaincode) ViewBuildingOwnership(ctx contractapi.TransactionContextInterface, buildingID string) (string, error) {
	// TODO: 구현
	return "", nil
}

// 건물 소유권 교환 함수
func (t *RealEstateChaincode) TradeBuildingOwnership(ctx contractapi.TransactionContextInterface, buyerID, buildingID string) error {
	// TODO: 구현
	return nil
}

func main() {
	cc, err := contractapi.NewChaincode(new(RealEstateChaincode))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting RealEstateChaincode: %s", err)
=======
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
>>>>>>> 5bab1385e2c6b2598e48f5ebf1d48e820fb27eb4
	}

	property.OwnerID = newOwnerID
