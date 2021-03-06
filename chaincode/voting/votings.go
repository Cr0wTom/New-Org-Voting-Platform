// Voting chaincode version 0.2 

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the vote structure, with 4 properties.  Structure tags are used by encoding/json library
type Asset_Votes struct {
	Vote string `json:"vote"`
	OwnerDesc string `json:"ownerdesc"`
	Voting string `json:"voting"`
}

type Votings struct {
	Name string `json:"name"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryVoter" {
		return s.queryVoter(APIstub, args)
	} else if function == "queryVotings" {
		return s.queryVotings(APIstub, args)
	} else if function == "queryAllVottings" {
		return s.queryAllVottings(APIstub)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "queryAllVotes" {
		return s.queryAllVotes(APIstub)
	} else if function == "doVoting" {
		return s.doVoting(APIstub, args)
	} else if function == "createVoting" {
		return s.createVoting(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// TODO: Query only voting categories to present in new platform
func (s *SmartContract) queryVotings(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	votingAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(votingAsBytes)
}

// Query votes
func (s *SmartContract) queryVoter(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	voterAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(voterAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	voters := []Asset_Votes{
	    Asset_Votes{Vote: "YES", OwnerDesc: "AUTh", Voting:"Test Voting"},
		Asset_Votes{Vote: "NO", OwnerDesc: "Almerys", Voting:"Test Voting"},
		Asset_Votes{Vote: "YES", OwnerDesc: "Eolas", Voting:"Test Voting"},
		Asset_Votes{Vote: "YES", OwnerDesc: "Nuro", Voting:"Test Voting"},
	}

	votings := []Votings{
		Votings{Name: "Test1"},
		Votings{Name: "Test2"},
		Votings{Name: "Test3"},
		Votings{Name: "Test4"},
	}

	i := 0
	for i < len(voters) {
		fmt.Println("i is ", i)
		voterAsBytes, _ := json.Marshal(voters[i])
		APIstub.PutState("VOTER"+strconv.Itoa(i), voterAsBytes)
		fmt.Println("Added", voters[i])
		i = i + 1
	}

	j := 0
	for j < len(votings) {
		fmt.Println("j is ", j)
		votingAsBytes, _ := json.Marshal(votings[j])
		APIstub.PutState("VOTING"+strconv.Itoa(j), votingAsBytes)
		fmt.Println("Added", votings[j])
		j = j + 1
	}

	return shim.Success(nil)
}


func (s *SmartContract) queryAllVotes(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "VOTER0"
	endKey := "VOTER999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllVotes:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAllVottings(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "VOTING0"
	endKey := "VOTING999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllVottings:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) createVoting(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	votingAsBytes, _ := APIstub.GetState(args[0])
	name := Votings{}



	json.Unmarshal(votingAsBytes, &name)
	// if (name.Name == "VOTINGNAME"){
		fmt.Println("Voting accepted!")
		name.Name = args[1]
//	}else{
//		return shim.Error("Voting Exists! You can only create the same voting once!")
//	}


	votingAsBytes, _ = json.Marshal(name)
	APIstub.PutState(args[0], votingAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) doVoting(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	voterAsBytes, _ := APIstub.GetState(args[0])
	vote := Asset_Votes{}



	json.Unmarshal(voterAsBytes, &vote)
	//if (vote.Vote == "VOTER"){
		fmt.Println("You havent voted before, vote accepted!")
		vote.Vote = args[1]
		vote.Voting = args[2]
		vote.OwnerDesc = args[3]
	//}else{
	//	return shim.Error("Already voted before!, you can only vote once")
	//}


	voterAsBytes, _ = json.Marshal(vote)
	APIstub.PutState(args[0], voterAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
