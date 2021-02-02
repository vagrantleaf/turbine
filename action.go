package main

import (
	"encoding/json"
	//"log"
	//"os"
	"path/filepath"
	//"bufio"
	"io/ioutil"
	"fmt"
)


type Action struct {
	Name string
	NodeType string
	Command string
}

var actions []Action

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


