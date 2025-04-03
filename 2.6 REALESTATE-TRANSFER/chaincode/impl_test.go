package chaincode_test

import (
	"sort"
	"fmt"
	"encoding/json"
	"realestatetransfer/chaincode"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	//Finance
	oneBank = `{"Id":"1","Name":"OneBank","Token":"1000000"}`
	tworieBank = `{"Id":"2","Name":"TworieBank","Token":"2000000"}`

	//Company
	threeGroup = `{"Id":"3","Name":"ThreeGroup","Token":"300000"}`
	fourStar = `{"Id":"4","Name":"FourStar","Token":"40000000"}`

	//Participant
	alice = `{"Id":"101","Name":"Alice","Token":"100"}`
	alice_zero = `{"Id":"101","Name":"Alice","Token":"0"}`
	bob = `{"Id":"102","Name":"Bob","Token":"100"}`
	bob_200 = `{"Id":"102","Name":"Bob","Token":"200"}`
	carol = `{"Id":"103","Name":"carol","Token":"100"}`
	dave = `{"Id":"104","Name":"dave","Token":"100"}`
	eve = `{"Id":"105","Name":"eve","Token":"100"}`
	fayithe = `{"Id":"106","Name":"fayithe","Token":"100"}`
	grace = `{"Id":"107","Name":"grace","Token":"0"}`
	heidi = `{"Id":"108","Name":"heidi","Token":"0"}`
	ivan = `{"Id":"109", "Name":"ivan", "Token":"0"}`
	justin = `{"Id":"110","Name":"justin","Token":"50"}`
	mallory = `{"Id":"111","Name":"mallory","Token":"10"}`


	emptyParticipant = "[]"
	oneParticipant = "["+alice+"]"
	twoParticipants = "["+alice+","+bob+"]"

	timestamp = `"2020-04-18T12:34:56Z"`

	//AtomTokenTransfer
	atomtokentransfer1 = `{"Id":"1", "ParticipantId":"3", "TransferType":"Withdrawn", "TokenAmount":"100"}`
	atomtokentransfer2 = `{"Id":"2", "ParticipantId":"3", "TransferType":"Withdrawn", "TokenAmount":"100"}`
	atomtokentransfer3 = `{"Id":"3", "ParticipantId":"3", "TransferType":"Withdrawn", "TokenAmount":"100"}`
	atomtokentransfer4 = `{"Id":"4", "ParticipantId":"107", "TransferType":"Deposit", "TokenAmount":"100"}`
	atomtokentransfer5 = `{"Id":"5", "ParticipantId":"108", "TransferType":"Deposit", "TokenAmount":"100"}`
	atomtokentransfer6 = `{"Id":"6", "ParticipantId":"109", "TransferType":"Deposit", "TokenAmount":"100"}`
	atomtokentransfer7 = `{"Id":"7", "ParticipantId":"105", "TransferType":"Withdrawn", "TokenAmount":"8"}`
	atomtokentransfer8 = `{"Id":"8", "ParticipantId":"106", "TransferType":"Withdrawn", "TokenAmount":"2"}`
	atomtokentransfer9 = `{"Id":"9", "ParticipantId":"103", "TransferType":"Deposit", "TokenAmount":"5"}`
	atomtokentransfer10 = `{"Id":"10", "ParticipantId":"104", "TransferType":"Deposit", "TokenAmount":"5"}`
	atomtokentransfer11 = `{"Id":"11", "ParticipantId":"110", "TransferType":"Withdrawn", "TokenAmount":"7"}`
	atomtokentransfer12 = `{"Id":"12", "ParticipantId":"111", "TransferType":"Withdrawn", "TokenAmount":"3"}`
	atomtokentransfer13 = `{"Id":"13", "ParticipantId":"105", "TransferType":"Deposit", "TokenAmount":"8"}`
	atomtokentransfer14 = `{"Id":"14", "ParticipantId":"106", "TransferType":"Deposit", "TokenAmount":"2"}`

	//TransactionLedger
	transactionLedger1 = `{"Id":"1","TransactionType":"Withdrawn","TransactionIds":["1","2","3"],"AmountToken":"300"}`
	transactionLedger2 = `{"Id":"2","TransactionType":"Deposit","TransactionIds":["4","5","6"],"AmountToken":"300"}`
	transactionLedger3 = `{"Id":"3","TransactionType":"Withdrawn","TransactionIds":["7","8"],"AmountToken":"10"}}`
	transactionLedger4 = `{"Id":"3","TransactionType":"Deposit","TransactionIds":["9","10"],"AmountToken":"10"}}`
	transactionLedger5 = `{"Id":"4","TransactionType":"Withdrawn","TransactionIds":["11","12"],"AmountToken":"10"}}`
	transactionLedger6 = `{"Id":"4","TransactionType":"Deposit","TransactionIds":["13","14"],"AmountToken":"10"}}`

	//Realestate
	realestate1 = `{"Id":"1","Name":"eileen-garden","OwnerId":"101","IsTrading":"true","TransactionPrice":"200","Timestamp":`+timestamp+`}`
	realestate1b = `{"Id":"1","Name":"eileen-garden","OwnerId":"102","IsTrading":"true","TransactionPrice":"200","Timestamp":`+timestamp+`}`
	realestate2 = `{"Id":"2","Name":"maetan-hillstate","OwnerId":"101","IsTrading":"true","TransactionPrice":"100","Timestamp":`+timestamp+`}`
	realestate2b = `{"Id":"2","Name":"maetan-hillstate","OwnerId":"102","IsTrading":"true","TransactionPrice":"100","Timestamp":`+timestamp+`}`
	realestate3 = `{"Id":"3","Name":"woncheon-jugong","OwnerId":"103","OwnerList":{"103":"25","104":"25"},"IsTrading":"false","TransactionPrice":"50","Timestamp":`+timestamp+`}`
	realestate4 = `{"Id":"4","Name":"maetan-jugong","OwnerId":"102","TransactionList":["2"]","IsTrading":"true","TransactionPrice":"10","Timestamp":`+timestamp+`}`

	oneRealestate = "["+realestate1+"]"
	twoRealestates = "["+realestate1+","+realestate2+"]"

	one = `"1"`
	two = `"2"`
	three = `"3"`
	four = `"4"`
	five = `"5"`
	six = `"6"`
	seven = `"7"`
	eight = `"8"`
	nine = `"9"`
	ten = `"10"`
	threeAndFour = `{"OwnerList":{"103":"5","104":"5"}}`
	fiveAndSix = `{"OwnerList":{"105":"3","106":"7"}}`
	hundred  = `"100"`
	oneOOne = `"101"`
	oneOTwo = `"102"`
	oneOThree = `"103"`
	oneOFour = `"104"`
	oneOFive = `"105"`
	oneOSix = `"106"`
	oneOSeven = `"107"`
	oneOEight = `"108"`
	oneONine = `"109"`
	twoHundred = `"200"`
)

func sortbyMap(m map[string] interface{}) string {
	keys := make([]string, 0, len(m))

	for _, k  := range keys {
		fmt.Printf("%s", k)
		keys = append(keys, k)
	}

	sort.Strings(keys)
	sortData, err := json.Marshal(keys)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	fmt.Printf("%s", sortData)
	return string(sortData)
}


// test
func responseOK(res pb.Response) func() bool {
	return func() bool { return res.Status < shim.ERRORTHRESHOLD }
}

// response Fail
func responseFail(res pb.Response) func() bool {
	return func() bool { return res.Status >= shim.ERRORTHRESHOLD }
}

// Convert function name & arguments into a byte format that MockStub accept
func getBytes(function string, args ...string) [][]byte {
	bytes := make([][]byte, 0, len(args)+1)
	bytes = append(bytes, []byte(function))
	for _, s := range args {
		bytes = append(bytes, []byte(s))
	}
	return bytes
}

// Ok1: normal init
func TestInit_Ok1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) {
		res := stub.MockInit(util.GenerateUUID(), nil)
		assert.Condition(t, responseOK(res))
	}
}

// NG1: unknown method Invoke
func TestInvoke_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))

	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("Unknown"))
		   assert.Condition(t, responseFail(res))
	   }
}

// OK1: success
func TestAddParticipant_OK(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("ListParticipants", one))
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, emptyParticipant, string(res.Payload))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("ListParticipants", oneOOne))
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, oneParticipant, string(res.Payload))
	   }
}

// NG1: less args
func TestAddParticipant_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant"))
		   assert.Condition(t, responseFail(res))
	   }
}

// NG2 illegal JSON args
func TestAddParticipant_NG2(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", "Not yet"))
		   assert.Condition(t, responseFail(res))
	   }
}

// Ok1 : 1 Participant
func TestListParticipants_OK1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("ListParticipants"))
		   assert.Condition(t, responseOK(res))
		   t.Logf("%s", res.Payload)
		   assert.JSONEq(t, oneParticipant, string(res.Payload))
	   }
}

// NG1:not arg check
func TestCheckParticipant_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("CheckParticipant"))
		   assert.Condition(t, responseFail(res))
	   }
}

// NG1:not arg listparticipant 
func TestListParticipants_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("ListParticipants", one))
		   assert.Condition(t, responseFail(res))
	   }
}

// Ok2 : 2 Participants
func TestListParticipants_OK2(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("ListParticipants"))
		   assert.Condition(t, responseOK(res))
		   t.Logf("%s", twoParticipants)
		   t.Logf("%s", res.Payload)
		   assert.JSONEq(t, twoParticipants, string(res.Payload))
	   }
}

// OK1: a single Realestate
func TestAddRealestate_OK1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("GetRealestate", one))
		   assert.Condition(t, responseFail(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate1))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("ListRealestates", one))
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, oneRealestate, string(res.Payload))
	   }
}

// NG1: a single Realestate no owner
func TestAddRealestate_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("GetRealestate", one))
		   assert.Condition(t, responseFail(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate1))
		   assert.Condition(t, responseFail(res))
	   }
}

// NG1 illegal JSON args
func TestAddRealestate_NG2(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", "Not yet"))
		   assert.Condition(t, responseFail(res))
	   }
}

// OK2: two Realestates
func TestListRealestates_OK2(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("GetRealestate", one))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate1))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate2))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("ListRealestates"))
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, twoRealestates, string(res.Payload))
	   }
}

// OK1: change owner alice to bob
func TestUpdateRealestate_OK1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate1))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("UpdateRealestate", realestate1b))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetRealestate", one))
		   t.Logf("%s", res.Payload)
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, realestate1b, string(res.Payload))
	   }
}

// NG1: specified realestate not exist
func TestUpdateRealestate_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		  res := stub.MockInvoke(util.GenerateUUID(), getBytes("UpdateRealestate", realestate1b))
		  assert.Condition(t, responseFail(res))
	   }
}

// NG1: realestate not exist
func TestTransferRealestate_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferRealestate", oneOOne, oneOTwo))
		   assert.Condition(t, responseFail(res))
	   }
}

// NG2: owner not exist1
func TestTransferRealestate_NG2(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate1))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferRealestate", oneOOne, oneOTwo))
		   assert.Condition(t, responseFail(res))
	   }
}

// NG3: less args
func TestTransferRealestate_NG3(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferRealestate", one))
		   assert.Condition(t, responseFail(res))
	   }
}

// NG3: bad args
func TestTransferRealestate_NG4(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferRealestate", "not yet"))
		   assert.Condition(t, responseFail(res))
	   }
}

// NG3: bad args
func TestValidateRealestate_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("ValidateRealestate", "not yet"))
		   assert.Condition(t, responseFail(res))
	   }
}

// OK2: transfer from Alice to Bob
func TestTransferRealestate_OK(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate2))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferRealestate", oneOOne, oneOTwo, two))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneOOne))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneOTwo))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetRealestate", two))
		   t.Logf("%s", res.Payload)
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, realestate2b, string(res.Payload))
	   }
}

// NG5: Not enough TokenAmount
func TestTransferRealestate_NG5(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate1))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferRealestate", oneOOne, oneOTwo, one))
		   assert.Condition(t, responseFail(res))
	   }
}


// NG6: Not IsTrading
func TestTransferRealestate_NG6(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate3))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferRealestate", oneOTwo, oneOOne, three))
		   assert.Condition(t, responseFail(res))
	   }
}

// NG1: transfer token from Alice to Bob not enough Token
func TestTransferToken_NG1(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferToken", oneOOne, oneOTwo, twoHundred))
		   assert.Condition(t, responseFail(res))
	   }
}

// OK1: CoOwnerRealestate
func TestCoOwnerRealestate_OK(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", carol))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", dave))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddRealestate", realestate3))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetRealestate", three))
		   t.Logf("%s", res.Payload)

		   //assert.Condition(t, responseOK(res))
	   }
}

// OK1: transfer token from Alice to Bob Add TransactionLedger
func TestTransferToken_OK2(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferToken", oneOOne, oneOTwo, hundred))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneOOne))
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, alice_zero, string(res.Payload))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneOTwo))
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, bob_200, string(res.Payload))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetTransactionLedger", one))
		   t.Logf("%s", res.Payload)
	   }
}


// OK1: transfer token from Alice to Bob
func TestTransferToken_OK(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", alice))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", bob))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("TransferToken", oneOOne, oneOTwo, hundred))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneOOne))
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, alice_zero, string(res.Payload))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneOTwo))
		   assert.Condition(t, responseOK(res))
		   assert.JSONEq(t, bob_200, string(res.Payload))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetTransactionLedger", one))
		   t.Logf("%s", res.Payload)
	   }
}

// OK1: Multi Token Transfer
func TestAtomTokenTransfer_OK(t *testing.T) {
	stub := shim.NewMockStub("realestatetransfer", new(chaincode.RealestateTransferCC))
	if assert.NotNil(t, stub) &&
	   assert.Condition(t, responseOK(stub.MockInit(util.GenerateUUID(), nil))) {
		   res := stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", threeGroup))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", grace))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", heidi))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddParticipant", ivan))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("WithdrawnToken", three, hundred))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("WithdrawnToken", three, hundred))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("WithdrawnToken", three, hundred))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("DepositToken", oneOSeven, hundred))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("DepositToken", oneOEight, hundred))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("DepositToken", oneONine, hundred))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddTransactionLedger", transactionLedger1))
		   assert.Condition(t, responseOK(res))
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("AddTransactionLedger", transactionLedger2))
		   assert.Condition(t, responseOK(res))

		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", three))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneOSeven))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneOEight))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetParticipant", oneONine))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetAtomTokenTransfer", one))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetAtomTokenTransfer", two))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetAtomTokenTransfer", three))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetAtomTokenTransfer", four))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetAtomTokenTransfer", five))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetAtomTokenTransfer", six))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetTransactionLedger", one))
		   t.Logf("%s", res.Payload)
		   res = stub.MockInvoke(util.GenerateUUID(), getBytes("GetTransactionLedger", two))
		   t.Logf("%s", res.Payload)
	   }
}



