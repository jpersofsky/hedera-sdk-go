package hedera

import (
	"github.com/pkg/errors"
	"time"

	protobuf "github.com/golang/protobuf/proto"
	"github.com/hashgraph/hedera-sdk-go/v2/proto"
)

// AccountInfo is info about the account returned from an AccountInfoQuery
type AccountInfo struct {
	AccountID                      AccountID
	ContractAccountID              string
	IsDeleted                      bool
	ProxyAccountID                 AccountID
	ProxyReceived                  Hbar
	Key                            Key
	Balance                        Hbar
	GenerateSendRecordThreshold    Hbar
	GenerateReceiveRecordThreshold Hbar
	ReceiverSigRequired            bool
	ExpirationTime                 time.Time
	AutoRenewPeriod                time.Duration
	TokenRelationships             []*TokenRelationship
}

func accountInfoFromProtobuf(pb *proto.CryptoGetInfoResponse_AccountInfo) (AccountInfo, error) {
	pubKey, err := keyFromProtobuf(pb.Key)
	if err != nil {
		return AccountInfo{}, err
	}

	tokenRelationship := make([]*TokenRelationship, len(pb.TokenRelationships))

	if pb.TokenRelationships != nil {
		for i, relationship := range pb.TokenRelationships {
			singleRelationship := tokenRelationshipFromProtobuf(relationship)
			tokenRelationship[i] = &singleRelationship
		}
	}

	return AccountInfo{
		AccountID:                      accountIDFromProtobuf(pb.AccountID),
		ContractAccountID:              pb.ContractAccountID,
		IsDeleted:                      pb.Deleted,
		ProxyAccountID:                 accountIDFromProtobuf(pb.ProxyAccountID),
		ProxyReceived:                  HbarFromTinybar(pb.ProxyReceived),
		Key:                            pubKey,
		Balance:                        HbarFromTinybar(int64(pb.Balance)),
		GenerateSendRecordThreshold:    HbarFromTinybar(int64(pb.GenerateSendRecordThreshold)),
		GenerateReceiveRecordThreshold: HbarFromTinybar(int64(pb.GenerateReceiveRecordThreshold)),
		ReceiverSigRequired:            pb.ReceiverSigRequired,
		TokenRelationships:             tokenRelationship,
		ExpirationTime:                 timeFromProtobuf(pb.ExpirationTime),
	}, nil
}

func (info AccountInfo) toProtobuf() ([]byte, error) {

	tokenRelationship := make([]*proto.TokenRelationship, len(info.TokenRelationships))

	for i, relationship := range info.TokenRelationships {
		singleRelationship := relationship.toProtobuf()
		tokenRelationship[i] = singleRelationship
	}

	pb, err := protobuf.Marshal(&proto.CryptoGetInfoResponse_AccountInfo{
		AccountID:                      info.AccountID.toProtobuf(),
		ContractAccountID:              info.ContractAccountID,
		Deleted:                        info.IsDeleted,
		ProxyAccountID:                 info.ProxyAccountID.toProtobuf(),
		ProxyReceived:                  info.ProxyReceived.tinybar,
		Key:                            info.Key.toProtoKey(),
		Balance:                        uint64(info.Balance.tinybar),
		GenerateSendRecordThreshold:    uint64(info.GenerateSendRecordThreshold.tinybar),
		GenerateReceiveRecordThreshold: uint64(info.GenerateReceiveRecordThreshold.tinybar),
		ReceiverSigRequired:            info.ReceiverSigRequired,
		TokenRelationships:             tokenRelationship,
		ExpirationTime:                 timeToProtobuf(info.ExpirationTime),
	})

	if err != nil {
		return pb, errors.Wrap(err, "error serializing account info")
	} else {
		return pb, nil
	}
}
