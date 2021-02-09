package main

import (
	"os/exec"
)

type ActionAvailableFn func() bool
type ActionCommandFn func()

type Action struct {
	Name        string
	Description string
	NodeType    string
	CommandFn   ActionCommandFn
}

var actions []Action

func RegisterActions() {
	RegisterActivePortScanAction()
	RegisterPassivePortScanAction()
}

func RegisterAction(action Action) {
	actions = append(actions, action)
}

func RegisterPassivePortScanAction() {
	action := Action{
		"Passive port scan",
		"Use Shodan to retrieve port scan information.",
		"IP",
		nil,
	}
	actions = append(actions, action)
}

