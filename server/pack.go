package server

import (
	"encoding/json"
)

const (
	PackError     = "error"
	PackSuccess   = "success"
	PackNotDefine = "not_define"
)

//Pack :
type Pack struct {
	Status string          `json:"stt"`
	Data   json.RawMessage `json:"data"`
}

func NewPack(stt string, dt []byte) *Pack {
	var res = &Pack{Status: stt, Data: dt}
	return res
}
