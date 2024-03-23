/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
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
	}
}
