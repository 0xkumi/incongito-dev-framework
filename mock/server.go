package mock

import (
	"github.com/incognitochain/incognito-chain/blockchain"
	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/incognitokey"
	"github.com/incognitochain/incognito-chain/wire"
	peer2 "github.com/libp2p/go-libp2p-peer"
)

type Server struct {
	BlockChain *blockchain.BlockChain
}

func (s *Server) PushBlockToAll(block common.BlockInterface, isBeacon bool) error {
	return nil
}

func (s *Server) PushMessageToAll(message wire.Message) error {
	return nil
}

func (s *Server) PushMessageToPeer(message wire.Message, id peer2.ID) error {
	return nil
}

func (s *Server) GetNodeRole() string {
	return ""
}
func (s *Server) EnableMining(enable bool) error {
	return nil
}
func (s *Server) IsEnableMining() bool {
	return true
}
func (s *Server) GetChainMiningStatus(chain int) string {
	return ""
}
func (s *Server) GetPublicKeyRole(publicKey string, keyType string) (int, int) {
	var beaconBestState blockchain.BeaconBestState
	err := beaconBestState.CloneBeaconBestStateFrom(s.BlockChain.GetBeaconBestState())
	if err != nil {
		return -2, -1
	}
	for shardID, pubkeyArr := range beaconBestState.ShardPendingValidator {
		keyList, _ := incognitokey.ExtractPublickeysFromCommitteeKeyList(pubkeyArr, keyType)
		found := common.IndexOfStr(publicKey, keyList)
		if found > -1 {
			return 1, int(shardID)
		}
	}
	for shardID, pubkeyArr := range beaconBestState.ShardCommittee {
		keyList, _ := incognitokey.ExtractPublickeysFromCommitteeKeyList(pubkeyArr, keyType)
		found := common.IndexOfStr(publicKey, keyList)
		if found > -1 {
			return 2, int(shardID)
		}
	}

	keyList, _ := incognitokey.ExtractPublickeysFromCommitteeKeyList(beaconBestState.BeaconCommittee, keyType)
	found := common.IndexOfStr(publicKey, keyList)
	if found > -1 {
		return 2, -1
	}

	keyList, _ = incognitokey.ExtractPublickeysFromCommitteeKeyList(beaconBestState.BeaconPendingValidator, keyType)
	found = common.IndexOfStr(publicKey, keyList)
	if found > -1 {
		return 1, -1
	}

	keyList, _ = incognitokey.ExtractPublickeysFromCommitteeKeyList(beaconBestState.CandidateBeaconWaitingForCurrentRandom, keyType)
	found = common.IndexOfStr(publicKey, keyList)
	if found > -1 {
		return 0, -1
	}

	keyList, _ = incognitokey.ExtractPublickeysFromCommitteeKeyList(beaconBestState.CandidateBeaconWaitingForNextRandom, keyType)
	found = common.IndexOfStr(publicKey, keyList)
	if found > -1 {
		return 0, -1
	}

	keyList, _ = incognitokey.ExtractPublickeysFromCommitteeKeyList(beaconBestState.CandidateShardWaitingForCurrentRandom, keyType)
	found = common.IndexOfStr(publicKey, keyList)
	if found > -1 {
		return 0, -1
	}

	keyList, _ = incognitokey.ExtractPublickeysFromCommitteeKeyList(beaconBestState.CandidateShardWaitingForNextRandom, keyType)
	found = common.IndexOfStr(publicKey, keyList)
	if found > -1 {
		return 0, -1
	}

	return -1, -1
}
func (s *Server) GetIncognitoPublicKeyRole(publicKey string) (int, bool, int) {
	return 0, true, 0
}
func (s *Server) GetMinerIncognitoPublickey(publicKey string, keyType string) []byte {
	return nil
}