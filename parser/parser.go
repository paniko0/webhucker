package parser

import "encoding/json"

type Received struct {
	RedirectUri string `json:"redirectUri"`
	Resource    string
}

func Parse(receivedJson string) *Received {
	res := &Received{}
	json.Unmarshal([]byte(receivedJson), &res)
	return res
}
