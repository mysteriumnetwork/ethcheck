package rpc

import (
	"encoding/json"
	"fmt"
)

type ErrorObject struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Response struct {
	Version string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *ErrorObject    `json:"error,omitempty"`
	ID      RequestID       `json:"id"`
}

func (e *ErrorObject) Error() string {
	return fmt.Sprintf(
		"RPC call error. code: %d, message: %s, data: %v",
		e.Code,
		e.Message,
		e.Data,
	)
}

func (r *Response) UnmarshalResult(dst interface{}) error {
	if r.Error != nil {
		return r.Error
	}

	// No actual marshalling happens because RawMessage just holds original
	// JSON subtree
	resBytes, err := r.Result.MarshalJSON()
	if err != nil {
		return err
	}

	return json.Unmarshal(resBytes, dst)
}
