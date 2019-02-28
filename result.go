package main

import "encoding/json"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func MarshalToJson(r Result) ([]byte, error) {
	return json.Marshal(&r)
}

func UnMarshalFromJson(data []byte) (*Result, error) {
	r := &Result{}

	err := json.Unmarshal(data, r)

	if err != nil {
		return nil, err
	}

	return r, nil
}
