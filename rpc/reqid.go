package rpc

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const (
	RequestIDKindNull = iota
	RequestIDKindNumber
	RequestIDKindString
)

var ErrorBadIDType = errors.New("bad RPC request ID type: must be null or string or integer")

type RequestID struct {
	value interface{}
}

func (id RequestID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *RequestID) UnmarshalJSON(b []byte) error {
	var valuePreview interface{}
	err := json.Unmarshal(b, &valuePreview)
	if err != nil {
		return err
	}

	switch v := valuePreview.(type) {
	case string:
		id.value = v
	case nil:
		id.value = nil
	case float64:
		var intValue int64
		err = json.Unmarshal(b, &intValue)
		if err != nil {
			return err
		}
		id.value = intValue
	default:
		return ErrorBadIDType
	}

	return nil
}

func (id RequestID) Kind() int {
	switch id.value.(type) {
	case nil:
		return RequestIDKindNull
	case int64:
		return RequestIDKindNumber
	case string:
		return RequestIDKindString
	}
	panic(ErrorBadIDType)
}

func (id RequestID) Value() interface{} {
	return id.value
}

func (id RequestID) AsInt64() int64 {
	return id.value.(int64)
}

func (id RequestID) AsString() string {
	return id.value.(string)
}

func (id RequestID) IsZero() bool {
	return id.value == nil
}

func (id RequestID) String() string {
	return fmt.Sprintf("%v", id.value)
}

func RequestIDFromInt64(id int64) RequestID {
	return RequestID{
		value: id,
	}
}

func RequestIDFromString(id string) RequestID {
	return RequestID{
		value: id,
	}
}

func NewUniqueRequestID() (RequestID, error) {
	reqUUID, err := uuid.NewRandom()
	if err != nil {
		return RequestID{}, err
	}

	return RequestID{
		value: reqUUID.String(),
	}, nil
}
