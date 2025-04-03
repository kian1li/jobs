package chaincode

import (
	"realestatetransfer"
	"encoding/json"
	"strings"
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/jinzhu/inflection"
)

func checkLen(logger *shim.ChaincodeLogger, expected int, args []string) error {
	if len(args) < expected {
		mes:= fmt.Sprintf(
			"not enough number of arguments: %d given, %d expected",
			len(args),
			expected,
		)
		logger.Warning(mes)
		return errors.New(mes)
	}
	return nil
}

type RealestateTransferCC struct {
}

func (this *RealestateTransferCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger := shim.NewLogger("realestatetransfer")
	logger.Info("chaincode initialized")
	return shim.Success([]byte{})
}

func (this *RealestateTransferCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger := shim.NewLogger("realestatetransfer")

	//sample of API use: show t X timestamp
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get TX timestamp: %s", err))
	}
	logger.Infof(
		"Invoke called: Tx ID = %s, timestamp = %s",
		stub.GetTxID(),
		timestamp,
	)
	var (
		fcn string
		args []string
	)
	fcn, args = stub.GetFunctionAndParameters()
	logger.Infof("function name = %s", fcn)

	switch fcn {
	//adds a new Participant
	case "AddParticipant":
		//check arguments length
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		// ummarshal
		participant := new(realestatetransfer.Participant)
		err := json.Unmarshal([]byte(args[0]), participant)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal Participant JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.AddParticipant(stub, participant)
		if err != nil {
			return shim.Error(err.Error())
		}

		//return success value
		return shim.Success([]byte{})

	// get an existing participant
	case "GetParticipant":
		// checks arguments length
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		var id string
		err := json.Unmarshal([]byte(args[0]), &id)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st argument: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		participant, err := this.GetParticipant(stub, id)
		if err != nil {
			return shim.Error(err.Error())
		}

		//marshal
		b, err := json.Marshal(participant)
		if err != nil {
			mes := fmt.Sprintf("failed to marshal participant: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		//return a success value
		return shim.Success(b)

	// updates an existing Paricipant
	case "UpdateParticipant":
		// checks arguments length
		if err := checkLen(logger, 1, args); err != err {
			return shim.Error(err.Error())
		}

		//unmarshal
		participant := new(realestatetransfer.Participant)
		err := json.Unmarshal([]byte(args[0]), participant)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal Participant JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.UpdateParticipant(stub, participant)
		if err != nil {
			return shim.Error(err.Error())
		}

		//return success value
		return shim.Success([]byte{})


	//lend token an existing Particicant
	case "TransferToken":
		// check arguments length
		if err := checkLen(logger, 3, args); err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		var fromParticipantId, toParticipantId, tokenAmount string
		err := json.Unmarshal([]byte(args[0]), &fromParticipantId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[1]), &toParticipantId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 2nd arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[2]), &tokenAmount)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 3th arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = this.TransferToken(stub, fromParticipantId, toParticipantId, tokenAmount)
		if err != nil {
			return shim.Error(err.Error())
		}

	// lists ListParticipants
	case "ListParticipants":
		ownlist, err := this.ListParticipants(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

		//marshal
		b, err := json.Marshal(ownlist)
		if err != nil {
			mes := fmt.Sprintf("failed to marshal Ownlist: %s", err.Error())

			logger.Warning(mes)
			return shim.Error(mes)
		}
		//return success value
		return shim.Success(b)

	// AddAtomTokenTransfer
	case "AddAtomTokenTransfer":
		//check arguments length
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		// ummarshal
		atomTokenTransfer := new(realestatetransfer.AtomTokenTransfer)
		err := json.Unmarshal([]byte(args[0]), atomTokenTransfer)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal AtomTokenTransfer JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.AddAtomTokenTransfer(stub, atomTokenTransfer)
		if err != nil {
			return shim.Error(err.Error())
		}

		//return success value
		return shim.Success([]byte{})

	// AddTransactionLedgerbyAttr		
	case "AddAtomTokenTransferbyAttr":
		// check arguments length
		if err := checkLen(logger, 3, args); err != nil {
			return shim.Error(err.Error())
		}
		// unmarshal
		var participantId, transferType, tokenAmount string
		err := json.Unmarshal([]byte(args[0]), &participantId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[1]), &transferType)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 2nd arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[2]), &tokenAmount)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 3th arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = this.AddAtomTokenTransferbyAttr(stub, participantId, transferType, tokenAmount)
		if err != nil {
			return shim.Error(err.Error())
		}

	// get an existing TransactionLedger
	case "GetAtomTokenTransfer":
		// checks arguments length
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		var id string
		err := json.Unmarshal([]byte(args[0]), &id)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st argument: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		atomTokenTransfer, err := this.GetAtomTokenTransfer(stub, id)
		if err != nil {
			return shim.Error(err.Error())
		}

		//marshal
		b, err := json.Marshal(atomTokenTransfer)
		if err != nil {
			mes := fmt.Sprintf("failed to marshal realestate: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		//return a success value
		return shim.Success(b)

	// WithdrawnToken
	case "WithdrawnToken":
		// check arguments length
		if err := checkLen(logger, 2, args); err != nil {
			return shim.Error(err.Error())
		}
		// unmarshal
		var participantId, tokenAmount, atomTokenTransferId string
		err := json.Unmarshal([]byte(args[0]), &participantId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[1]), &tokenAmount)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 2th arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		atomTokenTransferId, err = this.WithdrawnToken(stub, participantId, tokenAmount)
		if err != nil {
			mes := fmt.Sprintf("failed to WithdrawnToken %s %s", atomTokenTransferId, err)
			logger.Warning(mes)
			return shim.Error(err.Error())
		}

	// DepositToken
	case "DepositToken":
		// check arguments length
		if err := checkLen(logger, 2, args); err != nil {
			return shim.Error(err.Error())
		}
		// unmarshal
		var participantId, tokenAmount, atomTokenTransferId string
		err := json.Unmarshal([]byte(args[0]), &participantId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[1]), &tokenAmount)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 2th arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		atomTokenTransferId, err = this.DepositToken(stub, participantId, tokenAmount)
		if err != nil {
			mes := fmt.Sprintf("failed to WithdrawnToken %s %s", atomTokenTransferId, err)
			logger.Warning(mes)
			return shim.Error(err.Error())
		}

	// AddTransactionLedger
	case "AddTransactionLedger":
		//check arguments length
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		// ummarshal
		transactionLedger := new(realestatetransfer.TransactionLedger)
		err := json.Unmarshal([]byte(args[0]), transactionLedger)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal TransactionLedger JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.AddTransactionLedger(stub, transactionLedger)
		if err != nil {
			return shim.Error(err.Error())
		}

		//return success value
		return shim.Success([]byte{})

	// AddTransactionLedgerbyAttr		
	case "AddTransactionLedgerbyAttr":
		// check arguments length
		if err := checkLen(logger, 3, args); err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		var transactionIds []string
		var transactionType, tokenAmount string
		err = json.Unmarshal([]byte(args[0]), &transactionType)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1th arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err := json.Unmarshal([]byte(args[1]), transactionIds)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 2nd arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[2]), &tokenAmount)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 3th arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = this.AddTransactionLedgerbyAttr(stub, transactionType , tokenAmount ,transactionIds...)
		if err != nil {
			return shim.Error(err.Error())
		}

	// get an existing TransactionLedger
	case "GetTransactionLedger":
		// checks arguments length
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		var id string
		err := json.Unmarshal([]byte(args[0]), &id)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st argument: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		transactionLedger, err := this.GetTransactionLedger(stub, id)
		if err != nil {
			return shim.Error(err.Error())
		}

		//marshal
		b, err := json.Marshal(transactionLedger)
		if err != nil {
			mes := fmt.Sprintf("failed to marshal realestate: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		//return a success value
		return shim.Success(b)


	// adds a new Realestate
	case "AddRealestate":
		// checks arguments length
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		realestate := new(realestatetransfer.Realestate)
		err := json.Unmarshal([]byte(args[0]), realestate)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal Realestate JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.AddRealestate(stub, realestate)
		if err != nil {
			return shim.Error(err.Error())
		}

		// return success value
		return shim.Success([]byte{})

	//list realestate
	case "ListRealestates":
		realestates, err := this.ListRealestates(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

		// marshal
		b, err := json.Marshal(realestates)
		if err != nil {
			mes := fmt.Sprintf("failed to marshal Realestates: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		// return a success value
		return shim.Success(b)

	// get an existing Realestate
	case "GetRealestate":
		// checks arguments length
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		var id string
		err := json.Unmarshal([]byte(args[0]), &id)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st argument: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		realestate, err := this.GetRealestate(stub, id)
		if err != nil {
			return shim.Error(err.Error())
		}

		//marshal
		b, err := json.Marshal(realestate)
		if err != nil {
			mes := fmt.Sprintf("failed to marshal realestate: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		//return a success value
		return shim.Success(b)

	// updates an existing Realestate
	case "UpdateRealestate":
		// checks arguments length
		if err := checkLen(logger, 1, args); err != err {
			return shim.Error(err.Error())
		}

		//unmarshal
		realestate := new(realestatetransfer.Realestate)
		err := json.Unmarshal([]byte(args[0]), realestate)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal Realestate JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.UpdateRealestate(stub, realestate)
		if err != nil {
			return shim.Error(err.Error())
		}

		//return success value
		return shim.Success([]byte{})

	//transfer an existing Realestate to an existing Participant
	case "TransferRealestate":
		// checks arguments length
		if err := checkLen(logger, 3, args); err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		var fromParticipantId, toParticipantId, realestateId string
		err := json.Unmarshal([]byte(args[0]), &fromParticipantId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st arguments: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[1]), &toParticipantId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 2nd argument: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}
		err = json.Unmarshal([]byte(args[2]), &realestateId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 3rd argument: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.TransferRealestate(stub, fromParticipantId, toParticipantId, realestateId)
		if err != nil {
			return shim.Error(err.Error())
		}

		//returns success value
		return shim.Success([]byte{})

	}
	// function name is unknown
	mes := fmt.Sprintf("Unknown method: %s", fcn)
	logger.Warning(mes)
	return shim.Error(mes)
}

// method implements

// Adds a new Participant
func (this *RealestateTransferCC) AddParticipant(stub shim.ChaincodeStubInterface, participant *realestatetransfer.Participant) error {
//	return errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("AddParticipant: Id = %s", participant.Id)

	// check owner exists
	found, err := this.CheckParticipant(stub, participant.Id)
	if err != nil {
		return err
	}
	if found {
		mes := fmt.Sprintf("Participant Id = %s already exists", participant.Id)
		logger.Warning(mes)
		return errors.New(mes)
	}
	// converts to JSON
	b, err := json.Marshal(participant)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	// create composite key
	key, err := stub.CreateCompositeKey("Participant", []string{participant.Id})
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// stores state DB
	err =  stub.PutState(key, b)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// returns successfully
	return nil
}

//Get specified Participant
func (this *RealestateTransferCC) GetParticipant(stub shim.ChaincodeStubInterface, id string) (*realestatetransfer.Participant, error) {
	//return nil, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("GetParticipant: Id = %s", id)

	// create composite key
	key, err := stub.CreateCompositeKey("Participant", []string{id})
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// load state DB
	jsonBytes, err := stub.GetState(key)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}
	if jsonBytes == nil {
		mes := fmt.Sprintf("Participant with Id = %s was not found", id)
		logger.Warning(mes)
		return nil, errors.New(mes)
	}

	// unmarshal
	participant := new(realestatetransfer.Participant)
	err = json.Unmarshal(jsonBytes, participant)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// return success
	return participant, nil
}

//Update the content of the specified Participant
func (this *RealestateTransferCC) UpdateParticipant(stub shim.ChaincodeStubInterface, participant *realestatetransfer.Participant) error {
	//return errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("UpdateParticipant: participant = %+v", participant)

	// check existence of the specified Realestate
	found, err := this.CheckParticipant(stub, participant.Id)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	if !found {
		mes := fmt.Sprintf("Participant with Id = %s does not exist", participant.Id)
		logger.Warning(mes)
		return errors.New(mes)
	}
	// validate the Realestate
	ok , err := this.ValidateParticipant(stub, participant)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	if !ok {
		mes := "Validate of the Participant failed"
		logger.Warning(mes)
		return errors.New(mes)
	}

	// create composite key
	key, err := stub.CreateCompositeKey("Participant", []string{participant.Id})
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// converts to JSON
	b, err := json.Marshal(participant)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// store State DB
	err = stub.PutState(key, b)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// return success
	return nil
}

//checks existence of the specified Participant
func (this *RealestateTransferCC) CheckParticipant(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	//return false, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("CheckParticipant: Id = %s", id)

	// create composite key
	key, err := stub.CreateCompositeKey("Participant", []string{id})
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}

	//load State DB
	jsonBytes, err := stub.GetState(key)
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}

	//return successful
	return jsonBytes != nil, nil
}

//validate the specified Participant
func (this *RealestateTransferCC) ValidateParticipant(stub shim.ChaincodeStubInterface, participant *realestatetransfer.Participant) (bool, error) {
	//return false, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("ValidateParticipant: Id = %s", participant.Id)

	//check existence Participant
	found, err := this.CheckParticipant(stub, participant.Id)
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}
	// return success
	return found, nil
}

// Lists Participants
func (this *RealestateTransferCC) ListParticipants(stub shim.ChaincodeStubInterface) ([]*realestatetransfer.Participant, error) {
	//return nil, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Info("ListParticipants")

	// query returns iterator
	iter, err := stub.GetStateByPartialCompositeKey("Participant", []string{})
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	//close iterator
	defer iter.Close()
	participants := []*realestatetransfer.Participant{}

	//loop iterator
	for iter.HasNext() {
		kv, err := iter.Next()
		if err != nil {
			logger.Warning(err.Error())
			return nil, err
		}
		participant := new(realestatetransfer.Participant)
		err = json.Unmarshal(kv.Value, participant)
		if err != nil {
			logger.Warning(err.Error())
			return nil, err
		}
		participants = append(participants, participant)
	}

	// returns lists
	if len(participants) > 1 {
		logger.Infof("%d %s found", len(participants), inflection.Plural("Participant"))
	} else {
		logger.Infof("%d %s found", len(participants), "Participant")
	}
	return participants, nil
}

// Transfer the specified Token from the specified Participant to the specified Participant
func (this *RealestateTransferCC) TransferToken(stub shim.ChaincodeStubInterface, fromParticipantId string, toParticipantId string, tokenAmount string) error {
	//return errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("TransferRealestate: Token Amount = %s, from Participant- %s to Participant - %s", tokenAmount, fromParticipantId, toParticipantId)
	withdrawnAtomTokenTransferId, err := this.WithdrawnToken(stub, fromParticipantId, tokenAmount)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	depositAtomTokenTransferId, err:= this.DepositToken(stub, toParticipantId, tokenAmount)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	err = this.AddTransactionLedgerbyAttr(stub, "Withdrawn", withdrawnAtomTokenTransferId, tokenAmount)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	err = this.AddTransactionLedgerbyAttr(stub, "Deposit", depositAtomTokenTransferId, tokenAmount)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	// return success
	return nil
}

// Adds a new AtomTokenTransfer
func (this *RealestateTransferCC) AddAtomTokenTransfer(stub shim.ChaincodeStubInterface, atomTokenTransfer *realestatetransfer.AtomTokenTransfer) error {
//	return errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("AddAtomTokenTransfer %s", atomTokenTransfer.Id)

	// check owner exists
	found, err := this.CheckTransactionLedger(stub, atomTokenTransfer.Id)
	if err != nil {
		return err
	}
	if found {
		mes := fmt.Sprintf("AtomTokenTransfer Id = %s already exists", atomTokenTransfer.Id)
		logger.Warning(mes)
		return errors.New(mes)
	}
	// converts to JSON
	b, err := json.Marshal(atomTokenTransfer)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	// create composite key
	key, err := stub.CreateCompositeKey("AtomTokenTransfer", []string{atomTokenTransfer.Id})
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// stores state DB
	err =  stub.PutState(key, b)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// returns successfully
	return nil
}

//checks existence of the specified Participant
func (this *RealestateTransferCC) CheckAtomTokenTransfer(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	//return false, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("CheckAtomTokenTransfer: Id = %s", id)

	// create composite key
	key, err := stub.CreateCompositeKey("AtomTokenTransfer", []string{id})
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}

	//load State DB
	jsonBytes, err := stub.GetState(key)
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}

	//return successful
	return jsonBytes != nil, nil
}

//validate the specified Participant
func (this *RealestateTransferCC) ValidateAtomTokenTransfer(stub shim.ChaincodeStubInterface, atomTokenTransfer *realestatetransfer.AtomTokenTransfer) (bool, error) {
	//return false, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("ValidateAtomTokenTransfer: Id = %s", atomTokenTransfer.Id)

	//check existence AtomTokenTransfer
	found, err := this.CheckAtomTokenTransfer(stub, atomTokenTransfer.Id)
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}
	// return success
	return found, nil
}

// Lists AtomTokenTransfers
func (this *RealestateTransferCC) ListAtomTokenTransfers(stub shim.ChaincodeStubInterface) ([]*realestatetransfer.AtomTokenTransfer, error) {
	//return nil, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Info("ListAtomTokenTransfers")

	// query returns iterator
	iter, err := stub.GetStateByPartialCompositeKey("AtomTokenTransfer", []string{})
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	//close iterator
	defer iter.Close()
	atomTokenTransfers := []*realestatetransfer.AtomTokenTransfer{}

	//loop iterator
	for iter.HasNext() {
		kv, err := iter.Next()
		if err != nil {
			logger.Warning(err.Error())
			return nil, err
		}
		atomTokenTransfer := new(realestatetransfer.AtomTokenTransfer)
		err = json.Unmarshal(kv.Value, atomTokenTransfer)
		if err != nil {
			logger.Warning(err.Error())
			return nil, err
		}
		atomTokenTransfers = append(atomTokenTransfers, atomTokenTransfer)
	}

	// returns lists
	if len(atomTokenTransfers) > 1 {
		logger.Infof("%d %s found", len(atomTokenTransfers), inflection.Plural("AtomTokenTransfer"))
	} else {
		logger.Infof("%d %s found", len(atomTokenTransfers), "AtomTokenTransfer")
	}
	return atomTokenTransfers, nil
}

//Get specified AtomTokenTransfer
func (this *RealestateTransferCC) GetAtomTokenTransfer(stub shim.ChaincodeStubInterface, id string) (*realestatetransfer.AtomTokenTransfer, error) {
	//return nil, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("GetAtomTokenTransfer: Id = %s", id)

	// create composite key
	key, err := stub.CreateCompositeKey("AtomTokenTransfer", []string{id})
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// load state DB
	jsonBytes, err := stub.GetState(key)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}
	if jsonBytes == nil {
		mes := fmt.Sprintf("AtomTokenTransfer with Id = %s was not found", id)
		logger.Warning(mes)
		return nil, errors.New(mes)
	}

	// unmarshal
	atomTokenTransfer := new(realestatetransfer.AtomTokenTransfer)
	err = json.Unmarshal(jsonBytes, atomTokenTransfer)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// return success
	return atomTokenTransfer, nil
}


// Add TransactionLedger by Attributes
func (this *RealestateTransferCC) AddAtomTokenTransferbyAttr(stub shim.ChaincodeStubInterface, participantId string, transferType string, tokenAmount string) error {
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("AddAtomTokenTransferbyAttr: %s - transfertype %s - tokenamount %s",participantId, transferType, tokenAmount)
	atomTokenTransfer := new(realestatetransfer.AtomTokenTransfer)
	atomTokenTransfers, err := this.ListAtomTokenTransfers(stub)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	atomTokenTransfer.Id = strconv.Itoa(len(atomTokenTransfers) + 1)
	atomTokenTransfer.ParticipantId = participantId
	atomTokenTransfer.TransferType = transferType
	atomTokenTransfer.TokenAmount = tokenAmount
	this.AddAtomTokenTransfer(stub, atomTokenTransfer)
	return nil
}

// Add TransactionLedger by WithdrawnToken 
func (this *RealestateTransferCC) WithdrawnToken(stub shim.ChaincodeStubInterface, participantId string, tokenAmount string) (string, error) {
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("WithdrawnToken: %s - tokenamount %s",participantId, tokenAmount)
	// get Participant by participantId
	withdrawnParticipant := new(realestatetransfer.Participant)
	withdrawnParticipant, err := this.GetParticipant(stub, participantId)
	if withdrawnParticipant.Token < tokenAmount {
		logger.Infof("This Participant-%s enough Token", participantId)
		return "", err
	}
	// updates token field
	withdrawnParticipantTokenAmount, _ := strconv.Atoi(withdrawnParticipant.Token)
	withdrawnTokenAmount, _ := strconv.Atoi(tokenAmount)
	withdrawnParticipant.Token = strconv.Itoa(withdrawnParticipantTokenAmount - withdrawnTokenAmount)
	this.UpdateParticipant(stub, withdrawnParticipant)

	atomTokenTransfer := new(realestatetransfer.AtomTokenTransfer)
	atomTokenTransfers, err := this.ListAtomTokenTransfers(stub)
	if err != nil {
		logger.Warning(err.Error())
		return "" ,err
	}
	atomTokenTransfer.Id = strconv.Itoa(len(atomTokenTransfers) + 1)
	atomTokenTransfer.ParticipantId = participantId
	atomTokenTransfer.TransferType = "Withdrawn"
	atomTokenTransfer.TokenAmount = tokenAmount
	this.AddAtomTokenTransfer(stub, atomTokenTransfer)
	return atomTokenTransfer.Id, nil
}

// Add TransactionLedger by Deposit
func (this *RealestateTransferCC) DepositToken(stub shim.ChaincodeStubInterface, participantId string, tokenAmount string) (string, error) {
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("DepositToken: %s - tokenamount %s",participantId, tokenAmount)
	// get Participant by participantId
	depositParticipant := new(realestatetransfer.Participant)
	depositParticipant, _ = this.GetParticipant(stub, participantId)
	// updates token field
	depositParticipantTokenAmount, _ := strconv.Atoi(depositParticipant.Token)
	depositTokenAmount, _ := strconv.Atoi(tokenAmount)
	depositParticipant.Token = strconv.Itoa(depositParticipantTokenAmount + depositTokenAmount)
	this.UpdateParticipant(stub, depositParticipant)

	atomTokenTransfer := new(realestatetransfer.AtomTokenTransfer)
	atomTokenTransfers, err := this.ListAtomTokenTransfers(stub)
	if err != nil {
		logger.Warning(err.Error())
		return "", err
	}
	atomTokenTransfer.Id = strconv.Itoa(len(atomTokenTransfers) + 1)
	atomTokenTransfer.ParticipantId = participantId
	atomTokenTransfer.TransferType = "Deposit"
	atomTokenTransfer.TokenAmount = tokenAmount
	this.AddAtomTokenTransfer(stub, atomTokenTransfer)
	return atomTokenTransfer.Id, nil
}

// Adds a new TransactionLedger
func (this *RealestateTransferCC) AddTransactionLedger(stub shim.ChaincodeStubInterface, transactionLedger *realestatetransfer.TransactionLedger) error {
//	return errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("AddTransactionLedger %s", transactionLedger.Id)

	// check owner exists
	found, err := this.CheckTransactionLedger(stub, transactionLedger.Id)
	if err != nil {
		return err
	}
	if found {
		mes := fmt.Sprintf("TransactionLedger Id = %s already exists", transactionLedger.Id)
		logger.Warning(mes)
		return errors.New(mes)
	}
	// converts to JSON
	b, err := json.Marshal(transactionLedger)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	// create composite key
	key, err := stub.CreateCompositeKey("TransactionLedger", []string{transactionLedger.Id})
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// stores state DB
	err =  stub.PutState(key, b)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// returns successfully
	return nil
}

//checks existence of the specified Participant
func (this *RealestateTransferCC) CheckTransactionLedger(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	//return false, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("CheckTransactionLedger: Id = %s", id)

	// create composite key
	key, err := stub.CreateCompositeKey("TransactionLedger", []string{id})
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}

	//load State DB
	jsonBytes, err := stub.GetState(key)
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}

	//return successful
	return jsonBytes != nil, nil
}

//validate the specified Participant
func (this *RealestateTransferCC) ValidateTransactionLedger(stub shim.ChaincodeStubInterface, transactionLedger *realestatetransfer.TransactionLedger) (bool, error) {
	//return false, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("ValidateTransactionLedger: Id = %s", transactionLedger.Id)

	//check existence Participant
	found, err := this.CheckParticipant(stub, transactionLedger.Id)
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}
	// return success
	return found, nil
}

// Lists Participants
func (this *RealestateTransferCC) ListTransactionLedgers(stub shim.ChaincodeStubInterface) ([]*realestatetransfer.TransactionLedger, error) {
	//return nil, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Info("ListTransactionLedgers")

	// query returns iterator
	iter, err := stub.GetStateByPartialCompositeKey("TransactionLedger", []string{})
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	//close iterator
	defer iter.Close()
	transactionLedgers := []*realestatetransfer.TransactionLedger{}

	//loop iterator
	for iter.HasNext() {
		kv, err := iter.Next()
		if err != nil {
			logger.Warning(err.Error())
			return nil, err
		}
		transactionLedger := new(realestatetransfer.TransactionLedger)
		err = json.Unmarshal(kv.Value, transactionLedger)
		if err != nil {
			logger.Warning(err.Error())
			return nil, err
		}
		transactionLedgers = append(transactionLedgers, transactionLedger)
	}

	// returns lists
	if len(transactionLedgers) > 1 {
		logger.Infof("%d %s found", len(transactionLedgers), inflection.Plural("TransactionLedger"))
	} else {
		logger.Infof("%d %s found", len(transactionLedgers), "TransactionLedger")
	}
	return transactionLedgers, nil
}

//Get specified TransactionLedger
func (this *RealestateTransferCC) GetTransactionLedger(stub shim.ChaincodeStubInterface, id string) (*realestatetransfer.TransactionLedger, error) {
	//return nil, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("GetTransactionLedger: Id = %s", id)

	// create composite key
	key, err := stub.CreateCompositeKey("TransactionLedger", []string{id})
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// load state DB
	jsonBytes, err := stub.GetState(key)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}
	if jsonBytes == nil {
		mes := fmt.Sprintf("TransactionLedger with Id = %s was not found", id)
		logger.Warning(mes)
		return nil, errors.New(mes)
	}

	// unmarshal
	transactionLedger := new(realestatetransfer.TransactionLedger)
	err = json.Unmarshal(jsonBytes, transactionLedger)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// return success
	return transactionLedger, nil
}


// Add TransactionLedger by Attributes
func (this *RealestateTransferCC) AddTransactionLedgerbyAttr(stub shim.ChaincodeStubInterface, transactionType string, amountToken string, transactionIds ...string) error {
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("AddTransactionLedger: %s - AmountToken %s", strings.Join(transactionIds, ", "), amountToken)
	transactionLedger := new(realestatetransfer.TransactionLedger)
	transactionLedgers, err := this.ListTransactionLedgers(stub)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	transactionLedger.Id = strconv.Itoa(len(transactionLedgers) + 1)
	transactionLedger.TransactionType = transactionType
	transactionLedger.TransactionIds = transactionIds
	transactionLedger.AmountToken = amountToken
	this.AddTransactionLedger(stub, transactionLedger)
	return nil
}

// Adds a new Realestate
func (this *RealestateTransferCC) AddRealestate(stub shim.ChaincodeStubInterface, realestate *realestatetransfer.Realestate) error {
	//return errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("AddRealestate: Id = %s", realestate.Id)

	// create composite key
	key, err := stub.CreateCompositeKey("Realestate", []string{realestate.Id})
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	//check Realestate exists
	found, err := this.CheckRealestate(stub, realestate.Id)
	if found {
		mes := fmt.Sprintf("Realestate with Id = %s already exists", realestate.Id)
		logger.Warning(mes)
		return errors.New(mes)
	}

	// validate the Realestate
	ok, err := this.ValidateRealestate(stub, realestate)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	if !ok {
		mes := "Validate of the Realestate failed"
		logger.Warning(mes)
		return errors.New(mes)
	}

	// converts to JSON
	b, err := json.Marshal(realestate)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// store State DB
	err = stub.PutState(key, b)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	//return successful
	return nil
}

//check existence of the specified Realestate
func (this *RealestateTransferCC) CheckRealestate(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	//return false, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("CheckRealestate: Id = %s", id)

	// creates a composite key
	key, err := stub.CreateCompositeKey("Realestate", []string{id})
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}

	// load State DB
	jsonBytes, err := stub.GetState(key)
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}

	//return success
	return jsonBytes != nil, nil
}

//validate the specified Realestate
func (this *RealestateTransferCC) ValidateRealestate(stub shim.ChaincodeStubInterface, realestate *realestatetransfer.Realestate) (bool, error) {
	//return false, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("ValidateRealestate: Id = %s", realestate.Id)

	//check existence OwnerId
	found, err := this.CheckParticipant(stub, realestate.OwnerId)
	if err != nil {
		logger.Warning(err.Error())
		return false, err
	}
	// return success
	return found, nil
}

//Get specified Realestate
func (this *RealestateTransferCC) GetRealestate(stub shim.ChaincodeStubInterface, id string) (*realestatetransfer.Realestate, error) {
	//return nil, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("GetRealestate: Id = %s", id)

	// create composite key
	key, err := stub.CreateCompositeKey("Realestate", []string{id})
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// load state DB
	jsonBytes, err := stub.GetState(key)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}
	if jsonBytes == nil {
		mes := fmt.Sprintf("Realestate with Id = %s was not found", id)
		logger.Warning(mes)
		return nil, errors.New(mes)
	}

	// unmarshal
	realestate := new(realestatetransfer.Realestate)
	err = json.Unmarshal(jsonBytes, realestate)
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// return success
	return realestate, nil
}

//Update the content of the specified Realestate
func (this *RealestateTransferCC) UpdateRealestate(stub shim.ChaincodeStubInterface, realestate *realestatetransfer.Realestate) error {
	//return errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("UpdateRealestate: realestate = %+v", realestate)

	// check existence of the specified Realestate
	found, err := this.CheckRealestate(stub, realestate.Id)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	if !found {
		mes := fmt.Sprintf("Realestate with Id = %s does not exist", realestate.Id)
		logger.Warning(mes)
		return errors.New(mes)
	}
	// validate the Realestate
	ok , err := this.ValidateRealestate(stub, realestate)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	if !ok {
		mes := "Validate of the Realestate failed"
		logger.Warning(mes)
		return errors.New(mes)
	}

	// create composite key
	key, err := stub.CreateCompositeKey("Realestate", []string{realestate.Id})
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// converts to JSON
	b, err := json.Marshal(realestate)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// store State DB
	err = stub.PutState(key, b)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// return success
	return nil
}

// Lists Realestate
func (this *RealestateTransferCC) ListRealestates(stub shim.ChaincodeStubInterface) ([]*realestatetransfer.Realestate, error) {
	//return nil, errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Info("ListRealestates")

	// excute return iterator
	iter, err := stub.GetStateByPartialCompositeKey("Realestate", []string{})
	if err != nil {
		logger.Warning(err.Error())
		return nil, err
	}

	// close iterator this method
	defer iter.Close()

	// loops iterator
	realestates := []*realestatetransfer.Realestate{}
	for iter.HasNext() {
		kv, err := iter.Next()
		if err != nil {
			logger.Warning(err.Error())
			return nil, err
		}
		realestate := new(realestatetransfer.Realestate)
		err = json.Unmarshal(kv.Value, realestate)
		if err != nil {
			logger.Warning(err.Error())
			return nil, err
		}
		realestates = append(realestates, realestate)
	}

	// return success
	if len(realestates) > 1 {
		logger.Infof("%d %s found", len(realestates), inflection.Plural("Realestate"))
	} else {
		logger.Infof("%d %s found", len(realestates), "Realestate")
	}
	return realestates, nil
}

// Transfer the specified Realestate to the specified Owner
func (this *RealestateTransferCC) TransferRealestate(stub shim.ChaincodeStubInterface, fromParticipantId string, toParticipantId string, realestateId string) error {
	//return errors.New("Not yet")
	logger := shim.NewLogger("realestatetransfer")
	logger.Infof("TransferRealestate: Realestate Id = %s, fromParticipant - %s toParticipant - %s", realestateId, fromParticipantId, toParticipantId)

	// get specified fromParticipant (err return it not exist)
	//fromParticipant, err := this.GetParticipant(stub, fromParticipantId)
	//if err != nil {
	//	logger.Warning(err.Error())
	//	return err
	//}

	// get specified toParticipant (err return it not exist)
	toParticipant, err := this.GetParticipant(stub, toParticipantId)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// get specified Realestate (err return it not exist)
	realestate, err := this.GetRealestate(stub, realestateId)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}
	// if realestate.IsTrading false quit transaction
	isTrading, err := strconv.ParseBool(realestate.IsTrading)
	if err != nil {
		logger.Infof("IsTrading is format convertfail")
		logger.Warning(err.Error())
		return err
	}
	if !isTrading {
		logger.Infof("This Realestate not Trading now %s", realestateId)
		return err
	}

	// updates token field
	intToTokenAmount, _ := strconv.Atoi(toParticipant.Token)
	intTransactionPrice, _ := strconv.Atoi(realestate.TransactionPrice)


	// if Participant Token < realestate.TransactionPrice
	if intToTokenAmount < intTransactionPrice {
		logger.Infof("%s is not enoungh Token", toParticipant.Id)
		return err
	}
	this.TransferToken(stub,toParticipant.Id, fromParticipantId, realestate.TransactionPrice)

	// updates OwnerId field
	realestate.OwnerId = toParticipant.Id

	// store the update Realestate back State DB
	err = this.UpdateRealestate(stub, realestate)
	if err != nil {
		logger.Warning(err.Error())
		return err
	}

	// return success
	return nil
}
