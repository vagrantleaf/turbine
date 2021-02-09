package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	NodeType string
	Name     string
}

func (node Node) Serialise() {
	res, err := json.Marshal(node)

	if err != nil {
		Log(fmt.Sprintf("%s", err))
	}

	Log(string(res))
}
