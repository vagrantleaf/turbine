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

func RegisterActivePortScanAction() {
	if IsActivePortScanAvailable() {
		action := Action{
			"Active port scan",
			"Comprehensive TCP port scan with service discovery.",
			"IP",
			ActivePortScanCommand,
		}
		actions = append(actions, action)
	}
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

func IsActivePortScanAvailable() bool {
	_, err := exec.LookPath("nmap")
	if err != nil {
		Log("WARNING: Active Port Scan action not available: couldn't find 'nmap'.")
		return false
	}
	return true
}

func ActivePortScanCommand() {
	Log("Running active port scan")
	ActivePortScanCompleted()
}

func ActivePortScanCompleted() {

}

/*
func DeserialiseActions() {
	actionsDir := "data/actions"
	files, dirErr := ioutil.ReadDir(actionsDir)
	if dirErr != nil {
		Log(fmt.Sprintf("Error reading '%s' directory: %s", actionsDir, dirErr.Error()))
		return
	}

	var jsonFiles = 0
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		jsonFiles++

		actionPath := filepath.Join(actionsDir, file.Name())
		data, fileErr := ioutil.ReadFile(actionPath)
		if fileErr != nil {
			Log(fmt.Sprintf("Couldn't read JSON file: %s (%s)", actionPath, fileErr.Error()))
			continue
		}

		var action Action
		unmarshalErr := json.Unmarshal([]byte(data), &action)
		if unmarshalErr != nil {
			Log(fmt.Sprintf("Couldn't unmarshal JSON file: %s (%s)", actionPath, unmarshalErr.Error()))
			continue
		} else {
			actions = append(actions, action)
		}
	}

	Log(fmt.Sprintf("Loaded %d / %d actions.", len(actions), jsonFiles))
}
*/

