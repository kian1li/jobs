package realestatetransfer

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)


type Participant struct {
	Id string
	Name string
	Token string
}

type AtomTokenTransfer struct {
	Id string
	ParticipantId string
	TransferType string
	TokenAmount string
}

type TransactionLedger struct {
	Id string
	TransactionType string
	TransactionIds []string
	AmountToken string
}

type Realestate struct {
	Id string
	Name string
	OwnerId string
	OwnerList map[string]string
	OwnerListArray []AtomTokenTransfer
	TransactionList []string
	IsTrading string
	TransactionPrice string
	Timestamp time.Time
}

type RealestateTransfer interface {
	// Participant method interface
	AddParticipant(shim.ChaincodeStubInterface, *Participant) error
	GetParticipant(shim.ChaincodeStubInterface, string) (*Participant, error)
	UpdateParticipant(shim.ChaincodeStubInterface, *Participant) error
	CheckParticipant(shim.ChaincodeStubInterface, string) (bool, error)
	ValidateParticipant(shim.ChaincodeStubInterface, *Participant) (bool, error)
	ListParticipants(shim.ChaincodeStubInterface) ([]*Participant, error)

	TransferToken(shim.ChaincodeStubInterface, string, string, string) error

	// AtomTokenTransfer method interface
	AddAtomTokenTransfer(shim.ChaincodeStubInterface, *AtomTokenTransfer) error
	AddAtomTokenTransferAttr(shim.ChaincodeStubInterface, string, string, ...string) error
	GetAtomTokenTransfer(shim.ChaincodeStubInterface, string) (*AtomTokenTransfer, error)
	WithdrawnToken(shim.ChaincodeStubInterface, string, string) (string, error)
	DepositToken(shim.ChaincodeStubInterface, string, string) (string, error)
	ListAtomTokenTransfers(shim.ChaincodeStubInterface) ([]*AtomTokenTransfer, error)

	// TransactionLedger method interface
	GenerateId(shim.ChaincodeStubInterface) []byte
	AddTransactionLedger(shim.ChaincodeStubInterface, *TransactionLedger) error
	AddTransactionLedgerbyAttr(shim.ChaincodeStubInterface, string, string, string) error
	CheckTransactionLedger(shim.ChaincodeStubInterface, string) (bool, error)
	ValidateTransactionLedger(shim.ChaincodeStubInterface, *TransactionLedger) (bool, error)
	ListTransactionLedgers(shim.ChaincodeStubInterface) ([]*TransactionLedger, error)
	GetTransactionLedger(shim.ChaincodeStubInterface, string) (*TransactionLedger, error)

	// Realestate method interface
	AddRealestate(shim.ChaincodeStubInterface, *Realestate) error
	CheckRealestate(shim.ChaincodeStubInterface, string) (bool, error)
	ValidateRealestate(shim.ChaincodeStubInterface, *Realestate) (bool, error)
	GetRealestate(shim.ChaincodeStubInterface, string) (*Realestate, error)
	UpdateRealestate(shim.ChaincodeStubInterface, *Realestate) error
	ListRealestates(shim.ChaincodeStubInterface) ([]*Realestate, error)

	TransferRealestate(stub shim.ChaincodeStubInterface, fromParticipantId string, toParticipantId string, realestateId string) error
	//TransferRealestateNToN(stub shim.ChaincodeStubInterface, fromIds []string, toIds []string, tokenAmounts []string, realestateId string) error
}
