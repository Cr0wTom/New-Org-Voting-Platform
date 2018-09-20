package main

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("mylogger")

type VoteCategory struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type PersonalInfo struct {
	Organization string `json:"org"`
	Name string `json:"username"`
}

func UpdateVote(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering UpdateVote")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for vote update")
	}

	var voteId = args[0]
	var status = args[1]

	laBytes, err := stub.GetState(voteId)
	if err != nil {
		logger.Error("Could not fetch vote category from ledger", err)
		return nil, err
	}
	var voteCategory VoteCategory
	err = json.Unmarshal(laBytes, &voteCategory)
	voteCategory.Status = status

	laBytes, err = json.Marshal(&lvoteCategory)
	if err != nil {
		logger.Error("Could not marshal vote category post update", err)
		return nil, err
	}

	err = stub.PutState(voteId, laBytes)
	if err != nil {
		logger.Error("Could not save vote category post update", err)
		return nil, err
	}

	var customEvent = "{eventType: 'voteCategoryUpdate', description:" + voteId + "' Successfully updated status'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully updated vote category")
	return nil, nil

}

func CreateNewVote(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering CreateNewVote")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for new vote creation")
	}

	var voteId = args[0]
	var voteInput = args[1]

	err := stub.PutState(voteId, []byte(voteInput))
	if err != nil {
		logger.Error("Could not save new vote to ledger", err)
		return nil, err
	}

	var customEvent = "{eventType: 'voteCreation', description:" + voteId + "' Successfully created'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully saved new vote.")
	return nil, nil

}

func GetVote(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering GetVote")

	if len(args) < 1 {
		logger.Error("Invalid number of arguments")
		return nil, errors.New("Missing vote ID")
	}

	var voteId = args[0]
	bytes, err := stub.GetState(voteId)
	if err != nil {
		logger.Error("Could not fetch vote with id "+voteId+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}


func (t *VotingChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "GetVote" {
		return GetVote(stub, args)
	}
	return nil, nil
}

func (t *VotingChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreateNewVote" {
		username, _ := GetCertAttribute(stub, "username")
		role, _ := GetCertAttribute(stub, "role")
		if role == "admin" {
			return CreateNewVote(stub, args)
		} else {
			return nil, errors.New(username + " with role " + role + " does not have access to create new vote")
		}

	}
	return nil, nil
}

func main() {

	lld, _ := shim.LogLevel("DEBUG")
	fmt.Println(lld)

	logger.SetLevel(lld)
	fmt.Println(logger.IsEnabledFor(lld))

	err := shim.Start(new(VotingChaincode))
	if err != nil {
		logger.Error("Could not start VotingChaincode")
	} else {
		logger.Info("VotingChaincode successfully started")
	}

}
