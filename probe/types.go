package probe

import (
	"strconv"
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

