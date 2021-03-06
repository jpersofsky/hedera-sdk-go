package main

import (
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func main() {
	var client *hedera.Client
	var err error

	if os.Getenv("HEDERA_NETWORK") == "previewnet" {
		client = hedera.ClientForPreviewnet()
	} else {
		client, err = hedera.ClientFromConfigFile(os.Getenv("CONFIG_FILE"))

		if err != nil {
			client = hedera.ClientForTestnet()
		}
	}

	configOperatorID := os.Getenv("OPERATOR_ID")
	configOperatorKey := os.Getenv("OPERATOR_KEY")

	if configOperatorID != "" && configOperatorKey != "" && client.GetOperatorPublicKey().Bytes() == nil {
		operatorAccountID, err := hedera.AccountIDFromString(configOperatorID)
		if err != nil {
			panic(err)
		}

		operatorKey, err := hedera.PrivateKeyFromString(configOperatorKey)
		if err != nil {
			panic(err)
		}

		client.SetOperator(operatorAccountID, operatorKey)
	}

	initialAdminKeys := make([]hedera.PrivateKey, 3)
	for i, _  := range initialAdminKeys {
		key, err := hedera.GeneratePrivateKey()
		if err != nil {
			panic(err)
		}
		initialAdminKeys[i] = key
	}

	keyList := hedera.KeyListWithThreshold(2)
	for _, key  := range initialAdminKeys {
		keyList.Add(key.PublicKey())
	}

	topicTx, err := hedera.NewTopicCreateTransaction().
		SetTopicMemo("demo topic").
		SetAdminKey(keyList).
		FreezeWith(client)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {
		println("Signing ConsensusTopicCreateTransaction with key ", initialAdminKeys[i].String())
		topicTx.Sign(initialAdminKeys[i])
	}

	response, err := topicTx.Execute(client)
	if err != nil {
		panic(err)
	}

	receipt, err := response.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	topicID := *receipt.TopicID

	println("Created new topic ",  topicID.String(), " with 2-of-3 threshold key as adminKey.")

	newAdminKeys := make([]hedera.PrivateKey, 4)
	for i, _  := range newAdminKeys {
		key, err := hedera.GeneratePrivateKey()
		if err != nil {
			panic(err)
		}
		newAdminKeys[i] = key
	}

	keyList = hedera.KeyListWithThreshold(3)
	for _, key  := range newAdminKeys {
		keyList.Add(key.PublicKey())
	}

	topicUpdate, err := hedera.NewTopicUpdateTransaction().
		SetTopicID(topicID).
		SetTopicMemo("updated demo topic").
		SetAdminKey(keyList).
		FreezeWith(client)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {
		println("Signing ConsensusTopicCreateTransaction with initial admin key ", initialAdminKeys[i].String())
		topicUpdate.Sign(initialAdminKeys[i])
	}

	for i := 0; i < 3; i++ {
		println("Signing ConsensusTopicCreateTransaction with new admin key ", newAdminKeys[i].String())
		topicUpdate.Sign(newAdminKeys[i])
	}

	response, err = topicUpdate.Execute(client)
	if err != nil {
		panic(err)
	}

	receipt, err = response.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	println("Updated topic ", topicID.String(), " with 3-of-4 threshold key as adminKey");

	topicInfo, err := hedera.NewTopicInfoQuery().
		SetTopicID(topicID).
		Execute(client)
	if err != nil {
		panic(err)
	}

	println(topicInfo.Memo)
}
