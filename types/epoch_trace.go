package types

import (
	"github.com/ethereum/go-ethereum/common"
	web3gotypes "github.com/openweb3/web3go/types"
	"github.com/ukeyiogavin/go-conflux-sdk/types/cfxaddress"
)

type EpochTrace struct {
	CfxTraces        []*LocalizedTrace                     `json:"cfxTraces"`
	EthTraces        []*web3gotypes.LocalizedTrace         `json:"ethTraces"`
	MirrorAddressMap map[common.Address]cfxaddress.Address `json:"mirrorAddressMap"`
}
