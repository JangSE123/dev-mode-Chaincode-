package main

import (
   "encoding/json"
   "fmt"
   "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// User 구조체는 사용자 정보를 나타냅니다.
type User struct {
   UserID       string `json:"userID"`
   Name         string `json:"name"`
   BuildingName string `json:"buildingName"`
   Description  string `json:"description"`
}

// BuildingOwnership 체인코드의 구조체 정의를 확장합니다.
type BuildingOwnership struct {
   contractapi.Contract
}

// RegisterUser는 새로운 사용자를 등록합니다.
func (t *BuildingOwnership) RegisterUser(ctx contractapi.TransactionContextInterface, userID, name, buildingName, description string) error {
   fmt.Println("BuildingOwnership RegisterUser")

   user := User{
      UserID:       userID,
      Name:         name,
      BuildingName: buildingName,
      Description:  description,
   }

   userBytes, err := json.Marshal(user)
   if err != nil {
      return err
   }

   return ctx.GetStub().PutState(userID, userBytes)
}

// ViewBuilding는 부동산 정보를 조회합니다.
func (t *BuildingOwnership) ViewBuilding(ctx contractapi.TransactionContextInterface, buildingName string) (string, error) {
   userBytes, err := ctx.GetStub().GetState(buildingName)
   if err != nil {
      return "", fmt.Errorf("failed to get state for %s: %v", buildingName, err)
   }

   if userBytes == nil {
      return "", fmt.Errorf("Building not found %s", buildingName)
   }

   user := User{}
   err = json.Unmarshal(userBytes, &user)
   if err != nil {
      return "", fmt.Errorf("unmarshal error: %v", err)
   }

   userInfo, err := json.Marshal(user)
   if err != nil {
      return "", fmt.Errorf("marshal error: %v", err)
   }

   return string(userInfo), nil
}

func main() {
   cc, err := contractapi.NewChaincode(new(BuildingOwnership))
   if err != nil {
      panic(err.Error())
   }
   if err := cc.Start(); err != nil {
      fmt.Printf("Error starting BuildingOwnership chaincode: %s", err)
   }
}

// // TransferBuildingOwnership은 부동산 소유권을 이전합니다.
// func (t *BuildingOwnership) TransferBuildingOwnership(ctx contractapi.TransactionContextInterface, BuildingID, newOwnerID string) error {
// 	BuildingBytes, err := ctx.GetStub().GetState(BuildingID)
// 	if err != nil {
// 		return err
// 	}

// 	if BuildingBytes == nil {
// 		return fmt.Errorf("Building %s does not exist", BuildingID)
// 	}

// 	Building := Building{}
// 	err = json.Unmarshal(BuildingBytes, &Building)
// 	if err != nil {
// 		return err
// 	}

// 	Building.OwnerID = newOwnerID