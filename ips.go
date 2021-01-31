package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"bufio"
	"io/ioutil"
)

type IPNode struct {
	Ip    string
	State string
	Note  string
	Node
}

var ipNodes []IPNode

func SerialiseIPs() {
	res, _ := json.Marshal(ipNodes)

	ipsFilePath := filepath.Join(project, "ips.json")

	file, fileErr := os.Create(ipsFilePath)
	if fileErr != nil {
		log.Fatal(fileErr)
		return
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(string(res))
	writer.Flush()
}

func DeserialiseIPs() {
	ipsFilePath := filepath.Join(project, "ips.json")

	data, fileErr := ioutil.ReadFile(ipsFilePath)
	if fileErr != nil {
		log.Fatal(fileErr)
		return
	}

	unmarshalErr := json.Unmarshal([]byte(data), &ipNodes)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
		return
	}
	
	// Add all IPs nodes to the panel after we load them from disk.
	for _, ipNode := range(ipNodes) {
		noteTag := '.'
		if len(ipNode.Note) > 0 {
			noteTag = 'N'
		}

		ipsView.AddItem(ipNode.Ip, ipNode.State, noteTag, func() {
			ShowActions(ipNode)
		})
	}
}

func CreateIPsView() {
	ipsView = tview.NewList()
	ipsView.SetBorder(true)
	ipsView.SetTitle("IPs")
	ipsView.SetInputCapture(IPsInputHandler)
	AddRibbonEntry(ipsView.Box, "IPs", OnIPsViewFocused, OnIpsViewUnfocused)
}

func OnIPsViewFocused() {
	AddHotkeyEntry("Add IP", 'a', OnAddIPHotkey)
	AddHotkeyEntry("Delete IP", 'd', func(){ Log("Delete IP") })
}

func OnIpsViewUnfocused() {
	RemoveHotkeyEntry("Add IP")
	RemoveHotkeyEntry("Delete IP")
}

func OnAddIPHotkey() {
	ShowInputField("IP to add: ", 15, func(key tcell.Key){
		CloseInputField(key)
		app.SetFocus(ipsView)
	})
}

func IPsInputHandler(event *tcell.EventKey) *tcell.EventKey {
	event = RibbonInputHandler(event)
	if event == nil {
		return nil
	}

	//if event.Key() == tcell.KeyEnter {
	//	return nil
	//}

	return event
}

func AddIPNode(ip string, state string, note string) {
	ipNode := IPNode{
		Ip: ip, 
		State: state, 
		Note: note,
		Node: Node{
			NodeType: "IP",
			Name: ip, 
		},
	}
	ipNodes = append(ipNodes, ipNode)

	noteTag := '.'
	if len(note) > 0 {
		noteTag = 'N'
	}

	ipsView.AddItem(ip, state, noteTag, func() {
		ShowActions(ipNode)
	})
}

func ShowActions(ipNode IPNode) {
	Log("Showing actions")

	modal := tview.NewModal().
		SetText("Action to run").
		AddButtons([]string{"Nmap", "Reverse DNS lookup"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			Log(buttonLabel)
		})
	app.SetFocus(modal)
}
