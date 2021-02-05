package main

import (
	"fmt"
	"os"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

var (
	project      string
	app          *tview.Application
	pages        *tview.Pages
	actionsModal *ActionsModal
	ribbonView   *tview.TextView
	logView      *tview.TextView
	ipsView      *tview.List
	portsView    *tview.TreeView
	portsRoot    *tview.TreeNode
	hotkeysView  *tview.TextView
	inputField   *tview.InputField
	inputFieldAdded bool
	flexH        *tview.Flex
	flexV        *tview.Flex
)

type OnInputFieldClosedCallback func(tcell.Key)

func CreateLogView() {
	logView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	logView.SetBorder(true)
	logView.SetTitle("Log")
}

func CreatePortsView() {
	portsRoot = tview.NewTreeNode("192.168.0.1")
	portsView = tview.NewTreeView().SetRoot(portsRoot).SetCurrentNode(portsRoot)
	portsView.SetBorder(true)
	portsView.SetTitle("Ports")
	portsView.SetInputCapture(RibbonInputHandler)
	AddRibbonEntry(portsView.Box, "Ports", nil, nil)
}

func CreateInputField() {
	inputFieldAdded = false
	inputField = tview.NewInputField().
		SetLabel("input: ").
		SetFieldWidth(15)
}

func ShowInputField(label string, fieldWidth int, fn OnInputFieldClosedCallback) {
	if inputFieldAdded == false {
		flexV.RemoveItem(hotkeysView)
		flexV.AddItem(inputField, 1, 1, false)
		inputField.SetLabel(label)
		inputField.SetDoneFunc(fn)
		app.SetFocus(inputField)
		inputFieldAdded = true
	}
}

func CloseInputField(key tcell.Key) {
	if inputFieldAdded == true {
		flexV.RemoveItem(inputField)
		flexV.AddItem(hotkeysView, 1, 1, false)
		inputFieldAdded = false
	}
}

func Log(text string) {
	fmt.Fprintf(logView, "%s\n", text)
}

func CreateLayout() *tview.Flex {
	CreateLogView()
	CreateHotkeysView()
	CreateRibbonView()
	CreateIPsView()
	CreatePortsView()
	CreateInputField()
	UpdateRibbon()

	flexV = tview.NewFlex().SetDirection(tview.FlexRow)
	flexH = tview.NewFlex()

	flexH.AddItem(ipsView, 21, 1, false)
	flexH.AddItem(portsView, 0, 2, false)

	flexV.AddItem(ribbonView, 1, 1, false)
	flexV.AddItem(flexH, 0, 1, false)
	flexV.AddItem(logView, 10, 1, false)

	flexV.AddItem(hotkeysView, 1, 1, false)

	return flexV
}

func LoadProject() {
	DeserialiseIPs()
}

func SaveProject() {
	SerialiseIPs()
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: turbine <project_name>")
		return
	}

	project = os.Args[1]

	// TODO: Check that the project name is valid for the OS.
	os.Mkdir(project, 0700)

	app = tview.NewApplication()

	flex := CreateLayout()

	port80 := tview.NewTreeNode("80")
	port8080 := tview.NewTreeNode("8080")
	portsRoot.AddChild(port80)
	portsRoot.AddChild(port8080)

	actionsModal = CreateActionsModal()

	pages = tview.NewPages().
		AddPage("background", flex, true, true).
		AddPage("actionsModal", actionsModal.flex, true, false) // Actions modal starts hidden.

	LoadProject()
	RegisterActions()

	if err := app.SetRoot(pages, true).SetFocus(ipsView).Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}

	SaveProject()
}
