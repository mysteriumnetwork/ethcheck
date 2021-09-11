package probe

import (
	"strconv"

	"github.com/mysteriumnetwork/ethcheck/util"
)

type BlockNumberResult uint64

func (r BlockNumberResult) String() string {
	return "0x" + strconv.FormatUint(uint64(r), 16)
}

func (r *BlockNumberResult) UnmarshalText(text []byte) error {
	parsed, err := strconv.ParseUint(string(text), 0, 64)
	if err != nil {
		return err
	}

	*r = BlockNumberResult(parsed)
	return nil
}

func (r BlockNumberResult) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

type ShortBlockInfo struct {
	Timestamp BlockNumberResult `json:"timestamp"`
}

type RandomEthCallParamObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}

func NewRandomEthCallParamObject() (*RandomEthCallParamObject, error) {
	randomHexString, err := util.RandomHexString(20)
	if err != nil {
		return nil, err
	}
	return &RandomEthCallParamObject{
		Data: "0xe617aaac0000000000000000000000007cadbf6d95c81754b84057f2eb634acc2fd8e4330000000000000000000000007119442c7e627438deb0ec59291e31378f88dd06",
		From: "0x0000000000000000000000000000000000000000",
		To:   "0x" + randomHexString,
	}, nil
}
