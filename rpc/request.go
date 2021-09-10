package rpc

type Request struct {
	ID      RequestID   `json:"id,omitempty"`
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method,omitempty"`
	Params  interface{} `json:"params,omitempty"`
}

func NewRequest(id RequestID, method string, params interface{}) *Request {
	return &Request{
		ID:      id,
		Version: "2.0",
		Method:  method,
		Params:  params,
	}
}

func NewUniqueRequest(method string, params interface{}) (*Request, error) {
	id, err := NewUniqueRequestID()
	if err != nil {
		return nil, err
	}

	return NewRequest(id, method, params), nil
}
