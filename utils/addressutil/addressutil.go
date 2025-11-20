package addressutil

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/ukeyiogavin/go-conflux-sdk/types/cfxaddress"
	"github.com/ukeyiogavin/go-conflux-sdk/utils"
)

// EtherAddressToCfxAddress converts an ethereum address to a cfx address, it will change top 4 bit of eth address to 0b0001
func EtherAddressToCfxAddress(ethAddr common.Address, isContract bool, networkID uint32) cfxaddress.Address {
	if isContract {
		ethAddr[0] = ethAddr[0]&0x8f | 0x80
	} else {
		ethAddr[0] = ethAddr[0]&0x1f | 0x10
	}
	return cfxaddress.MustNewFromCommon(ethAddr, networkID)
}

func PubkeyToAddress(publicKey string, networkId uint32) (cfxaddress.Address, error) {
	commAddress, err := utils.PublicKeyToCommonAddress(publicKey)
	if err != nil {
		return cfxaddress.Address{}, errors.WithStack(err)
	}
	return cfxaddress.NewFromCommon(commAddress, networkId)
}
