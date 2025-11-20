package types

import (
	"fmt"
	"testing"

	"github.com/ukeyiogavin/go-conflux-sdk/utils"
)

func TestUnmarshal(t *testing.T) {

	cases := []string{
		`{"amount":"0x64","unlockBlockNumber":200}`,
		`{"amount":"0x64","unlockBlockNumber":"0x64"}`,
		fmt.Sprintf(`{"amount":"0x64","unlockBlockNumber":%v}`, ^uint64(0)),
	}

	for _, c := range cases {
		v := VoteStakeInfo{}
		e := utils.JsonUnmarshal([]byte(c), &v)
		if e != nil {
			t.Fatal(e)
		}
		// fmt.Println(v)
	}
}
