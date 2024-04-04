package types

import (
	"encoding/json"
	"fmt"
	"time"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/gogoproto/proto"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"

	"github.com/forbole/juno/v5/utils"
)

// Validator contains the data of a single validator
type Validator struct {
	ConsAddr   string
	ConsPubKey string
}

// NewValidator allows to build a new Validator instance
func NewValidator(consAddr string, consPubKey string) *Validator {
	return &Validator{
		ConsAddr:   consAddr,
		ConsPubKey: consPubKey,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// CommitSig contains the data of a single validator commit signature
type CommitSig struct {
	Height           int64
	ValidatorAddress string
	VotingPower      int64
	ProposerPriority int64
	Timestamp        time.Time
}

// NewCommitSig allows to build a new CommitSign object
func NewCommitSig(validatorAddress string, votingPower, proposerPriority, height int64, timestamp time.Time) *CommitSig {
	return &CommitSig{
		Height:           height,
		ValidatorAddress: validatorAddress,
		VotingPower:      votingPower,
		ProposerPriority: proposerPriority,
		Timestamp:        timestamp,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// Block contains the data of a single chain block
type Block struct {
	Height          int64
	Hash            string
	TxNum           int
	TotalGas        uint64
	ProposerAddress string
	Timestamp       time.Time
}

// NewBlock allows to build a new Block instance
func NewBlock(
	height int64, hash string, txNum int, totalGas uint64, proposerAddress string, timestamp time.Time,
) *Block {
	return &Block{
		Height:          height,
		Hash:            hash,
		TxNum:           txNum,
		TotalGas:        totalGas,
		ProposerAddress: proposerAddress,
		Timestamp:       timestamp,
	}
}

// NewBlockFromTmBlock builds a new Block instance from a given ResultBlock object
func NewBlockFromTmBlock(blk *tmctypes.ResultBlock, totalGas uint64) *Block {
	return NewBlock(
		blk.Block.Height,
		blk.Block.Hash().String(),
		len(blk.Block.Txs),
		totalGas,
		utils.ConvertValidatorAddressToBech32String(blk.Block.ProposerAddress),
		blk.Block.Time,
	)
}

// -------------------------------------------------------------------------------------------------------------------

// Tx represents an already existing blockchain transaction
type Tx struct {
	*tx.Tx
	*sdk.TxResponse
	Messages []*Message
}

// NewTx allows to create a new Tx instance from the given txResponse
func NewTx(txResponse *sdk.TxResponse, tx *tx.Tx, messages []*Message) (*Tx, error) {
	return &Tx{
		Tx:         tx,
		TxResponse: txResponse,
		Messages:   messages,
	}, nil
}

// MapTransaction allows to build a new Tx instance from the given txResponse and Cosmos transaction
func MapTransaction(txResponse *sdk.TxResponse, tx *tx.Tx, accountParser AccountAddressParser, cdc codec.Codec) (*Tx, error) {
	var messages []*Message
	for i, msg := range tx.GetMsgs() {
		message, err := MapMessage(txResponse.TxHash, txResponse.Height, i, msg, accountParser, cdc)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return NewTx(txResponse, tx, messages)
}

// FindEventByType searches inside the given tx events for the message having the specified index, in order
// to find the event having the given type, and returns it.
// If no such event is found, returns an error instead.
func (tx Tx) FindEventByType(index int, eventType string) (sdk.StringEvent, error) {
	for _, ev := range tx.Logs[index].Events {
		if ev.Type == eventType {
			return ev, nil
		}
	}

	return sdk.StringEvent{}, fmt.Errorf("no %s event found inside tx with hash %s", eventType, tx.TxHash)
}

// FindAttributeByKey searches inside the specified event of the given tx to find the attribute having the given key.
// If the specified event does not contain a such attribute, returns an error instead.
func (tx Tx) FindAttributeByKey(event sdk.StringEvent, attrKey string) (string, error) {
	for _, attr := range event.Attributes {
		if attr.Key == attrKey {
			return attr.Value, nil
		}
	}

	return "", fmt.Errorf("no event with attribute %s found inside tx with hash %s", attrKey, tx.TxHash)
}

// Successful tells whether this tx is successful or not
func (tx Tx) Successful() bool {
	return tx.TxResponse.Code == 0
}

// -------------------------------------------------------------------------------------------------------------------

// Message represents the data of a single message
type Message struct {
	TxHash    string
	Index     int
	Type      string
	Value     string
	Addresses []string
	Height    int64
}

// NewMessage allows to build a new Message instance
func NewMessage(txHash string, index int, msgType string, value string, addresses []string, height int64) *Message {
	return &Message{
		TxHash:    txHash,
		Index:     index,
		Type:      msgType,
		Value:     value,
		Addresses: addresses,
		Height:    height,
	}
}

// MapMessage allows to build a new Message instance from the given tx data, index and Cosmos message
func MapMessage(txHash string, txHeight int64, index int, msg sdk.Msg, accountParser AccountAddressParser, cdc codec.Codec) (*Message, error) {
	messageBz, err := cdc.MarshalJSON(msg)
	if err != nil {
		return nil, err
	}

	messageValue := string(messageBz)
	if msgReceivePacket, ok := msg.(*channeltypes.MsgRecvPacket); ok {
		trimMessageString := utils.TrimLastChar(string(messageBz))
		trimDataString := string(msgReceivePacket.Packet.Data)[1:]
		messageValue = fmt.Sprintf("%s,%s", trimMessageString, trimDataString)
	}

	accountAddresses, err := extractAccountAddresses(messageBz, accountParser)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		txHash,
		index,
		proto.MessageName(msg),
		messageValue,
		accountAddresses,
		txHeight,
	), nil
}

// extractAccountAddresses extracts the account addresses from the given message bytes
func extractAccountAddresses(messageBz []byte, isAccountAddress AccountAddressParser) ([]string, error) {
	var messageData map[string]interface{}
	err := json.Unmarshal(messageBz, &messageData)
	if err != nil {
		return nil, err
	}

	var addresses []string
	for key, value := range messageData {
		if stringValue, ok := value.(string); ok {
			if isAccountAddress(key) {
				addresses = append(addresses, stringValue)
			}
		}
	}

	return addresses, nil
}
